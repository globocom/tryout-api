package challenge

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"../db"
	"github.com/cezarsa/form"
	"github.com/gosimple/slug"
)

const token string = "acbea31b589a270ec856569d9f0b6c88c23bb6a96c66ac5bcb1f7f54d12b1d69"

type Path struct {
	Path       string `json:"path"`
	HttpStatus int    `json:"http_status"`
	Throughput int    `json:"throughput"`
	Input      string `json:"input"`
	Output     string `json:"output"`
}

type Challenge struct {
	Slug  string `bson:"_id"`
	Name  string `json:"name"`
	Paths []Path `json:"paths"`
}

func Create(c Challenge) error {
	c.Slug = slug.Make(c.Name)
	return db.Coll("challenge").Insert(&c)
}

func (c *Challenge) URL() string {
	return fmt.Sprintf("/challenge/%s/try")
}

func (c *Challenge) Start(repo string) error {
	err := c.createApp()
	if err != nil {
		return err
	}
	err = c.setRepo(repo)
	if err != nil {
		return err
	}
	return c.deployChallenge()
}

func (c *Challenge) createApp() error {
	baseURL := "http://tsuru.globoi.com"
	v := url.Values{}
	v.Set("name", c.Slug)
	v.Set("platform", "python")
	v.Set("plan", "large")
	v.Set("teamOwner", "bigdata")
	v.Set("pool", "bigdata")
	v.Set("description", "app create by tryout api")
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/apps", baseURL), strings.NewReader(v.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", fmt.Sprintf("bearer %s", token))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("Wrong status code. Want 201. Got: %d", resp.StatusCode)
	}
	return nil
}

func (c *Challenge) setRepo(repo string) error {
	baseURL := "http://tsuru.globoi.com"
	env := map[string]string{"Name": "REPO", "Value": repo}
	envs := map[string]interface{}{
		"Envs": []map[string]string{env},
	}
	v, err := form.EncodeToValues(&envs)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/apps/%s/env", baseURL, c.Slug), strings.NewReader(v.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", fmt.Sprintf("bearer %s", token))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Wrong status code. Want 200. Got: %d", resp.StatusCode)
	}
	return nil
}

func (c *Challenge) deployChallenge() error {
	baseURL := "http://tsuru.globoi.com"
	v := url.Values{}
	v.Set("image", "tryout/agent")
	v.Set("origin", "image")
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/apps/%s/deploy", baseURL, c.Slug), strings.NewReader(v.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", fmt.Sprintf("bearer %s", token))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		respBody, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("Wrong status code. Want 200. Got: %d\n Error: %s", resp.StatusCode, respBody)
	}
	return nil
}
