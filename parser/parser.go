package parser

import (
	"fmt"
	"time"

	"github.com/asticode/go-astilectron"
	bootstrap "github.com/asticode/go-astilectron-bootstrap"
	"gorm.io/gorm"
	"topfrag.org/tfparser/database"
	"topfrag.org/tfparser/models"
)

type timestamp time.Time

type Parser struct {
	version   int
	started   bool
	startedAt *timestamp
	updatedAt *timestamp
	ticker    *time.Ticker
	channel   chan bool
}

var TP Parser

type ParserUtilsChecker interface {
	HashFile(path string) ([]byte, error)
	GetFiles(path string) ([]string, error)
	GetParserVersion() int
}

type parserUtilsCheck struct{}

func (pu parserUtilsCheck) HashFile(path string) ([]byte, error) {
	hash, err := HashFile(path)
	return hash, err
}

func (pu parserUtilsCheck) GetFiles(path string) ([]string, error) {
	filePaths, err := GetFiles(path)
	return filePaths, err
}

func (pu parserUtilsCheck) GetParserVersion() int {
	return GetParserVersion()
}

func AddOrUpdateMatch(
	path string,
) (match *models.Match, err error) {
	fmt.Printf("AddOrUpdateMatch path: %s\n", path)
	// get file hash first
	var hash []byte
	hash, err = HashFile(path)
	if err != nil {
		return
	}

	// check if match was parsed before
	match, err = database.DB.GetMatchByHash(hash)
	currentVersion := GetParserVersion()

	// match version didn't change, no need to parse
	if match != nil && ((match.ParserVersion == currentVersion) && (match.FilePath == path)) {
		fmt.Printf("Match version didn't change\n")
		return
	}

	mapname := ""
	score1 := 0
	score2 := 0
	parsed_match = ParseMatch(path)
	fmt.Printf("Parsing %s - %d : %d\n", parsed_match.mapname, parsed_match.score1, parsed_match.score2)

	if err == gorm.ErrRecordNotFound {
		match, err = database.DB.CreateMatch(currentVersion, hash, path, mapname, score1, score2)
		fmt.Printf("added to db\n")
	} else {
		match, err = database.DB.UpdateMatch(currentVersion, hash, path, mapname, score1, score2)
		fmt.Printf("updated on db\n")
	}

	return
}

func StartParser(w *astilectron.Window, path string) (result bool, err error) {
	TP.version = 1 // bump this every time you update parsing rules to force all matches to reparse
	fmt.Println("debug10 - Parser attrs")
	result = true
	parserPeriod := time.Second * 10
	TP.ticker = time.NewTicker(parserPeriod)
	RunParser(w, path) // first run is immediate
	fmt.Println("debug11")
	TP.channel = make(chan bool)
	go func() {
		for {
			select {
			case <-TP.channel:
				return
			case t := <-TP.ticker.C:
				//TODO: check dir and add files to parse queue
				fmt.Printf("parsing tick: %s\n", t)
				RunParser(w, path)
			}
		}
	}()
	return
}

func ParseMatch(path string) (mapname string, score1 int, score2 int) {
	mapname = "test_map"
	score1 = 16
	score2 = 0
	return
}

type empty struct{}

func RunParser(w *astilectron.Window, path string) {
	now := time.Now()
	fmt.Println("debug9 - StartParser")
	TP.started = true
	TP.startedAt = (*timestamp)(&now)
	TP.updatedAt = (*timestamp)(&now)
	files, _ := GetFiles(path) // get list of files
	N := len(files)

	// I have no idea what I'm doing
	sem := make(chan empty, N)
	for _, filePath := range files {
		go func(filePath string) {
			msg := fmt.Sprintf("parsing %s\n", filePath)
			bootstrap.SendMessage(w, "parsing", filePath)
			fmt.Print(msg)
			AddOrUpdateMatch(filePath)
			sem <- empty{}
		}(filePath)
	}
	for i := 0; i < N; i += 1 {
		<-sem
	}
}
