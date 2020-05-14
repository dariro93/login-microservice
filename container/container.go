package container

import (
	"fmt"
	"sync"
)

type Container struct {
	Instances map[string]interface{}
}

var containerInstance *Container
var once sync.Once

func RetrieveContainer() *Container {
	once.Do(func() {
		containerInstance = &Container{
			Instances: make(map[string]interface{}),
		}
	})
	return containerInstance
}

func (c *Container) RetrieveInstanceByType(key string) (interface{}, error) {
	for k, instc := range c.Instances {
		if key == k {
			return instc, nil
		}
	}
	return nil, fmt.Errorf("There is not such instance %s", key)
}

func (c *Container) RegisterInstance(key string, impl interface{}) {
	c.Instances[key] = impl
}
