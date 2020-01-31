package main

import (
	"context"
	"github.com/souvikhaldar/grpc-tutorial/add"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial(":8192", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Error in dialing in: ", err)
	}
	defer conn.Close()

	client := add.NewAddClient(conn)
	result, err := client.Adder(context.Background(), &add.AdderRequest{FirstNum: 1, SecondNum: 2})
	if err != nil {
		log.Println("Erron in calling the procedure: ", err)
		return
	}
	log.Println("Result of summation: ", result.GetAddResult())

}
