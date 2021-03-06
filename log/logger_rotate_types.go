package log

import (
	"os"
	"sync"
)

const (
	defaultMaxSize int64 = 250
	megabyte       int64 = 1024 * 1024
)

// Rotate struct for each instance of Rotate
type Rotate struct {
	FileName string
	Rotate   *bool
	MaxSize  int64

	size   int64
	output *os.File
	mu     sync.Mutex
}
