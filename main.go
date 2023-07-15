package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"

	"gopkg.in/src-d/go-billy.v4/memfs"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/storage/memory"
)

func main() {
	f := memfs.New()
	_, err := git.Clone(memory.NewStorage(), f, &git.CloneOptions{
		URL: "https://github.com/winnietthepooh/blog.git",
	})
	if err != nil {
		log.Fatal(err)
	}
	files, err := f.ReadDir("./src/content/projects")
	if err != nil {
		log.Fatal(err)
	}

	for _, fi := range files {
		fmt.Println(fi.Name())
		file, err := f.Open("./src/content/projects/" + fi.Name())
		if err != nil {
			panic(err)
		}
		content, err := io.ReadAll(file)
		if err != nil {
			log.Fatal(err)
		}
		regx := regexp.MustCompile("---\ntitle: (.+)\ndescription: (.+)\ntags:\n((?:  - \"(.+)\"\n)+)---")
		data := regx.FindString(string(content))
		data = strings.ReplaceAll(data, "---", "")
		data = strings.ReplaceAll(data, "\"", "")
		lines := strings.Split(data, "\n")
		var tags []string
		var title string
		var description string
	main:
		for line := range lines {
			line := line
			if strings.HasPrefix(lines[line], "tags") {
				for line := line + 1; line < len(lines); line++ {
					tag := strings.ReplaceAll(lines[line], "  - ", "")
					tags = append(tags, tag)
				}
				break main
			} else if strings.HasPrefix(lines[line], "title") {
				titleSplit := strings.Split(lines[line], ":")
				title = titleSplit[1]
			} else if strings.HasPrefix(lines[line], "description") {
				descSplit := strings.Split(lines[line], ":")
				description = descSplit[1]
			}
		}
		url := "https://caden32.com/projects/project/" + strings.ToLower(strings.Split(fi.Name(), ".")[0])
		fmt.Println(tags, title, description, url)
	}
	http.HandleFunc("/.well-known/webfinger", func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Query().Has("resource") {
			fmt.Fprintf(writer, "%s", "{\"subject\":\"acct:caden@social.caden32.com\",\"links\":[{\"rel\":\"http://webfinger.net/rel/profile-page\",\"type\":\"text/html\",\"href\":\"https://social.caden32.com/caden\"},{\"rel\":\"self\",\"type\":\"application/activity+json\",\"href\":\"https://social.caden32.com/caden\",\"properties\":{\"https://www.w3.org/ns/activitystreams#type\":\"Person\"}}]}")
		} else {
			writer.WriteHeader(http.StatusNotFound)
			return
		}
	})
	http.HandleFunc("/caden", func(writer http.ResponseWriter, request *http.Request) {

	})
	// http.ListenAndServe(":8080", nil)
}

type Post struct {
	URL   string
	Title string
	Desc  string
	Tags  string
}
