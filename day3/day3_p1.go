package main

import (
    "fmt"
    "os"
	"log"
	"bufio"
	"strconv"
)

func addCounts(zeroCounts, onesCounts []int, num string) {
	for i, c := range num {
		if c == '1' {
			onesCounts[i]++;
		} else {
			zeroCounts[i]++;
		}
	}
}

func main() {
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

	size := len(scanner.Text())
	zeroCounts := make([]int, size)
	onesCounts := make([]int, size)
	addCounts(zeroCounts, onesCounts, scanner.Text())

	

    for scanner.Scan() {
		addCounts(zeroCounts, onesCounts, scanner.Text())
	}

	gammaStr := ""
	epsilonStr := ""
	for i, _ := range zeroCounts {
		if zeroCounts[i] > onesCounts[i] {
			gammaStr += "0"
			epsilonStr += "1"
		} else {
			gammaStr += "1"
			epsilonStr += "0"
		}
	}

	gamma, err := strconv.ParseInt(gammaStr, 2, 64)  
 	if err != nil {  
		log.Fatal(err)
 	} 
	
	epsilon, err := strconv.ParseInt(epsilonStr, 2, 64)  
 	if err != nil {  
		log.Fatal(err)
 	} 

	fmt.Printf("Gamma, Epsilon: %d, %d\n", gamma, epsilon)
	fmt.Printf("Power Consumption: %d\n", gamma * epsilon)
}