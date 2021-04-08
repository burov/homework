package main

import (
	"fmt"
	"testing"
	"time"
)

func TestDummy(t *testing.T) {
	var (
		winner = make(chan string)
	)

	go func() {
		//TODO)
		wp := PlayGame()
		winner <- wp
	}()

	select {
	case p := <-winner:
		if p == "" {
			t.Errorf("winner is undefined")
			return
		}
		fmt.Printf("Winner is %s\n", p)

	case <-time.After(5 * time.Second):
		t.Error("We're waiting too long, it seems like something wrong with the code")
	}

}
