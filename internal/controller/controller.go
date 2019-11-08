package controller

import (
	"fmt"
	"net/http"
	"os"
	"github.com/matassp/cloud-lab/internal/hashes"
	"github.com/gorilla/mux"
)

func Sha256Handler(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	h, _ := hashes.GetHash("Sha256", username)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, h)
}

func GithubUsernameHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, os.Getenv("GITHUB_USERNAME"))
}
