package release

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"
)

type ReleaseMeta struct {
	ArchName string `json:"archName"`
	Branch string `json:"branch"`
	CommitSha string `json:"commitSha"`
	BuildDate time.Time `json:"buildDate"`
	InstalledDate time.Time `json:"installedDate"`
}

func GetReleaseInfoFromJsonFile(filePath string) (ReleaseMeta, error) {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		return ReleaseMeta{}, err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var releaseMeta ReleaseMeta
	err = json.Unmarshal(byteValue, &releaseMeta)
	if err != nil {
		return ReleaseMeta{}, err
	}

	return releaseMeta, nil
}