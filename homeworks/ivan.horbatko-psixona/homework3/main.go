package main

import (
	"flag"
	"fmt"
	"net"

	"golang_2019/homeworks/ivan.horbatko-psixona/homework3/game"
)

func main() {
	opponent := flag.String("opponent", "none", "type of opponent: human, remote")
	connect := flag.String("connect", "none", "ip of host name to join the game")
	flag.Parse()

	if *opponent == "none" && *connect == "none" {
		fmt.Println("Please select correct opponent type or connect")
		return
	}

	if *opponent != "none" && *connect != "none" {
		fmt.Println("Cannot be sent both at same time")
		return
	}

	if *opponent != "none" {
		if *opponent == "remote" {
			fmt.Println("Listen on port 8080 for opponent connect")
			ln, err := net.Listen("tcp", ":8080")
			if err != nil {
				panic(err)
			}
			conn, err := ln.Accept()
			if err != nil {
				panic(err)
			}
			player1 := game.NewRemoteNetworkPlayer(game.O, conn)
			player2 := game.NewLocalPlayer(game.X, conn)
			game.NewTicTacToeGame(player1, player2).Run()
		} else if *opponent == "human" {
			player1 := game.NewHuman_Player(game.X)
			player2 := game.NewHuman_Player(game.O)
			game.NewTicTacToeGame(player1, player2).Run()
		}

	}
	if *connect != "none" {
		conn, err := net.Dial("tcp", *connect)
		if err != nil {
			panic(err)
		}
		player1 := game.NewLocalPlayer(game.O, conn)
		player2 := game.NewRemoteNetworkPlayer(game.X, conn)
		game.NewTicTacToeGame(player1, player2).Run()
	}
}
