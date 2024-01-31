package config

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/staketab/go-mina-payout/internal/vars"
	"gopkg.in/yaml.v3"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

type Config struct {
	Graphql struct {
		Endpoints []string `yaml:"endpoints"`
	} `yaml:"graphql"`
	Pool struct {
		BlockProducerPublicKey string  `yaml:"block_producer_public_key"`
		PayFromAddress         string  `yaml:"pay_from_address"`
		PayFromPrivateKey      string  `yaml:"pay_from_private_key"`
		PayMemo                string  `yaml:"pay_memo"`
		TxFee                  float64 `yaml:"tx_fee"`
		PayThreshold           float64 `yaml:"pay_threshold"`
	} `yaml:"pool"`
	PoolRates struct {
		CommissionRate          float64 `yaml:"commission_rate"`
		MFCommissionRate        float64 `yaml:"mf_commission_rate"`
		O1CommissionRate        float64 `yaml:"01_commission_rate"`
		InvestorsCommissionRate float64 `yaml:"investors_commission_rate"`
	} `yaml:"pool_rates"`
	Data struct {
		Blockberry struct {
			Enable bool   `yaml:"enable"`
			Source string `yaml:"source"`
			APIKey string `yaml:"api_key"`
		} `yaml:"blockberry"`
	} `yaml:"data"`
}

func ReadConfig() (Config, error) {
	usr, err := user.Current()
	if err != nil {
		return Config{}, fmt.Errorf("failed to get current user: %s", err)
	}
	filePath := filepath.Join(usr.HomeDir, vars.ConfigFilePath)
	data, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}, fmt.Errorf("error reading config file: %s", err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return Config{}, fmt.Errorf("error unmarshalling config file: %s", err)
	}

	return config, nil
}

func ReadCSV(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rawRecords, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var records []string
	for _, record := range rawRecords {
		if len(record) > 0 {
			records = append(records, record[0])
		}
	}

	return records, nil
}

func FindMatchingAddress(filePath, targetAddress string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "|")
		if len(parts) == 2 && parts[0] == targetAddress {
			return parts[1], nil
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return targetAddress, nil
}
