package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

// This hold the state of the player
type Player struct {
	health int
}

// Note that we do not have to create an initialization for the mutex value
func NewPlayer() *Player {
	return &Player{
		health: 100,
	}
}

func StartGameUI(player *Player) {
	ticker := time.NewTicker(time.Second)
	for {
		os.Stdout.Sync()
		fmt.Printf("Player health: %v\t\r", player.health)
		<-ticker.C
	}
}

func StartGameLoop(player *Player) {
	for {
		timer := time.NewTimer(time.Millisecond * 600)
		player.health -= rand.Intn(20)
		if player.health <= 0 {
			fmt.Println("Game Over")
			return
		}
		<-timer.C
	}
}

func SyncMem() {
	player := NewPlayer()
	go StartGameUI(player)
	StartGameLoop(player)
}
