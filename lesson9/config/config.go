package config

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
)

var (
	dbUrl      = flag.String("dbUrl", "", "should be like: postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable")
	jaegerUrl  = flag.String("jaegerUrl", "", "should be like: http://jaeger:16686")
	sentryUrl  = flag.String("sentryUrl", "", "should be like: http://sentry:9000")
	kafkaUrl   = flag.String("kafkaBroker", "", "should be like: kafka:9092")
	port       = os.Getenv("MY_VAR_PORT")
	someAppId  = os.Getenv("MY_VAR_APP_ID")
	someAppKey = os.Getenv("MY_VAR_APP_KEY")
	configFile = flag.String("myConfigFile", "", "should contain path to the file")
)

type MyConfigInternalStructure struct {
	X int `json:"x"`
	Y int `json:"y"`
	Z int `json:"z"`
}

type MyConfig struct {
	A        int                          `json:"a"`
	B        string                       `json:"b"`
	Internal []*MyConfigInternalStructure `json:"internal"`
}

func ParseAndValidate() {
	flag.Parse()
	if *dbUrl != "" {
		_, err := url.ParseRequestURI(*dbUrl)
		if err != nil {
			panic(err)
		}
		fmt.Println(*dbUrl)
	}
	if *jaegerUrl != "" {
		_, err := url.ParseRequestURI(*jaegerUrl)
		if err != nil {
			panic(err)
		}
		fmt.Println(*jaegerUrl)
	}
	if *sentryUrl != "" {
		_, err := url.ParseRequestURI(*sentryUrl)
		if err != nil {
			panic(err)
		}
		fmt.Println(*sentryUrl)
	}
	if *kafkaUrl != "" {
		_, err := url.ParseRequestURI(*kafkaUrl)
		if err != nil {
			panic(err)
		}
		fmt.Println(*kafkaUrl)
	}

	if port != "" {
		fmt.Println(port)
	}

	if someAppId != "" {
		fmt.Println(someAppId)
	}

	if someAppKey != "" {
		fmt.Println(someAppKey)
	}

	if *configFile != "" {
		/*JSON config file should be like:
		{
		        "a": 5,
		        "b": "asd",
		        "internal": [
		                {
		                        "x": 1,
		                        "y": 2,
		                        "z": 3
		                },
		                {
		                        "x": 4,
		                        "y": 5,
		                        "z": 6
		                }
		        ]
		 }
		*/
		myFile, err := os.ReadFile(*configFile)
		if err != nil {
			log.Fatal(errors.New("config file not opened"))
		}
		myConf := new(MyConfig)
		err = json.Unmarshal(myFile, &myConf)
		if err != nil {
			log.Fatal(errors.New("config file didn't contain valid json"))
		}
		data, err := json.MarshalIndent(myConf, " ", "\t")
		if err != nil {
			log.Fatal(errors.New("config can't be writen"))
		}
		fmt.Println(string(data))
	}
}
