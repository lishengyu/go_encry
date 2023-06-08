package http

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	. "main/decry"
	"mime/multipart"
	"net/http"
	"time"
)

var HttpServerMux *http.ServeMux

func HttpServerInit() {
	HttpServerMux = http.NewServeMux()
}

func HttpAddHandle(url string, handle func(w http.ResponseWriter, r *http.Request)) {
	if url == "" {
		fmt.Printf("url is nil\n")
		return
	}

	HttpServerMux.HandleFunc(url, handle)
	fmt.Printf(">>>>>>>>add URL:%s\n", url)
}

func HttpServerRoutine() {
	addrinfo := "127.0.0.1:8082"

	server := &http.Server{
		Addr:         addrinfo,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		Handler:      HttpServerMux,
	}

	fmt.Printf(">>>>Start HttpSeverRoutine:%s\n", addrinfo)
	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("err:%v\n", err)
	}
}

func RequestTest() {
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	msg := "this is a test.\n"

	part, err := writer.CreateFormFile("file", "parameter")
	if err != nil {
		fmt.Printf("create form file failed:%v\n", err)
		return
	}
	_, err = io.WriteString(part, msg)
	if err != nil {
		fmt.Printf("io copy failed:%v\n", err)
		return
	}
	writer.Close()

	urlPrefix := "http://127.0.0.1:8082/test"
	req, err := http.NewRequest("POST", urlPrefix, &buffer)
	if err != nil {
		fmt.Printf("Post failed:%v\n", err)
		return
	}

	req.Header.Set("Content-Type", "application/json;charset=utf-8")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Client Do failed:%v\n", err)
	}

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Read failed:%v\n", err)
		return
	}
	resp.Body.Close()

	fmt.Printf("http url send ok:%s\n", urlPrefix)
}

func HandleTest(w http.ResponseWriter, r *http.Request) {
	key := "2f0df6e04b1c2e60e29221e3659f7bf1"
	req, err := ioutil.ReadAll(r.Body)
	fmt.Printf("req(string):%s\nreq(byte):%v\n", string(req), req)
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}
	deBytes := AesDecrypt(req, []byte(key), "CBC")
	fmt.Printf("Decry:%s\n", deBytes)
}

func ClientRequestTest() {
	for {
		time.Sleep(5 * time.Second)
		RequestTest()
	}
}
