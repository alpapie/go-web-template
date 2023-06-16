package helper

import (
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter,s string, Data interface{}) error{
	page , err :=template.ParseFiles("template/pages/"+s+".page.tmpl")
	if err!=nil{
		return err
	}
	return page.Execute(w,Data)
}

func ErrorPage(w http.ResponseWriter, r *http.Request , s string) error{
	page , err :=template.ParseFiles()
	if err!=nil{
		return err
	}
	switch s {
	case  "405":
		w.WriteHeader(http.StatusMethodNotAllowed)
	case  "404":
		w.WriteHeader(http.StatusNotFound)
	case  "500":
		w.WriteHeader(http.StatusInternalServerError)
	}
	return page.Execute(w,nil)
}