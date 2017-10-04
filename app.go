package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"github.com/irvana/hmactraining/handler"
)

func main()  {
	router := httprouter.New()

	router.GET("/get/data", handler.HandlerGetData)
	router.GET("/get/gm/stat", handler.HandlerGetGMStat)

	http.ListenAndServe(":9000", router)
}