package domain

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func Test_SuccessResponse(t *testing.T) {
	r := struct {
		Name       string `json:"name"`
		Age        int    `json:"age"`
		Year       int    `json:"year"`
		University string `json:"university"`
	}{
		Name:       "Leonardo",
		Age:        23,
		Year:       2020,
		University: "University of Antioquia",
	}

	want := Response{
		Status: SuccessStatus,
		Result: r,
		Error:  nil,
	}

	got := SuccessResponse(r)

	assert.Equal(t, got, want)
}

func Test_ErrorResponse(t *testing.T) {
	e := Error{
		Message: "Something wrong",
		Details: "An error occurs",
	}
	want := Response{
		Status: ErrorStatus,
		Result: nil,
		Error:  &e,
	}

	got := ErrorResponse(e)

	assert.Equal(t, got, want)
}
