package service

import (
	"github.com/mv-kan/the-school-project/entity"
	"github.com/mv-kan/the-school-project/repo"
	"gorm.io/gorm"
)

func NewRoom(db *gorm.DB) IRoomService {
	pupilRepo := repo.New[entity.Pupil](db)
	roomRepo := repo.New[entity.Room](db)
	roomTypeRepo := repo.New[entity.RoomType](db)
	return roomService{service: service[entity.Room]{r: roomRepo}, pupilRepo: pupilRepo, roomTypeRepo: roomTypeRepo}
}

type IRoomService interface {
	IService[entity.Room]
	FindRoomType(room_id int) (entity.RoomType, error)
	// check available space in room
	FindAvailableSpace(room_id int) (int, error)
	// get all residents in a room
	FindAllResidents(room_id int) ([]entity.Pupil, error)
}

type roomService struct {
	service[entity.Room]
	roomTypeRepo repo.IRepository[entity.RoomType]
	pupilRepo    repo.IRepository[entity.Pupil]
}

func (s roomService) FindAvailableSpace(room_id int) (int, error) {
	// check for existing
	_, err := s.r.Find(room_id)
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

func (s roomService) FindAllResidents(room_id int) ([]entity.Pupil, error) {
	// check for exising
	_, err := s.r.Find(room_id)
	if err != nil {
		return nil, err
	}
	pupils, err := s.pupilRepo.FindAll()
	if err != nil {
		return nil, err
	}
	residents := make([]entity.Pupil, 0)
	for _, pupil := range pupils {
		if pupil.RoomID.Valid && pupil.RoomID.Int32 == int32(room_id) {
			residents = append(residents, pupil)
		}
	}
	return residents, nil
}

func (s roomService) FindRoomType(room_id int) (entity.RoomType, error) {
	room, err := s.r.Find(room_id)
	if err != nil {
		return entity.RoomType{}, err
	}
	roomType, err := s.roomTypeRepo.Find(room.RoomTypeID)
	if err != nil {
		return entity.RoomType{}, err
	}
	return roomType, nil
}
