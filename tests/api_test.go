package tests

import (
	"io"
	"net/http"
	"testing"
)

func TestCommitSurvey(t *testing.T) {

	r, _ := http.NewRequest("POST", "/commitSurvey")
}
