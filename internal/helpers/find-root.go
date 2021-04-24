package helpers

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

func FindRoot(file string) (string, string, error) {
	for dir := filepath.Dir(file); dir != "/"; dir = filepath.Dir(dir) {
		samba_dir := filepath.Join(dir, ".$samba")

		if stat, err := os.Stat(samba_dir); err == nil {
			if stat.IsDir() {
				if root, err := getSambaRoot(samba_dir); err == nil {
					if rel, err := filepath.Rel(dir, file); err == nil {
						return root, rel, nil
					} else {
						return "", "", err
					}
				} else {
					return "", "", err
				}
			}
		}
	}

	return "", "", errors.New("failed find root")
}

func getSambaRoot(dir string) (string, error) {
	dataFile := filepath.Join(dir, "remote-open.id.txt")
	if _, err := os.Stat(dataFile); os.IsNotExist(err) {
		log.Printf("will create info file at %s", dataFile)
		id, err := uuid.NewRandom()
		if err != nil {
			return "", fmt.Errorf("failed create uuid: %v", err)
		}

		if err := ioutil.WriteFile(dataFile, []byte(id.String()), 0644); err != nil {
			return "", fmt.Errorf("failed write uuid file at %s: %v", dataFile, err)
		}
	}

	if bs, err := ioutil.ReadFile(dataFile); err != nil {
		return "", fmt.Errorf("failed read uuid file at %s: %v", dataFile, err)
	} else {
		return string(bs), nil
	}
}
