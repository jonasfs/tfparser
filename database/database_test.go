package database

import (
	"bytes"
	"errors"
	"os"
	"testing"

	"github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm/logger"
	"topfrag.org/tfparser/models"
	"topfrag.org/tfparser/parser"

	"gorm.io/gorm"
)

var hashFileMock func(path string) ([]byte, error)
var getFilesMock func(dirPath string) (filePaths []string, err error)
var getParserVersionMock func() int

type puCheckMock struct{}

func (p puCheckMock) HashFile(path string) ([]byte, error) {
	return hashFileMock(path)
}

func (p puCheckMock) GetFiles(dirPath string) (filePaths []string, err error) {
	filePaths, err = getFilesMock(dirPath)
	return
}

func (p puCheckMock) GetParserVersion() int {
	return getParserVersionMock()
}

func TestMain(m *testing.M) {
	DB.db, DB.err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	DB.db.AutoMigrate(&models.Settings{}, &models.Match{})
	os.Exit(m.Run())
}

func TestSettings(t *testing.T) {
	var settings *models.Settings
	var err error

	settings, err = DB.Settings()
	if settings != nil {
		t.Error("Expected settings to be nil, got non nil instead")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Errorf("Expected record not found error, got %s", err)
	}

	DB.InitSettings()
	settings, err = DB.Settings()
	if settings == nil {
		t.Error("Expected settings to be not nil, got nil instead")
	}
	settings.DemoPath = "testPath"
	DB.SaveSettings(settings)
	settings2, _ := DB.Settings()
	if settings2.DemoPath != settings.DemoPath {
		t.Error("Expected settings.DemoPath to be ", settings.DemoPath, " and got ", settings2.DemoPath)
	}
}

func TestCreateMatch(t *testing.T) {
	hash := []byte("123")
	testPath := "testPath"
	match, err := DB.CreateMatch(1, hash, testPath, "map_name", 16, 3)
	if err != nil {
		t.Error("Got error during match creation: ", err)
	}
	if match.FilePath != testPath {
		t.Errorf("Expected FilePath '%s', got %s", testPath, match.FilePath)
	}
	var oneMatch models.Match
	var matches []models.Match
	result := DB.db.Find(&matches)
	if result.RowsAffected != 1 {
		t.Errorf("Expected to have a total of 1 match but got %d instead ", result.RowsAffected)
	}

	result = DB.db.Where(&models.Match{FileHash: hash})
	result.First(&oneMatch)

	matchCheck, err := DB.GetMatchByHash(hash)

	if err != nil {
		t.Error("Got error during match creation: ", err)
	}
	if matchCheck == nil {
		t.Error("Got nil match")
	} else {
		if matchCheck.FilePath != testPath {
			t.Errorf("Expected FilePath '%s', got %s", testPath, matchCheck.FilePath)
		}
		if !bytes.Equal(matchCheck.FileHash, hash) {
			t.Errorf("Expected FileHash '%x', got %x", hash, matchCheck.FileHash)
		}
	}

	_, err = DB.CreateMatch(1, hash, testPath, "map_name", 16, 3)
	result = DB.db.Find(&matches)
	sqliteErr := err.(sqlite3.Error)

	if !errors.Is(sqliteErr.ExtendedCode, sqlite3.ErrConstraintUnique) {
		t.Errorf("Expected sqlite3.ErrConstraintUnique, got %s", err)
	}
	if result.RowsAffected != 1 {
		t.Errorf("Expected to have a total of 1 match but got %d instead ", result.RowsAffected)
	}
}

func TestAddOrUpdateMatch(t *testing.T) {
	puMock := puCheckMock{}
	path := "testpath"
	testHash := []byte("345")
	var match *models.Match

	hashFileMock = func(path string) ([]byte, error) {
		return testHash, nil
	}

	getParserVersionMock = func() int {
		return parser.GetParserVersion()
	}

	// New match case
	match, _ = DB.AddOrUpdateMatch(puMock, path)

	if match == nil {
		t.Error("Expected a new match, got nil")
	} else {

		if match.FilePath != path {
			t.Errorf("Expected path to be '%s', got '%s'", path, match.FilePath)
		}
		if !bytes.Equal(match.FileHash, testHash) {
			t.Error("Expected hash to be", testHash, " got ", match.FileHash)
		}
	}
	var matches []models.Match
	result := DB.db.Find(&matches)
	if result.RowsAffected != 2 {
		t.Errorf("Expected to have a total of 2 matches but got %d instead ", result.RowsAffected)
	}

	// Update path case
	newPath := "newpath"
	DB.AddOrUpdateMatch(puMock, newPath)
	matchCheck, _ := DB.GetMatchByHash(testHash)

	if matchCheck == nil {
		t.Error("Expected an existing match, got nil")

	} else {

		result = DB.db.Find(&matches)
		if result.RowsAffected != 2 {
			t.Errorf("Expected to have a total of 2 matches but got %d instead ", result.RowsAffected)
		}

		if matchCheck.FilePath != newPath {
			t.Errorf("Expected path to be '%s', got '%s'", newPath, matchCheck.FilePath)
		}
		if !bytes.Equal(matchCheck.FileHash, testHash) {
			t.Error("Expected hash to be", testHash, " got ", matchCheck.FileHash)
		}
	}

	// Update version case

	getParserVersionMock = func() int {
		return 2
	}
	DB.AddOrUpdateMatch(puMock, newPath)
	matchCheck, _ = DB.GetMatchByHash(testHash)

	if matchCheck == nil {
		t.Error("Expected an existing match, got nil")

	} else {

		result = DB.db.Find(&matches)
		if result.RowsAffected != 2 {
			t.Errorf("Expected to have a total of 2 matches but got %d instead ", result.RowsAffected)
		}

		if matchCheck.FilePath != newPath {
			t.Errorf("Expected path to be '%s', got '%s'", newPath, matchCheck.FilePath)
		}
		if !bytes.Equal(matchCheck.FileHash, testHash) {
			t.Errorf("Expected hash to be '%s', got '%s'", testHash, matchCheck.FileHash)
		}
		if matchCheck.ParserVersion != 2 {
			t.Errorf("Expected parser version to be '%d', got '%d'", 2, matchCheck.ParserVersion)
		}
	}

}
