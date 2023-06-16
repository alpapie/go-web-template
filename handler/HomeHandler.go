package handler

import (
	"fmt"
	"goupie-tracker/helper"
	"goupie-tracker/middlewares"
	"goupie-tracker/models"
	"net/http"
)

var Url =map[string]string{
	"artist":"https://groupietrackers.herokuapp.com/api/artists",
	"event": "https://groupietrackers.herokuapp.com/api/relation",
	"location":"https://groupietrackers.herokuapp.com/api/locations",
	"relation": "https://groupietrackers.herokuapp.com/api/date/",
}
var Artists []models.Artist
var Events	models.Index


func Index(w http.ResponseWriter, r *http.Request) {
	ok , PageError:= middlewares.CheckRequest(r,"/","get")
	if !ok{
		helper.ErrorPage(w,r,PageError)
		return
	}
	iserr :=GetData()
	if !iserr {
		fmt.Fprintf(w,"good", Artists[0])
		fmt.Fprintf(w,"good", Events.Indexes[0].DatesLocations)
	} else {
		// fmt.Fprintf(w, string("error to get data"),iserr)
		helper.ErrorPage(w,r,"500")
		return
	}
}

func GetData() bool {

	err := helper.GetJson(Url["artist"], &Artists)
	err1 := helper.GetJson(Url["event"], &Events)

	return err !=nil || err1 !=nil
} 