package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

// This hold the state of the player
type Player struct {
	mu     sync.RWMutex
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

		// (Read) Lock the struct to ensure no routine has access to change it's value
		player.mu.RLock()

		fmt.Printf("Player health: %v\t\r", player.health)

		player.mu.RUnlock()

		<-ticker.C
	}
}

func StartGameLoop(player *Player) {
	for {
		timer := time.NewTimer(time.Millisecond * 600)

		// (Write) Lock the struct before modifying it's value to ensure no routine is currently reading
		player.mu.Lock()

		player.health -= rand.Intn(20)
		if player.health <= 0 {
			fmt.Println("Game Over")
			return
		}

		player.mu.Unlock()
		<-timer.C
	}
}

func SyncMem() {
	player := NewPlayer()
	go StartGameUI(player)
	StartGameLoop(player)
}
