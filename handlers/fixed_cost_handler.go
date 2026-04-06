package handlers

import (
	"encoding/json"
	"net/http"
	"sistem-laba/models"
	"sistem-laba/store"
	"strconv"
)

func GetBiayaTetap(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	data := store.GetAllBiayaTetap()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func SaveBiayaTetap(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost && r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var b models.BiayaTetap
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var saved models.BiayaTetap
	var err error
	if b.ID == 0 {
		saved, err = store.CreateBiayaTetap(b)
	} else {
		saved, err = store.UpdateBiayaTetap(b)
	}

	if err != nil {
		http.Error(w, "Gagal menyimpan biaya tetap", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(saved)
}

func DeleteBiayaTetap(w http.ResponseWriter, r *http.Request) {
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

	if err := store.DeleteBiayaTetap(uint(id)); err != nil {
		http.Error(w, "Gagal menghapus biaya tetap", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Biaya tetap berhasil dihapus"})
}
