//
package main

import (
	"fmt"
	"log"
	"github.com/JLuis23/average"
	"github.com/sirupsen/logru"
)
 func main (){
	lines, err := datafile.GetStrings("votes.txt")
	if err !=  nil {
		log.Fatal(err)
	}
	counts := make(map[string]int)
	for _, line := range lines {
				counts [line]++
	}
	for name, count := range counts {
		fmt.Printf("votes for %s: %d \n", name, count)
	}
	}

 