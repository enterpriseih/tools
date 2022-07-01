package main

import (
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var filename string

type HarborRobot struct {
	Name  string `json:"name"`
	Token string `json:"token"`
}

func init() {
	flag.StringVar(&filename, "f", "", "-f config.json")
	flag.Parse()
}

func main() {
	jsonFile, err := os.Open(filename)
	log.Println(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var robot HarborRobot

	harbor := "harbor.com"

	json.Unmarshal(byteValue, &robot)
	log.Println(robot)
	dockercfgJSONContent, _ := handleDockerCfgJSONContent(robot.Name, robot.Token, "", harbor)
	log.Println(string(dockercfgJSONContent))

	dockercfgJSONContent64 := base64.StdEncoding.EncodeToString(dockercfgJSONContent)
	log.Println(dockercfgJSONContent64)
	k8sYaml := fmt.Sprintf(`
apiVersion: v1
data:
  .dockerconfigjson: %s
kind: Secret
metadata:
  name: harbor-ops
  namespace: ops
type: kubernetes.io/dockerconfigjson
`, dockercfgJSONContent64)
	log.Println(k8sYaml)
}

func handleDockerCfgJSONContent(username, password, email, server string) ([]byte, error) {
	dockercfgAuth := DockerConfigEntry{
		Username: username,
		Password: password,
		Email:    email,
		Auth:     encodeDockerConfigFieldAuth(username, password),
	}

	dockerCfgJSON := DockerConfigJSON{
		Auths: map[string]DockerConfigEntry{server: dockercfgAuth},
	}

	return json.Marshal(dockerCfgJSON)
}

func encodeDockerConfigFieldAuth(username, password string) string {
	fieldValue := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(fieldValue))
}

// DockerConfigJSON represents a local docker auth config file
// for pulling images.
type DockerConfigJSON struct {
	Auths DockerConfig `json:"auths"`
	// +optional
	HttpHeaders map[string]string `json:"HttpHeaders,omitempty"`
}

// DockerConfig represents the config file used by the docker CLI.
// This config that represents the credentials that should be used
// when pulling images from specific image repositories.
type DockerConfig map[string]DockerConfigEntry

type DockerConfigEntry struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email,omitempty"`
	Auth     string `json:"auth,omitempty"`
}
