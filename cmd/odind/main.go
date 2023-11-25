package main

import (
	"context"

	"github.com/maticnetwork/heimdall/cmd/odind/service"
)

func main() {
	service.NewHeimdallService(context.Background(), nil)
}
