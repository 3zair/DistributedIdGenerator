package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Conf struct {
	App *struct {
		WorkID       int64  `json:"work_id"` //0-31
		StartTimeStr string `json:"start_time_str"`
		Port         int    `json:"port"`
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

	if C.App.WorkID > 31 {
		fmt.Printf("invalid workID: %d,valid workID 0-31 \n", C.App.WorkID)
		return
	}

	fmt.Printf("epochStr: %s, port: %d, workID: %d\n", C.App.StartTimeStr, C.App.WorkID, C.App.Port)
}
