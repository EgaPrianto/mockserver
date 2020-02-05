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
	addr := ":1234"
	if len(os.Args) > 1 {
		if os.Args[1] != "" {
			addr = ":" + os.Args[1]
		}
	}

	rResp, err := os.Open("response.yml")
	if err.Error() == "no such file or directory" {
		fmt.Printf("Please add response.yml file in current working directory!")
	} else {
		gojebug.CheckErr(err)
	}
	gojebug.CheckErr(err)
	rRespByte, err := ioutil.ReadAll(rResp)
	gojebug.CheckErr(err)
	handler := ResponseCustom{}

	err = yaml.Unmarshal(rRespByte, &handler)
	gojebug.CheckErr(err)

	if len(os.Args) > 2 {
		if os.Args[2] != "" {
			responsebody := os.Args[2]

			rBody, err := os.Open(responsebody)
			gojebug.CheckErr(err)
			rBodyByte, err := ioutil.ReadAll(rBody)
			gojebug.CheckErr(err)

			handler.Body = string(rBodyByte)
		}
	}

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
	Status int                 `yaml:"Status"`
	Body   string              `yaml:"Body"`
	Header map[string][]string `yaml:"Header"`
}

func (rs ResponseCustom) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	uv := w.Header()
	gojebug.PrintRequest(*r)

	w.WriteHeader(rs.Status)
	for k, values := range rs.Header {
		for _, v := range values {
			uv.Set(k, v)
		}
	}
	_, err := w.Write([]byte(rs.Body))

	gojebug.CheckErr(err)
}
