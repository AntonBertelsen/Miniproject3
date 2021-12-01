package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"
)

var id int64
var name string

var frontend *Frontend

func main() {

	fmt.Println(" ---- Setting up connection ---- ")
	frontend = newFrontend()

	// Create a unique ID to distinguish the bidder + the maximum bid they are willing to make
	rand.Seed(time.Now().UnixNano())
	id = int64(rand.Intn(math.MaxInt64))

	// Prompt the bidder to enter name:
	fmt.Println("Enter your name:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name = scanner.Text()

	// create the client instance and await confirmation of join
	fmt.Printf("\n\n--- WELCOME TO THE AUCTION HOUSE %v---\n", strings.ToUpper(name))

	go func() {

		fmt.Printf("Connected to %v replicas", len(frontend.replicationManagers))

		for {
			time.Sleep(500 * time.Millisecond)
			c := exec.Command("clear")
			c.Stdout = os.Stdout
			c.Run()

			outcome, err := frontend.heartbeat()

			if err == nil && outcome != nil {
				if !outcome.Ended {
					fmt.Printf("=================================================================\n")
					fmt.Printf("                     AUCTION IS IN PROGRESS!                     \n")
					fmt.Printf("-----------------------------------------------------------------\n\n")
					fmt.Printf("%v seconds remaining", getSecondsRemaining(time.UnixMicro(outcome.EndTime)))
					fmt.Println()
					fmt.Printf("Current highest bid: %v\n", outcome.HighestBid)
					fmt.Printf("By: %v\n", outcome.HighestBidderName)
					fmt.Println()
					if outcome.HighestBidderID != id {
						fmt.Printf("You are not the highest bidder! :( \n")
						fmt.Printf("Press [ENTER] to bid %v \n\n \n", outcome.HighestBid+50)
					} else {
						fmt.Printf("You are the highest bidder! :)\n\n\n\n")
					}
					fmt.Printf("(connected to %v replicas)\n", len(frontend.replicationManagers))
					fmt.Printf("=================================================================\n")
				} else {
					fmt.Printf("=================================================================\n")
					fmt.Printf("                         AUCTION IS OVER!                        \n")
					fmt.Printf("-----------------------------------------------------------------\n\n")

					if outcome.HighestBidderID != id {
						fmt.Printf("You did NOT win the auction! ಠ╭╮ಠ\n")
					} else {
						fmt.Printf("You WON the auction! \\ (•◡•) /\n\n\n")
					}
				}
			}
		}
	}()

	go func() {
		for {
			// Wait for user input
			fmt.Scanln()

			// Bid
			frontend.bid(name, id)
		}
	}()

	wait := make(chan struct{})
	<-wait
}

func getSecondsRemaining(t time.Time) int {
	currentTime := time.Now()
	difference := t.Sub(currentTime)

	return int(difference.Seconds())
}

func check(err error, message string) {
	if err != nil {
		if message != "" {
			log.Printf("%v :: %v\n", message)
		} else {
			log.Printf("Exception")
		}
	}
}
