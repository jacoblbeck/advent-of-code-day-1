package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	body, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	ints, err := ReadInts(strings.NewReader(string(body)))
	if err != nil {
		log.Fatal(err)
	}

	count := DetermineIncrease(ints)

	fmt.Println(count)

}

// ReadInts reads whitespace-separated ints from r. If there's an error, it
// returns the ints successfully read so far as well as the error value.
func ReadInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var result []int
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result = append(result, x)
	}
	return result, scanner.Err()
}

func DetermineIncrease(ints []int) int {
	temp := ints[0]
	var count int
	for _, num := range ints {
		if num > temp {
			count++
		}
		temp = num
	}

	return count
}
