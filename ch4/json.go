/*
Go语言中将结构体slice转为JSON的过程叫编组（marshaling）。通过json.Marshal函数完成
逆操作是json.UnMarshal

另有基于流的json.Encoder和json.Decoder
*/
package main

import (
	"fmt"
	"encoding/json"
	"log"
)

type Movie struct {
	Title string
	Year int `json:"released"`
	Color bool `json:"color,omitempty"` // 空值或零值时不生成到json中
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
}

func main() {
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	// 缩进格式
	data, err = json.MarshalIndent(movies, "", "    ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	// 只解析一部分成员
	var titles []struct{
		Title string
		Year int
	}
	if err = json.Unmarshal(data, &titles); err != nil{
		log.Fatalf("Json unmarshaling failed: %s", err)
	}
	fmt.Println(titles)
	fmt.Println("test汉字")
}
