package service

import (
	"testing"

	testingdb "github.com/mv-kan/the-school-project/testing-db"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// test room service function FindAvailableSpace
func TestRoomStatService_FindAvailableSpace(t *testing.T) {
	db := connectToDB()
	roomService := NewRoomStat(db)
	availableSpace, err := roomService.FindAvailableSpace(testingdb.TestRoomID)
	require.Nil(t, err)
	assert.Equal(t, testingdb.TestAvailableSpace, availableSpace)
}

func TestRoomStatService_FindAllResidents(t *testing.T) {
	db := connectToDB()
	roomService := NewRoomStat(db)
	residents, err := roomService.FindAllResidents(testingdb.TestRoomID)
	require.Nil(t, err)
	assert.Equal(t, testingdb.TestResidentID, residents[0].ID)
}

func TestRoomStatService_FindRoomType(t *testing.T) {
	db := connectToDB()
	roomService := NewRoomStat(db)
	roomType, err := roomService.FindRoomType(testingdb.TestRoomID)
	require.Nil(t, err)
	assert.Equal(t, testingdb.TestRoomTypeID, roomType.ID)
}
