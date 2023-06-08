package main

import (
	. "main/http"
)

func main() {
	HttpServerInit()
	HttpAddHandle("/test", HandleTest)

	go ClientRequestTest()

	HttpServerRoutine()
}
