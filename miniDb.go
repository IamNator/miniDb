package main

import (
	"errors"
	"fmt"
	"strings"
)

//############################
type cow struct {
	name string
}

func (myCow cow) Eat() {
	fmt.Println("grass")
}

func (myCow cow) Move() {
	fmt.Println("walk")
}

func (myCow cow) Speak() {
	fmt.Println("moo")
}

//################################

type bird struct {
	name string
}

func (myBird bird) Eat() {
	fmt.Println("worms")
}

func (myBird bird) Move() {
	fmt.Println("fly")
}

func (myBird bird) Speak() {
	fmt.Println("peep")
}

//####################################

type snake struct {
	name string
}

func (mySnake snake) Eat() {
	fmt.Println("mice")
}

func (mySnake snake) Move() {
	fmt.Println("slither")
}

func (mySnake snake) Speak() {
	fmt.Println("hsss")
}

//##########################################

type Animal interface {
	Eat()
	Move()
	Speak()
}

func createAnimal(animalName string, animalType string) (Animal, error) {
	var err error

	switch animalType {
	case "cow":
		return cow{name: animalName}, nil
	case "bird":
		return bird{name: animalName}, nil
	case "snake":
		return snake{name: animalName}, nil
	default:
		err = errors.New("Please enter a valid animal type")
	}
	return Animal(nil), err
}

func query(animalName string, animalAttr string, user_created_animal_list map[string]Animal) error {
	var err error

	if user_created_animal_list[animalName] == nil {
		return errors.New("Animal name not found")
	}

	switch {
	case strings.Contains(animalAttr, "eat"):
		user_created_animal_list[animalName].Eat()
	case strings.Contains(animalAttr, "move"):
		user_created_animal_list[animalName].Move()
	case strings.Contains(animalAttr, "speak"):
		user_created_animal_list[animalName].Speak()
	default:
		err = errors.New("Please enter a valid name and behave e.g > query bingo eat ")
	}
	return err
}

func actOnUserInput(rawUserInput []string, user_created_animal_list map[string]Animal) error {

	var err error
	command, name, attr := strings.ToLower(rawUserInput[0]), strings.ToLower(rawUserInput[1]), strings.ToLower(rawUserInput[2])

	switch command {

	case "newanimal":
		user_created_animal_list[name], err = createAnimal(name, attr)
		if err != nil {
			return err
		}
		fmt.Println("Created it!")
	case "query":
		err = query(name, attr, user_created_animal_list)
		if err != nil {
			return err
		}
	case "help":
		fmt.Println("\nfor adding a new animal \n> newanimal <animal_name> <cow|bird|snake>\nTo check an animals' attribute \n> query <animal_name> <eat|move|speak>\n")
	default:
		return errors.New("Please enter a valid command")
	}

	return nil
}

func getUserInput(userinput []string) {
	for i := 0; i < 3; i++ {
		_, err := fmt.Scan(&userinput[i])
		for err != nil {
			fmt.Printf("error :%+v\n", err)
			_, err = fmt.Scan(&userinput[i])
		}
	}
}

func main() {
	userCreatedAnimalList := make(map[string]Animal)
	user_input := make([]string, 3)

	fmt.Println("Please enter commands (enter > help h h  for help) :")

	for {
		fmt.Print("> ")

		getUserInput(user_input)

		err := actOnUserInput(user_input, userCreatedAnimalList)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}
