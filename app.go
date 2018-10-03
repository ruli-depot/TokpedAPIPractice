package main

import (
	"log"
	"net/http"
	elastic "practiceTokpedElastic/config"

	product "practiceTokpedElastic/product"

	"github.com/julienschmidt/httprouter"
	"github.com/tokopedia/panics"
)

func main() {
	elastic.Init()

	// Set routes
	router := httprouter.New()

	router.GET("/dummy", panics.CaptureHTTPRouterHandler(product.GetAllProduct))
	router.GET("/dummy/:type", panics.CaptureHTTPRouterHandler(product.GetAllType))

	log.Println("listening to localhost:12345")
	log.Fatal(http.ListenAndServe(":12345", router))
}
