package main

import (
	"encoding/base64"
	"fmt"
)

func main()  {
	rawData := base64.StdEncoding.EncodeToString([]byte("RawData"))
	
	fmt.Println(rawData)

	decoded, err := base64.StdEncoding.DecodeString(rawData)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(decoded))
}