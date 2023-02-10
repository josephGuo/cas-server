package admin

import (
	"encoding/json"
	"net/http"

	"github.com/josephGuo/cas-server/tools"
	"github.com/josephGuo/cas-server/types"
)

var (
	config *types.Config
)

// SupportServices adds support for services related admin feature
func SupportServices(cfg *types.Config) {
	config = cfg
}

// HandleServices handles `/admin/services` request
func HandleServices(w http.ResponseWriter, r *http.Request) {
	tools.LogAdmin("Service listing requested")

	js, err := json.Marshal(config)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")

	w.Write(js)
}
