package service

import (
	"lf/goLiveStreaming/internal/model"
	"lf/goLiveStreaming/internal/repository"
)

type KeyService interface {
	AuthStreamingKey(name, key string) (*model.Keys, error)
}

type keyService struct {
	keysRepository repository.KeysRepository
}

func NeyKeyService(repo repository.KeysRepository) KeyService {
	return &keyService{
		keysRepository: repo,
	}
}

func (sk *keyService) AuthStreamingKey(name, key string) (*model.Keys, error) {
	return sk.keysRepository.FindStreamKey(name, key)
}
