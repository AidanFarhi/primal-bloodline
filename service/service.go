package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"primalbl/config"
)

type ContactService struct {
	Config config.Config
}

func NewContactService(c config.Config) ContactService {
	return ContactService{c}
}

func (cs ContactService) SendMessage(name, number, message string) error {
	finalMessage := name + "\n" + number + "\n" + message
	values := url.Values{
		"phone":   {cs.Config.PrimalBloodlinePhoneNumber},
		"message": {finalMessage},
		"key":     {cs.Config.TextbeltAPIKey},
	}
	resp, err := http.PostForm(cs.Config.TextbeltURL, values)
	if err != nil {
		return err
	}
	var responseData map[string]any
	err = json.NewDecoder(resp.Body).Decode(&responseData)
	if err != nil {
		return err
	}
	fmt.Println(responseData)
	return nil
}
