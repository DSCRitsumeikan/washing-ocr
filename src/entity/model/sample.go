package model

import (
	"errors"
	"fmt"
)

type Sample struct {
	Sample string
}

func NewSample(sample string) (*Sample, error) {
	err := checkSample(sample)
	if err != nil {
		return nil, fmt.Errorf("failed checkSample: %w", err)
	}

	return &Sample{
		Sample: sample,
	}, nil
}

func checkSample(sample string) (err error) {
	if len(sample) == 0 {
		return errors.New("sample string is empty")
	}
	return
}
