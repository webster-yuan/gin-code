package di

import (
	"errors"
	"sync"
)

// Container 依赖注入容器
type Container struct {
	services sync.Map
}

var (
	container *Container
	once      sync.Once
)

// GetContainer 获取全局容器实例
func GetContainer() *Container {
	once.Do(func() {
		container = &Container{}
	})
	return container
}

// Register 注册服务
func (c *Container) Register(name string, service interface{}) {
	c.services.Store(name, service)
}

// Get 获取服务
func (c *Container) Get(name string) (interface{}, error) {
	service, ok := c.services.Load(name)
	if !ok {
		return nil, errors.New("服务不存在: " + name)
	}
	return service, nil
}

// MustGet 获取服务，如果不存在则panic
func (c *Container) MustGet(name string) interface{} {
	service, err := c.Get(name)
	if err != nil {
		panic(err)
	}
	return service
}
