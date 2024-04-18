package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/tibeahx/growers-dairy/app/common"
	"github.com/tibeahx/growers-dairy/app/types"
)

func (h *Handler) Strains(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	strains, err := h.service.Strains(ctx)
	if err != nil {
		common.JSONErr(w, err, http.StatusNoContent)
		return
	}
	common.JSON(w, strains)
}

func (h *Handler) CreateStrain(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		common.JSONErr(w, err, http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var req types.CreateStrainReq
	if err := json.Unmarshal(body, &req); err != nil {
		common.JSONErr(w, err, http.StatusBadRequest)
		return
	}
	newStrain, err := h.service.CreateStrain(ctx, req)
	if err != nil {
		common.JSONErr(w, err, http.StatusInternalServerError)
		return
	}
	common.JSON(w, newStrain)
}

func (h *Handler) StrainByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := common.GetParamsFromUrL(r, "id")
	if err != nil {
		common.JSONErr(w, err, http.StatusBadRequest)
		return
	}
	strain, err := h.service.StrainByID(ctx, id)
	if err != nil {
		common.JSONErr(w, err, http.StatusInternalServerError)
		return
	}
	common.JSON(w, strain)
}

func (h *Handler) UpdateStrain(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		common.JSONErr(w, err, http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var req types.UpdateStrainReq
	if err := json.Unmarshal(body, &req); err != nil {
		common.JSONErr(w, err, http.StatusBadRequest)
		return
	}
	updatedStrain, err := h.service.UpdateStrain(ctx, req)
	if err != nil {
		common.JSONErr(w, err, http.StatusInternalServerError)
		return
	}
	common.JSON(w, updatedStrain)
}

func (h *Handler) DeleteStrainByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := common.GetParamsFromUrL(r, "id")
	if err != nil {
		common.JSONErr(w, err, http.StatusBadRequest)
		return
	}
	if err := h.service.DeleteStrainByID(ctx, id); err != nil {
		common.JSONErr(w, err, http.StatusInternalServerError)
		return
	}
	common.JSON(w, "sucess")
}
