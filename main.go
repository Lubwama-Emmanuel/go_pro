package main

import (
	"errors"
	"log"
)

// import (
// 	"encoding/json"
// 	"log"
// 	// "github.com/Lubwama-Emmanuel/test1/helpers"
// )

// // var numPool = 1000

// // func CalculateValue(intChan chan int) {
// // 	num := helpers.RandomNumber(numPool)
// // 	intChan <- num
// // }

// // func main() {
// // 	intChan := make(chan int)
// // 	defer close(intChan)

// // 	go CalculateValue(intChan)

// // 	number := <-intChan
// // 	log.Println(number)
// // }

// type Person struct {
// 	FirstName string `json:"name"`
// 	Age       int    `json:"age"`
// }

// func main() {
// 	myJson := `
// 	[
// 		{
// 			"name": "Rex",
// 			"age": 24
// 		},
// 		{
// 			"name": "Emmanuel",
// 			"age": 12
// 		}
// 	]`

// 	var unmarshalled []Person

// 	err := json.Unmarshal([]byte(myJson), &unmarshalled)
// 	if err != nil {
// 		log.Println("An Error occured", err)

// 	}

// 	// log.Printf("unmarshalled %v", unmarshalled)

// 	var mySlice []Person

// 	var m1 Person
// 	m1.FirstName = "Lubwama"
// 	m1.Age = 11

// 	mySlice = append(mySlice, m1)

// 	var m2 Person
// 	m2.FirstName = "Dalton"
// 	m2.Age = 10

// 	mySlice = append(mySlice, m2)

// 	newJson, err := json.MarshalIndent(mySlice, "", "    ")
// 	if err != nil {
// 		log.Println("An Error occured", err)
// 	}

// 	log.Println(string(newJson))

// }

func main() {
	result, err := divide(100.0, 10.0)
	if err != nil {
		log.Println("An Error occured", err)
	}
	log.Println(result)
}

func divide(x, y float32) (float32, error) {
	var result float32

	if y == 0 {
		return result, errors.New("cannot divide by zero")
	}

	result = x / y
	return result, nil
}
