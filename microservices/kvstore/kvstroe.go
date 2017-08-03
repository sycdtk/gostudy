package main

import (
	"fmt"
	"net/http"
	"net/url"
	"sync"
)

var kvStore map[string]string
var mu *sync.RWMutex

func main() {
	kvStore = make(map[string]string)
	mu = new(sync.RWMutex)

	http.HandleFunc("/get", get)
	http.HandleFunc("/set", set)
	http.HandleFunc("/remove", remove)
	http.HandleFunc("/list", list)

	http.ListenAndServe(":9901", nil)
}

func get(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		values, err := url.ParseQuery(r.URL.RawQuery)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Error:", err)
			return
		}

		if len(values.Get("key")) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Error:", "wrong input key.")
			return
		}

		mu.Lock()
		if value, ok := kvStore[values.Get("key")]; ok {
			fmt.Fprint(w, value)
			return
		} else {
			fmt.Fprint(w, "Empty")
			return
		}
		mu.Unlock()

	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Error:Only GET accepted.")
	}
}

func set(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		values, err := url.ParseQuery(r.URL.RawQuery)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Error:", err)
			return
		}

		if len(values.Get("key")) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Error:", "wrong input key.")
			return
		}

		mu.Lock()
		kvStore[values.Get("key")] = values.Get("value")
		mu.Unlock()

		fmt.Fprint(w, "Success.")

	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Error:Only POST accepted.")
	}
}

func remove(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		values, err := url.ParseQuery(r.URL.RawQuery)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Error:", err)
			return
		}

		if len(values.Get("key")) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Error:", "wrong input key.")
			return
		}

		mu.Lock()
		delete(kvStore, values.Get("key"))
		mu.Unlock()

		fmt.Fprint(w, "Success.")

	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Error:Only POST accepted.")
	}
}

func list(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		mu.RLock()
		for k, v := range kvStore {
			fmt.Fprint(w, k, ":", v)
		}
		mu.RUnlock()
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Error:Only GET accepted.")
	}
}
