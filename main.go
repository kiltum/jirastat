package main

import (
	"flag"
	"fmt"
	jira "github.com/andygrunwald/go-jira"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
	"strings"
	"time"
)

func main() {
	viper.AutomaticEnv()
	viper.SetConfigName("jirastat")
	viper.AddConfigPath(".config")
	viper.AddConfigPath(".")
	viper.ReadInConfig()

	flag.String("js_host", "", "JIRA endpoint")
	flag.String("js_user", "", "Username")
	flag.String("js_pass", "", "Password")
	flag.String("js_project", "IT", "Password")
	flag.String("js_verb", "no", "Be verbose? Yes/No")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	if viper.GetString("js_host") == "" {
		log.Fatalln("js_host cannot be empty")
	}

	verb := strings.ToLower(viper.GetString("js_verb"))

	if verb == "yes" {
		log.Printf("Will use %s as jira endpoint\n", viper.GetString("js_host"))
	}

	if viper.GetString("js_user") == "" {
		log.Fatalln("js_user cannot be empty")
	}

	if viper.GetString("js_pass") == "" {
		log.Fatalln("js_pass cannot be empty")
	}

	if viper.GetString("js_project") == "" {
		log.Fatalln("js_project cannot be empty")
	}

	tp := jira.BasicAuthTransport{
		Username: viper.GetString("js_user"),
		Password: viper.GetString("js_pass"),
	}

	if verb == "yes" {
		log.Printf("Try to login as %s to %s\n", viper.GetString("js_user"), viper.GetString("js_host"))
	}

	var client, err = jira.NewClient(tp.Client(), viper.GetString("js_host"))
	if err != nil {
		log.Fatalf("Cannot login as %s to %s\n", viper.GetString("js_user"), viper.GetString("js_host"))
	}

	list, _, err := client.Project.GetList()

	if err != nil {
		log.Fatalf("Cannot get projects list: %v\n", err)
	}

	found := "no"

	for _, p := range *list {
		if verb == "yes" {
			log.Printf("Found project %v - %v\n", p.Key, p.Name)
		}

		if p.Key == viper.GetString("js_project") {
			found = "yes"
		}
	}

	if found == "no" {
		log.Fatalf("Project %s not found", viper.GetString("js_project"))
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

	if verb == "yes" {
		log.Printf("%d issues found.\n", len(issues))
	}

	cre := make(map[string]int) // maps for created and updated dated
	upd := make(map[string]int)

	for _, i := range issues {

		t := time.Time(i.Fields.Created) // convert go-jira.Time to time.Time for manipulation
		created := t.Format("2006-01-02")
		t = time.Time(i.Fields.Updated)
		updated := t.Format("2006-01-02")

		cre[created] = cre[created] + 1

		upd[updated] = upd[updated] + 1

		fmt.Printf("Issue Key: %s %s/%s\nIssue Summary: %s\nStatus: %s\n\n", i.Key, created, updated, i.Fields.Summary, i.Fields.Status.Name)

	}

	for k, v := range cre {
		fmt.Printf("%s\t%d\t%d\n", k, v, upd[k])
	}

}
