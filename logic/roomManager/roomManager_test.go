package roommanager

import (
	"testing"

	"blackjack.com/room"
)

// TestAddRoom calls RoomManager.AddRoom checking a valid return value
func TestAddRoom(t *testing.T) {
	rm := RoomManager{}
	room := rm.AddRoom(room.NewRoom("", "", false))
	if room == nil {
		t.Fatal("room should not be nil")
	}

	if &rm.rooms[0] == &room {
		t.Fatalf("The rooms pointer should be alike. %v == %v ", &room, &rm.rooms[0])
	}

	if len(rm.rooms) < 1 {
		t.Fatal("The RoomManager should have one room")
	}
}

// TestGetPublicRooms calls RoomManager.GetPublicRooms checking a valid return value
func TestGetPublicRooms(t *testing.T) {
	rm := RoomManager{}
	rm.AddRoom(room.NewRoom("", "", false))
	rm.AddRoom(room.NewRoom("", "123", true))
	if len(rm.GetPublicRooms()) > 1 {
		t.Fatalf("The roomManager should returns just one public room. rm.GetPublicRooms()==%v", rm.GetPublicRooms())
	}
}

// TestFindFirstOrDefault calls RoomManager.FindFirtsOrDefault checking a valid return value
func TestFindFirstOrDefault(t *testing.T) {
	rm := RoomManager{}
	rm.AddRoom(room.NewRoom("", "", false))
	rm.AddRoom(room.NewRoom("", "123", true))
	if rm.FindFirtsOrDefault(RoomPredicate(*room.NewRoom("01", "", true))) == nil {
		t.Fatal("rm.FindFirtsOrDefault(RoomPredicate(*room.NewRoom(\"01\", \"\", true))) should return the room")
	}

	room1 := rm.FindFirtsOrDefault(RoomPredicate(*room.NewRoom("02", "124", true)))
	if room1 != nil {
		t.Fatalf("Room %v should be nil", room1)
	}
	room1 = rm.FindFirtsOrDefault(RoomPredicate(*room.NewRoom("02", "123", true)))
	if room1 == nil || room1.Code != "02" {
		t.Fatalf("%v should not be nil, and should return 02 as a code", room1)
	}
}
