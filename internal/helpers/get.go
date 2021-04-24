package helpers

import (
	"io/ioutil"
	"log"
	"sync"
)

var cache map[string]string
var lock sync.RWMutex

func init() {
	cache = make(map[string]string)
}

func FindDriveById(id string) string {
	lock.RLock()
	if letter, ok := cache[id]; ok {
		lock.RUnlock()
		return letter
	}
	lock.RUnlock()

	for _, drive := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		letter := string(drive)
		if data, err := ioutil.ReadFile(letter + ":/.$samba/remote-open.id.txt"); err == nil {
			log.Printf("try samba drive: %s:", letter)
			if string(data) == id {

				lock.Lock()
				defer lock.Unlock()
				cache[id] = letter

				return letter
			}
		}
	}
	return ""
}
