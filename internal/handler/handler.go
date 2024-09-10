package handler

import (
	"fmt"
	"io"
	"net/http"

	"github.com/yourusername/lassie-retriever/internal/lassie"
)

func RetrieveHandler(lassieClient *lassie.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cid := r.URL.Query().Get("cid")
		if cid == "" {
			http.Error(w, "Missing CID parameter", http.StatusBadRequest)
			return
		}

		data, err := lassieClient.Retrieve(cid)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to retrieve data: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/octet-stream")
		io.Copy(w, data)
	}
}
