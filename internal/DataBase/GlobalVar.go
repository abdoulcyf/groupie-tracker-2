package DataBase

import (
	"log/slog"
	"os"
)

// -------------------------------------
const mainApiUrl = `https://groupietrackers.herokuapp.com/api`

// ------------------------------------------------
// the structure of mainApiUrl.json
type UrlTable struct {
	UrlArtists   string `json:"artists"`
	UrlLocations string `json:"locations"`
	UrlDates     string `json:"dates"`
	UrlRelation  string `json:"relation"`
}

var myUrlTable UrlTable

// --------------------------------
// define the structure of Artist type for TableOfArtists in database
type Artist struct {
	ID        int
	Name      string
	ImageUrl  string
	CreationDate int
	FirstAlbum   string
	Members   []string
	Locations []string // Add Locations field
	Dates     []string
	Relation  map[string][]string
}


// ----------------------------------
// the structure of a database
type ArtistsDataBase struct {
	TableOfArtists   map[int]Artist
	TableOfMembers   map[int][]string
	TableOfLocations map[int][]string
	TableOfDates     map[int][]string
	TableOfRelation  map[int]map[string][]string
}

// ----------------------------
// an instance of database
var Mydb ArtistsDataBase

// -----------------logger-----------------------
// setup logger
var (
	logHandlerOpts = &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}

	logHandler = slog.NewTextHandler(os.Stderr, logHandlerOpts)
	logger     = slog.New(logHandler)
	LoggerExt  = slog.New(logHandler)
	errMsg     string
	logMsg     string
)

// ----------------------------------------
