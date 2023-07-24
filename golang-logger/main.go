package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	HelloCounter = 1
	logger       zerolog.Logger
)

func main() {
	servicename := "service-name"
	if sn := os.Getenv("SERVICE_NAME"); sn != "" {
		servicename = sn
	}

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Build a logger with default fields
	logger = log.With().Str("service", servicename).Str("node", hostname).Logger()

	// Counter
	number := 1

	for number <= 604800 {

		// call function
		Hello()

		// calling sleep method with 1 second
		time.Sleep(1 * time.Second)

		// increment counter
		number++
	}
}

func Name() string {
	rand.Seed(time.Now().Unix())
	names := []string{
		"ruan",
		"stefan",
		"frank",
		"peter",
		"samantha",
		"susan",
		"jeffrey",
		"adam",
		"juan",
		"james",
		"michelle",
	}
	n := rand.Int() % len(names)
	return names[n]
}

func Country() string {
	rand.Seed(time.Now().Unix())
	countries := []string{
		"south africa",
		"mexico",
		"new zealand",
		"australia",
		"germany",
		"italy",
		"spain",
		"france",
	}
	c := rand.Int() % len(countries)
	return countries[c]
}

func Age() int {
	rand.Seed(time.Now().Unix())
	ages := []int{
		24, 28, 29, 30, 34, 36, 38,
	}
	a := rand.Int() % len(ages)
	return ages[a]
}

func JobTitle() string {
	rand.Seed(time.Now().Unix())
	jobtitles := []string{
		"software engineer",
		"database administrator",
		"teacher",
		"mechanical engineer",
		"photographer",
		"designer",
		"paramedic",
		"auto mechanic",
		"writer",
		"plumber",
		"receptionist",
		"carpenter",
		"truck driver",
		"lawyer",
	}
	c := rand.Int() % len(jobtitles)
	return jobtitles[c]
}

func Message() string {
	rand.Seed(time.Now().Unix())
	messages := []string{
		"a exception occured on /api",
		"inserted entry into database successfully",
		"high latency detected on /api",
		"high cpu utilization on server",
		"lookup returned successfully to the client",
		"cache has been cleared",
		"database has been initialized",
		"unable to connect to database",
	}
	m := rand.Int() % len(messages)
	return messages[m]
}

func RandomString() int {
	rand.Seed(time.Now().UnixNano())
	upper := 9000
	lower := 1000
	number := rand.Intn(upper) + lower
	return number
}

func Hello() {
	country := Country()
	name := Name()
	age := Age()
	message := Message()
	jobtitle := JobTitle()
	randomstring := fmt.Sprintf("%d-%d-%d-%d", RandomString(), RandomString(), RandomString(), RandomString())
	// Log with service name
	logger.Info().Int("count", HelloCounter).Str("name", name).Str("country", country).Int("age", age).Str("jobtitle", jobtitle).Str("creditcard", randomstring).Msg(message)
	HelloCounter++
}
