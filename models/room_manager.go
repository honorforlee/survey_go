package models

type Room struct {
	IsPublic     bool
	MaxUserCount int32
	AccessCode   string
}

var rooms []*Room

func init() {
	rooms = make([]*Room, 0)
}

func NewRoom(isPubic bool, maxUserCount int32) (room *Room) {
	room = &Room{isPubic, maxUserCount, generateAccessCode()}
	rooms = append(rooms, room)
	return room
}

/**
	生成接入码
**/
func generateAccessCode() string {
	return "1234567890"
}
