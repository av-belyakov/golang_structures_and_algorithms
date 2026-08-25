package sqlite3updater

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"

	"gopkg.in/yaml.v3"

	"sqlite3databaseupdate/internal/mispapi"
	"sqlite3databaseupdate/internal/sqlite3api"
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

	tableName := "placeholder_misp"

	sqlDb := sqlite3api.New(app.Sqlite3.FileDbPath)
	if err := sqlDb.Start(ctx); err != nil {
		return err
	}

	fmt.Println("Получаем списко полей таблицы")
	columnsInfo, err := sqlDb.GetTableColumns(ctx, tableName)
	if err != nil {
		return err
	}

	var columnTimeIsExist, columnSourceIsExist bool
	for _, columnInfo := range columnsInfo {
		if columnInfo.Name == "datetime" {
			columnTimeIsExist = true
		}

		if columnInfo.Name == "source" {
			columnSourceIsExist = true
		}
	}

	fmt.Println("Добавляем новые колонки в таблицу, если их нет")
	if !columnTimeIsExist {
		if err := sqlDb.AlterTableAddColumn(ctx, tableName, "datetime", "TEXT"); err != nil {
			return err
		}
	}

	if !columnSourceIsExist {
		if err := sqlDb.AlterTableAddColumn(ctx, tableName, "source", "VARCHAR(8)"); err != nil {
			return err
		}
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

	fmt.Println("Тестово получаем одно событие из MISP")
	request, err := mispapi.NewRequest("misp-center.cloud.gcm", 80)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	/*
		сделать всё тоже самое что ниже но в цикле для всех event id из listEvents и на новом бекапе
	*/

	statusCode, raw, err := request.GetEvent(ctx, "135338")
	if err != nil {
		return err
	}

	fmt.Printf("Request status:%d\n", statusCode)
	if statusCode != http.StatusOK {
		return fmt.Errorf("error, status code %d", statusCode)
	}

	event := mispapi.MispElement{}
	if err := json.Unmarshal(raw, &event); err != nil {
		return err
	}

	fmt.Printf("Event:%+v\nOrganization name:%s\n", event.Event, event.Event.Org.Name)

	if source, ok := app.Organizations[event.Event.Org.Name]; ok {
		if err := sqlDb.UpdateColumnSource(ctx, tableName, 10936, source); err != nil {
			return err
		}
	}

	return nil
}
