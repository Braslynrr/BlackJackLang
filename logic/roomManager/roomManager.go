package roommanager

import (
	"errors"
	"fmt"

	"blackjack.com/player"
	"blackjack.com/room"
	"github.com/gorilla/websocket"
)

type RoomManager struct {
	rooms []*room.Room
}

// predicate for FindFirtsOrDefault method
type RoomLambda func(room.Room) bool

// RoomPredicate return a RoomLambda to check if the room is joinable
func RoomPredicate(r room.Room) RoomLambda {
	return func(room room.Room) bool {
		return room.CheckJoin(r)
	}
}

// isPublic verifies is the room is public and it's not playing
func isPublic(room room.Room) bool {
	return !room.PlayingNow && room.IsPublic()
}

// FindFirtsOrDefault works like the functional programming 'filter' method,
// but instead of returns a list, returns the firts element or nil
// using RoomLambda as predicate
func (roomManager *RoomManager) FindFirtsOrDefault(predicate RoomLambda) *room.Room {
	for _, room := range roomManager.rooms {
		if predicate(*room) {
			return room
		}
	}
	return nil
}

// GetAll works like the functional programing filter method using RoomLambda as predicate
func (roomManager *RoomManager) GetAll(predicate RoomLambda) []*room.Room {
	var list []*room.Room
	for _, room := range roomManager.rooms {
		if predicate(*room) {
			list = append(list, room)
		}
	}
	return list
}

// AddRoom adds a room to the roomManager list, also adds the code
func (roomManager *RoomManager) AddRoom(room *room.Room) *room.Room {
	room.Code = fmt.Sprintf("0%v", len(roomManager.rooms)+1)
	roomManager.rooms = append(roomManager.rooms, room)
	return room
}

// GetPublicRooms gets all public rooms
func (room *RoomManager) GetPublicRooms() []*room.Room {
	return room.GetAll(isPublic)
}

// ConnectToRoom connects the Ws with the player in its room
func (manager *RoomManager) ConnectToRoom(conn *websocket.Conn, player player.Player, room room.Room) bool {
	serverRoom := manager.FindFirtsOrDefault(RoomPredicate(room))
	return serverRoom.AllowConnection(player, conn)
}

// StartGame sets the game in a initial state and notifies to the next player to play
func (manager *RoomManager) StartGame(player player.Player, room room.Room) (bool, error) {
	serverRoom := manager.FindFirtsOrDefault(RoomPredicate(room))
	if serverRoom != nil {

		if serverRoom.IsPlayerHost(player) {
			err := serverRoom.StartGame(true)
			return true, err
		} else {
			return false, errors.New(fmt.Sprintf("Player is not the host of the room %v", serverRoom.Code))
		}
	}
	return false, errors.New(fmt.Sprintf("The Room %v does not exist, imposible to start a game", serverRoom.Code))
}
