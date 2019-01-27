package handler

import (
	"encoding/json"
	"net/http"

	"github.com/miry/wattx_top_coins/cmd/top_coins/app"
)

// CoinsHandler process top coins endpoint
type CoinsHandler struct {
	app *app.App
}

// NewCoinsHandler initialize VersionHandler object
func NewCoinsHandler(app *app.App) *CoinsHandler {
	return &CoinsHandler{app: app}
}

// Rank,	Symbol,	Price USD
type coinResp struct {
	Rank   int64   `json:"rank"`
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
}

type coinsResp struct {
	Data []coinResp `json:"data"`
}

// List build json result for top coins
func (h *CoinsHandler) List(w http.ResponseWriter, r *http.Request) {
	resp := coinsResp{
		Data: []coinResp{
			{1, "BTC", 6634.41},
			{2, "ETH", 370.237},
			{3, "XRP", 0.471636},
		},
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), 500)
		h.app.Logger.Error().Err(err).Msg("Could not render version")
	}
}
