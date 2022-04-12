package error_handler

import "net/http"

func NotFoundError(w http.ResponseWriter) {
	http.Error(w, "Method not supported", http.StatusNotFound)
}
