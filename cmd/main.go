package main

import (
	"log"
	"time"

	"github.com/sapvs/keeper"
	"github.com/sapvs/keeper/sample"
)

func main() {
	keeper := keeper.NewKeeper(keeper.WithInterval(1 * time.Second))
	err := keeper.Start()
	if err != nil {
		log.Fatalln("could not start keeper")
	}
	dummyResource := &sample.DummyResource{}
	err = keeper.AddResource(dummyResource)
	if err != nil {
		log.Fatalln("could not add resource")
	}

	time.Sleep(10 * time.Second)
}
