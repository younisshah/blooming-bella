package main

import (
	"github.com/younisshah/blooming-bella/bb"
	"log"
	"fmt"
)

func main() {
	bella, err := blooming_bella.NewBella(1000, 0.01)

	if err != nil {
		log.Fatal(err)
	}
	bella.Add(10)
	bella.Add(121)
	bella.Add(13)
	bella.Add(111)

	fmt.Println(bella.Test(10)) // => true
	fmt.Println(bella.Test(104)) // => false
	fmt.Println(bella.Test(110)) // => false
	fmt.Println(bella.Test(13)) // => true
}
