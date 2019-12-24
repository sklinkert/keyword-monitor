package main

import (
	"github.com/lfritz/env"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	customsearch "google.golang.org/api/customsearch/v1"
	"io/ioutil"
	"strings"
)

var config struct {
	keywords []string
	domain   string
	id       string
}

// Result is a search result
type Result struct {
	Position int64
	Result   *customsearch.Result
}

func doSearch(search *customsearch.CseListCall) (result Result) {
	start := int64(1)
	for start < 300 {
		search.Start(start)
		call, err := search.Do()
		if err != nil {
			log.Fatal(err)
		}

		position, csResult := findDomain(call.Items, start)
		if csResult != nil {
			result = Result{
				Position: position,
				Result:   csResult,
			}
			return
		}

		// No more search results
		if call.SearchInformation.TotalResults < start {
			return
		}
		start = start + 10
	}
	return
}

func findDomain(results []*customsearch.Result, start int64) (position int64, result *customsearch.Result) {
	for index, r := range results {
		if strings.Contains(r.Link, config.domain) {
			return int64(index) + start, r
		}
	}
	return 0, nil
}

func main() {
	e := env.New()
	e.String("ID", &config.id, "id of CSE")
	e.List("KEYWORDS", &config.keywords, ",", "keywords to check for")
	e.String("DOMAIN", &config.domain, "Domain to check for")
	if err := e.Load(); err != nil {
		log.WithError(err).Fatal("Parameters are missing")
	}

	const searchKeyFile = "search-key.json"
	data, err := ioutil.ReadFile(searchKeyFile)
	if err != nil {
		log.WithError(err).Fatal("Reading %q failed", searchKeyFile)
	}

	//Get the config from the json key file with the correct scope
	conf, err := google.JWTConfigFromJSON(data, "https://www.googleapis.com/auth/cse")
	if err != nil {
		log.WithError(err).Fatal("google.JWTConfigFromJSON failed")
	}

	// Initiate an http.Client. The following GET request will be
	// authorized and authenticated on the behalf of
	// your service account.
	client := conf.Client(oauth2.NoContext)
	cseService, err := customsearch.New(client)
	if err != nil {
		log.WithError(err).Fatal("customsearch.New() failed")
	}

	for _, query := range config.keywords {
		search := cseService.Cse.List(query)
		search.Cx(config.id)

		result := doSearch(search)
		clog := log.WithFields(log.Fields{
			"Keyword":       query,
			"TotalKeywords": len(config.keywords),
			"Domain":        config.domain,
		})
		if result.Position == 0 {
			clog.Info("Not found")
			continue
		}

		clog = clog.WithFields(log.Fields{
			"Position": result.Position,
			"Title":    result.Result.Title,
			"URL":      result.Result.Link,
		})
		clog.Info("Found")
	}
}
