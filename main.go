package main

import (
	"fmt"
	"log"

	"github.com/itchyny/volume-go"
)

func main() {
	vol, err := volume.GetVolume()
	if err != nil {
		log.Fatalf("get volume failed: %+v", err)
	}
	fmt.Printf("current volume: %d\n", vol)
}
