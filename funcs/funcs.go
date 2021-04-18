package funcs

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/levigross/grequests"
	"github.com/mainak90/GitGists/models"
	"net/http"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type RepoList struct {
	Repos []string `json:"repos"`
}

type HookRequest struct {
	Name   string   `json:"name"`
	Config models.Config   `json:"config"`
	Events []string `json:"events"`
	Active bool     `json:"active"`
}

var GITHUB_TOKEN = func() string {
	if value, ok := os.LookupEnv("GITHUB_TOKEN"); ok {
		return value
	}
	log.Println("No GITHUB_TOKEN env variable detected, trying without an oauth token..")
	return ""
}()
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

func (r *RepoList)NewRepos(url string) {
	repolist := r.Repos
	start := time.Now()
	ch := make(chan string)
	for _,repo := range repolist {
		go MakeRequest(url, repo, ch)
	}
	for range repolist {
		log.Println(<-ch)
	}
	log.Println("Elapsed ", time.Since(start).Seconds())
}

func MakeList(path string) []string {
	var list []string
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Failed to open file ", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		list = append(list, scanner.Text())
	}
	log.Println(list)
	return list
}

func MakeRequest(url string, repo string, ch chan<-string) {
	start := time.Now()
	var descmsg = fmt.Sprintf("This is the generic repo description for %s", repo)
	var bodyreq = models.RepoRequest{Name: repo, Description: descmsg, Homepage: "https://github.com", Private:false}
	var postbody, _ = json.Marshal(bodyreq)
	var jsonStr = []byte(postbody)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("Authorization", tokenize)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	secs := time.Since(start).Seconds()
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(body))
	ch <- fmt.Sprintf("%.2f elapsed with response : %d %s", secs, len(body), url)
}

func (c *HookRequest)InitConfig(hookurl, owner, repo string) *HookRequest {
	c.Name = "web"
	var events []string
	events = append(events, "push","pull_request")
	c.Events = events
	c.Active = true
	c.Config = models.Config{Url: hookurl, Content_Type: "json", Secret: GITHUB_TOKEN}
	return c
}

func (c *HookRequest)CreateWebHook(url string) {
	var postbody, _ = json.Marshal(c)
	var jsonStr = []byte(postbody)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("Authorization", tokenize)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("%s",string(body))
}

func (c *HookRequest) CreateWebHooks(path string){
	file, _ := ioutil.ReadFile(path)
	ch := make(chan string)
	var mapdata models.WebhookList
	_ = json.Unmarshal([]byte(file), &mapdata)
	for i := 0; i < len(mapdata.Entries); i++ {
		eachreq := c.InitConfig(mapdata.Entries[i].Hook, mapdata.Entries[i].Owner, mapdata.Entries[i].Name)
		url := fmt.Sprintf("https://api.github.com/repos/%s/%s/hooks", mapdata.Entries[i].Owner, mapdata.Entries[i].Name)
		go eachreq.CreateWebHookConq(url, ch)
	}

	for range mapdata.Entries {
		log.Println(<-ch)
	}
}

func (c *HookRequest)CreateWebHookConq(url string, ch chan<-string) {
	var postbody, _ = json.Marshal(c)
	var jsonStr = []byte(postbody)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("Authorization", tokenize)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("%s",string(body))
	ch <- fmt.Sprintf("Done! Webhook created for : %s", url)
}

func GetOrgRepos(org string) []byte {
	url := fmt.Sprintf("https://api.github.com/orgs/%s/repos", org)
	req, _ := http.NewRequest("GET", url,nil)
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("Authorization", tokenize)
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	return data
}

