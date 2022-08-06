package roommanager

import (
	"fmt"

	"blackjack.com/room"
)

type RoomManager struct {
	rooms []*room.Room
}

type RoomLambda func(room.Room) bool

func isPublic(room room.Room) bool {
	return !room.PlayingNow && room.IsPublic()
}

func (roomManager *RoomManager) FindFirtsOrDefault(predicate RoomLambda) *room.Room {
	for _, room := range roomManager.rooms {
		if predicate(*room) {
			return room
		}
	}
	return nil
}

func (roomManager *RoomManager) GetAll(predicate RoomLambda) []*room.Room {
	var list []*room.Room
	for _, room := range roomManager.rooms {
		if predicate(*room) {
			list = append(list, room)
		}
	}
	return list
}

func (roomManager *RoomManager) AddRoom(room *room.Room) *room.Room {
	room.Code = fmt.Sprintf("0%v", len(roomManager.rooms)+1)
	roomManager.rooms = append(roomManager.rooms, room)
	return room
}

func (room *RoomManager) GetPublicRooms() []*room.Room {
	return room.GetAll(isPublic)
}
