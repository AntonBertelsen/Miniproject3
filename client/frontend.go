package main

import (
	pb "Mandatory_exercise_3"
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

type Frontend struct {
	replicationManagers []ReplicationManager
	leader              *pb.AuctionServiceClient
}

type ReplicationManager struct {
	replicationManager pb.AuctionServiceClient
	isLeader           bool
}

func newFrontend() *Frontend {

	f := &Frontend{
		replicationManagers: make([]ReplicationManager, 0),
	}

	fmt.Println("CONNECTING!")
	go f.connectToReplica("localhost:4242")
	go f.connectToReplica("localhost:6969")
	go f.connectToReplica("localhost:0420")
	return f
}

func (f *Frontend) connectToReplica(ip string) {
	// Connect with replica's IP
	conn, err := grpc.Dial(ip, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect gRPC server :: %v", err)
	}
	defer conn.Close()

	fmt.Println("Connected to Replica Manager at ip: " + ip)

	// Check if the replica manager is the leader.
	newServiceClient := pb.NewAuctionServiceClient(conn)
	isLeader, err := newServiceClient.IsLeader(context.Background(), &pb.Void{})

	if err == nil {
		// Append the new replicationManager to slice
		f.replicationManagers = append(f.replicationManagers, ReplicationManager{
			replicationManager: newServiceClient,
			isLeader:           isLeader.Value,
		})

		// If the node is the leader, set the leader in the Frontend
		if isLeader.Value == true {
			f.leader = &newServiceClient
		}
	} else {
		fmt.Println("ERROR " + err.Error())
	}

	// keep the go-routine running to not close the connection
	wait := make(chan bool)
	<-wait
}

func (f *Frontend) bid(name string, id int64) {
	// Leader must be obtained from pointer before you can call functions on it
	leader := *f.leader

	if leader != nil {
		outcome, err := leader.Result(context.Background(), &pb.Void{})
		check(err, "")

		if outcome.HighestBidderID != id && !outcome.Ended {
			// Bid the highest bid + 50, or the highest bid the bidder is able to make
			amount := outcome.HighestBid + 50

			_, _ = leader.Bid(context.Background(), &pb.AmountMessage{
				Time:       0,
				Id:         id,
				BidderName: name,
				Amount:     amount,
			})
			//handle acknowledgement somehow
		}
	} else {
		fmt.Println("Not connected to leader")
	}
}

func (f *Frontend) heartbeat() (*pb.OutcomeMessage, error) {

	// Ask each replication node if it is the leader
	for i := 0; i < len(f.replicationManagers); i++ {
		isLeaderMsg, err := f.replicationManagers[i].replicationManager.IsLeader(context.Background(), &pb.Void{})
		if err != nil {
			log.Println("Unable to reach a replication node")
			if f.replicationManagers[i].isLeader {
				log.Println("The leader node has failed")
				f.leader = nil
			}
			f.replicationManagers = removeReplicationManager(f.replicationManagers, i)
		} else {
			isLeader := isLeaderMsg.Value
			f.replicationManagers[i].isLeader = isLeader
			if isLeader == true {
				f.leader = &f.replicationManagers[i].replicationManager
			}
		}
	}

	if f.leader != nil {
		// Leader must be obtained from pointer before you can call functions on it
		leader := *f.leader

		outcome, err := leader.Result(context.Background(), &pb.Void{})
		check(err, "")
		return outcome, nil
	} else {
		fmt.Println("Not connected to leader")
		return nil, errors.New("no leader")
	}
}

func removeReplicationManager(s []ReplicationManager, i int) []ReplicationManager {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
