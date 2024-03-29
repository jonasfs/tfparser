package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/asticode/go-astilectron"
	bootstrap "github.com/asticode/go-astilectron-bootstrap"

	"topfrag.org/tfparser/database"
	"topfrag.org/tfparser/models"
	"topfrag.org/tfparser/parser"
	"topfrag.org/tfparser/profile"
)

type ParserUtilsChecker interface {
	HashFile(path string) ([]byte, error)
	GetFiles(path string) ([]string, error)
	GetParserVersion() int
}

type parserUtilsCheck struct{}

func (pu parserUtilsCheck) hashFile(path string) ([]byte, error) {
	hash, err := parser.HashFile(path)
	return hash, err
}

func (pu parserUtilsCheck) getFiles(path string) ([]string, error) {
	filePaths, err := parser.GetFiles(path)
	return filePaths, err
}

type databaseChecker interface {
	Settings() (*models.Settings, error)
	InitSettings() *models.Settings
	SaveSettings(settings *models.Settings) bool
	GetLatestNickname(steamid uint64) (string, time.Time, error)
}

type databaseCheck struct{}

func (d databaseCheck) Settings() (*models.Settings, error) {
	db := database.DB
	return db.Settings()
}

func (d databaseCheck) InitSettings() *models.Settings {
	db := database.DB
	return db.InitSettings()
}

func (d databaseCheck) SaveSettings(settings *models.Settings) bool {
	db := database.DB
	return db.SaveSettings(settings)
}

func (d databaseCheck) GetLatestNickname(steamid uint64) (
	string, time.Time, error,
) {
	db := database.DB
	return db.GetLatestNickname(steamid)
}

var db databaseChecker

func handleMessages(w *astilectron.Window, m bootstrap.MessageIn) (payload interface{}, err error) {
	db = databaseCheck{}
	switch m.Name {
	case "checkFirstTime":
		payload = profile.CheckFirstTime(db)
		return
	case "validatePath":
		var path string
		if len(m.Payload) > 0 {
			if err = json.Unmarshal(m.Payload, &path); err != nil {
				payload = err.Error()
				return
			}
			if !filepath.IsAbs(path) {
				execPath, _ := os.Executable()
				execDir, _ := filepath.Abs(filepath.Dir(execPath))
				path = filepath.Join(execDir, path)
			}
			fmt.Printf("debug - path: %s\n", path)
			payload, err = parser.GetFiles(path)
			if err != nil {
				fmt.Printf("debug2 - error: %s\n", err)
				if os.IsNotExist(err) {
					payload = "os.IsNotExist"
				}
				return
			}
			fmt.Printf("debug3 - payload: %s\n", payload)
			var result bool
			result, err = profile.SetDemoPath(path, db)
			if err != nil {
				fmt.Printf("debug4 - error: %s\n", err)
				payload = "os.IsNotExist"
				return
			}

			fmt.Printf("debug6 - set demo path result: %t\n", result)
			fmt.Printf("debug7 - set demo path error: %s\n", err)
			return
		}
	case "setDemoPath":
		var path string
		if len(m.Payload) > 0 {
			if err = json.Unmarshal(m.Payload, &path); err != nil {
				payload = err.Error()
				return
			}
			payload, err = profile.SetDemoPath(path, db)
		}

	case "initializeParser":
		fmt.Println("debug8 - initializeParser message")
		settings, _ := db.Settings()
		path := settings.DemoPath
		payload, err = parser.StartParser(w, path)
		return
	case "getPlayerList":
		fmt.Println("debug12 - getPlayerList message")
		payload = parser.GetPlayerList()
		return
	case "setPlayerProfile":
		var steamid string
		if len(m.Payload) > 0 {
			if err = json.Unmarshal(m.Payload, &steamid); err != nil {
				payload = err.Error()
				return
			}
			var steamid64 uint64
			steamid64, err = strconv.ParseUint(string(steamid), 10, 64)
			if err == nil {
				payload, err = profile.SetPlayer(steamid64, db)
				if err != nil {
					payload = err.Error()
				}
			} else {
				payload = err.Error()
			}
		}
		return
	case "getPlayerProfile":
		var player map[string]interface{}
		player, err = profile.GetPlayer(db)
		if err == nil {
			payload = player
		}
		return
	case "getPlayerMatches":
		fmt.Printf("getPlayerMatches\n")
		var steamid string
		if len(m.Payload) > 0 {
			fmt.Printf("m.Payload: %s\n", m.Payload)
			if err = json.Unmarshal(m.Payload, &steamid); err != nil {
				payload = err.Error()
				return
			}
			fmt.Printf("steamid: %s\n", steamid)
			var steamid64 uint64
			steamid64, err = strconv.ParseUint(string(steamid), 10, 64)
			fmt.Printf("steamid64: %d\n", steamid64)
			if err == nil {
				payload = parser.GetPlayerMatches(steamid64)
				if err != nil {
					payload = err.Error()
				}
			} else {
				payload = err.Error()
			}
		}
		return
	}
	return
}
