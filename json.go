package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

var (
	jsonFilePath = "post.json"
	createPostJsonFilePath = "create_post.json"
	decoderPostJsonFilePath = "create_post_with_decoder.json"
)

type Post struct {
	Id int `json:"id"`
	Content string `json:"content"`
	Author Author `json:"author"`
	Comments []Comment `json:"comments"`
}

type Author struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	Id int `json:"id"`
	Content string `json:"content"`
	Author string `json:"author"`
}

func JsonUnmarshal() {
	jsonFile, err := os.Open(jsonFilePath)
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading JSON data:", err)
		return
	}

	var post Post
	err = json.Unmarshal(jsonData, &post)
	if err != nil {
		fmt.Println("Error unmarshal JSON data:", err)
		return
	}
	fmt.Println(post)
}

func JsonDecoder() (post Post, err error) {
	jsonFile, err := os.Open(jsonFilePath)
	if err != nil {
		//fmt.Println("Error open JSON file:", err)
		return
	}
	defer jsonFile.Close()

	decoder := json.NewDecoder(jsonFile)
	for { // 遍历JSON文件, 直到遇到EOF为止
		err = decoder.Decode(&post) // 将JSON数据解析至结构体
		if err == io.EOF {
			break
		}
		if err != nil {
			//fmt.Println("Error decoding JSON:", err)
			return
		}
		//fmt.Println(post)
	}
	return
}

func JsonCreate() {
	post := Post{
		Id: 1,
		Content: "Hello World",
		Author: Author {
			Id: 2,
			Name: "Sau Sheong",
		},
		Comments: []Comment{
			Comment{
				Id: 3,
				Content: "Have a great day!",
				Author: "Adam",
			},
			Comment{
				Id: 4,
				Content: "How are you today?",
				Author: "Betty",
			},
		},
	}

	output, err := json.MarshalIndent(&post, "", "\t\t")
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}

	err = ioutil.WriteFile(createPostJsonFilePath, output, 0644)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}
}

func JsonCreateWithDecoder() {
	post := Post{
		Id: 1,
		Content: "Hello again",
		Author: Author{
			Id: 2,
			Name: "Henry",
		},
		Comments: []Comment{
			Comment{
				Id: 3,
				Content: "Have a great day!",
				Author: "Adam",
			},
			Comment{
				Id: 4,
				Content: "How are you today?",
				Author: "Betty",
			},
		},
	}
	jsonFile, err := os.Create(decoderPostJsonFilePath)
	if err != nil {
		fmt.Println("Error create json decoder file:", err)
		return
	}
	encoder := json.NewEncoder(jsonFile)
	encoder.SetIndent("", "\t")
	err = encoder.Encode(&post)
	if err != nil {
		fmt.Println("Error encoder JSON to file:", err)
		return
	}
}