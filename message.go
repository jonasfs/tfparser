package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

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
	}
	return
}
