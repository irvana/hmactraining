package handler

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"fmt"
	"log"
	"io/ioutil"
)

const (
	TOME_BASE_URL = "https://tome.tokopedia.com"
)

func HandlerGetData(writer http.ResponseWriter, req *http.Request, p httprouter.Params){
	urlPath := "/v1/product/get_summary"
	url := fmt.Sprintf("%s%s?product_id=196704028", TOME_BASE_URL, urlPath)

	client := &http.Client{}
	req, err :=http.NewRequest(http.MethodGet, url, nil)

	resp, err := client.Do(req)

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println("")
	}

	writer.Write(body)
}
