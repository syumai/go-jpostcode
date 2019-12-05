package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/syumai/go-jpostcode"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		postalCode := strings.TrimPrefix(r.URL.Path, "/")
		if postalCode == "favicon.ico" {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintln(w, `{"status": "404", "message": "postal code is not found"}`)
			return
		}

		addr, err := jpostcode.Find(postalCode)
		if err != nil {
			switch err {
			case jpostcode.ErrNotFound:
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprintln(w, `{"status": "404", "message": "postal code is not found"}`)
				return
			case jpostcode.ErrInvalidArgument:
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintln(w, `{"status": "400", "message": "postal code is not valid"}`)
				return
			default:
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintln(w, `{"status": "500", "message": "internal server error"}`)
				return
			}
		}

		addrJSON, err := addr.ToJSON()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, `{"status": "500", "message": "internal server error"}`)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s\n", addrJSON)
	})

	port := "8080"
	fmt.Printf("listening on http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
