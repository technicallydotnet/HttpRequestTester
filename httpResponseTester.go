package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func main() {

	var httpType, URL, rawPostBody, urlFile string
	var numberOfTimes, delay int
	var Urls []string

	flag.StringVar(&httpType, "httptype", "get", "Use to set the type of HTTP request")
	flag.StringVar(&URL, "url", "http://golang.com", "The web address to test")
	flag.IntVar(&numberOfTimes, "number", 1, "The web address to test")
	flag.StringVar(&rawPostBody, "postbody", "{data: 'hello world!'}", "data to be sent if a post request is being made")
	flag.IntVar(&delay, "delay", 1, "Time to separate each HTTP request, in milliseconds")
	flag.StringVar(&urlFile, "urls", "na", "List of URLs in a file that will be tested sequentially (separated by commas)")

	flag.Parse()
	println("Performing " + httpType + " request at " + URL + " for " + strconv.Itoa(numberOfTimes) + " times")

	httpType = strings.ToLower(httpType)

	var statusMessages [47]string = [47]string{"ok", "created", "accepted", "no content", "multiple Choice", "moved Permanently", "found", "seeOther", "not Modified", "temporary Redirect", "permenant Redirect", "bad Request", "unauthorized", "forbidden", "not Found", "method Not Allowed", "not Accepted", "proxy Auth Required", "conflict", "gone", "length Required", "precondition Failed", "payload Too Large", "URI Too Long", "Unsupported Media Type", "Range Not Satisfiable", "Im A Teapot", "misdirected Response", "unprocessed Entity", "locked", "failed Dependency", "upgrade Required", "precondition Required", "too Many Requests", "request Header Fields Too Large", "unavailable For Legal Reasons", "interal Server Error", "not Implemented", "bad Gateway", "service Unavailable", "gateway Timeout", "http Version Not Supported", "variant Also Negotiates", "insufficient Storage", "loop Detected", "not Extended", "network Auth Required"}
	var statusMessagesCount [47]int = [47]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	var statusCodes [48]int = [48]int{200, 201, 202, 204, 300, 301, 302, 303, 304, 307, 308, 400, 401, 403, 404, 405, 406, 407, 408, 409, 410, 411, 412, 413, 414, 415, 416, 418, 421, 422, 423, 424, 426, 428, 429, 431, 451, 500, 501, 502, 503, 504, 505, 506, 507, 508, 510, 511}

	if urlFile != "na" {
		byteUrls, err := ioutil.ReadFile(urlFile) // the file is inside the local directory
		if err != nil {
			println("Error - " + err.Error())
		}
		Urls = strings.Split(string(byteUrls), ",")
		for j := 0; j < len(Urls); j++ {
			statusMessagesCount = performHttpRequest(httpType, numberOfTimes, URL, statusMessagesCount, statusCodes, rawPostBody, delay)
			printResults(statusMessagesCount, statusCodes, statusMessages)
			statusMessagesCount = [47]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		}
	} else {
		statusMessagesCount = performHttpRequest(httpType, numberOfTimes, URL, statusMessagesCount, statusCodes, rawPostBody, delay)
		printResults(statusMessagesCount, statusCodes, statusMessages)
	}

}

func addToResponseTally(statusCode int, statusMessagesCount [47]int, statusCodes [48]int) [47]int {
	for i := 0; i < len(statusMessagesCount); i++ {
		if statusCode == statusCodes[i] {
			statusMessagesCount[i]++
		}
	}
	return statusMessagesCount
}

func performHttpRequest(httpType string, numberOfTimes int, URL string, statusMessagesCount [47]int, statusCodes [48]int, rawPostBody string, delay int) [47]int {
	if httpType == "get" {
		for i := 0; i < numberOfTimes; i++ {
			response, err := http.Get(URL)
			if err != nil {
				println(err)
			}
			statusMessagesCount = addToResponseTally(int(response.StatusCode), statusMessagesCount, statusCodes)
			time.Sleep(time.Duration(delay) * time.Millisecond)
		}
	} else if httpType == "post" {
		postBody, _ := json.Marshal(rawPostBody)
		for i := 0; i < numberOfTimes; i++ {
			response, err := http.Post(URL, "application/json", bytes.NewBuffer(postBody))
			if err != nil {
				println(err)
			}
			statusMessagesCount = addToResponseTally(int(response.StatusCode), statusMessagesCount, statusCodes)
			time.Sleep(time.Duration(delay) * time.Millisecond)
		}
	} else {
		println("http type not recognised or supported")
	}
	return statusMessagesCount
}
func printResults(statusMessagesCount [47]int, statusCodes [48]int, statusMessages [47]string) {
	for i := 0; i < len(statusMessagesCount); i++ {
		if statusMessagesCount[i] != 0 {
			var plural string
			if statusMessagesCount[i] > 1 {
				plural = " times"
			} else {
				plural = " time"
			}
			println(strconv.Itoa(statusCodes[i]) + " " + statusMessages[i] + " occurred " + strconv.Itoa(statusMessagesCount[i]) + plural)
		}
	}
}
