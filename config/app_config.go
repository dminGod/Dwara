package config

import (
	"os"
	"runtime"
	"errors"
	"github.com/spf13/viper"
	"fmt"
	"flag"
	"strings"
)


type Http struct {
	HttpEnabled bool
	HttpPort int
}

type Dwara_config struct {
	LogFolder string
}

type DwConfig struct {
	Http Http
	Dwara_config Dwara_config
}


// Dont need to add the .toml in the name here
var configFile string = "dwara.toml"

// With trailing slash
var linuxConfigFolders []string = []string{"/etc/"}
var windowsConfigFolders []string = []string{"\\dwara\\"}
var ViConfig *viper.Viper



/*
	This module handles the application configuration

	configuration can be passed from the configuration file
	also can be set using flags.

	location of the configuration file will either be a standard location or will using command line flags.

 */


func GetConfig(){

//	viper.AutomaticEnv('')

	configFile, err := getConfigFile()

	if err != nil {

		fmt.Println("There were errors during fetching application config")
		for _, e := range err {

			fmt.Println(e.Error())
		}

		os.Exit(1)
	}

	ViConfig = viper.New()

	ViConfig.SetConfigFile(configFile)
	ViConfig.AutomaticEnv()

	verr := ViConfig.ReadInConfig()

	if verr != nil {

		fmt.Println("There was an error reading in configuration. Error : ", verr.Error())
	}

	// Show all configuration...
	fmt.Println(ViConfig.AllSettings())
}

func ShowConfig() {

	fmt.Println(ViConfig.AllSettings())

}


/*
	Check in common default locations if the config file is available
	Check if there is an environment variable set
	Check if there is a config file passed

	Command line flags will override environment variables and common loctions
		-- If passed and file not found then application will exit with 1

	Environment variables will override common locations
		-- If environment variables set and file not found app will exit with 1

	If flags and env variables not found then common locations will be used

	If config file not found in any of the places then app will exit with  code 1

 */
func getConfigFile() (retFilePath string, retErrors []error) {

	var flagConfigFile string
	flag.StringVar(&flagConfigFile, "config_file", "", "Customize the path of the configuration file using this flag.")
	flag.Parse()
	fmt.Println("The flag config file", flagConfigFile)

	envConfigFile := os.Getenv("DWARA_CONFIG_FILE")


	// If we are in windows, check the folders we generally put stuff in
	if runtime.GOOS == "windows" {

		winFolders := getAllWindowsDrives()

		if len(winFolders) == 0 {

			retErrors = append(retErrors, errors.New("Windows, no drives found when searching for config."))
		}

		FileChecking:
		// Get the drive letters, loop over them and check which folder seems to be correct
		for _, curDrive := range getAllWindowsDrives() {

			// Loop over all the common windows folders
			for _, curFolder := range windowsConfigFolders {

				curFile := curDrive + ":" + curFolder + configFile

				fmt.Println("Checking for ", curFile)

				if fileExists(curFile) {

					// Config file found
					retFilePath = curFile
					break FileChecking

				}
			}
		}
	}

	// If we are in linux, check the folders for that
	if runtime.GOOS == "linux" {

		for _, curFolder := range linuxConfigFolders {

			curFile := curFolder + configFile

			if fileExists(curFile) {

				retFilePath = curFile
				break
			}
		}
	}


	if len(flagConfigFile) > 0 {

		if fileExists(flagConfigFile) {

			retFilePath = flagConfigFile
		} else {

			retErrors = append(retErrors, errors.New("Unable to locate the config_file file '" + flagConfigFile + "' exiting."))
		}
	} else {

		if len(envConfigFile) > 0 {

			// In windows if you use set DWARA_CONFIG_FILE="C:\..." it adds the " to the string also, removing that
			envConfigFile = strings.Trim(envConfigFile, `"`)

			if fileExists(envConfigFile) {

				retFilePath = envConfigFile
			} else {

				retErrors = append(retErrors, errors.New("Unable to locate the config file that is set on environment variable DWARA_CONFIG_FILE " + envConfigFile + " exiting"))
			}
		}
	}

	// In the end if you still don't have anything you have an error
	if retFilePath == "" {

		retErrors = append(retErrors, errors.New("No configuration files named " + configFile + " found."))
	}

	return
}

func getAllWindowsDrives() (availDrives []string) {

	for _, drive := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ"{

		_, err := os.Open(string(drive)+":\\")

		if err == nil {
			availDrives = append(availDrives, string(drive))
		}
	}

	return
}


func fileExists(curFile string) (retBool bool) {

	_, err := os.Stat( curFile )

	if err == nil {

		retBool = true
	} else {

		fmt.Println("Error locating file, " + curFile, err.Error())
	}



	return
}

