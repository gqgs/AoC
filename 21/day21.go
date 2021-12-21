package main

import (
	"fmt"
	"os"
)

func solve() error {
	file, err := os.Open(os.Args[1])
	if err != nil {
		return err
	}
	defer file.Close()

	silver(4, 8)
	gold(6, 10)

	return nil
}

func genDice(start int) func() int {
	var i, c int
	i = start
	return func() int {
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

func split(player, player1Score, player1Position, player2Score, player2Position int, cache map[string][2]int) (int, int) {
	cacheKey := fmt.Sprint(player, player1Score, player1Position, player2Score, player2Position)
	if value, cached := cache[cacheKey]; cached {
		return value[0], value[1]
	}

	var player1Wins, player2Wins int
	if player%2 == 0 {
		for _, d := range diracDice() {
			move := d[0] + d[1] + d[2]
			local1Position := player1Position + move
			if local1Position > 10 {
				local1Position %= 10
				if local1Position == 0 {
					local1Position = 10
				}
			}
			local1Score := player1Score
			local1Score += local1Position
			if local1Score >= 21 {
				player1Wins++
				continue
			}
			p1Wins, p2Wins := split(player+1, local1Score, local1Position, player2Score, player2Position, cache)
			player1Wins += p1Wins
			player2Wins += p2Wins
		}
	} else {
		for _, d := range diracDice() {
			move := d[0] + d[1] + d[2]
			local2Position := player2Position + move
			if local2Position > 10 {
				local2Position %= 10
				if local2Position == 0 {
					local2Position = 10
				}
			}
			local2Score := player2Score
			local2Score += local2Position
			if local2Score >= 21 {
				player2Wins++
				continue
			}
			p1Wins, p2Wins := split(player+1, player1Score, player1Position, local2Score, local2Position, cache)
			player1Wins += p1Wins
			player2Wins += p2Wins
		}
	}

	cache[cacheKey] = [2]int{player1Wins, player2Wins}
	return player1Wins, player2Wins
}

func gold(startPlayer1, startPlayer2 int) {
	player := 0
	cache := make(map[string][2]int)
	p1wins, p2wins := split(player, 0, startPlayer1, 0, startPlayer2, cache)
	fmt.Printf("p1 wins: %d, p2 wins: %d\n", p1wins, p2wins)
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
				fmt.Println("p1 won", player1Score, player2Score, dice()-1)
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
				fmt.Println("p2 won", player1Score, player2Score, dice()-1)
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
