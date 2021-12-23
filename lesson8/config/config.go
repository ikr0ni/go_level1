package config

import (
	"flag"
	"fmt"
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
)

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
}
