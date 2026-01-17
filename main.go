package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type barang struct {
	Id int `json:"id"`
	Nama string `json:"nama"`
}

var data = []barang {
	{Id: 1, Nama: "Handphone"},
	{Id: 2, Nama: "Laptop"},
	{Id: 3, Nama: "Tablet"},
}

func main() {
	http.HandleFunc("/Barang", BarangHandler)
	http.ListenAndServe("/Barang/", BarangByIDHandler)

	fmt.Println("Server berjalan di http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func BarangHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		json.NewEncoder(w).Encode(data)
		return
	}

	if r.Method == http.MethodGet {
		var barang Barang
		json.NewEncoder(w).Encode(data)
		return
	}
}