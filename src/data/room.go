package data

type Room struct {
	RoomId   string   `json:"roomId"`
	HostId   string   `json:"hostId"`
	RoomName string   `json:"roomName"`
	DocId    string   `json:"docId"`
	Members  []string `json: members`
}
