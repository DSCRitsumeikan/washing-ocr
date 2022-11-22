package controller

import (
	"app/src/application/input"
	"app/src/controller/request"
	"app/src/controller/response"
	"net/http"
)

type sampleController struct {
	usecase input.SampleUsecase
}

func NewSampleController(usecase input.SampleUsecase) sampleController {
	return sampleController{
		usecase: usecase,
	}
}

func (sc sampleController) ShowSample(c Context) {
	sampleReq := request.SampleReq{}
	v, ok := c.GetQuery("sample")
	if ok == false {
		c.JSON(http.StatusBadRequest, "BadRequest")
		return
	}

	sampleReq.Sample = v

	sample, err := sc.usecase.Show(input.ShowSample{
		Sample: sampleReq.Sample,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, "InternalServerError")
		return
	}

	resData := response.SampleRes{
		Sample: sample.Sample,
	}

	c.JSON(http.StatusOK, resData)

}
