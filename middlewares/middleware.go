package middlewares

import (
	"net/http"
	"strings"
)

func CheckRequest(r *http.Request, path, methode string) (bool, string) {
	if strings.ToLower(r.Method) == methode && r.URL.Path == path{
		return true, ""
	} else if strings.ToLower(r.Method) != methode{
		return false, "405"
	} else {
		return false, "404"
	}
}
