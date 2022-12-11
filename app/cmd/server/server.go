package server

import (
	"fmt"
	"regin/docs"
	"regin/providers"
	"regin/routes"
	"sync"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RunServer(wg *sync.WaitGroup) {
	r := gin.Default()

	defer wg.Done()

	docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	builder, err := providers.Init()

	if err != nil {
		panic(fmt.Errorf("unable to boot the service providers %w", err))
	}

	routes.Init(r, builder)

	// v1 := r.Group("/api/v1")
	// {
	// 	v1.GET("/", example.GetApiV1Example)
	// 	// eg := v1.Group("/example")
	// 	// {
	// 	// 	eg.GET("/helloworld", Helloworld)
	// 	// }
	// }

	// routes.Init(r)

	r.Run()
	// fmt.Println("Done running the server...")
}
