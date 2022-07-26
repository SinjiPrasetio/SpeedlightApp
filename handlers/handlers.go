package handlers

import (
	"App/data"
	"net/http"

	"github.com/SinjiPrasetio/speedlight"
)

type Handlers struct {
	App    *speedlight.Speedlight
	Models data.Models
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	err := h.render(w, r, "home", nil, nil)
	if err != nil {
		h.App.ErrorLog.Println("Error rendering : ", err)
	}
}
