package profile

import (
	"errors"
	"testing"
  "time"

	"topfrag.org/tfparser/database"
	"topfrag.org/tfparser/models"
)

var settingsExistsMock func() (*models.Settings, error)

type dbCheckMock struct{}

func (d dbCheckMock) Settings() (*models.Settings, error) {
	return settingsExistsMock()
}

func (d dbCheckMock) InitSettings() *models.Settings {
	return database.DB.InitSettings()
}

func (d dbCheckMock) SaveSettings(settings *models.Settings) bool {
	return database.DB.SaveSettings(settings)
}

func (d dbCheckMock) GetLatestNickname(steamid uint64) (string, time.Time, error){
	return "test", time.Now(), nil
}

func TestCheckFirstTime(t *testing.T) {
	dbMock := dbCheckMock{}
	var check bool

	// expect true when settings were not found
	settingsExistsMock = func() (*models.Settings, error) {
		err := errors.New("test")
		return nil, err
	}
	check = CheckFirstTime(dbMock)
	if check != true {
		t.Error("Expected true, got ", check)
	}

	// expect true when settings were found
	// but every attribute is default
	settingsExistsMock = func() (*models.Settings, error) {
		var testSettings models.Settings
		return &testSettings, nil
	}
	check = CheckFirstTime(dbMock)
	if check != true {
		t.Error("Expected true, got ", check)
	}

	// expect true when settings were found
	// but demo path is not set
	settingsExistsMock = func() (*models.Settings, error) {
		var testSettings models.Settings
		testSettings.DemoPathSet = false
		testSettings.PlayerSet = true
		return &testSettings, nil
	}
	check = CheckFirstTime(dbMock)
	if check != true {
		t.Error("Expected true, got ", check)
	}

	// expect true when settings were found
	// but player is not set
	settingsExistsMock = func() (*models.Settings, error) {
		var testSettings models.Settings
		testSettings.DemoPathSet = true
		testSettings.PlayerSet = false
		return &testSettings, nil
	}
	check = CheckFirstTime(dbMock)
	if check != true {
		t.Error("Expected true, got ", check)
	}

	// expect false when settings were found
	// and demo path and player are set
	settingsExistsMock = func() (*models.Settings, error) {
		var testSettings models.Settings
		testSettings.DemoPathSet = true
		testSettings.PlayerSet = true
		return &testSettings, nil
	}
	check = CheckFirstTime(dbMock)
	if check != false {
		t.Error("Expected false, got ", check)
	}

}
