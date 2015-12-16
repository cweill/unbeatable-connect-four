package main

import (
	"bufio"
    "fmt"
    "os"
	"c4/game"
	"strconv"
)

func main() {
	var ai *game.AI
	g := game.New()
	if requestPlayWithAI() {
		ai = &game.AI{}
	}
	for {
		fmt.Println(g)
		fmt.Printf("Player %v's turn!\n", g.Turn)

		// Move.
		var col game.Column
		if ai != nil && g.Turn == game.Black {
			col = ai.ChooseMove(g)
		} else {
			col = requestMove()
		}
		gg, err := g.Move(col);
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}

		// Check end game.
		if gg.IsGameOver() {
			fmt.Println(gg)
			fmt.Printf("Player %v wins!\n", gg.Turn)
			return
		}

		// Next player's turn.
		g = gg.NextTurn()
	}	
}

func requestPlayWithAI() bool {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enable AI? [y/n]: ")
		scanner.Scan()
		switch scanner.Text() {
		case "y":
			return true
		case "n": 
			return false
		}
	}
}

func requestMove() game.Column {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter move [0-6]: ")
		scanner.Scan()
		text := scanner.Text()
		v, err := strconv.Atoi(text)
		if err != nil {
			continue
		}
		return game.Column(v)
	}
}