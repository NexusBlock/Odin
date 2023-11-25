package main

import (
	"context"

	"github.com/nexusblock/heimdall/cmd/odind/service"
)

func main() {
	service.NewHeimdallService(context.Background(), nil)
}
