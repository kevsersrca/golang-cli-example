package docker

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestDockerRun(t *testing.T) {
	imageName := "mongo:latest"
	containerName := fmt.Sprintf("image%d", rand.Int())

	s := ProvideService()
	_, err := s.ImagePull(imageName)
	if err != nil {
		t.Error(err)
	}

	container, err := s.Create(imageName, containerName)
	if err != nil {
		t.Error(err)
	}

	err = s.Start(container.ID)
	if err != nil {
		t.Error(err)
	}

	err = s.ContainerWait(container.ID)
	if err != nil {
		t.Error(err)
	}
}

func TestDockerList(t *testing.T) {
	s := ProvideService()
	_, err := s.List()
	if err != nil {
		t.Error(err)
	}
}

func TestDockerGet(t *testing.T) {
	s := ProvideService()
	container, err := s.List()
	if err != nil {
		t.Error(err)
	}
	_, err = s.Get(container[0].ID)
	if err != nil {
		t.Error(err)
	}
}

func TestDockerLogs(t *testing.T) {
	s := ProvideService()
	container, err := s.List()
	if err != nil {
		t.Error(err)
	}
	_, err = s.Get(container[0].ID)
	if err != nil {
		t.Error(err)
	}
}

func TestDockerStop(t *testing.T) {
	s := ProvideService()
	container, err := s.List()
	if err != nil {
		t.Error(err)
	}
	err = s.Stop(container[0].ID)
	if err != nil {
		t.Error(err)
	}
}

func TestDockerStart(t *testing.T) {
	s := ProvideService()
	container, err := s.List()
	if err != nil {
		t.Error(err)
	}
	err = s.Start(container[0].ID)
	if err != nil {
		t.Error(err)
	}
}

func TestDockerImageList(t *testing.T) {
	s := ProvideService()
	_, err := s.ImageList()
	if err != nil {
		t.Error(err)
	}
}
