// Program runner.  Configures routes and starts the webserver.
package main

import (
	"fmt"
	"github.com/codegangsta/martini"
	"log"
	"net/http"
)

func main() {
	m := martini.Classic()
	m.Use(martini.Static("../../app"))

	m.Post("/authorize", authorize)
	m.Get("/authorized", authorized)
	m.Get("/device", getDevices)
	m.Get("/alarm/:deviceId", getAlarms)
	m.Post("/alarm", setAlarm)
	m.Delete("/alarm/:deviceId/:alarmId", deleteAlarm)

	log.Println("running on port", configuration.App_port)
	http.ListenAndServe(fmt.Sprintf(":%v", configuration.App_port), m)
}
