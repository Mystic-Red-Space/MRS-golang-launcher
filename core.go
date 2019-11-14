package main

import (
	"encoding/json"
	"github.com/JoshuaDoes/go-yggdrasil"
	"github.com/google/uuid"
	"io/ioutil"
	"net/http"
	"os"
)

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

type mcfile struct {
	Name string `json:"name"`
	Dir  string `json:"dir"`
	Url  string `json:"url"`
	Md5  string `json:"md5"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type modpack struct {
	Name    string `json:"name"`
	Icon    string `json:"icon"`
	Profile string `json:"profile"`
}

func getJson(url string, format interface{}) {
	resp, _ := http.Get(url)
	res, _ := ioutil.ReadAll(resp.Body)
	_ = json.Unmarshal(res, &format)
}

func mclogin(id string, pass string) (*yggdrasil.AuthenticationResponse, *yggdrasil.Error) {
	var authClient *yggdrasil.Client
	if fileExists("clientToken.txt") {
		tokenFile, _ := ioutil.ReadFile("clientToken.txt")
		authClient = &yggdrasil.Client{ClientToken: string(tokenFile)}
	} else {
		tokenFile, _ := os.Create("clientToken.txt")
		clientToken := uuid.New()
		_, _ = tokenFile.WriteString(clientToken.String())
		_ = tokenFile.Close()
		authClient = &yggdrasil.Client{ClientToken: clientToken.String()}
	}
	resp, err := authClient.Authenticate(id, pass, "Minecraft", 1)
	return resp, err
}

func downloadFile(filepath string, url string) error {

    // Get the data
    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    // Create the file
    out, err := os.Create(filepath)
    if err != nil {
        return err
    }
    defer out.Close()

    // Write the body to file
    _, err = io.Copy(out, resp.Body)
    return err
}
