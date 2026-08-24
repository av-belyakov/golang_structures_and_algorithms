package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	sqlite3updater "golang_structures_and_algorithms/someusefulsubmodules/sqlite3databaseupdate/addnewcolums/cmd"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill, syscall.SIGINT)
	defer stop()

	go func() {
		<-ctx.Done()
		fmt.Println("Sqlite3Updater module is stop")

		stop()
	}()

	app, err := sqlite3updater.NewApp(sqlite3updater.Settings{
		ConfigPath:    "./configs/organizations.yml",
		Sqlite3DbPath: "./backup/backup.db",
		MispHost:      "misp-center.cloud.gcm",
	})
	if err != nil {
		log.Fatalln(err)
	}

	if err := app.Start(ctx); err != nil {
		log.Fatalln(err)
	}
}
