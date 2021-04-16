package funcs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/levigross/grequests"
	"github.com/mainak90/GitGists/models"
	"net/http"
	"io/ioutil"
	"log"
	"os"
)

var GITHUB_TOKEN = os.Getenv("GITHUB_TOKEN")
var tokenize = func (GITHUB_TOKEN string) string {
	return fmt.Sprintf("token %s", GITHUB_TOKEN)
}(GITHUB_TOKEN)
var requestOptions = &grequests.RequestOptions{Headers: map[string]string{"Accept": "application/vnd.github.v3+json", "Authorization": tokenize}}

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

func NewRepo(repoUrl, repo string) models.CreateResponse {
	var descmsg = fmt.Sprintf("This is the generic repo description for %s", repo)
	var bodyreq = models.RepoRequest{Name: repo, Description: descmsg, Homepage: "https://github.com", Private:false}
	var postbody, _ = json.Marshal(bodyreq)
	var jsonStr = []byte(postbody)
	req, err := http.NewRequest("POST", repoUrl, bytes.NewBuffer(jsonStr))
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("Authorization", tokenize)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return models.CreateResponse{Status:resp.Status, Body:string(body)}
}