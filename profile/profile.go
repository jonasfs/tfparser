package profile

import (
	"fmt"
	"time"

	"topfrag.org/tfparser/database"
	"topfrag.org/tfparser/models"
)

type DatabaseChecker interface {
	Settings() (*models.Settings, error)
	InitSettings() *models.Settings
	SaveSettings(settings *models.Settings) bool
	GetLatestNickname(steamid uint64) (string, time.Time, error)
}

type databaseCheck struct{}

func (db databaseCheck) settings() (*models.Settings, error) {
	return database.DB.Settings()
}

func (db databaseCheck) initSettings() *models.Settings {
	return database.DB.InitSettings()
}

func (db databaseCheck) saveSettings(settings *models.Settings) bool {
	return database.DB.SaveSettings(settings)
}

func (db databaseCheck) getLatestNickname(steamid uint64) (
	string, time.Time, error,
) {
	return database.DB.GetLatestNickname(steamid)
}

func CheckFirstTime(db DatabaseChecker) bool {
	var settings *models.Settings
	settings, _ = db.Settings()
	if settings == nil {
		return true
	}

	if !settings.DemoPathSet {
		return true
	}
	if !settings.PlayerSet {
		return true
	}
	return false
}

func SetDemoPath(path string, db DatabaseChecker) (result bool, err error) {
	result = false

	var settings *models.Settings
	settings, _ = db.Settings()
	if settings == nil {
		settings = db.InitSettings()
	}
	settings.DemoPath = path
	settings.DemoPathSet = true
	result = db.SaveSettings(settings)
	return
}

func SetPlayer(steamid uint64, db DatabaseChecker) (result bool, err error) {
	result = false

	var settings *models.Settings
	settings, _ = db.Settings()
	if settings == nil {
		fmt.Printf("debug 5 - settings == nil")
		settings = db.InitSettings()
	}
	settings.Player = steamid
	settings.PlayerSet = true
	result = db.SaveSettings(settings)
	return
}

func GetPlayer(db DatabaseChecker) (player map[string]interface{}, err error) {
	settings, err := db.Settings()
	if err == nil {
		player = make(map[string]interface{})
		player["steamid"] = fmt.Sprintf("%d", settings.Player)
		player["nickname"], _, err = db.GetLatestNickname(settings.Player)
	}
	return
}
