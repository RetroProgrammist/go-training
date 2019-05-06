package main

import (
	"fmt"
	"strconv"
	"collection/collection"
)

func main() {
	var collection collection.Collection //{len:0 firstEl:<nil> lastEl:<nil>}

	for i := 0; i < 10; i++ {
		collection.Add("element_" + strconv.Itoa(i))
	}

	fmt.Println("--------------------- First block ---------------------")
	collection.Print()
	collection.Remove(4)
	collection.Print()
	fmt.Println("\n--------------------- Second block ---------------------")
	fmt.Println(collection.Get(3))
	fmt.Println(collection.First())
	fmt.Println(collection.Last())
	fmt.Println(collection.Length())

}

