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
/**
 *
 * @api {GET} /product/:type Get All Documents from Index's Type
 * @apiName GetAllType
 * @apiGroup Product
 * @apiVersion  0.0.0
 *
 * @apiSuccess (200) {type} name description
 *
 * @apiParamExample  {type} Request-Example:
 * {
 *     property : value
 * }
 *
 *
 * @apiSuccessExample {type} Success-Response:
 * {
 *     property : value
 * }
 *
 *
 */
func GetAllType(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	start := time.Now()

	indexType := params.ByName("type")
	log.Println("indexType :", indexType)

	data, err := getAllType(indexType)
	if err != nil {
		log.Println("error :", err.Error())
		displayInfo(start, req, http.StatusInternalServerError)
		return
	}

	middleware.RenderHTTPResponse(w, data, "")
	displayInfo(start, req, http.StatusOK)
}
