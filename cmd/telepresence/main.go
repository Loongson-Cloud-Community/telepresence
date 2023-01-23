package main

import (
	"context"
	"github.com/telepresenceio/telepresence/v2/pkg/client/userd/trafficmgr"

	"github.com/telepresenceio/telepresence/v2/pkg/client/cli"
)

func main() {
	cli.Main(cli.InitContext(context.Background(), trafficmgr.NewSessionBuilder()))
}
