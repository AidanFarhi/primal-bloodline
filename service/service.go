package service

import "primalbl/config"

type ContactService struct {
	Config config.Config
}

func NewContactService(c config.Config) ContactService {
	return ContactService{c}
}

func (cs ContactService) SendMessage(name, email, message string) error {
	// finalMessage := name + "\n" + email + "\n" + message
	return nil
}
