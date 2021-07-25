package main

import (
	"fmt"
)

type Animal struct {
	food, locomotion, sound string
}

func (object Animal) Eat() {
	fmt.Println(object.food)
}

func (object Animal) Move() {
	fmt.Println(object.locomotion)
}

func (object Animal) Speak() {
	fmt.Println(object.sound)
}

func main() {
	animalClass := map[string]Animal{
		"cow":   {"grass", "walk", "moo"},
		"bird":  {"worms", "fly", "peep"},
		"snake": {"mice", "slither", "hsss"},
	}
	for {
		fmt.Print(">")
		animalName := "0"
		informationRequested := "0"
		fmt.Scan(&animalName, &informationRequested)
		if informationRequested == "eat" {
			animalClass[animalName].Eat()
		} else if informationRequested == "move" {
			animalClass[animalName].Move()
		} else if informationRequested == "speak" {
			animalClass[animalName].Speak()
		}
	}
}
