package docker

func GetService() *Service {
	return ProvideService()
}
