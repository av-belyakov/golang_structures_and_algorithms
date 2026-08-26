package sqlite3updater

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"

	"sqlite3databaseupdate/internal/mispapi"
	"sqlite3databaseupdate/internal/sqlite3api"
)

const (
	tableName        = "placeholder_misp"
	logFile          = "process.log"
	envMispTokenName = "GO_SQLITE3DATABASEUPDATE_MISPTOKEN"
	hostMisp         = "misp-center.cloud.gcm"
	portMisp         = 80
)

func NewApp(s Settings) (*ApplicationSettings, error) {
	settings := &ApplicationSettings{
		Organizations: map[string]string{},
		Sqlite3: Sqlite3Settings{
			FileDbPath: s.Sqlite3DbPath,
		},
	}

	// инициируем чтение файла .env
	if err := godotenv.Load(".env"); err != nil {
		return settings, err
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

	//проверяем наличие файла логов
	_, err := os.Stat(logFile)
	if err == nil {
		if err := os.Remove(logFile); err != nil {
			fmt.Println("error:", err)
		}
	}

	sqlDb := sqlite3api.New(app.Sqlite3.FileDbPath)
	if err := sqlDb.Start(ctx); err != nil {
		return err
	}

	fmt.Println("Получаем список полей таблицы для того что бы понять нужно ли добавлять новые колонки")
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

	if !columnTimeIsExist {
		fmt.Println("Добавляем колонку 'datetime'")
		if err := sqlDb.AlterTableAddColumn(ctx, tableName, "datetime", "TEXT"); err != nil {
			return err
		}
	}

	if !columnSourceIsExist {
		fmt.Println("Добавляем колонку 'source'")
		if err := sqlDb.AlterTableAddColumn(ctx, tableName, "source", "VARCHAR(8)"); err != nil {
			return err
		}
	}

	fmt.Println("Получаем список событий MISP")
	listEvents, err := sqlDb.GetAllEventsId(ctx)
	if err != nil {
		return err
	}

	fmt.Println("Выполняем поиск событий в MISP")
	request, err := mispapi.NewRequest(envMispTokenName, hostMisp, portMisp)
	if err != nil {
		return err
	}

	eventIdJsonError := []int64{}
	notFoundEventId := map[int64]int64{}
	var foundEventId int
	for eventId, caseId := range listEvents {
		statusCode, raw, err := request.GetEvent(ctx, time.Second*10, strconv.Itoa(int(eventId)))
		if err != nil {
			fmt.Println("error:", err)

			continue
		}

		if statusCode != http.StatusOK {
			fmt.Printf("status code %d for eventId:%d\n", statusCode, eventId)
			notFoundEventId[eventId] = caseId

			continue
		}

		event := mispapi.MispElement{}
		if err := json.Unmarshal(raw, &event); err != nil {
			fmt.Printf("error: %s (eventId:'%d')\n", err.Error(), eventId)

			eventIdJsonError = append(eventIdJsonError, eventId)

			continue
		}

		//fmt.Printf("Event:%+v\nOrganization name:%s\n", event.Event, event.Event.Org.Name)

		source, ok := app.Organizations[event.Event.Org.Name]
		if !ok {
			fmt.Printf("organization name '%s' not found\n", event.Event.Org.Name)

			continue
		}

		if err := sqlDb.UpdateColumnSource(ctx, tableName, eventId, source); err != nil {
			fmt.Println("error:", err)

			continue
		}

		foundEventId++
	}

	finalMessage := fmt.Sprintf(`
	Всего событий MISP в БД:%d
	Событий успешно обработанно:%d
	События которые не были найдены в MISP:%d
	Событий которые вызвали ошибки при конвертации из JSON:%d
	`,
		len(listEvents),
		foundEventId,
		len(notFoundEventId),
		len(eventIdJsonError),
	)

	fmt.Println(finalMessage)
	fmt.Printf("Подробно можно посмотреть в лог-файле '%s'\n", logFile)

	f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	strJsonError := strings.Builder{}
	fmt.Fprintf(&strJsonError, "События которые вызвали ошибки при конвертации из JSON:\n")
	for k, v := range eventIdJsonError {
		fmt.Fprintf(&strJsonError, "%d. %d\n", k+1, v)
	}

	strNotFoundEventId := strings.Builder{}
	fmt.Fprintf(&strNotFoundEventId, "События которые не были найдены в MISP:\n")
	count := 1
	for eventId, caseId := range notFoundEventId {
		fmt.Fprintf(&strNotFoundEventId, "%d. eventId:'%d', caseId:'%d'\n", count, eventId, caseId)
		count++
	}

	fmt.Fprintf(f, "%s\n\n%s\n\n%s\n", finalMessage, strJsonError.String(), strNotFoundEventId.String())

	return nil
}
