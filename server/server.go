package main

import (
	pb "Mandatory_exercise_3"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

var lamportTime int64

type AuctionHouse struct {
	pb.UnimplementedAuctionServiceServer
	bidders []Bidder
	highestBidder *Bidder
	endTime time.Time
	ended bool
}

func newServer() *AuctionHouse {
	s := &AuctionHouse{
		bidders: make([]Bidder, 0),
		highestBidder: &Bidder{
			id:         -1,
			name:       "No one has made a bid yet",
			highestBid: 0,
		},
		endTime: time.Now().Add(time.Minute * 5),
	}
	return s
}

type Bidder struct {
	id int64
	name string
	highestBid float64
}

func main() {

	lis, err := net.Listen("tcp", "localhost:5001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	server := newServer()

	pb.RegisterAuctionServiceServer(s, server)
	log.Printf("server listening at %v", lis.Addr())

	log.Printf("\n\n--- Action House ---")
	go func() {
		log.Println("NOW BIDDING!")
		time.Sleep(server.endTime.Sub(time.Now()))
		server.ended = true
		fmt.Println("THE AUCTION IS OVER!")
	}()
	s.Serve(lis)
}

func (auctionHouse *AuctionHouse) Bid(_ context.Context, message *pb.AmountMessage) (*pb.AcknowledgementMessage, error){

	// Determine the bidder
	var bidder *Bidder

	for i := 0; i < len(auctionHouse.bidders); i++ {
		if auctionHouse.bidders[i].id == message.Id {
			bidder = &auctionHouse.bidders[i]
		}
	}

	// If this is the bidder's first bid
	if bidder == nil {
		bidder = &Bidder {
			id:         message.Id,
			name:       message.BidderName,
			highestBid: message.Amount,
		}
		auctionHouse.bidders = append(auctionHouse.bidders, *bidder)
	}

	if auctionHouse.highestBidder.highestBid < message.Amount {
		auctionHouse.highestBidder = bidder
		bidder.highestBid = message.Amount
		return &pb.AcknowledgementMessage{}, nil
	}
	return &pb.AcknowledgementMessage{}, nil
}

func (auctionHouse *AuctionHouse) Result(_ context.Context, _ *pb.Void) (*pb.OutcomeMessage, error){

	return &pb.OutcomeMessage{
		Time:              lamportTime,
		HighestBid:        auctionHouse.highestBidder.highestBid,
		HighestBidderID:   auctionHouse.highestBidder.id,
		HighestBidderName: auctionHouse.highestBidder.name,
		Ended:             auctionHouse.ended,
		EndTime:		   auctionHouse.endTime.UnixMicro(),
	}, nil
}