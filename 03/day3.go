package main

import (
	"bufio"
	"os"
	"strconv"
)

func silver(numbers []string) {
	var count []int
	for _, number := range numbers {
		if cap(count) == 0 {
			count = make([]int, len(number))
		}
		for i, n := range number {
			if n == '1' {
				count[i]++
			}
		}
	}

	var gammaRateBinary, episilonRateBinary string
	for _, c := range count {
		if c > len(numbers)/2 {
			gammaRateBinary += "1"
			episilonRateBinary += "0"
		} else {
			gammaRateBinary += "0"
			episilonRateBinary += "1"
		}
	}

	gammaRate, _ := strconv.ParseInt(gammaRateBinary, 2, 0)
	episilonRate, _ := strconv.ParseInt(episilonRateBinary, 2, 0)

	println("silver:", gammaRateBinary, episilonRateBinary, gammaRate, episilonRate, gammaRate*episilonRate)
}

func gold(numbers []string) {
	oxygenRatingBinary := goldFilter(numbers, '0', '1', func(mcb0, mcb1 []string) bool {
		return len(mcb0) > len(mcb1)
	})
	co2RatingBinary := goldFilter(numbers, '1', '0', func(mcb0, mcb1 []string) bool {
		return len(mcb0) < len(mcb1)
	})

	oxygenRating, _ := strconv.ParseInt(oxygenRatingBinary, 2, 0)
	co2Rating, _ := strconv.ParseInt(co2RatingBinary, 2, 0)

	println("gold:", oxygenRatingBinary, co2RatingBinary, oxygenRating, co2Rating, oxygenRating*co2Rating)
}

func goldFilter(numbers []string, b0, b1 byte, cmpFunc func(mcb0, mcb1 []string) bool) string {
	var index int
	for {
		var mcb0, mcb1 []string
		for _, n := range numbers {
			switch n[index] {
			case b0:
				mcb0 = append(mcb0, n)
			case b1:
				mcb1 = append(mcb1, n)
			}
		}
		if cmpFunc(mcb0, mcb1) {
			numbers = mcb0
		} else {
			numbers = mcb1
		}

		if len(numbers) == 1 {
			return numbers[0]
		}

		index++
	}
}

func solve() error {
	file, err := os.Open("day3")
	if err != nil {
		return err
	}
	defer file.Close()

	var numbers []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numbers = append(numbers, scanner.Text())
	}

	silver(numbers)
	gold(numbers)

	return nil
}

func main() {
	if err := solve(); err != nil {
		panic(err)
	}
}
