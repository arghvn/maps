package main

import (
	"fmt"
	"sort"
)

func main() {

	actor := map[string]int{

		"Paltrow": 43,

		"Cruise": 53,

		"Redford": 79,

		"Diaz": 43,

		"Kilmer": 56,

		"Pacino": 75,

		"Ryder": 44,
	}

	// Store the keys in a slice

	var sortedActor []string

	for key := range actor {

		sortedActor = append(sortedActor, key)

	}

	// Sort the slice alphabetically

	sort.Strings(sortedActor)

	//Retrieve the keys from the slice and use them to look up the map values

	for _, name := range sortedActor {

		fmt.Printf("%s : %d years old\n", name, actor[name])

	}

}

//Inside if we can define a variable and then with; Let's separate it and evaluate it.
// Just for familiarity with such syntaxes on if Statements were raised in maps4

//Output :
//Cruise : 53 years old
//Diaz : 43 years old
//Kilmer : 56 years old
//Pacino : 75 years old
//Paltrow : 43 years old
//Redford : 79 years old
//Ryder : 44 years old
