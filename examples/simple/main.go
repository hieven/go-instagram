package main

import (
	"context"
	"fmt"

	instagram "github.com/hieven/go-instagram/src"
	"github.com/hieven/go-instagram/src/config"
)

func main() {
	cnf := &config.Config{
		Username: "USERNAME",
		Password: "PASSWORD",
	}

	ctx := context.Background()

	ig, _ := instagram.New(cnf)
	ig.Login(ctx)

	Aresp, _ := ig.Timeline().Feed(ctx, &instagram.TimelineFeedRequest{})
	fmt.Println(Aresp.Items[0].MediaOrAd)
}
