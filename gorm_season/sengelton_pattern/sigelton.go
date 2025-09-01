package singleton

//در دیزاین پترن singleton
//، هدف این است که از یک استراکت ( یا یک کلاس در زبان‌های شی‌گرا) تنها یک
//داشته باشیم instance یک دسترسی گلوبال وجود داشته باشد.

import (
	"sync"
)

// Interface برای قابلیت تست بهتر
type Database interface {
	Connect() string
	SetConnection(conn string)
}

type database struct {
	connection string
	mu         sync.RWMutex
}

var (
	instance *database
	once     sync.Once
)

func GetInstance() Database {
	once.Do(func() {
		instance = &database{connection: "default://localhost"}
	})
	return instance
}

func (d *database) Connect() string {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.connection
}

func (d *database) SetConnection(conn string) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.connection = conn
}
