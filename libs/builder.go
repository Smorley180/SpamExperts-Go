package libs

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/spf13/viper"
)

func Buildrequest(apiPath string) (response string) {
	// Builds the http request from config and passed in apiPath
	builder := strings.Builder{}

	builder.WriteString("https://")
	builder.WriteString(viper.GetString("username"))
	builder.WriteString(":")
	builder.WriteString(viper.GetString("password"))
	builder.WriteString("@")
	builder.WriteString(viper.GetString("hostname"))
	builder.WriteString(apiPath)
	builder.WriteString("/format/plain")

	return builder.String()
}

func Makerequest(requestPath string) (response string) {
	// Makes http request and handles error logging
	resp, err := http.Get(requestPath)

	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)
	return sb
}

func Arrayformat(items []string) (response string) {
	// Build Json formatted style array from string array and return as string
	toReturn := "["
	for i := 0; i < len(items); i++ {
		toReturn += "\"" + items[i] + "\", "
	}
	toReturn = strings.Trim(toReturn, ", ")
	toReturn += "]"
	return toReturn
}
