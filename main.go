package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/EgaPrianto/gojebug"
)

func main() {
	addr := ":8095"
	if len(os.Args) > 1 {
		if os.Args[1] != "" {
			addr = ":" + os.Args[1]
		}
	}
	responsebody := "response_body"
	if len(os.Args) > 2 {
		if os.Args[2] != "" {
			responsebody = os.Args[2]
		}
	}
	rHeader, err := os.Open("response_headers.yml")
	gojebug.CheckErr(err)
	rHeaderByte, err := ioutil.ReadAll(rHeader)
	gojebug.CheckErr(err)
	rHeaderMap := make(map[string][]string)

	err = yaml.Unmarshal(rHeaderByte, &rHeaderMap)
	gojebug.CheckErr(err)

	rBody, err := os.Open(responsebody)
	gojebug.CheckErr(err)
	rBodyByte, err := ioutil.ReadAll(rBody)
	gojebug.CheckErr(err)
	handler := ResponseCustom{respBody: rBodyByte}
	s := &http.Server{
		Addr:         addr,
		Handler:      handler,
		ReadTimeout:  310 * time.Second,
		WriteTimeout: 310 * time.Second,
	}
	fmt.Println("Listening on Port " + addr)
	gojebug.CheckErr(s.ListenAndServe())
}

type ResponseCustom struct {
	respCode int
	respBody []byte
}

func (rs ResponseCustom) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	uv := w.Header()
	gojebug.PrintRequest(*r)
	for k, values := range rs.responseHeader() {
		for _, v := range values {
			uv.Set(k, v)
		}
	}
	// w.WriteHeader(404)
	_, err := w.Write(rs.respBody)
	gojebug.CheckErr(err)
}

func (rs ResponseCustom) responseHeader() map[string][]string {

	rHeader, err := os.Open("response_headers.yml")
	gojebug.CheckErr(err)
	rHeaderByte, err := ioutil.ReadAll(rHeader)
	gojebug.CheckErr(err)
	rHeaderMap := make(map[string][]string)

	err = yaml.Unmarshal(rHeaderByte, &rHeaderMap)
	gojebug.CheckErr(err)
	return rHeaderMap
}
