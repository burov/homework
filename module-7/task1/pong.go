package main

import (
	"fmt"
	"time"
)

const (
	maxScore         = 15
	player1, player2 = "Player 1", "Player2"
)

func play(players [2]chan struct{}) (winner string) {
	var points [2]int
	_ = points

	return ""
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
