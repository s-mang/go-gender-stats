package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const slackURL = "https://slack.com/api/users.list?token=%s"

var slackAPIToken = os.Getenv("SLACK_API_TOKEN")

func getGopherNames() ([]string, error) {
	u := fmt.Sprintf(slackURL, slackAPIToken)
	resp, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var users SlackResp
	err = json.NewDecoder(resp.Body).Decode(&users)
	if err != nil {
		return nil, err
	}

	names := make([]string, 0, len(users.Members))
	for _, u := range users.Members {
		names = append(names, u.Profile.FirstName)
	}

	return names, nil

}

type SlackResp struct {
	OK      bool        `json:"ok"`
	Members []SlackUser `json:"members"`
}

type SlackUser struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Deleted bool   `json:"deleted"`
	Color   string `json:"color"`
	Profile struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		RealName  string `json:"real_name"`
		Email     string `json:"email"`
		Skype     string `json:"skype"`
		Phone     string `json:"phone"`
		Image24   string `json:"image_24"`
		Image32   string `json:"image_32"`
		Image48   string `json:"image_48"`
		Image72   string `json:"image_72"`
		Image192  string `json:"image_192"`
	} `json:"profile"`
	IsAdmin  bool `json:"is_admin"`
	IsOwner  bool `json:"is_owner"`
	Has2Fa   bool `json:"has_2fa"`
	HasFiles bool `json:"has_files"`
}
