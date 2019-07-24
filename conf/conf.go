package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type Conf struct {
	App *struct {
		WorkID   int64  `json:"work_id"`
		EpochStr string `json:"epoch_str"`
		Port     int    `json:"port"`
	}
}

var C *Conf

func init() {
	dir := "conf"
	for i := 0; i < 3; i++ {
		info, err := os.Stat(dir)
		if err == nil && info.IsDir() {
			break
		}

		dir = filepath.Join("..", dir)
	}

	content, err := ioutil.ReadFile(filepath.Join(dir, "conf.json"))
	if err != nil {
		fmt.Printf("get conf file contents error: %v\n", err)
		os.Exit(-1)
	}

	if err := json.Unmarshal(content, &C); err != nil {
		fmt.Printf("conf.json wrong format: %v\n", err)
		os.Exit(-1)
	}

	log.Printf("epochStr: %s, port: %d, workID: %d\n", C.App.EpochStr, C.App.WorkID, C.App.Port)
}
