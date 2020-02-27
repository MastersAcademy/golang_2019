package game

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

type Player interface {
	MakeMove(game *TicTacToeGame)
	Value() SlotValue
}

type BasePlayer struct {
	value SlotValue
}

func NewBasePlayer(v SlotValue) *BasePlayer {
	return &BasePlayer{
		value: v,
	}
}

func (b *BasePlayer) Value() SlotValue {
	return b.value
}

type LocalNetworkPlayer struct {
	conn net.Conn
	*BasePlayer
}

func NewLocalPlayer(value SlotValue, conn net.Conn) *LocalNetworkPlayer {
	return &LocalNetworkPlayer{
		conn:       conn,
		BasePlayer: NewBasePlayer(value),
	}
}

func (h *LocalNetworkPlayer) MakeMove(game *TicTacToeGame) {
	var slot int
	success := false
	for !success {
		fmt.Println("Please enter a value 1-9 to represent the slot to make your move")
		_, err := fmt.Scanf("%d", &slot)
		if err != nil {
			fmt.Println("Invalid slot selected")
			break
		}
		success = game.MakeMove(slot-1, h.value)
		if !success {
			fmt.Println("Cannot make a move in that slot")
		}
	}
	fmt.Fprintf(h.conn, strconv.Itoa(slot)+"\n")
}

type Human_Player struct {
	*BasePlayer
}

func NewHuman_Player(value SlotValue) *Human_Player {
	return &Human_Player{
		BasePlayer: NewBasePlayer(value),
	}
}

func (h *Human_Player) MakeMove(game *TicTacToeGame) {
	success := false
	for !success {
		fmt.Println("Please enter a value 1-9 to represent the slot to make your move")
		var slot int
		_, err := fmt.Scanf("%d", &slot)
		if err != nil {
			fmt.Println("Invalid slot selected")
			break
		}
		success = game.MakeMove(slot-1, h.value)
		if !success {
			fmt.Println("Cannot make a move in that slot")
		}
	}
}

type RemoteNetworkPlayer struct {
	conn net.Conn
	*BasePlayer
}

func NewRemoteNetworkPlayer(value SlotValue, conn net.Conn) *RemoteNetworkPlayer {
	return &RemoteNetworkPlayer{
		conn:       conn,
		BasePlayer: NewBasePlayer(value),
	}
}

func (h *RemoteNetworkPlayer) MakeMove(game *TicTacToeGame) {
	slot, err := bufio.NewReader(h.conn).ReadString('\n')
	if err != nil {
		// TODO remove this later
		panic(err)
	}
	strSlot, err := strconv.Atoi(strings.Trim(slot, " \n"))
	if err != nil {
		// TODO remove this later
		panic(err)
	}
	game.MakeMove(strSlot-1, h.Value())
}
