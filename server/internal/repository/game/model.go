package game

type Game struct {
	GameID  string   `json:"gameId"`
	Players []Player `json:"players"`
	Host    Player   `json:"host"`
	Word    Word     `json:"word"`
}

type Word struct {
	Undercover string `json:"undercover"`
	Normal     string `json:"normal"`
}

type Role string

const (
	Undercover Role = "UNDERCOVER"
	Normal     Role = "NORMAL"
	White      Role = "WHITE"
)

type Player struct {
	PlayerID string `json:"playerId"`
	Name     string `json:"name"`
	Role     Role   `json:"role"`
	IsReveal bool   `json:"isReveal"`
}
