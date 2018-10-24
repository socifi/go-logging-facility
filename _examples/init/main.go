package main

import (
	"encoding/json"
	"fmt"
	"github.com/socifi/go-logging-facility/handlers/init"
	_ "github.com/socifi/go-logging-facility/hook/dsn"
	"io/ioutil"
	"os"
)

type autoGenerated struct {
	/* Your config here... */
	Logging loginit.LogConfig `json:"logging"`
}

func main() {
	raw, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var c autoGenerated
	json.Unmarshal(raw, &c)
	fmt.Printf("%#v\n", c)
	logger := loginit.Init(c.Logging)
	logger.WithField("dsn", "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable").Info("trial")
}
