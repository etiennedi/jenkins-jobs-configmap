package main

import (
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

type Job struct {
	Git  Git    `yaml:"git"`
	Name string `yaml:"name"`
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

	for _, job := range input.Jobs {
		check(t.Execute(os.Stdout, job))
	}
}
