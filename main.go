package main

import (
	"fmt"
	jira "github.com/andygrunwald/go-jira"
	"log"
	"strings"
	"time"
)

func main() {
	tp := jira.BasicAuthTransport{
		Username: "v.kaloshin@crpt.ru",
		Password: "passwordwashere",
	}

	fmt.Printf("Login")
	var client, err = jira.NewClient(tp.Client(), "https://crptteam.atlassian.net")
	if err != nil {
		log.Fatal("cannot login")
	}

	u, _, err := client.User.Get("v.kaloshin")
	if err != nil {
		log.Fatal("cannot user get")
	}

	fmt.Printf("\nEmail: %v\nSuccess!\n", u.EmailAddress)

	list, _, err := client.Project.GetList()

	if err != nil {
		log.Fatal("Cannot get projects")
	}

	for _, p := range *list {
		fmt.Printf("%v %v\n", p.Key, p.Name)

	}

	var issues []jira.Issue

	appendFunc := func(i jira.Issue) (err error) {
		issues = append(issues, i)
		return err
	}
	// project="Infrastructure Tasks" and updatedDate>-30d order by updated DESC
	// createdDate
	err = client.Issue.SearchPages(fmt.Sprintf(`project=%s and updatedDate>-10d`, strings.TrimSpace("IT")), nil, appendFunc)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues found.\n", len(issues))

	for _, i := range issues {

		t := time.Time(i.Fields.Created) // convert go-jira.Time to time.Time for manipulation
		created := t.Format("2006-01-02")
		t = time.Time(i.Fields.Updated)
		updated := t.Format("2006-01-02")
		fmt.Printf("Issue Key: %s %s/%s\nIssue Summary: %s\nStatus: %s\n\n", i.Key, created, updated, i.Fields.Summary, i.Fields.Status.Name)
	}
}
