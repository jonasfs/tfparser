package database

import (
	"os"
	"path/filepath"

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

	DB.db.AutoMigrate(&models.Settings{}, &models.Match{})
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

func (db database) CreateMatch(
	parserVersion int,
	hash []byte,
	path string,
	mapname string,
	score1 int,
	score2 int,
) (match *models.Match, err error) {
	match = &models.Match{
		ParserVersion: parserVersion,
		FileHash:      hash,
		FilePath:      path,
		MapName:       mapname,
		FinalScore1:   score1,
		FinalScore2:   score2,
	}
	result := DB.db.Create(match)
	err = result.Error

	return
}

func (db database) UpdateMatch(
	parserVersion int,
	hash []byte,
	path string,
	mapname string,
	score1 int,
	score2 int,
) (match *models.Match, err error) {
	match, err = DB.GetMatchByHash(hash)
	if err != nil {
		return
	}
	match.FilePath = path
	match.MapName = mapname
	match.FinalScore1 = score1
	match.FinalScore2 = score2
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

/*
func (db database) AddOrUpdateMatch(
	p ParserUtilsChecker,
	path string,
) (match *models.Match, err error) {
	// get file hash first
	var hash []byte
	hash, err = p.HashFile(path)
	if err != nil {
		return
	}

	// check if match was parsed before
	match, err = DB.GetMatchByHash(hash)
	currentVersion := p.GetParserVersion()

	if match != nil && ((match.ParserVersion == currentVersion) && (match.FilePath == path)) {
		return match, err
	}

	mapname := ""
	score1 := 0
	score2 := 0
	mapname, score1, score2 = parser.ParseMatch(path) // TODO

	if err == gorm.ErrRecordNotFound {
		match, err = DB.CreateMatch(currentVersion, hash, path, mapname, score1, score2)
	} else {
		match, err = DB.UpdateMatch(currentVersion, hash, path, mapname, score1, score2)
	}

	return match, err
}
*/
