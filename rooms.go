package goteams

// This ia Room object
type Room struct {
	ID 				string 		`json:"id"`
	Title			string		`json:"title"`
	Type			string		`json:"type"`
	Locked			bool		`json:"isLocked"`
	TeamID			string		`json:"teamId"`
	LastActivity	string		`json:"lastActivity"`
}
