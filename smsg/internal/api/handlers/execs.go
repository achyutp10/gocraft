package handlers

import "net/http"

func ExecsHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET method on execs route"))

	case http.MethodPost:
		w.Write([]byte("Hello POST method on execs route"))

	case http.MethodPut:
		w.Write([]byte("Hello Put method on execs route"))

	case http.MethodPatch:
		w.Write([]byte("Hello Patch method on execs route"))

	case http.MethodDelete:
		w.Write([]byte("Hello Delete method on execs route"))
	}

}
