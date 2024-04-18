package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/tibeahx/growers-dairy/app/common"
	"github.com/tibeahx/growers-dairy/app/types"
)

func (h *Handler) GrowLogs(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	growLogs, err := h.service.GrowLogs(ctx)
	if err != nil {
		common.JSONErr(w, err, http.StatusNoContent)
		return
	}
	common.JSON(w, growLogs)
}

func (h *Handler) CreateNewGrowLog(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}
	defer r.Body.Close()

	var req types.CreateGrowLogReq
	if err := json.Unmarshal(body, &req); err != nil {
		common.JSONErr(w, err, http.StatusBadRequest)
		return
	}
	growLow, err := h.service.CreateGrowLog(ctx, req)
	if err != nil {
		common.JSONErr(w, err, http.StatusInternalServerError)
		return
	}
	common.JSON(w, growLow)
}

func (h *Handler) GrowLogByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := common.GetParamsFromUrL(r, "id")
	if err != nil {
		common.JSONErr(w, err, http.StatusBadRequest)
		return
	}
	growLog, err := h.service.GrowLogByID(ctx, id)
	if err != nil {
		common.JSONErr(w, err, http.StatusInternalServerError)
		return
	}
	common.JSON(w, growLog)
}

func (h *Handler) UpdateGrowLog(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		common.JSONErr(w, err, http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var req types.UpdateGrowLogReq
	if err := json.Unmarshal(body, &req); err != nil {
		common.JSONErr(w, err, http.StatusBadRequest)
		return
	}
	updatedGrowLog, err := h.service.UpdateGrowLog(ctx, req)
	if err != nil {
		common.JSONErr(w, err, http.StatusInternalServerError)
		return
	}
	common.JSON(w, updatedGrowLog)
}

func (h *Handler) DeleteGrowLogByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := common.GetParamsFromUrL(r, "id")
	if err != nil {
		common.JSONErr(w, err, http.StatusBadRequest)
		return
	}
	if err := h.service.DeleteGrowLogByID(ctx, id); err != nil {
		common.JSONErr(w, err, http.StatusInternalServerError)
		return
	}
	common.JSON(w, "sucess")
}
