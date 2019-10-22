package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
	"math/rand"
)

func ChooseRandPlanet(planets map[string]interface{}) interface{} {
	i := 0
	rand := rand.Intn(8)

	for planet := range planets {
		if i == rand {
			fmt.Println("Traveling to ", planet)
			return planets[planet]
		} else {
			i += 1
		}
	}
	return ""
}

func LookupPlanet(planets map[string]interface{}) interface{} {
	fmt.Println("Name the plant you would like to visit.")

	var planet string
	fmt.Scanln(&planet)
	//If 'ok' is true, or in this case returns something, assign it to value
	if val, ok := planets[strings.ToLower(planet)]; ok {
		return val
	} else {
		fmt.Println("That planet does not exist..")
		return LookupPlanet(planets)
	}
}

func ValidateInput(input string, planets map[string]interface{}) interface{} {
	//Go until a valid input of Y or N is received
	if input != "Y" && input != "N" {
		fmt.Println("Sorry, I didn't get that.")
		fmt.Scanln(&input)
		return ValidateInput(input, planets)
	} else if input == "N" {
		return LookupPlanet(planets)
	} else if input == "Y" {
		return ChooseRandPlanet(planets)
	}
	return ""
}

func main() {
	//Reads json file given from user, if error prints it
	planetJson, err := ioutil.ReadFile(strings.ToLower(os.Args[1]))
	if err != nil {
		fmt.Print(err)
	}
	//Maps the contents of the json file to planets variable
	var planets map[string]interface{}
	json.Unmarshal([]byte(planetJson), &planets)

	fmt.Println("Welcome to the Solar System!\nThere are 9 planets to explore.\nWhat is your name?")

	var name string
	fmt.Scanln(&name)

	fmt.Println("Nice to meet you, "+name+".\nLet's go on an adventure.")
	fmt.Println("Shall I randomly choose a plant for you to visit? (Y or N)")
	
	var answer string
	fmt.Scanln(&answer)

	fmt.Println(ValidateInput(answer, planets))
}