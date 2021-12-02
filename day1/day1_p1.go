package main

import (
    "fmt"
    "os"
	"log"
	"bufio"
	"strconv"
)


func main() {
	var numIncreases int = 0;
	args := os.Args
	if len(args) < 2 {
		log.Fatal("Missing a command line argument\n")
	}

    file, err := os.Open(args[1])
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		return
	}
    
	previous, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		log.Fatal(err)
	}

    for scanner.Scan() {
        current, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		if current > previous {
			numIncreases++;
		}
		
		previous = current
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

	fmt.Printf("Num increases: %d\n", numIncreases)
}