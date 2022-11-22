package interactor

import (
	"app/src/application/input"
	"app/src/application/output"
	"app/src/entity/model"
	"fmt"
)

type sampleInteractor struct {
}

func NewSampleInteractor() *sampleInteractor {
	return &sampleInteractor{}
}

func (s *sampleInteractor) Show(showSample input.ShowSample) (*output.ShowSample, error) {
	sample, err := model.NewSample(showSample.Sample)
	if err != nil {
		return nil, fmt.Errorf("failed repository.SampleRepository.Show: %w", err)
	}

	return &output.ShowSample{
		Sample: sample.Sample,
	}, nil
}
