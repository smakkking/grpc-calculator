package main

import (
	"context"
	"fmt"

	"grpc_calc/pkg/sdk/go/calcv1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial(":12345", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		panic("cant connect to server")
	}

	c := calcv1.NewCalculatorClient(conn)

	for {
		var a, b int64
		var operation rune
		fmt.Scanf("%d %c %d", &a, &operation, &b)

		switch operation {
		case '+':
			resp, err := c.Add(context.Background(), &calcv1.Request{A: a, B: b})
			if err != nil {
				fmt.Printf("%s\n", err.Error())
				continue
			}
			fmt.Println(resp.Result)

		case '-':
			resp, _ := c.Sub(context.Background(), &calcv1.Request{A: a, B: b})
			if err != nil {
				fmt.Printf("%s\n", err.Error())
				continue
			}
			fmt.Println(resp.Result)
		case '*':
			resp, _ := c.Mul(context.Background(), &calcv1.Request{A: a, B: b})
			if err != nil {
				fmt.Printf("%s\n", err.Error())
				continue
			}
			fmt.Println(resp.Result)
		case '/':
			resp, _ := c.Div(context.Background(), &calcv1.Request{A: a, B: b})
			if err != nil {
				fmt.Printf("%s\n", err.Error())
				continue
			}
			fmt.Println(resp.Result)
		default:
			fmt.Println("error: no such operation. Try again!")
		}
	}
}
