package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync/atomic"
	"time"
)

// This hold the state of the player
type Player struct {
	// mu sync.RWMutex
	// health int
	health int32
}

// Note that we do not have to create an initialization for the mutex value
func NewPlayer() *Player {
	return &Player{
		health: 100,
	}
}

func (p *Player) getHealth() int {
	// (Read) Lock the struct to ensure no routine has access to change it's value
	// p.mu.RLock()
	// defer p.mu.RUnlock()

	// return p.health
	return int(atomic.LoadInt32(&p.health))
}

func (p *Player) takeDamage(value int) {
	// (Write) Lock the struct before modifying it's value to ensure no routine is currently reading
	// p.mu.Lock()
	// defer p.mu.Unlock()

	// p.health -= value
	atomic.StoreInt32(&p.health, p.health-int32(value))
}

func StartGameUI(player *Player) {
	ticker := time.NewTicker(time.Second)
	for {
		os.Stdout.Sync()

		fmt.Printf("Player health: %v\t\r", player.getHealth())

		<-ticker.C
	}
}

func StartGameLoop(player *Player) {
	for {
		timer := time.NewTimer(time.Millisecond * 600)

		player.takeDamage(rand.Intn(20))
		if player.getHealth() <= 0 {
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
