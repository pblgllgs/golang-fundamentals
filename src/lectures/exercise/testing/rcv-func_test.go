//--Summary:
//  Copy your rcv-func solution to this directory and write unit tests.
//
//--Requirements:
//* Write unit tests that ensure:
//  - Health & energy can not go above their maximums
//  - Health & energy can not go below 0
//* If any of your  tests fail, make the necessary corrections
//  in the copy of your rcv-func solution file.
//
//--Notes:
//* Use `go test -v ./exercise/testing` to run these specific tests

package testing

import (
	"fmt"
	"testing"
)

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
		player.energy = 0
		fmt.Println("The consume is", energy, "->", "the new energy is:", player.energy)
	} else {
		player.energy -= energy
		fmt.Println("The consume is", energy, "->", "the new energy is:", player.energy)
	}
}

func TestNotMaximums(t *testing.T) {
	player := Player{name: "MAX", energy: 500, maxEnergy: 500, health: 100, maxHealth: 100}
	player.addEnergy(410)
	player.addHealth(410)
	if player.energy > 500 {
		t.Errorf("La energia no debe ser mayor que 500")
	} else if player.health > 100 {
		t.Errorf("La health no debe ser mayor que 100")
	}
}

func TestNotBeCero(t *testing.T) {
	player := Player{name: "MAX", energy: 500, maxEnergy: 500, health: 100, maxHealth: 100}
	player.consumeEnergy(510)
	player.consumeHealth(110)
	if player.energy < 0 {
		t.Errorf("La energia no debe ser menor que 0")
	} else if player.health < 0 {
		t.Errorf("La health no debe ser menor que 0")
	}
}
