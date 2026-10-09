package handler

import "edu-crm/storage"

type Handler struct {
	storage *storage.Storage 
}

func NewHandler(strg *storage.Storage) *Handler {
	return &Handler{
        storage: strg,
	}
}