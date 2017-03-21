package gopro

import (
	"log"
	"net/http"
	"os"
)

type BasicAuth struct {
	Password string
}

type GoPro struct {
	APIRequester *APIRequester
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
	gopro.APIRequester.Client = client
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
	gopro.APIRequester = &APIRequester{URL: "http://" + Ipaddr}

	if len(auth) == 1 {
		gopro.APIRequester.BasicAuth = &BasicAuth{Password: auth[0].(string)}
	} else {
		gopro.APIRequester.BasicAuth = &BasicAuth{Password: ""}
	}
	return gopro
}

func (gopro *GoPro) Status() (*http.Response, error) {
	resp, err := gopro.APIRequester.getWithPort("", 8080)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (gopro *GoPro) GetPowerStatus() (*http.Response, error) {
	resp, err := gopro.APIRequester.get("/bacpac/se")
	if err != nil {
		return nil, err
	}
	return resp, nil
}
