package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/tibeahx/growers-dairy/app/common"
	"github.com/tibeahx/growers-dairy/app/types"
)

func (h *Handler) LogEntries(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	entries, err := h.service.LogEntries(ctx)
	if err != nil {
		common.JSONErr(w, err, http.StatusInternalServerError)
		return
	}
	common.JSON(w, entries)
}

func (h *Handler) CreateLogEntry(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		common.JSONErr(w, err, http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var req types.CreateLogEntryReq
	if err := json.Unmarshal(body, &req); err != nil {
		common.JSONErr(w, err, http.StatusBadRequest)
		return
	}
	newEntry, err := h.service.CreateLogEntry(ctx, req)
	if err != nil {
		common.JSONErr(w, err, http.StatusBadRequest)
		return
	}
	common.JSON(w, newEntry)
}

func (h *Handler) LogEntryByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := common.GetParamsFromUrL(r, "id")
	if err != nil {
		common.JSONErr(w, err, http.StatusInternalServerError)
	}
	actualEntry, err := h.service.EntryByID(ctx, id)
	if err != nil {
		common.JSONErr(w, err, http.StatusBadRequest)
		return
	}
	common.JSON(w, actualEntry)
}

func (h *Handler) UpdateEntry(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		common.JSONErr(w, err, http.StatusInternalServerError)
	}
	var req types.UpdateLogEntryReq
	if err := json.Unmarshal(body, &req); err != nil {
		common.JSONErr(w, err, http.StatusInternalServerError)
	}
	updatedEntry, err := h.service.UpdateEntry(ctx, req)
	if err != nil {
		common.JSONErr(w, err, http.StatusBadRequest)
	}
	common.JSON(w, updatedEntry)
}

func (h *Handler) DeleteEntryByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := common.GetParamsFromUrL(r, "id")
	if err != nil {
		common.JSONErr(w, err, http.StatusBadRequest)
	}
	if err := h.service.DeleteEntryByID(ctx, id); err != nil {
		common.JSONErr(w, err, http.StatusBadRequest)
	}
	common.JSON(w, "success")
}
func (h *Handler) UploadToBucket(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		common.JSONErr(w, err, http.StatusBadRequest)
		return
	}
	logEntryID := r.FormValue("logentry_id")
	if len(logEntryID) == 0 {
		common.JSONErr(w, fmt.Errorf("missing value for logentry_id"), http.StatusBadRequest)
		return
	}
	file, header := h.service.GetFileFromForm(r)
	defer file.Close()

	url, err := h.service.UploadFileToS3(ctx, file, header)
	if err != nil {
		common.JSONErr(w, err, http.StatusInternalServerError)
		return
	}
	intID, _ := strconv.Atoi(logEntryID)
	err = h.service.LinkURLToEntry(ctx, r, intID, url)
	if err != nil {
		common.JSONErr(w, err, http.StatusInternalServerError)
	}
	common.JSON(w, "sucess")
}
