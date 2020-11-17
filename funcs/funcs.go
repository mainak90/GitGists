package funcs

import (
	"encoding/json"
	"github.com/levigross/grequests"
	"github.com/mainak90/GitGists/models"
	"io/ioutil"
	"log"
	"os"
)

var GITHUB_TOKEN = os.Getenv("GITHUB_TOKEN")
var requestOptions = &grequests.RequestOptions{Auth: []string{GITHUB_TOKEN, "x-oauth-basic"}}

func GetStats(url string) *grequests.Response {
	resp, err := grequests.Get(url, requestOptions)
	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}
	return resp
}

func CreateGithubGist(url string, args []string) *grequests.Response {
	description := args[0]
	var fileContents = make(map[string]models.File)
	for i := 1; i < len(args); i++ {
		dat, err := ioutil.ReadFile(args[i])
		if err != nil {
			log.Println("Error is: ", err)
			return nil
		}
		var file models.File
		file.Content = string(dat)
		fileContents[args[i]] = file
	}
	var gist = models.Gist{Description: description, Public: true, Files: fileContents}
	var postbody, _ = json.Marshal(gist)
	var requestoptions_copy = requestOptions
	requestoptions_copy.JSON = string(postbody)
	resp, err := grequests.Post(url, requestoptions_copy)
	if err != nil {
		log.Println("Error encountered: ", err)
	}
	return resp
}
