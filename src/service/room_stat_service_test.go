package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// test room service function FindAvailableSpace
func TestRoomStatService_FindAvailableSpace(t *testing.T) {
	db := connectToDB()
	roomService := NewRoomStat(db)
	availableSpace, err := roomService.FindAvailableSpace(testRoomID)
	require.Nil(t, err)
	assert.Equal(t, testAvailableSpace, availableSpace)
}

func TestRoomStatService_FindAllResidents(t *testing.T) {
	db := connectToDB()
	roomService := NewRoomStat(db)
	residents, err := roomService.FindAllResidents(testRoomID)
	require.Nil(t, err)
	assert.Equal(t, testResidentID, residents[0].ID)
}

func TestRoomStatService_FindRoomType(t *testing.T) {
	db := connectToDB()
	roomService := NewRoomStat(db)
	roomType, err := roomService.FindRoomType(testRoomID)
	require.Nil(t, err)
	assert.Equal(t, testRoomTypeID, roomType.ID)
}
