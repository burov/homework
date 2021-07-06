package waitfor

import (
	"fmt"
	"net"
	"time"

	"github.com/pkg/errors"
)

var ErrTimeout = errors.New("exited by timeout")

func It(host string, port int, timeout int) error {
	addr := fmt.Sprintf("%s:%d", host, port)
	ch := make(chan struct{})
	//errCh := make(chan error)
	go func() {
		step := 0
		for {
			fmt.Printf("Asking %s %2d of %2d\n", addr, step, timeout)
			step++
			c, err := net.Dial("tcp", addr)
			if err == nil {
				c.Close()
				ch <- struct{}{}
			}
			time.Sleep(time.Second)
		}
	}()
	select {
	case <-ch:
		return nil
	case <-time.After(time.Duration(timeout) * time.Second):
		return ErrTimeout
	}
}
