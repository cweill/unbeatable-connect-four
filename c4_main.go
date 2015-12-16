package main

import (
	"bufio"
    "fmt"
    "os"
	"c4/game"
	"strconv"
)

func main() {
	g := game.New()
	g.AI = requestPlayWithAI()
	for {
		fmt.Println(g)
		fmt.Printf("Player %v's turn!\n", g.Turn)
		v := requestMove()
		gg, err := g.Move(v);
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}
		if gg.IsGameOver() {
			fmt.Println(gg)
			fmt.Printf("Player %v wins!\n", gg.Turn)
			return
		}
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

func requestMove() int {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter move [0-6]: ")
		scanner.Scan()
		text := scanner.Text()
		v, err := strconv.Atoi(text)
		if err != nil {
			continue
		}
		return v
	}
}