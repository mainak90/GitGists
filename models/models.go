package models

type Repo struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Forks    int    `json:"forks"`
	Private  bool   `json:"private"`
}

type File struct {
	Content string `json:"content"`
}

type Gist struct {
	Description string          `json:"description"`
	Public      bool            `json:"public"`
	Files       map[string]File `json:"files"`
}

type RepoRequest struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Homepage string `json:"homepage"`
	Private  bool `json:"private"`
}

type CreateResponse struct {
	Status string `json:"status"`
	Body   string `json:"body"`
}

type Config struct {
	Url          string   `json:"url"`
	Content_Type string   `json:"content_type"`
	Secret 		 string   `json:"secret"`
}

type WebhookEach struct {
	Name  string `json:"name"`
	Owner string  `json:"owner"`
	Hook  string  `json:"hook"`
}

type WebhookList struct {
	Entries []WebhookEach `json:"entries"`
}