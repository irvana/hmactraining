package handler

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"fmt"
	"log"
	"io/ioutil"
	"encoding/json"
	hmacUtil "github.com/irvana/hmactraining/auth"
)

type (
	Product struct{
		DataProduct []DataProduct `json:"data"`
	}
	DataProduct struct {
		ProductId int `json:"product_id"`
	}
)


const (
	TOME_BASE_URL = "https://tome.tokopedia.com"
	GM_BASE_URL = "https://slicer-staging.tokopedia.com"
)

func HandlerGetGMStat(writer http.ResponseWriter, req *http.Request, p httprouter.Params){
	urlPath:="/gmstat/cube/gm_prod_graph/aggregate"
	content := "drilldown=date&cut=shop_id:67726|date:20160427-20160503&order=date:asc"
	url := fmt.Sprintf("%s%s?%s",GM_BASE_URL, urlPath, content)
	contentType := ""

	auth, contentMd5, time := hmacUtil.GetHMACKey(url, http.MethodGet, content, contentType)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("Authorization", auth)
	req.Header.Set("X-Date", time)
	req.Header.Set("Postman-Date", time)
	req.Header.Set("Content-MD5", contentMd5)
	req.Header.Set("X-Method", http.MethodGet)

	log.Println(req)
	resp, err := client.Do(req)

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
	}
	writer.Write(body)
}

func HandlerGetData(writer http.ResponseWriter, req *http.Request, p httprouter.Params){
	urlPath := "/v1/product/get_summary"
	url := fmt.Sprintf("%s%s?product_id=196704028",
		TOME_BASE_URL, urlPath)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)

	resp, err := client.Do(req)

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println("")
	}

	var prod Product
	json.Unmarshal(body, &prod)
	writer.Write([]byte(fmt.Sprintf("%+v",prod)))
}


