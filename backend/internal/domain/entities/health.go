package entities

import "errors"

type Stats struct {
	CPUPercent string
	CPUCore    int
}

type Health struct {
	Status  string
	Version string
	Stats   Stats
}

func NewHealth(status string, version string, stats Stats) (*Health, error) {
	if status == "" {
		return nil, errors.New("name is required")
	}
	if version == "" {
		return nil, errors.New("version is required")
	}
	return &Health{
		Status:  status,
		Version: version,
		Stats:   stats,
	}, nil
}
