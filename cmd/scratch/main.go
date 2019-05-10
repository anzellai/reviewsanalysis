package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"time"

	"github.com/anzellai/reviewsanalysis/pkg/scratch"
)

func main() {
	api := scratch.NewAPI()
	api.Limit = 10
	{
		err := api.GetTop10ProjectComments()
		if err != nil {
			fmt.Errorf("error fetching project comments: %v", err)
			return
		}
	}
	{
		err := api.GetAnalysis()
		if err != nil {
			fmt.Errorf("error analysing reviews: %v", err)
			return
		}
	}
	results, err := json.MarshalIndent(api, "", "  ")
	if err != nil {
		fmt.Errorf("error marshalling api: %v", err)
		return
	}
	t := time.Now()
	basepath := "./data"
	filename := path.Join(basepath, fmt.Sprintf("scratch_top10_project_comments_%d-%02d-%02dT%02d%02d.json",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(),
	))
	if _, err := os.Stat(basepath); os.IsNotExist(err) {
		os.Mkdir(basepath, 0700)
	}
	err = ioutil.WriteFile(filename, results, 0644)
	if err != nil {
		fmt.Errorf("error writing results on disk: %v", err)
		return
	}
	fmt.Println(string(results))
}
