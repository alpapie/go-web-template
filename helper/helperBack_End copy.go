package helper

import (
	"encoding/json"
	"net/http"
)

//******************* VERIF IF THE STRING IS AN INT*****************************************************
func IsInt(s string) bool{
	for _, v := range s {
		if v < '0' || v > '9'{
			return false
		}
	}
	return true
}

//***************************** GET JSON DATA ************************************************************
func GetJson(url string, model interface{}) error{
	response , err := http.Get(url)
	if err!=nil{
		return err
	}
	defer response.Body.Close()
	return json.NewDecoder(response.Body).Decode(model)
}

