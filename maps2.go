package main

import "fmt"

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

	for i := 1; i < 4; i++ {

		fmt.Printf("\nRUN NUMBER %d\n", i)

		for key, value := range actor {

			fmt.Printf("%s : %d years old\n", key, value)

		}

	}

}

//The best solution is to change the structure by first transferring the information
// to the structure of the slices and then relating them to the sort package , in maps3

//outputs :
//RUN NUMBER 1
//Redford : 79 years old
//Diaz : 43 years old
//Kilmer : 56 years old
//Pacino : 75 years old
//Ryder : 44 years old
//Paltrow : 43 years old
//Cruise : 53 years old

//RUN NUMBER 2
//Paltrow : 43 years old
//Cruise : 53 years old
//Redford : 79 years old
//Diaz : 43 years old
//Kilmer : 56 years old
//Pacino : 75 years old
//Ryder : 44 years old

//RUN NUMBER 3
//Cruise : 53 years old
//Redford : 79 years old
//Diaz : 43 years old
//Kilmer : 56 years old
//Pacino : 75 years old
//Ryder : 44 years old
//Paltrow : 43 years old
