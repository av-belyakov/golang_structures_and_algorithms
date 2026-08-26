package misp

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	"sqlite3databaseupdate/internal/mispapi"
)

func TestMain(m *testing.M) {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatalln(err)
	}

	m.Run()
}

func TestGetEvent(t *testing.T) {
	var (
		eventId string = "133513"

		envMispTokenName string = "GO_SQLITE3DATABASEUPDATE_MISPTOKEN"
		hostMisp         string = "misp-center.cloud.gcm"
		portMisp         int    = 80
	)

	request, err := mispapi.NewRequest(envMispTokenName, hostMisp, portMisp)
	if err != nil {
		log.Fatalln(err)
	}

	statusCode, raw, err := request.GetEvent(t.Context(), time.Second*10, eventId)
	assert.NoError(t, err)
	assert.Equal(t, statusCode, http.StatusOK)

	fmt.Println("RAW:", string(raw))

	event := mispapi.MispElement{}
	assert.NoError(t, json.Unmarshal(raw, &event))
}
