package docker

import (
	"context"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

type IService interface {
	List() ([]types.Container, error)
	ListAll() ([]types.Container, error)
	Get(id string) (types.ContainerJSON, error)
	Logs(id string) (io.ReadCloser, error)
	Stop(id string) error
	Start(id string) error
	Create(imageName string, containerName string) (container.ContainerCreateCreatedBody, error)
	ImagePull(imageName string) (io.Reader, error)
	ImageList() ([]types.ImageSummary, error)
	ContainerWait(id string) error
}

type Service struct {
	cli *client.Client
}

func ProvideService() *Service {
	return &Service{
		cli: connect(),
	}
}

// docker ps
func (s *Service) List() ([]types.Container, error) {
	return s.cli.ContainerList(context.Background(), types.ContainerListOptions{})
}

// docker ps --all
func (s *Service) ListAll() ([]types.Container, error) {
	return s.cli.ContainerList(context.Background(), types.ContainerListOptions{All: true})
}

// docker inspect
func (s *Service) Get(id string) (types.ContainerJSON, error) {
	return s.cli.ContainerInspect(context.Background(), id)
}

// docker logs
func (s *Service) Logs(id string) (io.ReadCloser, error) {
	out, err := s.cli.ContainerLogs(context.Background(), id, types.ContainerLogsOptions{
		ShowStdout: true,
	})
	io.Copy(os.Stdout, out)
	return out, err
}

// stop container
func (s *Service) Stop(id string) error {
	return s.cli.ContainerStop(context.Background(), id, nil)
}

//delete container
func (s *Service) Delete(id string) error {
	return s.cli.ContainerRemove(context.Background(), id, types.ContainerRemoveOptions{})
}

// start container
func (s *Service) Start(id string) error {
	return s.cli.ContainerStart(context.Background(), id, types.ContainerStartOptions{})
}

// create container
func (s *Service) Create(imageName string, containerName string) (container.ContainerCreateCreatedBody, error) {
	return s.cli.ContainerCreate(context.Background(), &container.Config{
		Image: imageName,
	}, nil, nil, nil, containerName)

}

// pull image
func (s *Service) ImagePull(imageName string) (io.Reader, error) {
	out, err := s.cli.ImagePull(context.Background(), imageName, types.ImagePullOptions{})
	io.Copy(os.Stdout, out)
	return out, err
}

// image list
func (s *Service) ImageList() ([]types.ImageSummary, error) {
	return s.cli.ImageList(context.Background(), types.ImageListOptions{})
}

// container wait
func (s *Service) ContainerWait(id string) error {
	statusCh, errCh := s.cli.ContainerWait(context.Background(), id, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			return err
		}
	case <-statusCh:
		return nil
	}
	return nil
}
