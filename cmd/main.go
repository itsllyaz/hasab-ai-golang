package main

import (
	"fmt"

	hasabai "github.com/itsllyaz/hasab-ai-golang"
)

var HasabApi = os.Getenv("HASAB-AI-API")

func main() {
	fmt.Println("*** HASAB-AI ***")
	client := hasabai.New(HasabApi)
	

	resp, err := client.TTSRecord(1)
	if err != nil{
		fmt.Println("ERROR: ", err)
		return
	}
	fmt.Println("RECORD: ", *resp)
}
