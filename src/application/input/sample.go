package input

import "app/src/application/output"

type SampleUsecase interface {
	Show(ShowSample) (*output.ShowSample, error)
}

type ShowSample struct {
	Sample string
}
