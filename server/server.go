package main

import (
	pb "Mandatory_exercise_3"
	"context"
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

var auctionHouseServer *AuctionHouse
var replicationServer *ReplicationServer

type AuctionHouse struct {
	pb.UnimplementedAuctionServiceServer
	bidders       []Bidder
	highestBidder *Bidder
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
	isLeader bool
	id       int64

	// Server is connected to all other replication clients as a client
	replicationClients []ReplicationClient
}

type ReplicationClient struct {
	replicationClient pb.ReplicationServiceClient
	isLeader          bool
	id                int64
}

func main() {

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

	log.Printf("\n\n--- Action House ---")
	go func() {
		log.Println("NOW BIDDING!")
		time.Sleep(auctionHouseServer.endTime.Sub(time.Now()))
		auctionHouseServer.ended = true
		fmt.Println("THE AUCTION IS OVER!")
	}()

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

	rand.Seed(time.Now().UnixNano())

	// 4242 is the initial leader
	s := &ReplicationServer{
		replicationClients: make([]ReplicationClient, 0),
		isLeader:           os.Args[1] == "4242",
		id:                 int64(rand.Intn(math.MaxInt64)),
	}

	fmt.Println("Connecting to other replication nodes")
	// If something goes wrong in the future it is probably here where we connect to our selves
	go s.connectToReplica("localhost:4242")
	go s.connectToReplica("localhost:6969")
	go s.connectToReplica("localhost:0420")

	fmt.Println(s)
	return s
}

func (r *ReplicationServer) connectToReplica(ip string) {
	// Connect with replica's IP
	conn, err := grpc.Dial(ip, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect gRPC server :: %v", err)
	}
	defer conn.Close()

	fmt.Println("Connected to Replica Manager at ip: " + ip)

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
		} else {
			fmt.Println(err.Error())
		}
		fmt.Println("asking!")
		time.Sleep(500 * time.Millisecond)
	}
	// Append the new replicationManager to slice
	r.replicationClients = append(r.replicationClients, ReplicationClient{
		replicationClient: newReplicationClient,
		isLeader:          newIsLeader,
		id:                newId,
	})

	fmt.Println(r)

	// keep the go-routine running to not close the connection
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

func (auctionHouse *AuctionHouse) Bid(_ context.Context, message *pb.AmountMessage) (*pb.AcknowledgementMessage, error) {

	// Determine the bidder
	var bidder *Bidder

	for i := 0; i < len(auctionHouse.bidders); i++ {
		if auctionHouse.bidders[i].id == message.Id {
			bidder = &auctionHouse.bidders[i]
		}
	}

	// If this is the bidder's first bid
	if bidder == nil {
		bidder = &Bidder{
			id:         message.Id,
			name:       message.BidderName,
			highestBid: message.Amount,
		}
		auctionHouse.bidders = append(auctionHouse.bidders, *bidder)
	}

	if auctionHouse.highestBidder.highestBid < message.Amount {
		auctionHouse.highestBidder = bidder
		bidder.highestBid = message.Amount
	}

	// Replicate the auctionHouse to the other nodes
	fmt.Println("Big was made!")
	fmt.Println("Calling replication")
	auctionHouse.replicateAuctionHouse()
	return &pb.AcknowledgementMessage{}, nil
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

func (auctionHouse *AuctionHouse) replicateAuctionHouse() {

	fmt.Println("Starting replication")

	// Replicate result
	for i := 0; i < len(replicationServer.replicationClients); i++ {
		fmt.Println("xxxxxxx")

		bidders := make([]*pb.Bidder, 0)
		// IF SOMETHING IS NOT WORKING LOOK HERE!
		for _, b := range auctionHouse.bidders {
			bidders = append(bidders, &pb.Bidder{
				Id:         b.id,
				Name:       b.name,
				HighestBid: b.highestBid,
			})
		}
		fmt.Println(bidders)
		_, err := replicationServer.replicationClients[i].replicationClient.Replicate(context.Background(), &pb.AuctionHouseMessage{
			Bidders:         bidders,
			HighestBidderID: 0,
			EndTime:         0,
			Ended:           false,
		})
		fmt.Println("done")
		check(err, "Something went wrong! Could not replicate")
	}
}

func (auctionHouse *AuctionHouse) IsLeader(_ context.Context, _ *pb.Void) (*pb.BooleanMessage, error) {
	return &pb.BooleanMessage{Value: replicationServer.isLeader}, nil
}

func (r *ReplicationServer) RequestInfo(_ context.Context, _ *pb.Void) (*pb.ReplicationManagerMessage, error) {
	return &pb.ReplicationManagerMessage{
		Id:       r.id,
		IsLeader: r.isLeader,
	}, nil
}

func (r *ReplicationServer) Replicate(_ context.Context, msg *pb.AuctionHouseMessage) (*pb.AcknowledgementMessage, error) {

	fmt.Println("REPLICATING!")
	newBidders := make([]Bidder, 0)

	fmt.Println("Length of msg.bidders")
	fmt.Println(len(msg.Bidders))
	fmt.Println(msg.Bidders)
	fmt.Println("--------")

	//potential problem
	for _, bidder := range msg.Bidders {
		fmt.Println("for loop!")
		newBidders = append(newBidders, Bidder{
			id:         bidder.Id,
			name:       bidder.Name,
			highestBid: bidder.HighestBid,
		})
	}
	fmt.Println("now down heere")

	var highestBidder *Bidder
	for i := 0; i < len(newBidders); i++ {
		fmt.Println("for loop 2")
		if newBidders[i].id == msg.HighestBidderID {
			highestBidder = &newBidders[i]
			fmt.Println("new highest bidder!")
		}
	}

	fmt.Println("Previous auction house:")
	fmt.Println(auctionHouseServer.bidders)

	auctionHouseServer = &AuctionHouse{
		UnimplementedAuctionServiceServer: pb.UnimplementedAuctionServiceServer{},
		bidders:                           newBidders,
		highestBidder:                     highestBidder,
		endTime:                           time.UnixMicro(msg.EndTime),
		ended:                             msg.Ended,
	}

	fmt.Println("New auction house:")
	fmt.Println(auctionHouseServer.bidders)

	return &pb.AcknowledgementMessage{}, nil
}

// backup replicant nodes constantly ping the leader to make sure it still exists. In case it does not the node
func (r *ReplicationServer) heartbeat() {
	// Request info from the client repeatedly to discover changes and to notice of the connection is lost
	for i := 0; i < len(r.replicationClients); i++ {
		isLeaderMsg, err := r.replicationClients[i].replicationClient.RequestInfo(context.Background(), &pb.Void{})
		if err != nil {
			log.Println("Unable to reach a replication node")
			if r.replicationClients[i].isLeader {
				log.Println("The leader node has failed")

				log.Println(r.replicationClients)

				log.Println("Now removing lost client:")
				// We remove the node
				r.replicationClients = removeReplicationClient(r.replicationClients, i)
				log.Println(r.replicationClients)

				// We now need to assign a new leader
				// Each replication node has a unique ID. Leadership is passed to the node with the highest ID
				// Since we already know the ID of every replication node, we simply assign ourselves as the leader
				// if we have the highest id with no need to contact the other nodes first.
				// Next time the other nodes ping us they will learn that we are the new leader.

				log.Println("Calling election")

				var highest int64
				for i2 := 0; i2 < len(r.replicationClients); i2++ {
					if r.replicationClients[i2].id > highest {
						highest = r.replicationClients[i2].id
					}
				}
				log.Println("Setting new leader")
				log.Println(highest)

				// If we have the highest index set as new leader
				if highest == r.id {
					log.Println("I am the new leader!")
					fmt.Println(auctionHouseServer)
					r.isLeader = true
				}
			} else {

				log.Println(r.replicationClients)
				log.Println("Now removing lost client:")
				r.replicationClients = removeReplicationClient(r.replicationClients, i)
				log.Println(r.replicationClients)
			}
		} else {
			r.replicationClients[i].isLeader = isLeaderMsg.IsLeader
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
