package main

import (
	"bufio"
	"./ai"
	"./game"
	"fmt"
	"os"
	"strconv"
)

func main() {
	g := game.New()
	aip := requestPlayWithAI()
	for {
		fmt.Println(g)
		fmt.Printf("Player %v's turn!\n", g.Turn)

		// Move.
		var col game.Column
		if aip != nil && g.Turn == aip.Player {
			col = aip.ChooseMove(g)
		} else {
			col = requestMove()
		}
		gg, err := g.Move(col)
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

func requestPlayWithAI() *ai.AI {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		// AI opponent?
		fmt.Print("Enable AI? [y/n]: ")
		scanner.Scan()
		switch scanner.Text() {
		case "y":
			return requestAIDifficulty()
		case "n":
			return nil
		}
	}
}

func requestAIDifficulty() *ai.AI {
	aip := &ai.AI{Player: game.Black}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		// AI difficulty?
		fmt.Print("AI difficulty [e(asy)/m(edium)/h(ard)/i(mpossible)]: ")
		scanner.Scan()
		switch scanner.Text() {
		case "e":
			aip.Difficulty = ai.Easy
		case "m":
			aip.Difficulty = ai.Medium
		case "h":
			aip.Difficulty = ai.Hard
		case "i":
			aip.Difficulty = ai.Impossible
		default:
			continue
		}
		return aip
	}
}

func requestMove() game.Column {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		// One-index columns for simpler input.
		fmt.Print("Enter move [1-7]: ")
		scanner.Scan()
		text := scanner.Text()
		col, err := strconv.Atoi(text)
		if err != nil {
			continue
		}
		return game.Column(col - 1)
	}
}
