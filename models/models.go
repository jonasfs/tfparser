package models

import (
	"time"
	"gorm.io/gorm"
)

// Settings is used to save user defined settings
type Settings struct {
	gorm.Model
	DemoPath    string `gorm:"default:\"\""`
	DemoPathSet bool   `gorm:"default:false"`
	Player      uint64 `gorm:"default:0"`
	PlayerSet   bool   `gorm:"default:false"`
}

// Player is ...
type Player struct {
	gorm.Model
	SteamID64	uint64 `gorm:"unique"`
}

// MatchPlayer is used on the match view
type MatchPlayer struct {
	SteamID64	uint64
	Player		Player `gorm:"foreignKey:SteamID64"`
	MatchID		uint
	Match		Match
	Nickname	string
}

// Match is ...
type Match struct {
	gorm.Model
	ParserVersion	int
	FileHash		[]byte `gorm:"unique"`
	FilePath		string
	FileDate		time.Time
	MapName			string
	Score1			int
	Score2			int
	Players			[]MatchPlayer
}
