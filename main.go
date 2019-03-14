package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"time"
)

const baseURL = "https://www.swedishbankers.se"

func main() {
	p := fmt.Println

	fetchClient := http.Client{
		Timeout: time.Second * 60,
	}

	req, err := http.NewRequest(http.MethodGet, baseURL+"/fraagor-vi-arbetar-med/clearingnummer/clearingnummer", nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := fetchClient.Do(req)
	if err != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	regexpPattern, parseErr := regexp.Compile("<a[^>]+href=\"([^\"]+)\"[^>]*><span>Clearingnummer - ")
	if parseErr != nil {
		log.Fatal(parseErr)
	}

	regexpResult := regexpPattern.FindSubmatch(body)
	if len(regexpResult) > 0 {
		pdfPath := string(regexpResult[1])
		pdfReq, pdfReqErr := http.NewRequest(http.MethodGet, baseURL+pdfPath, nil)
		if pdfReqErr != nil {
			log.Fatal(pdfReqErr)
		}

		pdfRes, pdfGetErr := fetchClient.Do(pdfReq)
		if pdfGetErr != nil {
			log.Fatal(pdfGetErr)
		}

		pdfBody, pdfReadErr := ioutil.ReadAll(pdfRes.Body)
		if pdfReadErr != nil {
			log.Fatal(pdfReadErr)
		}
		p(string(pdfBody))
	}
}
