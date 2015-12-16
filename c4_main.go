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
	for {
		fmt.Println(g)
		fmt.Printf("Player %v's turn!\n", g.Turn)
		v := requestMove()
		if err := g.Move(v); err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}
		if g.IsGameOver() {
			fmt.Println(g)
			fmt.Printf("Player %v wins!\n", g.Turn)
			return
		}
		g.NextTurn()
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