package auth

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"time"
	"strings"
)

const (
	KEY  = "web_service_v4"
)

func GetMD5Hash(input string) string{
	hasher := md5.New()
	hasher.Write([]byte(input))

	return hex.EncodeToString(hasher.Sum(nil))
}

func GetHMACKey(urlPath, method, content, contentType string) (string, string, string){

	//read https://wiki.tokopedia.net/Tokopedia_Web_Service_v4_Documentation
	contentMd5 := GetMD5Hash(content)
	currTime := time.Now().Format("Mon, 02 Jan 2006 15:04:05 -0700")

	strJoin := strings.Join([]string{method, contentMd5, contentType, currTime, urlPath}, "\n")

	signature := hmac.New(sha1.New, []byte(KEY))
	signature.Write([]byte(strJoin))

	hmacKey := base64.StdEncoding.EncodeToString(signature.Sum(nil))
	auth := "TKPD Tokopedia:"+hmacKey
	return  auth, contentMd5, currTime
}
