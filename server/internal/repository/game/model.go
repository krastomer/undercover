package game

type Game struct {
	GameID  string
	Players []Player
	Host    Player
	Word    Word
}

type Word struct {
	Undercover string
	Normal     string
}

type Role string

const (
	Undercover = "UNDERCOVER"
	Normal     = "NORMAL"
	White      = "WHITE"
)

type Player struct {
	PlayerID string
	Name     string
	Role     Role
	IsReveal bool
}
