package models

import (
	json "encoding/json"
)

type Response struct {

	Msg string
	Stat int 
	Data any 
}

func MakeResp(msg string, st int, dt any) *string {

	resp := &Response{
		Msg: msg, Stat: st, Data: dt,
	}

	respTxt, err := json.Marshal(resp);
	if err != nil {
		panic(err)
		return nil
	}

	validText := string(respTxt)

	return &validText
}
