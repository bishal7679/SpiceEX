package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"runtime/debug"

	"github.com/bishal7679/SpiceEx/internal/config"
	"github.com/echa/code/iata"
)

var app *config.AppConfig

// NewHelpers sets up app config for helpers
func NewHelpers(a *config.AppConfig) {
	app = a
}

func ClientError(w http.ResponseWriter, status int) {
	app.InfoLog.Println("Client error with status of", status)
	http.Error(w, http.StatusText(status), status)
}

func ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func CreateJSON(data map[iata.AirportCode]iata.Airport) {
	jsondata, err := json.MarshalIndent(data,"","  ")
	if err != nil {
		app.ErrorLog.Println("Error encoding json :(", err)
		return
	}
	file, err := os.Create("airport.json")
	if err != nil {
		app.ErrorLog.Println("Error creating json file :(", err)
		return
	}
	defer file.Close()

	_, err = file.Write(jsondata)
	if err != nil {
		app.ErrorLog.Println("Error writing json file :(", err)
		return
	}
}
