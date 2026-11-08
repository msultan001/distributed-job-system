package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/msultan001/distributed-job-system/internal/job"
)

type JoBHandler struct {
	service *job.JobService
}

func NewHandler(service *job.JobService) *JoBHandler {
	return &JoBHandler{
		service: service,
	}
}

func (h *JoBHandler) CreateJobPost(w http.ResponseWriter, r *http.Request) {
	var jo job.JobDTO
	if err := json.NewDecoder(r.Body).Decode(&jo); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	newJob := job.Job{
		ID:      uuid.New(),
		Type:    jo.Type,
		Payload: jo.Payload,
	}
	if err := h.service.CreateJob(r.Context(), &newJob); err != nil {
		http.Error(w, "could not create a job", http.StatusBadRequest)
		return
	}
}

func (h *JoBHandler) GetProcessedJob(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	fmt.Println(id)

}
