package vars

import (
	"log"
	"os"
)

var (
	InfoLog        = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile|log.Lmsgprefix)
	ErrorLog       = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile|log.Lmsgprefix)
	Binary         = "mina-pay"
	Version        = "1.0.0"
	ConfigPath     = ".mina-pay/"
	ConfigFilePath = ConfigPath + "config.yaml"
)

const (
	MinaBurnAddress              = "B62qiburnzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzmp7r7UN6X"
	Investors                    = "./data/Investors_Addresses.csv"
	Mina_Foundation_Addresses    = "./data/Mina_Foundation_Addresses.csv"
	O1_Labs_Addresses            = "./data/O1_Labs_Addresses.csv"
	Mina_Foundation_Replacements = "./data/Mina_Foundation_Replacements"

	SubstitutePayTo  = ".substitutePayTo"
	NegotiatedFees   = ".negotiatedFees"
	NegotiatedBurn   = ".negotiatedBurn"
	BurnSupercharged = ".burnSupercharged"
)
