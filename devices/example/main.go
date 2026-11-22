// this isn't supposed to be a real device, just testing things

package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync/atomic"
)

var on atomic.Bool

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RemoteAddr, r.Method, r.URL.RequestURI(), r.UserAgent())
		if r.URL.Path != "/state" {
			w.Write([]byte("<h1>example device</h1>"))
			return
		}
		if r.Method == http.MethodPost {
			var req struct{ State string }
			json.NewDecoder(r.Body).Decode(&req)
			if req.State == "on" || req.State == "off" {
			on.Store(req.State == "on")
			}
		}
		state := "off"
		if on.Load() {
			state = "on"
		}
		json.NewEncoder(w).Encode(map[string]string{"state": state})
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}