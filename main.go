package main

import (
	"flag"
	"fmt"
	"regin/app/cmd/server"
	"sync"
)

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
// func Helloworld(g *gin.Context) {
// 	g.JSON(http.StatusOK, "helloworld")
// }

// func HelloIndex(g *gin.Context) {

// 	var msgs []string

// 	msgs = append(msgs, "Foobar")

// 	g.JSON(http.StatusOK, gin.H{
// 		"message":  "this is V2",
// 		"messages": msgs,
// 	})
// }

var start *bool

func parseArguments() {
	start = flag.Bool("serve", false, "Run start the server")

	flag.Parse()
}

// `Hot Reloading` - to run this server in dev mode, install `pilu/fresh` package
// and run `fresh -- --serve` in the command line`
func main() {
	parseArguments()

	wg := sync.WaitGroup{}

	if *start {
		wg.Add(1)
		go server.RunServer(&wg)
	}

	fmt.Println("Waiting for everyone to be done")

	wg.Wait()
}
