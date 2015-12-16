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
	for !g.IsGameOver() {
		fmt.Println(g)
		v := requestMove()
		if err := g.Move(v); err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
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