package product

import (
	"context"
	"encoding/json"
	"log"
	"practiceTokpedElastic/config"

	elastic "gopkg.in/olivere/elastic.v5"
)

func getAllType(indexType string) (map[string]interface{}, error) {

	log.Println("source setup : ", config.CfgElastic)
	errResult := make(map[string]interface{}, 0)
	var result map[string]interface{}

	esURI := elastic.SetURL(config.CfgElastic.Host)

	// open new client to es
	es, err := elastic.NewClient(esURI,
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false),
	)
	if err != nil {
		log.Println("loc : product")
		log.Println("fn : getAllProduct")
		log.Println("error 1 : ", err.Error())
		return errResult, err
	}
	// ------------------------------

	// do some search
	esSearch, err := es.Search().Index("product").Timeout("5s").Do(context.Background()) // with timeout
	// esSearch, err := es.Search().Index("dummy").Do(context.Background())
	if err != nil {
		log.Println("loc : product")
		log.Println("fn : getAllProduct")
		log.Println("error 2 : ", err.Error())
		return errResult, err
	}

	// log.Println("es :", es)
	// log.Println("esSearch :", esSearch)
	byt, err := json.Marshal(esSearch)
	if err != nil {
		log.Println("loc : product")
		log.Println("fn : getAllProduct")
		log.Println("error 3 : ", err.Error())
		return errResult, err
	}

	parsed := make(map[string]interface{})
	err = json.Unmarshal(byt, &parsed)
	if err != nil {
		log.Println("loc : product")
		log.Println("fn : getAllProduct")
		log.Println("error 4 : ", err.Error())
		return errResult, err
	}
	result = parsed

	if len(result) < 1 {
		result = errResult
	}

	return result, nil

}
