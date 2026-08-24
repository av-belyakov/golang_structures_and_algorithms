package sqlite3updater

import (
	"context"
	"fmt"
	"golang_structures_and_algorithms/someusefulsubmodules/sqlite3databaseupdate/addnewcolums/internal/sqlite3api"
	"net/url"
	"os"

	"gopkg.in/yaml.v3"
)

func NewApp(s Settings) (*ApplicationSettings, error) {
	settings := &ApplicationSettings{
		Organizations: map[string]string{},
		Sqlite3: Sqlite3Settings{
			FileDbPath: s.Sqlite3DbPath,
		},
	}

	urlBase, err := url.Parse("http://" + s.MispHost)
	if err != nil {
		return settings, err
	}
	settings.MISP.UrlBase = urlBase

	cfgFile, err := os.ReadFile(s.ConfigPath)
	if err != nil {
		return settings, err
	}

	orgs := struct {
		Organization []struct {
			OrgName    string `yaml:"orgName"`
			SourceName string `yaml:"sourceName"`
		} `yaml:"ORGANIZATIONS"`
	}{}

	if err = yaml.Unmarshal(cfgFile, &orgs); err != nil {
		return settings, err
	}

	for _, org := range orgs.Organization {
		settings.Organizations[org.OrgName] = org.SourceName
	}

	return settings, nil
}

func (app *ApplicationSettings) Start(ctx context.Context) error {
	fmt.Println("UpdateSqlite3 module start")
	fmt.Println("Инициализация модуля взаимодействия с БД Sqlite3")

	sqlDb := sqlite3api.New(app.Sqlite3.FileDbPath)
	if err := sqlDb.Start(ctx); err != nil {
		return err
	}
	if err := sqlDb.Start(ctx); err != nil {
		return err
	}

	fmt.Println("Получаем список событий MISP")
	listEvents, err := sqlDb.GetAllEventsId(ctx)
	if err != nil {
		return err
	}

	fmt.Println("Всего событий MISP найдено:", len(listEvents))
	var num int
	for eventId, caseId := range listEvents {
		if num == 11 {
			break
		}

		num++
		fmt.Printf("%d. eventId:%d, caseId:%d\n", num, eventId, caseId)
	}

	return nil
}
