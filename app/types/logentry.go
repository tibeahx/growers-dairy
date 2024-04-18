package types

import (
	"time"
)

type LogEntry struct {
	ID        int       `json:"id"`
	EntryDate time.Time `json:"entry_date"`
	Comment   string    `json:"comment"`
	Photos    []string  `json:"photos"`
	GrowLogID int       `json:"growlog_id"`
}

type CreateLogEntryReq struct {
	EntryDate time.Time `json:"entry_date"`
	Comment   string    `json:"comment"`
	GrowLogID int       `json:"growlog_id"`
}

type CreateLogEntryResp struct {
	ID        int       `json:"id"`
	EntryDate time.Time `json:"entry_date"`
	Comment   string    `json:"comment"`
	GrowLogID int       `json:"growlog_id"`
}

type UpdateLogEntryReq struct {
	ID        int      `json:"id"`
	Comment   string   `json:"comment"`
	Photos    []string `json:"photos"`
	GrowLogID int      `json:"growlog_id"`
}

type UpdateLogEntryResp struct {
	ID        int      `json:"id"`
	Comment   string   `json:"comment"`
	Photos    []string `json:"photos"`
	GrowLogID int      `json:"growlog_id"`
}
