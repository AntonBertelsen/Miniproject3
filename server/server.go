package main

import (
	pb "Mandatory_exercise_3"
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"math"
	"math/rand"
	"net"
	"os"
	"time"
)

var lamportTime int64

/*
-> client bids!
-> frontend
-> auction house server on leader node
-> leader node (as a client) requests other replication managers (as servers) to update their data.
   the client (leader node) waits for acknowledgement response. once complete...
-> frontend
-> client (view updated with new data)
*/

// We store 2 services

// AuctionHouseServer is reachable by the clients / users of the auction house
var auctionHouseServer *AuctionHouse

// ReplicationServer is reachable by the replication nodes in the system. Each node is both a server
// and a client connected to every other replication node
var replicationServer *ReplicationServer

type AuctionHouse struct {
	pb.UnimplementedAuctionServiceServer
	bidders       []Bidder
	highestBidder *Bidder // Pointer to the highest bidder within the bidders array
	endTime       time.Time
	ended         bool
}

type Bidder struct {
	id         int64
	name       string
	highestBid float64
}

type ReplicationServer struct {
	pb.UnimplementedReplicationServiceServer
	isLeader           bool
	id                 int64
	replicationClients []ReplicationClient // Server is connected to all other replication clients as a client
}

// ReplicationClient is used to store each replication node from the perspective of the replication node server
// This means every node has every other node (including itself) stored as a replication client
type ReplicationClient struct {
	replicationClient pb.ReplicationServiceClient
	isLeader          bool
	id                int64
}

func main() {
	// Listen traffic on the port supplied as console line argument
	lis, err := net.Listen("tcp", "localhost:"+os.Args[1])
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", os.Args[1], err)
	}
	log.Printf("server listening at %v", lis.Addr())

	// Instantiate a grpc server with no services registered
	s := grpc.NewServer()

	// Instantiate, and register an AuctionHouse server
	auctionHouseServer = newAuctionHouseServer()
	pb.RegisterAuctionServiceServer(s, auctionHouseServer)

	// Instantiate, and register a Replication -server-
	replicationServer = newReplicationServer()
	pb.RegisterReplicationServiceServer(s, replicationServer)

	go func() {
		log.Printf("\n\n--- Action House ---")
		log.Printf("Clients can now bid. Auction ends in %v seconds\n=================\n\n", auctionHouseServer.endTime.Sub(time.Now()).Seconds())

		// Let this goroutine sleep for the duration of the auction
		time.Sleep(auctionHouseServer.endTime.Sub(time.Now()))

		// End the auction
		auctionHouseServer.ended = true
		log.Printf("The auction has ended!\n")
		log.Printf("The winner was %v bidding %v!\n=================\n\n", auctionHouseServer.highestBidder.name, auctionHouseServer.highestBidder.highestBid)
	}()

	// We assume that all replication nodes can stop functioning at any point. To accomplish this we 'ping' every node every 500ms to make sure it is still reachable
	go func() {
		for {
			time.Sleep(500 * time.Millisecond)
			replicationServer.heartbeat()
		}
	}()

	s.Serve(lis)
}

// Creates a replication server node storing connections to the other replication nodes
func newReplicationServer() *ReplicationServer {

	// Generate a random ID
	// Here we simply rely on how improbable it is that two clients generate the same ID
	rand.Seed(time.Now().UnixNano())
	randomID := int64(rand.Intn(math.MaxInt64))

	s := &ReplicationServer{
		replicationClients: make([]ReplicationClient, 0),
		isLeader:           os.Args[1] == "6001", // 6001 is redefined as the leader
		id:                 randomID,
	}

	fmt.Println("Connecting to other replication nodes")
	go s.connectToReplica("localhost:6001")
	go s.connectToReplica("localhost:6002")
	go s.connectToReplica("localhost:6003")
	go s.connectToReplica("localhost:6004")
	go s.connectToReplica("localhost:6005")
	return s
}

func (r *ReplicationServer) connectToReplica(ip string) {

	// Connect with replica's IP
	conn, err := grpc.Dial(ip, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect gRPC server :: %v", err)
	}
	defer conn.Close()

	log.Println("Attempting to reach Replica Manager at ip: " + ip)

	// Create a new replicationClient struct to be stored in the replication server struct
	newReplicationClient := pb.NewReplicationServiceClient(conn)

	// Check if the replica manager is the leader
	// Since the server may not exist yet, we keep asking until we get a response
	var newIsLeader bool
	var newId int64
	for {
		replicationClientMessage, err := newReplicationClient.RequestInfo(context.Background(), &pb.Void{})
		if err == nil {
			newIsLeader = replicationClientMessage.IsLeader
			newId = replicationClientMessage.Id
			break
		}
		// Retry until the connection is established
		time.Sleep(500 * time.Millisecond)
	}

	log.Println("Successfully reached the Replica Manager!")

	// Append the new ReplicationClient to the list of replicationClients stored in the ReplicationServer struct
	r.replicationClients = append(r.replicationClients, ReplicationClient{
		replicationClient: newReplicationClient,
		isLeader:          newIsLeader,
		id:                newId,
	})

	// Keep the go-routine running to not close the connection
	wait := make(chan bool)
	<-wait
}
func newAuctionHouseServer() *AuctionHouse {
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

func (auctionHouse *AuctionHouse) IsLeader(_ context.Context, _ *pb.Void) (*pb.BooleanMessage, error) {
	return &pb.BooleanMessage{Value: replicationServer.isLeader}, nil
}

func (auctionHouse *AuctionHouse) Bid(_ context.Context, message *pb.AmountMessage) (*pb.AcknowledgementMessage, error) {

	// Determine the bidder stored in our list of bidders from the supplied id
	var bidder *Bidder
	for i := 0; i < len(auctionHouse.bidders); i++ {
		if auctionHouse.bidders[i].id == message.Id {
			bidder = &auctionHouse.bidders[i]
		}
	}

	// If this is the bidder's first bid, add the bidder
	if bidder == nil {
		bidder = &Bidder{
			id:         message.Id,
			name:       message.BidderName,
			highestBid: message.Amount,
		}
		auctionHouse.bidders = append(auctionHouse.bidders, *bidder)
	}

	var err error
	// If the bid is the highest bid so far (which it should always unless someone made a bid at the same time that arrived earlier)
	if auctionHouse.highestBidder.highestBid < message.Amount {
		auctionHouse.highestBidder = bidder
		bidder.highestBid = message.Amount
	} else {
		err = errors.New("bid less than highest bid")
	}

	// Replicate the auctionHouse to the other nodes
	auctionHouse.requestAuctionHouseReplication()

	// Return an acknowledgement message
	return &pb.AcknowledgementMessage{}, err
}

func (auctionHouse *AuctionHouse) Result(_ context.Context, _ *pb.Void) (*pb.OutcomeMessage, error) {

	return &pb.OutcomeMessage{
		Time:              lamportTime,
		HighestBid:        auctionHouse.highestBidder.highestBid,
		HighestBidderID:   auctionHouse.highestBidder.id,
		HighestBidderName: auctionHouse.highestBidder.name,
		Ended:             auctionHouse.ended,
		EndTime:           auctionHouse.endTime.UnixMicro(),
	}, nil
}

func (auctionHouse *AuctionHouse) requestAuctionHouseReplication() {

	// In order to send the bidders through GRPC we need to reconstruct the array using protobuffer types
	bidders := make([]*pb.Bidder, 0)

	for _, b := range auctionHouse.bidders {
		bidders = append(bidders, &pb.Bidder{
			Id:         b.id,
			Name:       b.name,
			HighestBid: b.highestBid,
		})
	}

	// For every Replication Node stored as a client on this replica server
	for i := 0; i < len(replicationServer.replicationClients); i++ {

		// Call the replication RPC on the replication node server associated with the replication node client
		_, err := replicationServer.replicationClients[i].replicationClient.Replicate(context.Background(), &pb.AuctionHouseMessage{
			Bidders:         bidders,
			HighestBidderID: auctionHouse.highestBidder.id,
			EndTime:         auctionHouse.endTime.UnixMicro(),
			Ended:           auctionHouse.ended,
		})
		check(err, "Something went wrong! Could not replicate")
	}
}

func (r *ReplicationServer) Replicate(_ context.Context, msg *pb.AuctionHouseMessage) (*pb.AcknowledgementMessage, error) {

	// Reconstruct bidders array from the protobuffer types included in the msg
	newBidders := make([]Bidder, 0)

	for _, bidder := range msg.Bidders {
		newBidders = append(newBidders, Bidder{
			id:         bidder.Id,
			name:       bidder.Name,
			highestBid: bidder.HighestBid,
		})
	}

	// Grab a pointer to the highest bidder in the array
	var highestBidder *Bidder
	for i := 0; i < len(newBidders); i++ {
		if newBidders[i].id == msg.HighestBidderID {
			highestBidder = &newBidders[i]
		}
	}

	// Assign replicated changes to the auctionHouse
	auctionHouseServer.bidders = newBidders
	auctionHouseServer.highestBidder = highestBidder
	auctionHouseServer.endTime = time.UnixMicro(msg.EndTime)
	auctionHouseServer.ended = msg.Ended

	return &pb.AcknowledgementMessage{}, nil
}

func (r *ReplicationServer) RequestInfo(_ context.Context, _ *pb.Void) (*pb.ReplicationManagerMessage, error) {
	return &pb.ReplicationManagerMessage{
		Id:       r.id,
		IsLeader: r.isLeader,
	}, nil
}

func (r *ReplicationServer) heartbeat() {

	// For each replication client, make sure it does still exist
	for i := 0; i < len(r.replicationClients); i++ {

		// Request info from the client repeatedly to discover changes and to notice of the connection is lost
		nodeInfo, err := r.replicationClients[i].replicationClient.RequestInfo(context.Background(), &pb.Void{})
		if err != nil {
			// If an error has occurred it means we were unable to reach the replication node.

			log.Println("Unable to reach a replication node!")

			// Check if the unreachable node was the leader
			lostNodeWasLeader := r.replicationClients[i].isLeader

			// Remove the node
			log.Println("Now removing lost node from known replication nodes")
			r.replicationClients = removeReplicationClient(r.replicationClients, i)
			log.Printf("System now has %v replication nodes\n", len(r.replicationClients))

			if lostNodeWasLeader {

				log.Println("Unable to reach leader node")
				log.Println("A new leader must be assigned")

				// We now need to assign a new leader
				// Each replication node has a unique ID. Leadership is passed to the node with the highest ID
				// Since we already know the ID of every replication node, we simply assign ourselves as the leader
				// if we have the highest id with no need to contact the other nodes first.
				// Next time the other nodes ping us they will learn that we are the new leader.

				// Determine the node with the highest ID
				var highest int64
				for i2 := 0; i2 < len(r.replicationClients); i2++ {
					if r.replicationClients[i2].id > highest {
						highest = r.replicationClients[i2].id
					}
				}
				// If we have the highest index set as new leader
				if highest == r.id {
					log.Println("I am the new leader!")
					r.isLeader = true
				}
				log.Println("A new leader has been picked")
			}
		} else {
			// Set isLeader based on node info. This lets the node discover new leaders
			r.replicationClients[i].isLeader = nodeInfo.IsLeader
		}
	}
}

func removeReplicationClient(s []ReplicationClient, i int) []ReplicationClient {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
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
