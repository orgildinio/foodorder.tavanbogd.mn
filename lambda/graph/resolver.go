package graph

import (
	"sync"
)

type Resolver struct {
	mutex sync.Mutex
}
