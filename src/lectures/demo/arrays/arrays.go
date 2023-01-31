package main

import "fmt"

type Room struct {
	name    string
	cleaned bool
}

func checkCleanLiness(rooms [4]Room) {
	for i := 0; i < len(rooms); i++ {
		room := rooms[i]
		if room.cleaned {
			fmt.Println(room.name, "is clean!!")
		} else {
			fmt.Println(room.name, "is dirty!!")
		}
	}
}

func main() {
	rooms := [...]Room{
		{name: "one", cleaned: true},
		{name: "two", cleaned: false},
		{name: "tree", cleaned: false},
		{name: "four", cleaned: true},
	}
	checkCleanLiness(rooms)
}
