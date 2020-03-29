package main

import (
	"encoding/json"
	"fmt"
)

type Repo struct {
	Author   string    `json:"author"`
	Github   string    `json:"github"`
	Projects []Project `json:"projects"`
}

type Project struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func main() {
	r := Repo{
		Author: "Emerson MX",
		Github: "https://github.com/emersonmx",
		Projects: []Project{
			Project{
				Name: "tmplt",
				Url:  "https://github.com/emersonmx/tmplt",
			},
			Project{
				Name: "dotfiles",
				Url:  "https://github.com/emersonmx/dotfiles",
			},
		},
	}

	j, _ := json.Marshal(r)
	fmt.Println(string(j))

	js := []byte(`{"author":"Emerson MX","github":"https://github.com/emersonmx","projects":[{"name":"tmplt","url":"https://github.com/emersonmx/tmplt"},{"name":"dotfiles","url":"https://github.com/emersonmx/dotfiles"}]}`)

	r2 := Repo{}
	fmt.Println(json.Unmarshal(js, &r2))
	fmt.Println(r2)
}
