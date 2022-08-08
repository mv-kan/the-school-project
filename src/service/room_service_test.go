package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// test room service function FindAvailableSpace
func TestRoomService_FindAvailableSpace(t *testing.T) {
	db := connectToDB()
	roomService := NewRoom(db)
	availableSpace, err := roomService.FindAvailableSpace(testRoomID)
	require.Nil(t, err)
	assert.Equal(t, testAvailableSpace, availableSpace)
}

func TestRoomService_FindAllResidents(t *testing.T) {
	db := connectToDB()
	roomService := NewRoom(db)
	residents, err := roomService.FindAllResidents(testRoomID)
	require.Nil(t, err)
	assert.Equal(t, testResidentID, residents[0].ID)
}

func TestRoomService_FindRoomType(t *testing.T) {
	db := connectToDB()
	roomService := NewRoom(db)
	roomType, err := roomService.FindRoomType(testRoomID)
	require.Nil(t, err)
	assert.Equal(t, testRoomTypeID, roomType.ID)
}
