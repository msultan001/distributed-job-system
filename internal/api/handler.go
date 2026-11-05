package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/msultan001/distributed-job-system/internal/job"
)

type JoBHandler struct{}

func NewHandler() *JoBHandler {
	return &JoBHandler{}
}

func (h *JoBHandler) CreateJobPost(w http.ResponseWriter, r *http.Request) {
	var jo job.JobDTO
	if err := json.NewDecoder(r.Body).Decode(&jo); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	newJob := job.Job{
		ID:      10,
		Type:    jo.Type,
		Payload: jo.Payload,
	}
	fmt.Println(newJob)
	// w.Write([]byte("that was real quick"))
}

func (h *JoBHandler) GetProcessedJob(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	fmt.Println(id)

}
