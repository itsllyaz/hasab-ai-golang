package hasabai

import (
	"fmt"
	"net/http"
)

// ebxLBWjhW3tmYUtrnZoRDB8P74kFR9xL

var data =  map[string]any{
	"file": "test1.mp3", 
	"translate": true, 
	"language": "amh", 
}
func GetTranscription(){
	fmt.Println("bobobo")	
}


func PostRequest(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "boo for life")
}
