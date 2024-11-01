package freeport

import (
	"fmt"
	"net"
	"sync"
)

func getFreePort() (int, func() error, error) {
	l, err := net.Listen("tcp", `:0`)
	if err != nil {
		return 0, noop, fmt.Errorf(`unable to find free port: %w`, err)
	}

	close := func() error {
		if err := l.Close(); err != nil {
			return fmt.Errorf(`unable to close listener after finding free port: %w`, err)
		}
		return nil
	}

	return l.Addr().(*net.TCPAddr).Port, close, err
}

// GetFreePort asks the kernel for a free open port that is ready to use.
func GetFreePort() (int, error) {
	port, close, err := getFreePort()
	if err != nil {
		return 0, err
	}

	if err := close(); err != nil {
		return 0, err
	}

	return port, nil
}

// GetPort is deprecated and included for backwards compatibility.
// It works like MustGetFreePort.
func GetPort() int {
	return MustGetFreePort()
}

// MustGetFreePort is like GetFreePort but panics on error.
func MustGetFreePort() int {
	port, close, err := getFreePort()
	if err != nil {
		panic(err)
	}

	if err := close(); err != nil {
		panic(err)
	}

	return port
}

// GetFreePorts is like GetFreePort but gets multiple ports.
func GetFreePorts(n int) ([]int, error) {
	ports := make([]int, n)
	close := make([]func() error, n)

	var err error
	for i := 0; i < n; i++ {
		var er error
		ports[i], close[i], er = getFreePort()
		err = wrapif(err, er)
	}

	for _, c := range close {
		err = wrapif(err, c())
	}

	return ports, err
}

func wrapif(err, er error) error {
	switch {
	case err == nil && er == nil:
		return nil
	case err != nil && er == nil:
		return err
	case err == nil && er != nil:
		return er
	}

	return fmt.Errorf(`%w: %s`, err, er)
}

func noop() error {
	return nil
}

// GetFreePortsFromRange get free ports from a start to an end
func GetFreePortsFromRange(start int, end int) ([]int, error) {
	count := end - start + 1
	var ports []int
	portChan := make(chan int, count)
	errChan := make(chan error, count)

	go func() {
		wg := sync.WaitGroup{}
		for i := start; i < end-1; i++ {
			wg.Add(1)
			go func(port int) {
				defer wg.Done()
				addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("0.0.0.0:%d", port))
				if err != nil {
					errChan <- err
					return
				}

				l, err := net.ListenTCP("tcp", addr)
				if err != nil {
					errChan <- err
					return
				}
				defer l.Close()
				portChan <- l.Addr().(*net.TCPAddr).Port
			}(i)
		}
		wg.Wait()
		close(portChan)
		close(errChan)
	}()

	for err := range errChan {
		if err != nil {
			return nil, err
		}
	}

	for port := range portChan {
		ports = append(ports, port)
	}

	return ports, nil
}
