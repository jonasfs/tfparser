package models

import (
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
	SteamID64 uint64 `gorm:"unique"`
}

// MatchPlayer is used on the match view
type MatchPlayer struct {
	SteamID64 uint64
}

// Match is ...
type Match struct {
	gorm.Model
	ParserVersion int
	FileHash      []byte `gorm:"unique"`
	FilePath      string
	MapName       string
	FinalScore1   int
	FinalScore2   int
	//Team1			[]MatchPlayer
	//Team2			[]MatchPlayer
}
