package service

import (
	"github.com/mv-kan/the-school-project/entity"
	"github.com/mv-kan/the-school-project/repo"
	"gorm.io/gorm"
)

func NewRoomStat(db *gorm.DB) IRoomStatService {
	pupilRepo := repo.New[entity.Pupil](db)
	roomRepo := repo.New[entity.Room](db)
	roomTypeRepo := repo.New[entity.RoomType](db)
	return roomStatService{roomRepo: roomRepo, pupilRepo: pupilRepo, roomTypeRepo: roomTypeRepo}
}

type IRoomStatService interface {
	FindRoomType(room_id int) (entity.RoomType, error)
	// check available space in room
	FindAvailableSpace(room_id int) (int, error)
	// get all residents in a room
	FindAllResidents(room_id int) ([]entity.Pupil, error)
}

type roomStatService struct {
	roomRepo     repo.IRepository[entity.Room]
	roomTypeRepo repo.IRepository[entity.RoomType]
	pupilRepo    repo.IRepository[entity.Pupil]
}

func (s roomStatService) FindAvailableSpace(room_id int) (int, error) {
	// check for existing
	_, err := s.roomRepo.Find(room_id)
	if err != nil {
		return -1, err
	}
	roomType, err := s.FindRoomType(room_id)
	if err != nil {
		return -1, err
	}
	residents, err := s.FindAllResidents(room_id)
	if err != nil {
		return -1, err
	}
	return roomType.MaxOfResidents - len(residents), nil
}

func (s roomStatService) FindAllResidents(room_id int) ([]entity.Pupil, error) {
	// check for exising
	_, err := s.roomRepo.Find(room_id)
	if err != nil {
		return nil, err
	}
	pupils, err := s.pupilRepo.FindAll()
	if err != nil {
		return nil, err
	}
	residents := make([]entity.Pupil, 0)
	for _, pupil := range pupils {
		if pupil.RoomID != nil && *pupil.RoomID == room_id {
			residents = append(residents, pupil)
		}
	}
	return residents, nil
}

func (s roomStatService) FindRoomType(room_id int) (entity.RoomType, error) {
	room, err := s.roomRepo.Find(room_id)
	if err != nil {
		return entity.RoomType{}, err
	}
	roomType, err := s.roomTypeRepo.Find(room.RoomTypeID)
	if err != nil {
		return entity.RoomType{}, err
	}
	return roomType, nil
}
