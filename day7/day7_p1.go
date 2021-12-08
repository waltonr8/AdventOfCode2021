package main

import (
    "fmt"
    "os"
	"log"
	"bufio"
	"strconv"
	"strings"
	"sort"
)

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
	half := len(nums)/2
	total := 0
	for i, num := range nums {
		if i < half {
			total += nums[half] - num
		} else {
			total += num - nums[half]
		}
	}

	// if there are two midpoints (ie len of array is even) check if the other midpoint is better
	if len(nums) % 2 == 0 && nums[half] != nums[half + 1] {
		total2 := 0
		for i, num := range nums {
			if i < half {
				total2 += nums[half + 1] - num
			} else {
				total2 += num - nums[half + 1]
			}
		}

		if (total2 < total) {
			total = total2
		}
	}

	fmt.Printf("Nums: %v\n", nums)
	fmt.Printf("Total: %d\n", total)
}