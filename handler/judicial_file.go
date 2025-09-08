package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/IsraelTeo/api-registry-court-files-MVP/dto"
	"github.com/IsraelTeo/api-registry-court-files-MVP/model"
	"github.com/IsraelTeo/api-registry-court-files-MVP/service"
	"github.com/gorilla/mux"
)

type JudicialFileHandler struct {
	service service.JudicialFileService
}

func NewJudicialFileHandler(service service.JudicialFileService) *JudicialFileHandler {
	return &JudicialFileHandler{service: service}
}

func (h *JudicialFileHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "invalid ID to integer conversion", http.StatusBadRequest)
		return
	}

	file, err := h.service.GetByID(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(file)
}

func (h *JudicialFileHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	files, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(files)
}

func (h *JudicialFileHandler) Create(w http.ResponseWriter, r *http.Request) {
	var fileDTO dto.JudicialFile
	if err := json.NewDecoder(r.Body).Decode(&fileDTO); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	createdFile, err := h.service.Create(&fileDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdFile)
}

func (h *JudicialFileHandler) Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}

	var file model.JudicialFile
	if err := json.NewDecoder(r.Body).Decode(&file); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	updatedFile, err := h.service.Update(uint(id), &file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(updatedFile)
}

func (h *JudicialFileHandler) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *JudicialFileHandler) AddPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fileID, _ := strconv.Atoi(params["id"])
	personID, _ := strconv.Atoi(params["personId"])

	if err := h.service.AddPerson(uint(fileID), uint(personID)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *JudicialFileHandler) AddLawyer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fileID, _ := strconv.Atoi(params["id"])
	lawyerID, _ := strconv.Atoi(params["lawyerId"])

	if err := h.service.AddLawyer(uint(fileID), uint(lawyerID)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
