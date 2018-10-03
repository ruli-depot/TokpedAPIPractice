package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func RenderHTTPResponse(w http.ResponseWriter, response interface{}, callback string) {
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Println("loc : middleware")
		log.Println("fn : renderHTTPResponse")
		log.Println("error 1 : ", err.Error())
	}

	t := time.Now().Add(15 * time.Minute)
	now := time.Now()

	if callback != "" {
		w.Header().Add("Content-Type", "text/javascript")
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
	}

	w.Header().Set("Cache-Control", "public, max-age=900")
	w.Header().Set("Expires", t.Format(time.RFC1123))
	w.Header().Set("Date", now.Format(time.RFC1123))

	if callback != "" {
		fmt.Fprintf(w, "%s(%s)", callback, jsonResponse)
	} else {
		w.Write(jsonResponse)
	}
}
