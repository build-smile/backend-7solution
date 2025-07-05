package main

import (
	"context"
	"github.com/build-smile/backend-7solution/httpserve"
	"github.com/build-smile/backend-7solution/infrastructure"
	"github.com/build-smile/backend-7solution/tasks"
)

func main() {

	infrastructure.InitConfig()
	infrastructure.ConnectDB()
	// Set up context with cancel for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	tasks.StartUserCountLogger(ctx, infrastructure.MongoDB)
	httpserve.Run()

}
