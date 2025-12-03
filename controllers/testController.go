package controllers

import (
	http "net/http"
	io "io"
	fmt "fmt"
	json "encoding/json"
	models "meenu_be/models"
)

func TestFunc( w http.ResponseWriter, _ *http.Request ) {

	resp := &models.Response{
		Msg: "some string",
		Stat: 1,
		Data: "ayo",
	}


	fmt.Println("got in test")
	jsonTxt, err := json.Marshal(resp);
	if err != nil {
		fmt.Println("error marshalling")
		return;
	}
	fmt.Println(string(jsonTxt))
	io.WriteString(w, string(jsonTxt))
}

