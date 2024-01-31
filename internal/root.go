package internal

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/staketab/go-mina-payout/config"
	ff "github.com/staketab/go-mina-payout/internal/init"
	"github.com/staketab/go-mina-payout/internal/vars"
	"github.com/staketab/go-mina-payout/internal/version"
	"gopkg.in/yaml.v3"
	"os"
	"os/user"
	"path/filepath"
)

var RootCmd = &cobra.Command{
	Use:   vars.Binary,
	Short: "Mina-payout script",
	Long:  "GO implementation of mina-payout script",
	Run: func(cmd *cobra.Command, args []string) {
		vars.InfoLog.Println("Welcome to:", vars.Binary)
	},
}

func InitConfig() {
	err := ff.CreateDirectory(vars.ConfigPath)
	if err != nil {
		vars.ErrorLog.Fatal(err)
	}
	err = ff.CreateConfigFile(vars.ConfigFilePath)
	if err != nil {
		vars.ErrorLog.Fatal(err)
	}
	err = ff.CreateFilesInDir(vars.ConfigPath)
	if err != nil {
		vars.ErrorLog.Fatal(err)
	}
}

func ReadTestConfig() {
	result, err := TestConfig()
	if err != nil {
		vars.ErrorLog.Println(err)
		return
	}
	if result.Pool.BlockProducerPublicKey != "" {
		vars.InfoLog.Fatal("Config is empty.")
	} else {
		vars.InfoLog.Println("Config read successfully!")
	}
}

func InitStart() {
	InitConfig()
	//ReadTestConfig()
}

func TestConfig() (config.Config, error) {
	usr, err := user.Current()
	if err != nil {
		return config.Config{}, fmt.Errorf("failed to get current user: %s", err)
	}
	filePath := filepath.Join(usr.HomeDir, vars.ConfigFilePath)

	data, err := os.ReadFile(filePath)
	if err != nil {
		return config.Config{}, fmt.Errorf("failed to read config file: %s", err)
	}

	var configs config.Config
	err = yaml.Unmarshal(data, &configs)
	if err != nil {
		return config.Config{}, fmt.Errorf("failed to unmarshal config: %s", err)
	}

	return configs, nil
}

func init() {
	RootCmd.AddCommand(initCommand)
	RootCmd.AddCommand(calcCommand)
	RootCmd.AddCommand(versionCommand)
	RootCmd.AddCommand(testCommand)
}

var initCommand = &cobra.Command{
	Use:   "init",
	Short: "Initialize the config",
	Long:  "Initialize the config",
	Run: func(cmd *cobra.Command, args []string) {
		InitStart()
	},
}

var calcCommand = &cobra.Command{
	Use:   "calc",
	Short: "Start calculating rewards",
	Long:  "Start calculating rewards",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var versionCommand = &cobra.Command{
	Use:   "version",
	Short: "Print the CLI version",
	Long:  "Print the CLI version",
	Run: func(cmd *cobra.Command, args []string) {
		version.GetVersion()
	},
}

var testCommand = &cobra.Command{
	Use:   "test",
	Short: "Test",
	Long:  "Test",
	Run: func(cmd *cobra.Command, args []string) {
		//address := "B62qqV16g8s744GHM6Dph1uhW4fggYwyvtDnVSoRUyYqNvTir3Rqqzx"
		//decode.ToMD5(address)
		//
		//inv, err := config.ReadCSV(vars.Investors)
		//if err != nil {
		//	fmt.Println("Error when reading CSV file:", err)
		//	return
		//}
		//
		//for _, record := range inv {
		//	println(record)
		//}

		targetAddress := "B62qpJZYLwCjH5Hafi9YiCGGgVhuoq9j6A47MxJG3qzH3nzS3pZZcnn"
		matchingAddress, err := config.FindMatchingAddress(vars.Mina_Foundation_Replacements, targetAddress)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println(matchingAddress)
		}

	},
}
