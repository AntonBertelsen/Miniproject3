package main

import (
	pb "Mandatory_exercise_3"
	"bufio"
	"context"
	"fmt"
	"google.golang.org/grpc"
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
var budget float64

func main() {
	fmt.Println("CONNECTING!")
	// Connect
	conn, err := grpc.Dial("localhost:5001", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	defer conn.Close()

	// Create a unique ID to distinguish the bidder + the maximum bid they are willing to make
	rand.Seed(time.Now().UnixNano())
	id = int64(rand.Intn(math.MaxInt64))
	budget = float64(1000+(rand.Intn(1000))) // max budget is between 1000 and 2000

	// Prompt the bidder to enter name:
	fmt.Println("Enter your name:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name = scanner.Text()

	client := pb.NewAuctionServiceClient(conn)

	// create the client instance and await confirmation of join
	fmt.Printf("\n\n--- WELCOME TO THE AUCTION HOUSE %v---\n", strings.ToUpper(name))

	waitChannel := make(chan struct{})

	go func() {
		for{
			time.Sleep(500 * time.Millisecond)
			c := exec.Command("clear")
			c.Stdout = os.Stdout
			c.Run()

			outcome, err := client.Result(context.Background(),&pb.Void{})
			if err != nil {
				fmt.Println(err.Error())
			}

			if !outcome.Ended {
				fmt.Printf("=================================================================\n")
				fmt.Printf("                     AUCTION IS IN PROGRESS!                     \n")
				fmt.Printf("-----------------------------------------------------------------\n\n")
				fmt.Printf("%v seconds remaining", getSecondsRemaining(time.UnixMicro(outcome.EndTime)))
				fmt.Println()
				fmt.Printf("Current highest bid: %v\n",outcome.HighestBid)
				fmt.Printf("By: %v\n",outcome.HighestBidderName)
				fmt.Println()
				if outcome.HighestBidderID != id {
					fmt.Printf("You are not the highest bidder! :( \n")
					fmt.Printf("Press [ENTER] to bid %v \n\n \n", outcome.HighestBid + 50)
				} else {
					fmt.Printf("You are the highest bidder! :)\n\n \n\n")
				}
				fmt.Printf("=================================================================\n")
			} else {
				fmt.Printf("=================================================================\n")
				fmt.Printf("                         AUCTION IS OVER!                        \n")
				fmt.Printf("-----------------------------------------------------------------\n\n")

				if outcome.HighestBidderID != id {
					fmt.Printf("You did NOT win the auction! ಠ╭╮ಠ\n")
				} else {
					fmt.Printf("You WON the auction! \\ (•◡•) /\n\n \n")
				}
			}
		}
	}()

	go func() {
		for {
			fmt.Scanln()
			// Ask what the highest bid is
			outcome, err := client.Result(context.Background(),&pb.Void{})
			if err != nil {
				fmt.Println(err.Error())
			}

			if outcome.HighestBidderID != id && !outcome.Ended{
				// Bid the highest bid + 50, or the highest bid the bidder is able to make
				amount := outcome.HighestBid + 50

				_, _ = client.Bid(context.Background(), &pb.AmountMessage{
					Time:       0,
					Id:         id,
					BidderName: name,
					Amount:     amount,
				})
				//TODO: handle acknowledgement somehow
			}
		}
	}()
	<-waitChannel
}

func getSecondsRemaining(t time.Time) int {
	currentTime := time.Now()
	difference := t.Sub(currentTime)

	return int(difference.Seconds())
}