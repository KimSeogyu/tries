package trie

import (
	"github.com/ethereum/go-ethereum/rlp"
	"log"
)

type Node interface {
	Hash() []byte
	Raw() []any
}

func Hash(node Node) []byte {
	if IsEmptyNode(node) {
		return EmptyNodeHash
	}

	return node.Hash()
}

func Serialize(node Node) []byte {
	var raw any
	if IsEmptyNode(node) {
		raw = EmptyNodeRaw
	} else {
		raw = node.Raw()
	}

	b, err := rlp.EncodeToBytes(raw)
	if err != nil {
		log.Fatal(err)
	}

	return b
}
