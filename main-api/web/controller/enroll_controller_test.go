package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mv-kan/the-school-project/entity"
	"github.com/mv-kan/the-school-project/logger"
	"github.com/mv-kan/the-school-project/service"
	testingdb "github.com/mv-kan/the-school-project/testing-db"
	"github.com/stretchr/testify/require"
)

func TestEnrollController_Enroll(t *testing.T) {
	// init controller service
	db := connectToDB()
	log := logger.New()
	enrollServ := service.NewEnroll(db)
	enrollCtrl := NewEnroll(log, enrollServ)
	pupilToEnroll := entity.Pupil{
		Name:          "Test",
		Surname:       "Something",
		SchoolClassID: testingdb.TestSchoolClassID,
	}
	// test enroll
	pupilJSON, err := json.Marshal(pupilToEnroll)
	require.Nil(t, err)

	req, err := http.NewRequest("POST", "/enroll", bytes.NewReader(pupilJSON))
	require.Nil(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(enrollCtrl.Enroll)

	handler.ServeHTTP(rr, req)

	require.Equal(t, http.StatusOK, rr.Code)

}
