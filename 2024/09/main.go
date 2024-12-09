package main

import (
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
)

func decode(line string) []string {
	var decoded []string
	for i, c := range line {
		n, _ := strconv.Atoi(string(c))
		if i%2 == 0 {
			digit := fmt.Sprint(i / 2)
			for range n {
				decoded = append(decoded, digit)
			}
		} else {
			for range n {
				decoded = append(decoded, ".")
			}
		}
	}
	return decoded
}

func silver(line string) int {
	decoded := decode(line)
	left := 0
	right := len(decoded) - 1
	for left < right {
		if decoded[left] != "." {
			left++
			continue
		}
		if decoded[right] == "." {
			right--
			continue
		}

		decoded[left], decoded[right] = decoded[right], decoded[left]
	}
	var total int
	for key, value := range decoded {
		if value == "." {
			break
		}
		n, _ := strconv.Atoi(value)
		total += key * n
	}

	return total
}

type Block struct {
	ID      int
	Size    int
	Moved   bool
	MovedTo int
}

type Free struct {
	ID   int
	IDS  []int
	Size int
}

func gold(line string) int {
	var blocks []*Block
	var frees []*Free
	for i, c := range line {
		n, _ := strconv.Atoi(string(c))
		if i%2 == 0 {
			blocks = append(blocks, &Block{
				ID:   i / 2,
				Size: n,
			})
		} else {
			frees = append(frees, &Free{
				ID:   i / 2,
				Size: n,
			})
		}
	}
Continue:
	for _, block := range slices.Backward(blocks) {
		for _, free := range frees {
			if block.ID <= free.ID {
				continue Continue
			}

			if free.Size-len(free.IDS) >= block.Size {
				for range block.Size {
					free.IDS = append(free.IDS, block.ID)
				}
				block.Moved = true
				block.MovedTo = free.ID
				continue Continue
			}
		}
	}

	memory := make([]int, 0, len(blocks)+len(frees))
	for i := range min(len(blocks), len(frees)) {
		block := blocks[i]
		var id int
		if !block.Moved {
			id = block.ID
		}
		for range block.Size {
			memory = append(memory, id)
		}

		free := frees[i]
		for _, id := range free.IDS {
			memory = append(memory, id)
		}
		for range free.Size - len(free.IDS) {
			memory = append(memory, 0)
		}
	}

	var total int
	for i, value := range memory {
		total += i * value
	}

	return total
}

func solve() error {
	file, err := os.Open(os.Args[1])
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	line := string(data)
	println(silver(line))
	println(gold(line))

	return nil
}

func main() {
	if err := solve(); err != nil {
		panic(err)
	}
}
