package sqliteinteraction

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"runtime"

	"github.com/av-belyakov/golang_structures_and_algorithms/interfaces"
)

// New инициализирует новый модуль взаимодействия с API TheHive
// при инициализации возращается канал для взаимодействия с модулем, все
// запросы к модулю выполняются через данный канал
func New(path string, logging interfaces.Logger) (*SqliteApiOptions, error) {
	options := SqliteApiOptions{logger: logging}

	if path == "" {
		return &options, errors.New("the path to the database file should not be empty")
	}

	options.path = path

	return &options, nil
}

func (opts *SqliteApiOptions) Start(ctx context.Context) (chan<- ChanApiSqlite, error) {
	chanListene := make(chan ChanApiSqlite)

	sqldb, err := sql.Open("sqlite3", opts.path)
	if err != nil {
		return chanListene, err
	}

	if sqldb.Ping() != nil {
		return chanListene, err
	}

	opts.db = sqldb

	go func() {
		for {
			select {
			case <-ctx.Done():
				return

			case msg := <-chanListene:
				opts.route(routeSettings{
					command:      msg.Command,
					taskId:       msg.TaskID,
					service:      msg.Service,
					data:         msg.Data,
					chanResponse: msg.ChanResponse,
				})
			}
		}
	}()

	return chanListene, nil
}

// Route маршрутизатор обработки запросов
func (opts *SqliteApiOptions) route(settings routeSettings) {
	if settings.taskId == "" || settings.command == "" {
		_, f, l, _ := runtime.Caller(0)
		opts.logger.Send("error", fmt.Sprintf(" 'the sql query cannot be processed, the command and the task ID must not be empty' %s:%d", f, l-1))

		return
	}

	switch settings.command {
	case "insert section tags":
		go opts.handlerSectionInsertTags(settings.taskId, settings.service, settings.data, settings.chanResponse)
	case "insert section creater":
		go opts.handlerSectionInsertCreater(settings.taskId, settings.service, settings.data, settings.chanResponse)
	case "select section tags":
		go opts.handlerSectionSelectTags(settings.taskId, settings.service, settings.data, settings.chanResponse)
	case "select section creater":
		go opts.handlerSectionSelectCreater(settings.taskId, settings.service, settings.data, settings.chanResponse)
	}
}

// handlerSectionInsertTags обработчик секции добавления информации о тегах
func (opts *SqliteApiOptions) handlerSectionInsertTags(taskId, service string, data []byte, chanResponse chan<- ChanOutputApiSqlite) {
	response := ChanOutputApiSqlite{TaskID: taskId}

	stmt := opts.prepareTableExecutedCommands()

	//порядок: id, service, binary_data
	result, err := stmt.Exec(taskId, service, data)
	if err != nil {
		_, f, l, _ := runtime.Caller(0)
		opts.logger.Send("error", fmt.Sprintf(" '%s' %s:%d", err.Error(), f, l-1))

		chanResponse <- response

		return
	}

	fmt.Println("func 'handlerSectionInsertTags' RESPONSE:", result)

	//
	// Вообще лучше сделать что бы тип chanResponse
	// соответствовал определенному интерфейсу и далее описать этот
	// интерфейс в commoninterfaces.interfaces.go
	// так будет полее гибко
	//

	response.Status = true
	chanResponse <- response
}

// handlerSectionInsertCreater обработчик секции добавления информации о создаваемых, новых данных
func (apisqlite *SqliteApiOptions) handlerSectionInsertCreater(taskId, service string, data []byte, chanResponse chan<- ChanOutputApiSqlite) {

}

// handlerSectionSelectTags обработчик секции запроса информации о тегах
func (apisqlite *SqliteApiOptions) handlerSectionSelectTags(taskId, service string, data []byte, chanResponse chan<- ChanOutputApiSqlite) {

}

// handlerSectionSelectCreater обработчик секции запроса информации о создаваемых, новых данных
func (apisqlite *SqliteApiOptions) handlerSectionSelectCreater(taskId, service string, data []byte, chanResponse chan<- ChanOutputApiSqlite) {

}

func (opts *SqliteApiOptions) prepareTableExecutedCommands() *sql.Stmt {
	stmt, err := opts.db.Prepare("INSERT INTO table_executed_commands (id, service, command, name, description) values(?,?,?,?,?)")
	if err != nil {
		_, f, l, _ := runtime.Caller(0)
		opts.logger.Send("error", fmt.Sprintf(" '%s' %s:%d", err.Error(), f, l-1))

		return nil
	}

	return stmt
}
