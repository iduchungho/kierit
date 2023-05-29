package session

import (
	"github.com/gofiber/fiber/v2/middleware/session"
	"sync"
)

var store *session.Store
var lock = &sync.Mutex{}

func GetStore() *session.Store {
	if store == nil {
		lock.Lock()
		defer lock.Unlock()
		if store == nil {
			store = session.New()
		} else {
			return store
		}
	}
	return store
}
