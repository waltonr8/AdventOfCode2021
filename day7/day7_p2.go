package main

import (
    "fmt"
    "os"
	"log"
	"bufio"
	"strconv"
	"strings"
	"sort"
	"math"
)

func getFuelCostForMove(move int) int {
	cost := 0
	for i := 0; i <= move; i++ {
		cost += i
	}

	return cost
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

	s := strings.Split(scanner.Text(), ",")

	nums := make([]int, 0)
	
	for _, numS := range s {
		num, err := strconv.ParseInt(numS, 10, 64)  
 		if err != nil {  
			log.Fatal(err)
 		} 
	
		nums = append(nums, int(num))
	}

	sort.Ints(nums)
	mean := 0
	for _, num := range nums {
		mean += num
	}

	fmt.Printf("Mean as float: %f\n", float64(mean) / float64(len(nums)))
	mean = int(math.Round(float64(mean) / float64(len(nums))))
	mean = mean - 1
	fmt.Printf("Mean: %d\n", mean)

	total := 0
	for _, num := range nums {
		total += getFuelCostForMove(int(math.Abs(float64(num - mean))))
	}

	fmt.Printf("Total: %d\n", total)
}