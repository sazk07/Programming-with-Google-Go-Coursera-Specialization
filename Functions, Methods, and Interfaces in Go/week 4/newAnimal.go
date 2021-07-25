package main

import (
	"fmt"
)

type animalProperties struct {
	food       string
	locomotion string
	noise      string
}

type Animal interface {
	Eat()
	Move()
	Speak()
}

func (ani animalProperties) Eat() {
	fmt.Println(ani.food)
	return
}

func (ani animalProperties) Move() {
	fmt.Println(ani.locomotion)
	return
}

func (ani animalProperties) Speak() {
	fmt.Println(ani.noise)
	return
}

func main() {
	animalMap := make(map[string]animalProperties)
	animalMap["cow"] = animalProperties{"grass", "walk", "moo"}
	animalMap["bird"] = animalProperties{"worms", "fly", "peep"}
	animalMap["snake"] = animalProperties{"mice", "slither", "hsss"}
	var genralAni Animal
	for {
		var command, requestAni, requestType string
		fmt.Print(">")
		fmt.Scan(&command, &requestAni, &requestType)
		if command == "query" {
			genralAni = animalMap[requestAni]
			switch requestType {
			case "eat":
				genralAni.Eat()
			case "move":
				genralAni.Move()
			case "speak":
				genralAni.Speak()
			}
		}
		if command == "newanimal" {
			animalMap[requestAni] = animalMap[requestType]
			fmt.Println("Created it!")
		}
	}
}
