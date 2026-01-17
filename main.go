package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type barang struct {
	Id   int    `json:"id"`
	Nama string `json:"nama"`
}

var data = []barang{
	{Id: 1, Nama: "Handphone"},
	{Id: 2, Nama: "Laptop"},
	{Id: 3, Nama: "Tablet"},
}

func main() {
	http.HandleFunc("/Barang", BarangHandler)
	http.HandleFunc("/Barang/", BarangByIDHandler)

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
		var Barang barang
		json.NewEncoder(w).Encode(data)
		Barang.Id = len(data) + 1
		data = append(data, Barang)
		json.NewEncoder(w).Encode(Barang)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func BarangByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/Barang/"):]
	id, _ := strconv.Atoi(idStr)

	for i, b := range data {
		if b.Id == id {

			if r.Method == http.MethodGet {
				var update barang
				json.NewDecoder(r.Body).Decode(&update)
				data[i].Nama = update.Nama
				json.NewEncoder(w).Encode(data[i])
				return
			}

			if r.Method == http.MethodDelete {
				data = append(data[:i], data[i+1:]...)
				w.Write([]byte("Data berhasil dihapus"))
				return
			}
		}
	}

	w.WriteHeader(http.StatusNotFound)
}
