package main

import "fmt"

type Space struct {
	occupied bool
}

type ParkingSlot struct {
	spaces []Space
}

func occupySpace(lot *ParkingSlot, spaceNum int) {
	lot.spaces[spaceNum-1].occupied = true
}

func (lot *ParkingSlot) occupySpace(spaceNum int) {
	lot.spaces[spaceNum-1].occupied = true
}

func (lot *ParkingSlot) vacateSpace(spaceNum int) {
	lot.spaces[spaceNum-1].occupied = false
}

func main() {

	lot := ParkingSlot{spaces: make([]Space, 4)}
	fmt.Println(lot)
	lot.occupySpace(1)

	occupySpace(&lot, 2)
	fmt.Println(lot)

	lot.vacateSpace(2)
	fmt.Println(lot)
}
