package version

import (
	"encoding/json"
	"fmt"
	"github.com/staketab/go-mina-payout/internal/vars"
)

func GetVersion() {
	data := struct {
		Version string `json:"version"`
	}{
		Version: vars.Version,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		vars.ErrorLog.Println("Error:", err)
		return
	}

	fmt.Println(string(jsonData))
}
