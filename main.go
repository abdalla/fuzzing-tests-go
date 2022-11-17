package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Request struct {
	Values []int `json:"values"`
}

func GetHighest(w http.ResponseWriter, r *http.Request) {
	var req Request

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	high := 0

	for _, v := range req.Values {
		if v > high {
			high = v
		}
	}

	//Producing error
	// if high == 75 {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	w.Write([]byte("Something went wrong"))
	// }

	fmt.Fprintf(w, "%d", high)
}
