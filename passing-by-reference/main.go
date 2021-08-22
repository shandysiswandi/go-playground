package main

import (
	"encoding/json"
	"log"
)

type someStruct struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

type someStruct2 struct {
	Data   interface{} `json:"data"`
	Stacks []string    `json:"stacks"`
}

func main() {
	var aData = []byte(`{"error":false,"message":"owh"}`)
	var a someStruct

	err := someFunc(aData, &a)
	if err != nil {
		log.Fatalln("oops", err)
	}
	log.Println(a)

	var a2Data1 = []byte(`{"data":[],"stacks":["satu"]}`)
	var a2Data2 = []byte(`{"data":false,"stacks":["satu"]}`)
	var a2Data3 = []byte(`{"data":"ada","stacks":["satu"]}`)
	var a2 someStruct2

	err = someFunc(a2Data1, &a2)
	if err != nil {
		log.Fatalln("oops", err)
	}
	log.Println(a2)

	err = someFunc(a2Data2, &a2)
	if err != nil {
		log.Fatalln("oops", err)
	}
	log.Println(a2)

	err = someFunc(a2Data3, &a2)
	if err != nil {
		log.Fatalln("oops", err)
	}
	log.Println(a2)

	log.Println("success..")
}

func someFunc(data []byte, i interface{}) error {
	return json.Unmarshal(data, i)
}
