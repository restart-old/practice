package kits

import (
	"sync"

	"github.com/df-plus/kit"
)

func init() {
	kits = make(map[string]kit.Kit)
}

var kits map[string]kit.Kit
var kitsMu sync.RWMutex

func Register(k kit.Kit) {
	kitsMu.Lock()
	defer kitsMu.Unlock()
	kits[k.Name()] = k
}
func ByName(name string) (kit.Kit, bool) {
	kitsMu.RLock()
	defer kitsMu.RUnlock()
	k, ok := kits[name]
	return k, ok
}
