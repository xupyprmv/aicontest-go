package structs

import (
	"time"
)

type Game struct {
	GID       string
	START     time.Time
	FINISH    time.Time
	BOTID     []string
	RESULTMAP map[string]float64
	LISTING   string
	NOTES     []string
}
