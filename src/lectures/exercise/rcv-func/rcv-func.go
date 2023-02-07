//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package main

import "fmt"

type Player struct {
	health    int
	maxHealth int
	energy    int
	maxEnergy int
	name      string
}

const (
	MAX_ENERGY = 500
	MAX_HEALTH = 100
)

func (player *Player) addHealth(health int) {
	if health+player.health > MAX_HEALTH {
		player.health = MAX_HEALTH
	} else {
		player.health += health
		fmt.Println("add", health, "of health ->", "the new health is:", player.health)
	}

}

func (player *Player) consumeHealth(health int) {
	if health > player.health {
		player.health = 0
		fmt.Println("The damage is", health, "->", "the new health is:", player.health)
	} else {
		player.health -= health
		fmt.Println("The damage is", health, "->", "the new health is:", player.health)
	}
}

func (player *Player) addEnergy(energy int) {
	if energy+player.energy > MAX_ENERGY {
		player.energy = MAX_ENERGY
	} else {
		player.energy += energy
		fmt.Println("add", energy, "of energy ->", "the new energy is:", player.energy)
	}
}

func (player *Player) consumeEnergy(energy int) {
	if energy > player.energy {
		player.health = 0
	} else {
		player.energy -= energy
		fmt.Println("The consume is", energy, "->", "the new energy is:", player.energy)
	}
}

func main() {

	player1 := Player{name: "MAX", energy: 500, maxEnergy: 500, health: 100, maxHealth: 100}
	player1.consumeHealth(99)
	// player1.info()
	player1.addHealth(10)
	// player1.info()
	player1.consumeEnergy(20)
	// player1.info()
	player1.addEnergy(10)
	// player1.info()
	player1.consumeHealth(9999)
}
