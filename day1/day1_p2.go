package main

import (
    "fmt"
    "os"
	"log"
	"bufio"
	"strconv"
)


func main() {
	numIncreases := 0;
	args := os.Args
	if len(args) < 2 {
		fmt.Printf("Missing a command line argument\n")
		return
	}

    file, err := os.Open(args[1])
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

	scanner := bufio.NewScanner(file)
    
	var previous int64 = 0
	sum := [3]int64{0, 0, 0}
    for i := 0; scanner.Scan(); i++ {
        current, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		
		sum[i % 3] = current
		total := sum[0] + sum[1] + sum[2]
		if i >= 3 && total > previous {
			// calculated the previous already since i is at least 3 so we have an increase
			numIncreases++
		} 

		previous = total
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

	fmt.Printf("Num increases: %d\n", numIncreases)
}