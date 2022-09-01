package controller

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/mv-kan/the-school-project/entity"
	"github.com/mv-kan/the-school-project/logger"
	"github.com/mv-kan/the-school-project/repo"
	"github.com/mv-kan/the-school-project/service"
	testingdb "github.com/mv-kan/the-school-project/testing-db"
	"github.com/mv-kan/the-school-project/web/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRoomController_GetRoomType(t *testing.T) {
	// init service and controller
	db := connectToDB()
	log := logger.New()

	roomServ := service.New(repo.New[entity.Room](db))
	roomStatServ := service.NewRoomStat(db)
	roomCtrl := NewRoom(log, roomServ, roomStatServ)

	// init request and handler
	path := "/rooms/{id}/type"
	r := *mux.NewRouter()
	r.HandleFunc(path, roomCtrl.GetRoomType).Methods(http.MethodGet)

	req, err := http.NewRequest("GET", fmt.Sprintf("/rooms/%d/type", testingdb.TestRoomID), nil)
	require.Nil(t, err)
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)
	// assert response
	assert.Equal(t, http.StatusOK, rr.Code)
	roomType, err := utils.ParseJSONNoValidator[entity.RoomType](rr.Result().Body)
	require.Nil(t, err)
	assert.Equal(t, testingdb.TestRoomTypeID, roomType.ID)
}

func TestRoomController_GetAvailableSpace(t *testing.T) {
	// init service and controller
	db := connectToDB()
	log := logger.New()

	roomServ := service.New(repo.New[entity.Room](db))
	roomStatServ := service.NewRoomStat(db)
	roomCtrl := NewRoom(log, roomServ, roomStatServ)

	// init request and handler
	path := "/rooms/{id}/available-space"
	r := *mux.NewRouter()
	r.HandleFunc(path, roomCtrl.GetAvailableSpace).Methods(http.MethodGet)

	req, err := http.NewRequest("GET", fmt.Sprintf("/rooms/%d/available-space", testingdb.TestRoomID), nil)
	require.Nil(t, err)
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)
	// assert response
	assert.Equal(t, http.StatusOK, rr.Code)
	availableSpaceMap, err := utils.ParseJSONNoValidator[map[string]int](rr.Result().Body)
	require.Nil(t, err)
	availableSpace := availableSpaceMap["availableSpace"]
	assert.Equal(t, testingdb.TestAvailableSpace, availableSpace)
}

func TestRoomController_GetAllResidents(t *testing.T) {
	// init service and controller
	db := connectToDB()
	log := logger.New()

	roomServ := service.New(repo.New[entity.Room](db))
	roomStatServ := service.NewRoomStat(db)
	roomCtrl := NewRoom(log, roomServ, roomStatServ)

	// init request and handler
	path := "/rooms/{id}/all-residents"
	r := *mux.NewRouter()
	r.HandleFunc(path, roomCtrl.GetAllResidents).Methods(http.MethodGet)

	req, err := http.NewRequest("GET", fmt.Sprintf("/rooms/%d/all-residents", testingdb.TestRoomID), nil)
	require.Nil(t, err)
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)
	// assert response
	assert.Equal(t, http.StatusOK, rr.Code)
	residents, err := utils.ParseJSONNoValidator[[]entity.Pupil](rr.Result().Body)
	require.Nil(t, err)
	assert.Equal(t, testingdb.TestPupilInDB.ID, residents[0].ID)
}
