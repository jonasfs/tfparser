package database

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"topfrag.org/tfparser/models"
)

type database struct {
	filename string
	db       *gorm.DB
	err      error
}

var DB database

func init() {
	DB.filename = "tfparse.db"
	executablePath, _ := os.Executable()
	dir, _ := filepath.Abs(filepath.Dir(executablePath))
	path := filepath.Join(dir, DB.filename)
	DB.db, DB.err = gorm.Open(sqlite.Open(path), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if DB.err != nil {
		panic("failed to connect database")
	}

	DB.db.AutoMigrate(
		&models.Settings{},
		&models.Match{},
		&models.MatchPlayer{},
		&models.Player{},
	)
}

func (db database) Settings() (settings *models.Settings, err error) {
	settings = &models.Settings{}
	result := DB.db.First(settings)
	if result.Error != nil {
		settings = nil
	}
	err = result.Error
	return
}

func (db database) InitSettings() *models.Settings {
	settings := models.Settings{}
	DB.db.Create(&settings)
	return &settings
}

func (db database) SaveSettings(settings *models.Settings) bool {
	DB.db.Save(settings)
	return true
}

/* Match functions */

func (db database) CreateMatch(
	parserVersion int,
	hash []byte,
	path string,
	date time.Time,
	mapname string,
	score1 int,
	score2 int,
	participants map[uint64]map[string]interface{},
) (match *models.Match, err error) {
	match = &models.Match{
		ParserVersion: parserVersion,
		FileHash:      hash,
		FilePath:      path,
		FileDate:      date,
		MapName:       mapname,
		Score1:        score1,
		Score2:        score2,
	}
	result := DB.db.Create(match)
	err = result.Error
	if err == nil {
		for steamid, player := range participants {
			DB.CreateMatchPlayer(
				match.ID,
				steamid,
				player["nickname"].(string),
			)
		}
	}

	return
}

func (db database) UpdateMatch(
	parserVersion int,
	hash []byte,
	path string,
	date time.Time,
	mapname string,
	score1 int,
	score2 int,
	participants map[uint64]map[string]interface{},
) (match *models.Match, err error) {
	match, err = DB.GetMatchByHash(hash)
	if err != nil {
		return
	}
	match.FilePath = path
	match.FileDate = date
	match.MapName = mapname
	match.Score1 = score1
	match.Score2 = score2
	match.ParserVersion = parserVersion
	DB.db.Save(match)
	return
}

func (db database) GetMatchByHash(hash []byte) (*models.Match, error) {
	var err error
	result := DB.db.Where(&models.Match{FileHash: hash})
	err = result.Error
	if err != nil {
		return nil, err
	}
	var match models.Match
	result.First(&match)
	err = result.Error
	if err != nil {
		return nil, err
	}
	return &match, err
}

func (db database) GetPlayerMatches(steamid uint64) map[string]map[string]interface{} {

	matches := make(map[string]map[string]interface{})
	query := fmt.Sprintf(`
		SELECT
			matches.id,
			matches.file_date,
			matches.map_name,
			matches.score1,
			matches.score2
		FROM matches INNER JOIN match_players
		ON matches.id = match_players.match_id
		WHERE match_players.steam_id64 = %d
		ORDER BY matches.file_date DESC
	`, steamid)
	rows, err := DB.db.Raw(query).Rows()
	defer rows.Close()
	if err == nil {
		for rows.Next() {
			match := make(map[string]interface{})
			var id uint
			var score1, score2 int
			var timestamp time.Time
			var map_name string
			rows.Scan(&id, &timestamp, &map_name, &score1, &score2)
			idString := fmt.Sprintf("%d", id)
			match["id"] = idString
			match["timestamp"] = timestamp
			match["mapName"] = map_name
			match["score1"] = score1
			match["score2"] = score2
			matches[idString] = match
		}
	}
	return matches
}

/* Player functions */

func (db database) CreatePlayer(
	steamid uint64,
) (player *models.Player, err error) {
	player = &models.Player{
		SteamID64: steamid,
	}
	result := DB.db.Create(player)
	err = result.Error
	return
}

func (db database) GetPlayerBySteamID(steamid uint64) (*models.Player, error) {
	var err error
	result := DB.db.Where(&models.Player{SteamID64: steamid})
	err = result.Error
	if err != nil {
		return nil, err
	}
	var player models.Player
	result.First(&player)
	err = result.Error
	if err != nil {
		return nil, err
	}
	return &player, err
}

func (db database) GetPlayers() map[string]map[string]interface{} {
	result := DB.db.Model(&models.Player{})
	players := make(map[string]map[string]interface{})
	rows, err := result.Rows()
	defer rows.Close()
	if err == nil {
		for rows.Next() {
			var player_object models.Player
			DB.db.ScanRows(rows, &player_object)
			player := make(map[string]interface{})
			stringSteamID := strconv.FormatUint(player_object.SteamID64, 10)
			player["steamid"] = stringSteamID
			player["nickname"], player["last"], err = DB.GetLatestNickname(
				player_object.SteamID64)
			query2 := fmt.Sprintf(`
				SELECT
					COUNT(*)
				FROM match_players
				WHERE match_players.steam_id64 = %d
			`, player_object.SteamID64)
			row2 := DB.db.Raw(query2).Row()
			var demos int
			row2.Scan(&demos)
			player["demos"] = demos
			players[stringSteamID] = player
		}
	}
	return players
}

func (db database) CreateMatchPlayer(
	matchID uint,
	steamID uint64,
	nickname string,
) (match_player *models.MatchPlayer, err error) {
	match_player = &models.MatchPlayer{
		SteamID64: steamID,
		MatchID:   matchID,
		Nickname:  nickname,
	}
	result := DB.db.Create(match_player)
	err = result.Error

	return
}

func (db database) GetLatestNickname(
	steamid uint64,
) (nickname string, date time.Time, err error) {
	query := fmt.Sprintf(`
		SELECT
			match_players.nickname,
			matches.file_date
		FROM matches INNER JOIN match_players
		ON matches.id = match_players.match_id
		WHERE match_players.steam_id64 = %d
		ORDER BY matches.file_date DESC LIMIT 1
	`, steamid)

	row := DB.db.Raw(query).Row()
	row.Scan(&nickname, &date)
	return
}
