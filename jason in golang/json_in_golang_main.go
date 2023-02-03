package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Something struct {
	Name    string    `json:"name"`
	Married bool      `json:"married"`
	Age     float64   `json:"age"`
	Address []Address `json:"address"`
	Marks   []int     `json:"marks"`
}
type Address struct {
	Street int    `json:"street"`
	City   string `json:"city"`
}

func main() {
	jsonFile, err := os.Open("something.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	jsonByteValues, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	 var something Something
	something := map[string]interface{}{}
	
	

	json.Unmarshal(jsonByteValues, &something) 
	log.Println(something)
	fmt.Print(string(jsonByteValues)) /
	newJsonByteValues, err := json.Marshal(something)
	if err != nil {
	 	log.Fatal(err)
	 }
	 os.WriteFile("some.json",newJsonByteValues)

}
func main4() {
	
	user := map[int]interface{}{}
	user[1] = "name"
	user[2] = true
	 fmt.Println(user[2])
	users := map[string]bool{}
	users["golang"] = true

	value, ok := users["rahul"]
	if ok == false {
		fmt.Println("value not there")
	}

	fmt.Println(users["rahul"])
	fmt.Println("Hi its my branch")

}
