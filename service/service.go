package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"primalbl/config"
	"primalbl/model"
	"primalbl/repo"
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

type CatService struct {
	CatRepository repo.CatRepository
}

func NewCatService(cr repo.CatRepository) CatService {
	return CatService{
		CatRepository: cr,
	}
}

func (cs CatService) GetAllCats() []model.Kitten {
	return cs.CatRepository.GetAllCats()
}

func (cs CatService) GetCatByReferenceName(catReferenceName string) model.Kitten {
	return cs.CatRepository.GetCatByReferenceName(catReferenceName)
}
