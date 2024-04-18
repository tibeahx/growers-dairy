package storage

import (
	"errors"
	"time"
)

const (
	queryTimeout = 10 * time.Second
)

var (
	errGrowLogNotFound = errors.New("growlog not found")
	errScanningRows    = errors.New("error scanning rows")
	errStrainNotFound  = errors.New("strain not found")
	errEntryNotFound   = errors.New("logentry not found")
)
