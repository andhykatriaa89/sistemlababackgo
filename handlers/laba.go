package handlers

import (
	"encoding/json"
	"net/http"
	"sistem-laba/models"
	"sistem-laba/store"
	"strconv"
	"time"
)

type HitungRequest struct {
	Pendapatan float64  `json:"pendapatan"`
	Modal      float64  `json:"modal"`
	Catatan    string   `json:"catatan"`
	Items      []string `json:"items"`
}

func HitungLaba(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req HitungRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	labaBersih := req.Pendapatan - req.Modal

	margin := 0.0
	if req.Pendapatan > 0 {
		margin = (labaBersih / req.Pendapatan) * 100
	}

	status := "Untung"
	if labaBersih < 0 {
		status = "Rugi"
	} else if labaBersih == 0 {
		status = "Impas"
	}

	transaksi := models.Transaksi{
		Pendapatan: req.Pendapatan,
		Modal:      req.Modal,
		LabaBersih: labaBersih,
		Margin:     margin,
		Status:     status,
		Catatan:    req.Catatan,
		Items:      req.Items,
		CreatedAt:  time.Now(),
	}

	saved, err := store.CreateTransaksi(transaksi)
	if err != nil {
		http.Error(w, "Gagal menyimpan", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(saved)
}

func GetTransaksi(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	data := store.GetAllTransaksi()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func DeleteTransaksi(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := store.DeleteTransaksi(uint(id)); err != nil {
		http.Error(w, "Gagal menghapus", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Berhasil dihapus"})
}
