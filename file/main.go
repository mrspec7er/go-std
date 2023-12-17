package main

import (
	"fmt"
	"os"
)

func main()  {
	header := "name, age, networth"
	data := "\nRobert, 45, $.25B" + "\nAxe, 30, $.230M"

	createFile("data.csv", header)
	modifyFile("data.csv", data)
}

func createFile(name string, value string)  {
	file, err := os.OpenFile("./file/" + name, os.O_CREATE|os.O_WRONLY, 0777)

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()
	file.WriteString(value)
}

func modifyFile(name string, value string)  {
	file, err := os.OpenFile("./file/" + name, os.O_RDWR|os.O_APPEND, 0777)

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()
	file.WriteString(value)
}