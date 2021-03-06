package main

import (
	"log"
	"os"

	"github.com/Jeffail/gabs"
	"github.com/Murmuration-Labs/bitscreen"
	"github.com/ipfs/go-cid"
)

func main() {
	bitscreen.MaybeCreateBitscreen()

	proposal, err := gabs.ParseJSONBuffer(os.Stdin)
	if err != nil {
		log.Fatalf("Unable to parse proposal JSON: %s", err)
	}

	for _, path := range [][]string{
		[]string{"Proposal", "PieceCID", "/"},
		[]string{"Proposal", "Label"},
		[]string{"Ref", "Root", "/"},
		[]string{"PayloadCID", "/"},
	} {
		c, err := cid.Parse(proposal.Search(path...).Data())
		// check only if found a valid CID
		if err == nil {
			if bitscreen.BlockCid(c) {
				os.Exit(1)
			}
		}
	}
}
