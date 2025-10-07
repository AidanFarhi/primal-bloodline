package repo

import (
	"encoding/json"
	"os"
	"primalbl/config"
	"primalbl/model"
)

type CatRepository struct {
	Cats map[string]model.Kitten
}

func NewCatRepository(conf config.Config) CatRepository {
	repo := CatRepository{
		Cats: map[string]model.Kitten{},
	}
	file, _ := os.Open(conf.JSONPath)
	defer file.Close()
	var kittens []model.Kitten
	json.NewDecoder(file).Decode(&kittens)
	for _, k := range kittens {
		repo.Cats[k.CatReferenceName] = k
	}
	return repo
}

func (cr CatRepository) GetAllCats() []model.Kitten {
	kittens := []model.Kitten{}
	for _, v := range cr.Cats {
		kittens = append(kittens, v)
	}
	return kittens
}

func (cr CatRepository) GetCatByReferenceName(referenceName string) model.Kitten {
	val, ok := cr.Cats[referenceName]
	if ok {
		return val
	}
	return model.Kitten{}
}
