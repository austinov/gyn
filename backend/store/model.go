package store

type Profile struct {
	Login    string `json:"login"`
	UserName string `json:"username"`
}

type Dictionary struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
