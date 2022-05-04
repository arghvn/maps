package main

//maps : a A map is a set of value-key data with a type data associative
//structure similar to a dictionary in Python.
//space for more explain for above sentences

//In Go, a map is a powerful, ingenious, and versatile data structure.
//Golang Maps is a collection of unordered pairs of key-value.
//It is widely used because it provides fast lookups and values that can retrieve, update or delete with the help of keys.

import "fmt"

func main() {

	//To create an empty map, use the 'make' :
	//make(map[key-type]val-type)

	m := make(map[string]int)

	//space for more explain for above sentences

	// Set key/value pairs using typical 'name[key] = val' syntax.
	//space for more explain for above sentences

	m["k1"] = 7
	m["k2"] = 13

	//Printing a map with e.g. 'fmt.Println' will show all of its key/value pairs.

	fmt.Println("map:", m)

	//Get a value for a key with 'name[key]'.

	v1 := m["k1"]

	fmt.Println("v1: ", v1)

	//The 'len' returns the number of key/value pairs when called on a map.

	fmt.Println("len:", len(m))

	// The 'delete' removes key/value pairs from a map.

	delete(m, "k2")

	fmt.Println("map:", m)

	//The optional second return value when getting a value from a map indicates if the key was present in the map.
	//This can be used to disambiguate between missing keys and key with zero values like '0' or "".

	_, prs := m["k2"]

	fmt.Println("prs:", prs)

	//we can also declare and initialize a new map in  the same line with this syntax.

	n := map[string]int{"foo": 1, "bar": 2}

	fmt.Println("map:", n)

}

// outputs :
//map: map[k1:7 k2:13]
//v1: 7
//len: 2
//map: map[k1:7]
//prs: false
//map: map[foo:1 bar:2]

//more details :

//As we can see at first glance, the definition of Map is similar to Slice,
// except that its syntax is defined as -key [map (make type-val] type) and is
// implemented in the above code inside the variable m.

//We can use this to set the key name[key] = val and We can use this to get the value of a key name[key]

//Like array and slice, the len function can also be used for maps so that you can count the number of keys.

//The internal delete function, which consists of two parameters
// (the first parameter to get the map and the second parameter of the key) gives us access to delete a specific key.

//It has been said before that if the type of value is set to, for example,
// int and a value is not specified (or did not exist), we will return 0!
//But if we really put 0 as a variable and we want to know how to know if we put this 0 or if the value is not returned.
//Sometimes mentioning a phrase called OK comma, which means that when you start receiving a value with the syntax [key [name],
// there are actually two outputs, and that second output tells us in boolean that this key has been set, or No.
//In the example above, an example is mentioned, but the name prs is used instead of ok.
//The last example of the above code shows how to define a map and set it to a line.
//We can do the iterate process with the keyword range on arrays, sections and maps with the help of for,
//but the important point is that according to the Go language design, this output is not ordered on the map, and each time randomly is shown .
