package main

import (
	"compress/flate"
	"fmt"
	"io"
	"os"

	"./pdf"
)

// const baseURL = "https://www.swedishbankers.se"

func main() {
	p := fmt.Println

	// fetchClient := http.Client{
	// 	Timeout: time.Second * 60,
	// }

	// req, err := http.NewRequest(http.MethodGet, baseURL+"/fraagor-vi-arbetar-med/clearingnummer/clearingnummer", nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// res, getErr := fetchClient.Do(req)
	// if err != nil {
	// 	log.Fatal(getErr)
	// }

	// body, readErr := ioutil.ReadAll(res.Body)
	// if readErr != nil {
	// 	log.Fatal(readErr)
	// }

	// regexpPattern, parseErr := regexp.Compile("<a[^>]+href=\"([^\"]+)\"[^>]*><span>Clearingnummer - ")
	// if parseErr != nil {
	// 	log.Fatal(parseErr)
	// }

	// regexpResult := regexpPattern.FindSubmatch(body)
	// if len(regexpResult) > 0 {
	// pdfPath := string(regexpResult[1])
	// p(pdfPath)
	// pdfReq, pdfReqErr := http.NewRequest(http.MethodGet, baseURL+pdfPath, nil)
	// if pdfReqErr != nil {
	// 	log.Fatal(pdfReqErr)
	// }

	// pdfRes, pdfGetErr := fetchClient.Do(pdfReq)
	// if pdfGetErr != nil {
	// 	log.Fatal(pdfGetErr)
	// }

	// pdfBody, pdfReadErr := ioutil.ReadAll(pdfRes.Body)
	// if pdfReadErr != nil {
	// 	log.Fatal(pdfReadErr)
	// }

	// document := pdf.Parse(pdfBody)
	// p(document.Version)

	// document = pdf.ParseString(pdfBody)

	// Free memory.
	// document = nil
	// }

	// pdfBody, err := ioutil.ReadFile("x.pdf")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	document := pdf.ParseFile("x.pdf")
	p(document)
	p("Hello")
}

func uncompress() {
	inputFile, err := os.Open("compressed")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer inputFile.Close()

	outputFile, err := os.Create("file.txt.decompressed")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer outputFile.Close()

	flateReader := flate.NewReader(inputFile)

	defer flateReader.Close()
	io.Copy(outputFile, flateReader)
}
