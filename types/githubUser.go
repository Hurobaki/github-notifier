package types

type GithubUser struct {
	Login  string `json:"login"`
	Id     int    `json:"id"`
	Avatar string `json:"avatar_url"`
	Url    string `json:"html_url"`
}
