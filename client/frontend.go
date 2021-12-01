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

	log.Println("CONNECTING!")
	go f.connectToReplica("localhost:6001")
	go f.connectToReplica("localhost:6002")
	go f.connectToReplica("localhost:6003")
	go f.connectToReplica("localhost:6004")
	go f.connectToReplica("localhost:6005")
	return f
}

func (f *Frontend) connectToReplica(ip string) {
	// Sometimes connecting fails for no obvious reason. We retry connecting 5 times if connection is not established yet.
	retry := 5
	for i := 0; i < retry; i++ {
		// Connect with replica's IP
		conn, err := grpc.Dial(ip, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Failed to connect gRPC server :: %v", err)
		}
		defer conn.Close()

		fmt.Println("Attempting to reach Replica Manager at ip: " + ip)

		// Check if the replica manager is the leader.
		newServiceClient := pb.NewAuctionServiceClient(conn)
		isLeader, err := newServiceClient.IsLeader(context.Background(), &pb.Void{})

		// If nothing is wrong
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

			// Stop attempting to reach the node since the connection has been established
			retry = 0
		} else {
			fmt.Println("Could not reach replication node: " + err.Error())
		}
	}
	// keep the go-routine running to not close the connection
	wait := make(chan bool)
	<-wait
}

func (f *Frontend) bid(name string, id int64) {
	// Leader must be obtained from pointer before you can call functions on it
	leader := *f.leader

	// If there is a known leader
	if leader != nil {
		outcome, err := leader.Result(context.Background(), &pb.Void{})
		if err == nil {
			if outcome.HighestBidderID != id && !outcome.Ended {
				// Bid the highest bid + 50, or the highest bid the bidder is able to make
				amount := outcome.HighestBid + 50

				_, err = leader.Bid(context.Background(), &pb.AmountMessage{
					Time:       0,
					Id:         id,
					BidderName: name,
					Amount:     amount,
				})
				if err != nil {
					log.Println(err)
				}
			}
		}
	} else {
		fmt.Println("Not connected to leader")
	}
}

func (f *Frontend) heartbeat() (*pb.OutcomeMessage, error) {

	// Ask each replication node if it is the leader
	for i := 0; i < len(f.replicationManagers); i++ {
		isLeaderMsg, err := f.replicationManagers[i].replicationManager.IsLeader(context.Background(), &pb.Void{})

		// If there is an error the replication node is unreachable
		if err != nil {
			//log.Println("Unable to reach a replication node")

			// If the leader node is unreachable, remove the reference to the leader node
			if f.replicationManagers[i].isLeader {
				//log.Println("The leader node has failed")
				f.leader = nil
			}

			// Remove the lost node from the slice
			f.replicationManagers = removeReplicationManager(f.replicationManagers, i)
		} else {
			isLeader := isLeaderMsg.Value
			f.replicationManagers[i].isLeader = isLeader
			if isLeader == true {
				f.leader = &f.replicationManagers[i].replicationManager
			}
		}
	}

	// Obtain the current outcome of the auction, so it can be displayed by the client UI
	if f.leader != nil {
		// Leader must be obtained from pointer before you can call functions on it
		leader := *f.leader

		outcome, _ := leader.Result(context.Background(), &pb.Void{})
		return outcome, nil
	} else {
		// Not connected to a leader
		return nil, errors.New("no leader")
	}
}

func removeReplicationManager(s []ReplicationManager, i int) []ReplicationManager {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
