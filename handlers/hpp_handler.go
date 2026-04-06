package handlers

import (
	"encoding/json"
	"net/http"
	"sistem-laba/models"
	"sistem-laba/store"
	"strconv"
)

// Bahan Baku Handlers
func GetBahanBaku(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	bahan := store.GetAllBahanBaku()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bahan)
}

func SaveBahanBaku(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost && r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var b models.BahanBaku
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var saved models.BahanBaku
	var err error
	if b.ID == 0 {
		saved, err = store.CreateBahanBaku(b)
	} else {
		saved, err = store.UpdateBahanBaku(b)
	}

	if err != nil {
		http.Error(w, "Gagal menyimpan bahan baku", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(saved)
}

func DeleteBahanBaku(w http.ResponseWriter, r *http.Request) {
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

	if err := store.DeleteBahanBaku(uint(id)); err != nil {
		http.Error(w, "Gagal menghapus bahan baku", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Bahan baku berhasil dihapus"})
}

// Produk & Resep Handlers
func GetProduk(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	produk := store.GetAllProduk()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(produk)
}

func CreateProduk(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var p models.Produk
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	saved, err := store.CreateProduk(p)
	if err != nil {
		http.Error(w, "Gagal menyimpan produk dan resep", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(saved)
}

func DeleteProduk(w http.ResponseWriter, r *http.Request) {
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

	if err := store.DeleteProduk(uint(id)); err != nil {
		http.Error(w, "Gagal menghapus produk", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Produk berhasil dihapus"})
}
