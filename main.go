package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("ready to go");
	res, err := http.Get("https://api.weatherapi.com/v1/current.json?key=23090c339d98438e93790043240504&q=bengaluru&aqi=yes");

	if(err != nil) {
		panic(err)
	}
	defer res.Body.Close()

	if(res.StatusCode != 200) {
		panic("weather api not available")
	}

	body, err := io.ReadAll(res.Body)

	if(err != nil) {
		panic(err)
	}

	fmt.Println(string(body));
}