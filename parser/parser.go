package parser

import (
	"fmt"
	"time"
	"os"

	"github.com/asticode/go-astilectron"
	bootstrap "github.com/asticode/go-astilectron-bootstrap"
	"gorm.io/gorm"
	"topfrag.org/tfparser/database"
	dem "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs"
)

type timestamp time.Time

// Parser is ...
type Parser struct {
	version   int
	started   bool
	startedAt *timestamp
	updatedAt *timestamp
	ticker    *time.Ticker
	channel   chan bool
}

// TODO: Add proper comment
var TP Parser

// TODO: Add proper comment
type UtilsChecker interface {
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
) (success bool) {
	success = false
	fmt.Printf("AddOrUpdateMatch path: %s\n", path)
	// get file hash first
	var hash []byte
	var err error
	hash, err = HashFile(path)
	if err != nil {
		return
	}

	// check if match was parsed before
	match, err := database.DB.GetMatchByHash(hash)
	currentVersion := GetParserVersion()

	// match version didn't change, no need to parse
	if match != nil && ((match.ParserVersion == currentVersion) && (match.FilePath == path)) {
		fmt.Printf("Match version didn't change\n")
		return
	}

	mapname := ""
	score1 := 0
	score2 := 0
	date := time.Now()
	participants := make(map[uint64]map[string]interface{})
	mapname, date, score1, score2, participants = ParseMatch(path)
	fmt.Printf("Parsing %s - %d : %d\n", mapname, score1, score2)

	if err == gorm.ErrRecordNotFound {
		match, err = database.DB.CreateMatch(
			currentVersion,
			hash,
			path,
			date,
			mapname,
			score1,
			score2,
			participants,
		)
		success = true
		fmt.Printf("match added to db\n")
	} else {
		match, err = database.DB.UpdateMatch(
			currentVersion,
			hash,
			path,
			date,
			mapname,
			score1,
			score2,
			participants,
		)
		success = true
		fmt.Printf("match updated on db\n")
	}
	return
}

func AddOrUpdatePlayer(player map[string]interface{}) {
	// check if steamid was seen before
	var steamid uint64
	steamid = player["steamid"].(uint64)
	_, err := database.DB.GetPlayerBySteamID(steamid)
	if err == gorm.ErrRecordNotFound {
		_, err = database.DB.CreatePlayer(steamid)
		fmt.Printf("player added to db\n")
	} else {
		fmt.Printf("player updated on db\n")
	}
}

func StartParser(w *astilectron.Window, path string) (result bool, err error) {
	TP.version = 1 // bump this every time you update parsing rules to force all matches to reparse
	fmt.Println("debug10 - Parser attrs")
	result = true
	parserPeriod := time.Second * 60
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

func ParseMatch(path string) (
	mapname string,
	date time.Time,
	score1 int,
	score2 int,
	players map[uint64]map[string]interface{},
) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	p := dem.NewParser(f)
	defer p.Close()
	header, _ := p.ParseHeader()
	mapname = header.MapName
	score1 = 0
	score2 = 0
	stat, _ := f.Stat()
	date = stat.ModTime()

	err = p.ParseToEnd()
	if err != nil {
		panic(err)
	}
	participants := p.GameState().Participants().AllByUserID()
	players = make(map[uint64]map[string]interface{})
	for _, s := range participants {
		if s.SteamID64 != uint64(0) {
			player := make(map[string]interface{})
			player["nickname"] = s.Name
			player["steamid"] = s.SteamID64
			AddOrUpdatePlayer(player)
			players[s.SteamID64] = player
		}
	}
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
			success := AddOrUpdateMatch(filePath)
			//Send message to client after parsing match
			if success {
				bootstrap.SendMessage(w, "parsed", nil)
			}
			sem <- empty{}
		}(filePath)
	}
	for i := 0; i < N; i += 1 {
		<-sem
	}
}

func GetPlayerList() (players map[string]map[string]interface{}) {
	fmt.Printf("debug13 - GetPlayerList")
	players = database.DB.GetPlayers()
	return
}
