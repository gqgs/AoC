package main

import (
	"crypto/sha256"
	"fmt"
	"os"
	"sort"

	"github.com/gqgs/AoC2021/generic"
)

func solve() error {
	file, err := os.Open(os.Args[1])
	if err != nil {
		return err
	}
	defer file.Close()

	// TODO: properly parse this from input file
	board := new(Board)
	// board.addAmphipods(Bronze, 2, 1)
	// board.addAmphipods(Desert, 2, 2)
	// board.addAmphipods(Desert, 2, 3)
	// board.addAmphipods(Amber, 2, 4)

	// board.addAmphipods(Copper, 4, 1)
	// board.addAmphipods(Copper, 4, 2)
	// board.addAmphipods(Bronze, 4, 3)
	// board.addAmphipods(Desert, 4, 4)

	// board.addAmphipods(Bronze, 6, 1)
	// board.addAmphipods(Bronze, 6, 2)
	// board.addAmphipods(Amber, 6, 3)
	// board.addAmphipods(Copper, 6, 4)

	// board.addAmphipods(Desert, 8, 1)
	// board.addAmphipods(Amber, 8, 2)
	// board.addAmphipods(Copper, 8, 3)
	// board.addAmphipods(Amber, 8, 4)

	board.addAmphipods(Desert, 2, 1)
	board.addAmphipods(Desert, 2, 2)
	board.addAmphipods(Desert, 2, 3)
	board.addAmphipods(Desert, 2, 4)

	board.addAmphipods(Amber, 4, 1)
	board.addAmphipods(Copper, 4, 2)
	board.addAmphipods(Bronze, 4, 3)
	board.addAmphipods(Amber, 4, 4)

	board.addAmphipods(Copper, 6, 1)
	board.addAmphipods(Bronze, 6, 2)
	board.addAmphipods(Amber, 6, 3)
	board.addAmphipods(Bronze, 6, 4)

	board.addAmphipods(Copper, 8, 1)
	board.addAmphipods(Amber, 8, 2)
	board.addAmphipods(Copper, 8, 3)
	board.addAmphipods(Bronze, 8, 4)

	println(silver(*board, 0))

	return nil
}

type Board struct {
	state [11][5]*Amphipods
}

const (
	Amber  int = 1
	Bronze int = 10
	Copper int = 100
	Desert int = 1000
)

type Amphipods struct {
	Type int
}

func (a Amphipods) isFinalRoom(x, y int) bool {
	if y == 0 {
		return false
	}
	return x == a.roomIndex()
}

func (a Amphipods) roomIndex() int {
	switch a.Type {
	case Amber:
		return 2
	case Bronze:
		return 4
	case Copper:
		return 6
	case Desert:
		return 8
	default:
		panic("")
	}
}

func (a Amphipods) moveCost(i int) int {
	return a.Type * i
}

func (b *Board) addAmphipods(atype, x, y int) {
	b.state[x][y] = &Amphipods{
		Type: atype,
	}
}

func (b Board) isFinal() bool {
	for _, x := range []int{2, 4, 6, 8} {
		for _, y := range []int{1, 2, 3, 4} {
			if b.state[x][y] == nil || !b.state[x][y].isFinalRoom(x, y) {
				return false
			}
		}
	}
	return true
}

func (b Board) String() string {
	var str string
	for y := 0; y < 5; y++ {
		for x := 0; x < 11; x++ {
			if b.state[x][y] == nil {
				if y == 0 || x == 2 || x == 4 || x == 6 || x == 8 {
					str += "."
				} else {
					str += "#"
				}
			} else {
				str += typeString(b.state[x][y].Type)
			}
		}
		str += "\n"
	}
	return str
}

func typeString(t int) string {
	switch t {
	case Amber:
		return "A"
	case Bronze:
		return "B"
	case Copper:
		return "C"
	case Desert:
		return "D"
	default:
		return "#"
	}
}

func (b Board) validMoves() [][2][3]int {
	var moves [][2][3]int
	var isBlockinDoor bool
<<<<<<< HEAD

=======
>>>>>>> 299dcab (Day 24)
	// at the front of room door
	for _, x := range []int{2, 4, 6, 8} {
		if b.state[x][0] != nil {
			isBlockinDoor = true
			for i := x + 1; i < 11 && b.state[i][0] == nil; i++ {
				if i != 2 && i != 4 && i != 6 && i != 8 {
					moves = append(moves, [2][3]int{
						{x, 0, b.state[x][0].moveCost(i - x)},
						{i, 0, 0},
					})
				}
			}

			for i := x - 1; i >= 0 && b.state[i][0] == nil; i-- {
				if i != 2 && i != 4 && i != 6 && i != 8 {
					moves = append(moves, [2][3]int{
						{x, 0, b.state[x][0].moveCost(x - i)},
						{i, 0, 0},
					})
				}
			}
		}
	}

	if isBlockinDoor {
		return moves
	}

	needsToMove := func(x, y int) bool {
		if b.state[x][y] == nil {
			return false
		}
		for dy := y; dy <= 4; dy++ {
			if b.state[x][dy] != nil && !b.state[x][dy].isFinalRoom(x, dy) {
				return true
			}
		}
		return !b.state[x][y].isFinalRoom(x, y)
	}

	// inside the room
	for _, x := range []int{2, 4, 6, 8} {
		for _, y := range []int{4, 3, 2, 1} {
			if needsToMove(x, y) {
				moves = append(moves, [2][3]int{
					{x, y, b.state[x][y].moveCost(1)},
					{x, y - 1, 0},
				})
			}
		}
	}

	// at the hallway
	for x := 0; x < 11; x++ {
		if b.state[x][0] != nil {
			if b.pathExists(next(x, b.state[x][0]), b.state[x][0]) {
				roomIndex := b.state[x][0].roomIndex()
				for _, y := range []int{4, 3, 2, 1} {
					if b.state[roomIndex][y] == nil {
						moves = append(moves, [2][3]int{
							{x, 0, b.state[x][0].moveCost(distance(x, roomIndex, 0, y))},
							{roomIndex, y, 0},
						})
						break
					}
				}
			}
		}
	}

	// filter illegal moves
	var i int
	for _, m := range moves {
		destX, destY := m[1][0], m[1][1]
		if b.state[destX][destY] != nil {
			continue
		}
		moves[i] = m
		i++
	}
	moves = moves[:i]

	return moves
}

func distance(x0, x1, y0, y1 int) int {
	return generic.Abs(x1-x0) + generic.Abs(y1-y0)
}

func next(x int, a *Amphipods) int {
	if x < a.roomIndex() {
		return x + 1
	}
	return x - 1
}

func (b Board) pathExists(x int, a *Amphipods) bool {
	if x < 0 || x >= 11 || b.state[x][0] != nil {
		return false
	}

	isFinalRoom := func(x, y int) bool {
		return b.state[x][y] != nil && b.state[x][y].isFinalRoom(x, y)
	}

	if x == a.roomIndex() {
		return (b.state[x][4] == nil && b.state[x][3] == nil && b.state[x][2] == nil && b.state[x][1] == nil) ||
			isFinalRoom(x, 4) && b.state[x][3] == nil && b.state[x][2] == nil && b.state[x][1] == nil ||
			isFinalRoom(x, 4) && isFinalRoom(x, 3) && b.state[x][2] == nil && b.state[x][1] == nil ||
			isFinalRoom(x, 4) && isFinalRoom(x, 3) && isFinalRoom(x, 2) && b.state[x][1] == nil
	}
	return b.pathExists(next(x, a), a)
}

var cache = make(map[[28]byte]Cache)

type Cache struct {
	Value int
	Valid bool
}

var globalMin = 1<<63 - 1

func silver(b Board, cost int) (int, bool) {
	boardString := sha256.Sum224([]byte(b.String()))

	if cost >= globalMin {
		return 0, false
	}

	if b.isFinal() {
		globalMin = generic.Min(globalMin, cost)
		return 0, true
	}

	if value, cached := cache[boardString]; cached {
		return value.Value, value.Valid
	}

	var costs []int
	moves := b.validMoves()
	if len(moves) == 0 {
		cache[boardString] = Cache{0, false}
		return 0, false
	}

	sort.Slice(moves, func(i, j int) bool {
		return moves[i][0][2] < moves[j][0][2]
	})

	for _, move := range moves {
		b.state[move[0][0]][move[0][1]], b.state[move[1][0]][move[1][1]] = b.state[move[1][0]][move[1][1]], b.state[move[0][0]][move[0][1]]
		subcost, exists := silver(b, cost+move[0][2])
		if exists {
			costs = append(costs, move[0][2]+subcost)
		}
		b.state[move[0][0]][move[0][1]], b.state[move[1][0]][move[1][1]] = b.state[move[1][0]][move[1][1]], b.state[move[0][0]][move[0][1]]
	}

	if len(costs) == 0 {
		cache[boardString] = Cache{0, false}
		return 0, false
	}

	min := generic.Min(costs...)
	cache[boardString] = Cache{min, true}
	return min, true
}

func main() {
	if err := solve(); err != nil {
		panic(err)
	}
}
