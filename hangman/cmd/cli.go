package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/imhotepio/letsgo_labs/hangman"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT)

	go func() {
		<-sigs
		fmt.Printf("\n\n 🙅‍♀️ No way!! Seriously??\n\n")
		os.Exit(0)
	}()

	g, err := hangman.NewGame("assets/words.txt")
	if err != nil {
		log.Fatal(err)
	}

	for !gameOver(g) {
		fmt.Printf("\nYour Word: %10s\n", string(g.Tally.Letters))
		c := prompt(g)
		if c != '\n' {
			_, err := g.Guess(c)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func gameOver(g *hangman.Game) bool {
	if g.Tally.Status == hangman.Won {
		fmt.Printf("\n👏  Noace!! You've just won\n\n")
		return true
	}

	if g.Tally.Status == hangman.Lost {
		fmt.Printf("\n😿  Meow! You've just lost. It was `%s\n\n", g.Tally.Word)
		return true
	}
	return false
}

func prompt(g *hangman.Game) rune {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("\n%10s [%d/%d]? ", "Your Guess", g.Tally.TurnsLeft, hangman.MaxGuesses)
	char, _, err := reader.ReadRune()
	if err != nil {
		panic(err)
	}
	return char
}
