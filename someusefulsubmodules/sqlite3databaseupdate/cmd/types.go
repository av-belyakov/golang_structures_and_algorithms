package sqlite3updater

import "net/url"

type Settings struct {
	ConfigPath    string
	Sqlite3DbPath string
	MispHost      string
}

type ApplicationSettings struct {
	Organizations map[string]string
	Sqlite3       Sqlite3Settings
	MISP          MISPSettings
}

type Sqlite3Settings struct {
	FileDbPath string
}

type MISPSettings struct {
	UrlBase *url.URL
	AuthKey string
	Verify  bool
}
