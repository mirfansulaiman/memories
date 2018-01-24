package main

import (
	"errors"
	"flag"
	"fmt"
	"log"

	"github.com/bovarysme/memories/crypto"
)

var bruteforce bool
var source, dest, ourID, theirID string

func init() {
	flag.BoolVar(&bruteforce, "bruteforce", false, "perform a key-recovery attack on the input file")
	flag.StringVar(&source, "source", "", "path to the input file (e.g. chat-1067048330)")
	flag.StringVar(&dest, "dest", "", "path to the output file")
	flag.StringVar(&ourID, "oid", "", "your MID (e.g. u529a3d0285ef0aa49e713aeac1d2bafb)")
	flag.StringVar(&theirID, "tid", "", "your chat partner's MID")

	flag.Parse()
}

func cmdBruteforce() error {
	if source == "" {
		return errors.New("Error: -source need to be set. See -help for more details.")
	}

	log.Printf("Performing a key-recovery attack on '%s'", source)
	err := crypto.Bruteforce(source)

	return err
}

func cmdDecrypt() error {
	if source == "" || ourID == "" || theirID == "" {
		return errors.New("Error: -source, -oid and -tid need to be set. See -help for more details.")
	}

	if dest == "" {
		dest = fmt.Sprintf("%s.sqlite", source)
	}

	log.Printf("Decrypting '%s' to '%s'\n", source, dest)
	err := crypto.Decrypt(source, dest, ourID, theirID)

	return err
}

func main() {
	var err error
	if bruteforce {
		err = cmdBruteforce()
	} else {
		err = cmdDecrypt()
	}

	if err != nil {
		log.Fatal(err)
	}
}
