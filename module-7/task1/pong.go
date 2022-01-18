package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	maxScore         = 15
	player1, player2 = "Player 1", "Player 2"
)

func play(players [2]chan struct{}) (winner string) {
	var points [2]int
	var wg sync.WaitGroup

	players[0] <- struct{}{}
	wg.Add(2)

	ping := func(player1, player2 chan struct{}) {
		for {
			if gameEnded(points) {
				wg.Done()
				return
			}
			<-player1
			if ballMissed() {
				fmt.Println("Ball lost by", player1)
				points[1]++
			}
			player2 <- struct{}{}
		}
	}

	go ping(players[0], players[1])
	go ping(players[1], players[0])

	for {
		if gameEnded(points) {
			wg.Wait()
			if points[0] > points[1] {
				return player1
			} else if points[0] < points[1] {
				return player2
			}
		}
	}
}

func ballMissed() bool {
	now := time.Now()
	fmt.Println(now.UnixNano())
	return now.UnixNano()%9 == 0
}

func gameEnded(points [2]int) bool {
	for _, p := range points {
		if p > maxScore {
			return true
		}
	}

	return false
}

func PlayGame() string {
	players := [2]chan struct{}{
		make(chan struct{}, 1),
		make(chan struct{}, 1),
	}

	defer func() {
		for _, p := range players {
			close(p)
		}
	}()

	return play(players)
}
