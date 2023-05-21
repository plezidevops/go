package main

import "fmt"

type Player struct {
	name        string
	health      int
	attackPower float64
}

func (player Player) getHealth() int {
	return player.health
}

func main() {
	player := Player{
		name:        "captain jack",
		health:      1,
		attackPower: 3.2,
	}

	fmt.Printf("The player health %d\n", player.getHealth())
	fmt.Printf("This is the player %+v\n", player)

}
