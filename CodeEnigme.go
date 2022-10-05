package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 1000

	readFile, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	readFile.Close()

	fmt.Print(lines[0]) //écrit le premier fragment
	fmt.Print(" ")
	fmt.Print(lines[len(lines)-1]) //écrit le deuxième fragment

	for i := 0; i < len(lines); i++ { //écrit le troisième fragment
		if lines[i] == "before" {
			y, _ := strconv.Atoi(lines[i+1])
			fmt.Print(" ")
			fmt.Print(lines[y-1])

		}
	}

	for i := 0; i < len(lines); i++ { //écrit le quatrième fragment
		if lines[i] == "now " {
			tempo := lines[i-1]
			fmt.Print(" ")
			fmt.Println(lines[int(tempo[1])/len(lines)-1])
		}
	}

	fmt.Println("Un numéro aleatoire entre 0 et 100 : ", rand.Intn(max-min)+min)
}
