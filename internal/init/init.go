package init

import (
	"fmt"
	"github.com/staketab/go-mina-payout/internal/vars"
	"os"
	"os/user"
	"path/filepath"
)

func CreateDirectory(path string) error {
	usr, err := user.Current()
	if err != nil {
		return fmt.Errorf("failed to controller current user: %s", err)
	}

	fullPath := filepath.Join(usr.HomeDir, path)

	err = os.MkdirAll(fullPath, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to cr directory: %s", err)
	}

	vars.InfoLog.Printf("Directory created: %s\n", fullPath)
	return nil
}

func EnsureDir(usr string, path string) error {
	fullPath := filepath.Join(usr, path)
	err := os.MkdirAll(fullPath, os.ModePerm)
	if err != nil {
		return fmt.Errorf("error creating directory: %v", err)
	}
	return nil
}

func CreateFilesInDir(path string) error {
	usr, err := user.Current()
	if err != nil {
		return fmt.Errorf("failed to controller current user: %s", err)
	}
	err = EnsureDir(usr.HomeDir, path)
	if err != nil {
		return err
	}
	fileNames := []string{
		vars.SubstitutePayTo,
		vars.NegotiatedFees,
		vars.NegotiatedBurn,
		vars.BurnSupercharged,
	}
	for _, fileName := range fileNames {
		filePath := filepath.Join(usr.HomeDir, path, fileName)
		file, err := os.Create(filePath)
		if err != nil {
			return fmt.Errorf("error creating file %s: %v", filePath, err)
		}
		file.Close()

		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			return fmt.Errorf("failed to create file: %s", filePath)
		} else {
			vars.InfoLog.Printf("File created successfully: %s\n", filePath)
		}
	}
	return nil
}

func CreateConfigFile(path string) error {
	usr, err := user.Current()
	if err != nil {
		return fmt.Errorf("failed to controller current user: %s", err)
	}
	filePath := filepath.Join(usr.HomeDir, path)
	content := []byte(`graphql:
  endpoints:
    - https://api.minascan.io/node/berkeley/v1/graphql
    - https://proxy.berkeley.minaexplorer.com/graphql
pool:
  block_producer_public_key:
  pay_from_address:
  pay_from_private_key:
  pay_memo: go_mina_pay_e70
  tx_fee: 0.001
  pay_threshold: 0.000000001
pool_rates:
  commission_rate: 0.05
  mf_commission_rate: 0.08
  01_commission_rate: 0.05
  investors_commission_rate: 0.08
data:
  blockberry:
    enable: true
    source: https://api.blockberry.one/mina-berkeley
    api_key:
  archive_db:
    enable: false
    source: http://127.0.0.1:5432/archive
`)

	err = os.WriteFile(filePath, content, 0644)
	if err != nil {
		return fmt.Errorf("failed to create config file: %s", err)
	}

	vars.InfoLog.Printf("Config file created successfully: %s\n", filePath)
	return nil
}
