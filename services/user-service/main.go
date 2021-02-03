// Code generated by candi v1.3.1.

package main

import (
	"fmt"
	"runtime/debug"

	"pkg.agungdp.dev/candi/codebase/app"
	"pkg.agungdp.dev/candi/config"

	service "monorepo/services/user-service/internal"
)

const serviceName = "user-service"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("\x1b[31;1mFailed to start %s service: %v\x1b[0m\n", serviceName, r)
			fmt.Printf("Stack trace: \n%s\n", debug.Stack())
		}
	}()

	cfg := config.Init(serviceName)
	defer cfg.Exit()

	srv := service.NewService(serviceName, cfg)
	app.New(srv).Run()
}
