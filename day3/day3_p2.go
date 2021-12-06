package main

import (
    "fmt"
    "os"
	"log"
	"bufio"
	"strconv"
)

func addCounts(zeroCounts []int, num string) {
	for i, c := range num {
		if c != '1' {
			zeroCounts[i]++;
		}
	}
}

func remove(s []string, i int) []string {
	return append(s[:i], s[i+1:]...)
}

func getMostCommon(zeroCounts []int, n int) string {
	half := n/2
	str := ""
	for i, _ := range zeroCounts {
		if zeroCounts[i] > half {
			str += "0"
		} else if zeroCounts[i] < half {
			str += "1"
		} else {
			str += "2"
		}
	}

	return str
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
	keepOxy := make([]string, 0)
	keepCO2 := make([]string, 0)
	keepOxy = append(keepOxy, scanner.Text())
	keepCO2 = append(keepCO2, scanner.Text())
	addCounts(zeroCounts, scanner.Text())

    for scanner.Scan() {
		addCounts(zeroCounts, scanner.Text())
		keepOxy = append(keepOxy, scanner.Text())
		keepCO2 = append(keepCO2, scanner.Text())
	}

	i := 0
	for !(len(keepOxy) <= 1) {
		zeroCounts = make([]int, size)
		for _, num := range keepOxy {
			addCounts(zeroCounts, num)
		}

		str := getMostCommon(zeroCounts, len(keepOxy))
		j := 0
		for j < len(keepOxy) {
			if (keepOxy[j][i] != str[i] && str[i] != '2') || (keepOxy[j][i] == '0' && str[i] == '2') {
				keepOxy = remove(keepOxy, j)
			} else {
				j++
			}
		}
		i++
	}
	//fmt.Printf("nums: %v\n", keepOxy)

	i = 0
	for !(len(keepCO2) <= 1) {
		zeroCounts = make([]int, size)
		for _, num := range keepCO2 {
			addCounts(zeroCounts, num)
		}
		str := getMostCommon(zeroCounts, len(keepCO2))
		j := 0
		for j < len(keepCO2) {
			if keepCO2[j][i] == str[i] || (keepCO2[j][i] == '1' && str[i] == '2') {
				keepCO2 = remove(keepCO2, j)
			} else {
				j++
			}
		}
		i++
	}
	//fmt.Printf("nums: %v\n", keepCO2)

	o, err := strconv.ParseInt(keepOxy[0], 2, 64)  
 	if err != nil {  
		log.Fatal(err)
 	} 
	
	c, err := strconv.ParseInt(keepCO2[0], 2, 64)  
 	if err != nil {  
		log.Fatal(err)
 	} 

	fmt.Printf("Oxygen, CO2: %d, %d\n", o, c)
	fmt.Printf("LifeSupport: %d\n", o * c)
}