package product

import (
	"log"
	"net/http"
	"practiceTokpedElastic/middleware"
	"time"

	"github.com/julienschmidt/httprouter"
)

func displayInfo(watcher time.Time, req *http.Request, status int) {
	timeInterval := time.Since(watcher)
	log.Println(status, "||", req.Method, "|| URI : ", req.RequestURI, " || time interval :", timeInterval)
}

//get all product
func GetAllProduct(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	start := time.Now()

	data, err := getAllProduct()
	if err != nil {
		log.Println("error :", err.Error())
		displayInfo(start, req, http.StatusInternalServerError)
		return
	}

	middleware.RenderHTTPResponse(w, data, "")
	displayInfo(start, req, http.StatusOK)
}

//get all type data
func GetAllType(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	start := time.Now()

	indexType := params.ByName("type")
	log.Println("indexType :", indexType)

	// data, err := getAllType()
	// if err != nil {
	// 	log.Println("error :", err.Error())
	// 	displayInfo(start, req, http.StatusInternalServerError)
	// 	return
	// }

	displayInfo(start, req, http.StatusOK)
}
