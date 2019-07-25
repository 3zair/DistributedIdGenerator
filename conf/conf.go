package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type Conf struct {
	App *struct {
		Name string `json:"name"`
		Port int    `json:"port"`
	}
	WorkId int64 `json:"work_id"` // [0, 31]
	Epoch  int64 `json:"epoch"`   // 时间戳
}

var C *Conf

func init() {
	content, err := ioutil.ReadFile("conf/conf.json")
	if err != nil {
		fmt.Printf("get conf file contents error: %v\n", err)
		os.Exit(-1)
	}

	if err := json.Unmarshal(content, &C); err != nil {
		fmt.Printf("conf.json wrong format: %v\n", err)
		os.Exit(-1)
	}

	if C.WorkId > 31 || C.WorkId < 0 {
		fmt.Printf("invalid workID: %d, valid workID 0-31\n", C.WorkId)
		os.Exit(-1)
	}

	if C.Epoch < 0 || C.Epoch > time.Now().Unix() {
		fmt.Printf("invalid epoch: %d\n", C.Epoch)
		os.Exit(-1)
	}

	fmt.Printf("conf:\nAppName: %s\nAppPort: %d\nWorkerId: %d\nEposh: %d\n", C.App.Name, C.App.Port, C.WorkId, C.Epoch)
	return
}
