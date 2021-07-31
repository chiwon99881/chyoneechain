package p2p

import (
	"encoding/json"
	"fmt"

	"github.com/chiwon99881/chyocoin/blockchain"
	"github.com/chiwon99881/chyocoin/utils"
)

// MessageKind is type of integer
type MessageKind int

const (
	// MessageNewestBlock is iota variables
	MessageNewestBlock MessageKind = iota
	// MessageAllBlocksRequest is iota variables
	MessageAllBlocksRequest
	// MessageAllBlocksResponse is iota variables
	MessageAllBlocksResponse
)

// Message is type of message struct
type Message struct {
	Kind    MessageKind
	Payload []byte
}

func makeMessage(kind MessageKind, payload interface{}) []byte {
	m := Message{
		Kind:    kind,
		Payload: utils.ToJSON(payload),
	}
	return utils.ToJSON(m)
}

func sendNewestBlock(p *peer) {
	b, err := blockchain.FindBlock(blockchain.Blockchain().NewestHash)
	utils.HandleError(err)
	m := makeMessage(MessageNewestBlock, b)
	p.inbox <- m
}

func handleMsg(m *Message, p *peer) {
	switch m.Kind {
	case MessageNewestBlock:
		var payload blockchain.Block
		err := json.Unmarshal(m.Payload, &payload)
		utils.HandleError(err)
		fmt.Println(payload)
	}
}
