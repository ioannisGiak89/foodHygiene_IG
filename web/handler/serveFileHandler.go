package handler

import (
	"net/http"
)

// FileServerHandler makes assets available through http
type FileServerHandler struct{}

// NewFileServerHandler creates a new FileServerHandler
func NewFileServerHandler() *FileServerHandler {
	return &FileServerHandler{}
}

// Init the FileServerHandler
func (h *FileServerHandler) Init() {
	// For js
	fs := http.FileServer(http.Dir("web/assets/js"))
	http.Handle("/assets/js/", http.StripPrefix("/assets/js/", fs))

	// For css
	fs = http.FileServer(http.Dir("web/assets/css"))
	http.Handle("/assets/css/", http.StripPrefix("/assets/css/", fs))
}
