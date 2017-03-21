package gogopro

import (
	"log"
	"net/http"
	"os"
)

type BasicAuth struct {
	Password string
}

type GoPro struct {
	Power *Power
}

var (
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func (gopro *GoPro) Init() (*GoPro, error) {
	client := &http.Client{
		CheckRedirect: nil,
	}
	gopro.Power.APIRequester.Client = client
	return gopro, nil
}

func (gopro *GoPro) initLogger() {
	Info = log.New(os.Stdout,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(os.Stdout,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(os.Stderr,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func CreateGoPro(Ipaddr string, auth ...interface{}) *GoPro {
	gopro := &GoPro{}
	APIRequester := &APIRequester{URL: "http://" + Ipaddr}

	if len(auth) == 1 {
		APIRequester.BasicAuth = &BasicAuth{Password: auth[0].(string)}
	} else {
		APIRequester.BasicAuth = &BasicAuth{Password: ""}
	}

	power := CreatePower(APIRequester).Init()
	gopro.Power = power

	return gopro
}
