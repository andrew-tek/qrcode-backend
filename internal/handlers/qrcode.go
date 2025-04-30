package handlers

import (
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/skip2/go-qrcode"
)

type QRRequest struct {
	Data string `json:"data"`
}

type QRResponse struct {
	Image string `json:"image"`
}

func GenerateQRCode(w http.ResponseWriter, r *http.Request) {
	var req QRRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	qrCode, err := qrcode.Encode(req.Data, qrcode.Medium, 256)
	if err != nil {
		http.Error(w, "Failed to generate QR code", http.StatusInternalServerError)
		return
	}

	encoded := base64.StdEncoding.EncodeToString(qrCode)
	response := QRResponse{Image: encoded}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}