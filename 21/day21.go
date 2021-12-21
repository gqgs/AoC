package main

import (
	"fmt"
	"os"
	"strconv"
)

func stringsToInts(strs []string) ([]int, error) {
	ints := make([]int, len(strs))
	for i := range strs {
		var err error
		ints[i], err = strconv.Atoi(strs[i])
		if err != nil {
			return nil, err
		}
	}
	return ints, nil
}

func solve() error {
	file, err := os.Open(os.Args[1])
	if err != nil {
		return err
	}
	defer file.Close()

	// silver(4, 8)
	gold(4, 8)

	return nil
}

func genDice(start int) func() int {
	var i, c int
	i = start
	return func() int {
		println("dice roll", i)
		c++
		i++
		return i - 1
	}
}

func diracDice() [][3]int {
	var rolls [][3]int
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			for k := 1; k <= 3; k++ {
				rolls = append(rolls, [3]int{i, j, k})
			}
		}
	}
	return rolls
}

func split(player, player1Score, player1Position, player2Score, player2Position int) (int, int) {
	var player1Wins, player2Wins int
	if player%2 == 0 {
		for _, d := range diracDice() {
			move := d[0] + d[1] + d[2]
			player1Position = player1Position + move
			if player1Position > 10 {
				player1Position %= 10
				if player1Position == 0 {
					player1Position = 10
				}
			}
			player1Score += player1Position
			if player1Score >= 21 {
				player1Wins++
				continue
			}
			p1Wins, p2Wins := split(player+1, player1Score, player1Position, player2Score, player2Position)
			fmt.Println("p1,p2", p1Wins, p2Wins)
			player1Wins += p1Wins
			player2Wins += p2Wins
		}
	} else {
		for _, d := range diracDice() {
			move := d[0] + d[1] + d[2]
			player2Position = player2Position + move
			if player2Position > 10 {
				player2Position %= 10
				if player2Position == 0 {
					player2Position = 10
				}
			}
			player2Score += player2Position
			if player2Score >= 21 {
				player2Wins++
				continue
			}
			p1Wins, p2Wins := split(player+1, player1Score, player1Position, player2Score, player2Position)
			player1Wins += p1Wins
			player2Wins += p2Wins
		}

	}
	return player1Wins, player2Wins
}

func gold(startPlayer1, startPlayer2 int) {
	player := 0
	p1wins, p2wins := split(player, 0, startPlayer1, 0, startPlayer2)
	fmt.Println("p1wins, p2wins", p1wins, p2wins)
}

func silver(startPlayer1, startPlayer2 int) {
	dice := genDice(1)
	player := 0
	var player1Score, player2Score int
	player1Position := startPlayer1
	player2Position := startPlayer2
	for {
		if player%2 == 0 {
			move := dice() + dice() + dice()
			player1Position = player1Position + move
			if player1Position > 10 {
				player1Position %= 10
				if player1Position == 0 {
					player1Position = 10
				}
			}
			player1Score += player1Position
			if player1Score >= 1000 {
				fmt.Println("p1 won", player1Score, player2Score)
				return
			}
		} else {
			move := dice() + dice() + dice()
			player2Position = player2Position + move
			if player2Position > 10 {
				player2Position %= 10
				if player2Position == 0 {
					player2Position = 10
				}
			}
			player2Score += player2Position
			if player2Score >= 1000 {
				fmt.Println("p2 won", player1Score, player2Score)
				return
			}
		}
		player++
	}
}

func main() {
	if err := solve(); err != nil {
		panic(err)
	}
}
