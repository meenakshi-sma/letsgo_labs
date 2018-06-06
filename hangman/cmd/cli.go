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
	var err error

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT)

	go func() {
		<-sigs
		fmt.Printf("\n\n 🙅‍♀️ No way!! Seriously??\n\n")
		os.Exit(0)
	}()

	var g *hangman.Game
	if g, err = hangman.NewGame("assets/words.txt"); err != nil {
		log.Fatal(err)
	}

	t := g.CurrentTally()
	for !gameOver(t) {
		fmt.Printf("\nYour Word: %10s\n", string(t.Letters))
		c := prompt(t)
		if c != '\n' {
			if t, err = g.Guess(c); err != nil {
				log.Fatal(err)
			}
		}
	}
}

func gameOver(t *hangman.Tally) bool {
	if t.Status == hangman.Won {
		fmt.Printf("\n👏  Noace!! You've just won\n\n")
		return true
	}

	if t.Status == hangman.Lost {
		fmt.Printf("\n😿  Meow! You've just lost. It was `%s\n\n", t.Word)
		return true
	}
	return false
}

func prompt(t *hangman.Tally) rune {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("\n%10s [%d/%d]? ", "Your Guess", t.TurnsLeft, hangman.MaxGuesses)
	char, _, err := reader.ReadRune()
	if err != nil {
		panic(err)
	}
	return char
}
