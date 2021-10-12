package registry

import "github.com/nestorivan/academy-go-q32021/service"

func (r *registry) NewFileService() service.FileService {
  return service.NewFileService()
}