package routes

import "net/http"

type Route struct {
	URI                    string
	Method                 string
	Fun                    func(http.ResponseWriter, *http.Request)
	RequiresAuthentication bool
}
