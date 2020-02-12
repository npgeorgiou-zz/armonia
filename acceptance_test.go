package logger

import (
	"encoding/json"
	util "lendme/golang-seeder-backend/entry_points"
	lib_http "lendme/golang-seeder-backend/lib"
	http_adapter "lendme/golang-seeder-backend/src/adapters/http"
	"lendme/golang-seeder-backend/src/core/entities"
	"lendme/golang-seeder-backend/src/di"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"
)

func setup() di.Di {
	os.Setenv("DB_DRIVER", "sqlite3")
	os.Setenv("DB_URI", ":memory:")
	os.Setenv("DB_NAME", "")

	di := util.CreateDefaultDi()
	di.GetPersistence().MigrateSchema()

	return di
}

func teardown() {
}

func TestRootUri(t *testing.T) {
	app := setup()

	response := hitAPI("GET", "/", "", app, t)

	assertSame(http.StatusOK, response.StatusCode, t)
	assertSame("Hi there, I log stuff!", response.ReadBody(), t)

	teardown()
}

func TestCreateUiLog(t *testing.T) {
	di := setup()

	input := http_adapter.CreateLogInput{
		Project: "wl",
		Level:   "INFO",
		Message: "You made a bug!",
		SentAt: 1577836800,
	}

	response := hitAPI("POST", "/create", jsonFrom(input), di, t)

	// Assert response.
	assertSame(http.StatusOK, response.StatusCode, t)

	// Assert db
	log, err := di.GetPersistence().FindFirstLog()

	assertNil(err, t)
	assertSame("INFO", log.Level, t)
	assertSame("wl", log.Project, t)
	assertSame("You made a bug!", log.Message, t)
	assertTrue(log.SentAt.Equal(time.Unix(1577836800, 0)), t)

	teardown()
}

func TestDeleteUiLog(t *testing.T) {
	di := setup()

	log := entities.Log{
		Project: "wl",
		Level:   "INFO",
		Message: "You made a bug!",
	}

	di.GetPersistence().SaveLog(&log)

	input := map[string]uint{"Id": log.ID}
	hitAPI("POST", "/delete", jsonFrom(input), di, t)

	log, err := di.GetPersistence().FindFirstLog()
	assertSame("Entity not found", err.Error(), t)

	teardown()
}

func hitAPI(method, url string, body string, di di.Di, t *testing.T) *lib_http.Response {
	router := util.CreateRouter(di)
	server := httptest.NewServer(router)
	defer server.Close()

	request, err := http.NewRequest(method, server.URL+url, strings.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	client := server.Client()
	response, err := client.Do(request)
	if err != nil {
		t.Fatal(err)
	}

	return (*lib_http.Response)(response)
}

func assertSame(expected interface{}, actual interface{}, t *testing.T) {
	if expected != actual {
		t.Errorf("Expected %v got %v", expected, actual)
	}
}

func assertTrue(actual bool, t *testing.T) {
	if actual != true {
		t.Errorf("Expected true got false")
	}
}

func assertNil(actual interface{}, t *testing.T) {
	if actual != nil {
		t.Errorf("Expected nil got %q", actual)
	}
}

func jsonFrom(any interface{}) string{
	body, err := json.Marshal(any)
	if err != nil {
		panic(err)
	}

	return string(body)
}