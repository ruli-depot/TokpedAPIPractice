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

	// product
	router.ServeFiles("/product-doc/*filepath", http.Dir("product/doc/"))
	router.GET("/product", panics.CaptureHTTPRouterHandler(product.GetAllProduct))
	router.GET("/product/:type", panics.CaptureHTTPRouterHandler(product.GetAllType))

	// test-index
	router.ServeFiles("/test-doc/*filepath", http.Dir("testIndex/doc/"))

	log.Println("listening to localhost:12345")
	log.Fatal(http.ListenAndServe(":12345", router))
}
