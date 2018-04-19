package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"

	yaml "gopkg.in/yaml.v2"
)

type Git struct {
	Enabled              bool   `yaml:"enabled"`
	RemoteName           string `yaml:"remoteName"`
	RemoteURL            string `yaml:"remoteURL"`
	JenkinsCredentialsID string `yaml:"jenkinsCredentialsID"`
	Branches             string `yaml:"branches"`
}

type SCMPolling struct {
	Enabled               bool   `yaml:"enabled"`
	Schedule              string `yaml:"schedule"`
	IgnorePostCommitHooks bool   `yaml:"ignorePostCommitHooks"`
}

type Timer struct {
	Enabled               bool   `yaml:"enabled"`
	Schedule              string `yaml:"schedule"`
	IgnorePostCommitHooks bool   `yaml:"ignorePostCommitHooks"`
}

type Triggers struct {
	SCMPolling SCMPolling `yaml:"scmPolling"`
	Timer      Timer      `yaml:"timer"`
}

type Job struct {
	Git                     Git      `yaml:"git"`
	Name                    string   `yaml:"name"`
	DisplayName             string   `yaml:"displayName"`
	Description             string   `yaml:"description"`
	DisableConcurrentBuilds bool     `yaml:"disableConcurrentBuilds"`
	Triggers                Triggers `yaml:"triggers"`
}

type Input struct {
	Jobs []Job `yaml:"jobs"`
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	var input Input

	jobsYamlFile, _ := filepath.Abs("./jobs.yaml")
	yamlFile, err := ioutil.ReadFile(jobsYamlFile)
	check(err)

	err = yaml.Unmarshal(yamlFile, &input)
	check(err)

	t, err := template.ParseFiles("./templates/job.xml")
	check(err)

	jobsFolder := filepath.Join(".", "jobs")
	os.MkdirAll(jobsFolder, os.ModePerm)

	for _, job := range input.Jobs {
		relativePath := "./jobs/" + job.Name + ".xml"
		path, _ := filepath.Abs(relativePath)
		jobXML, err := os.Create(path)
		check(err)
		check(t.Execute(jobXML, job))
		fmt.Printf("succesfully created job %s\n", relativePath)
	}
}
