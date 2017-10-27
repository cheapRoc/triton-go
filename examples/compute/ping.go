package main

import (
	"context"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	triton "github.com/joyent/triton-go"
	"github.com/joyent/triton-go/authentication"
	"github.com/joyent/triton-go/compute"
)

func main() {
	keyID := os.Getenv("SDC_KEY_ID")
	accountName := os.Getenv("SDC_ACCOUNT")
	keyMaterial := os.Getenv("SDC_KEY_MATERIAL")

	var signer authentication.Signer
	var err error

	if keyMaterial == "" {
		signer, err = authentication.NewSSHAgentSigner(keyID, accountName)
		if err != nil {
			log.Fatalf("Error Creating SSH Agent Signer: {{err}}", err)
		}
	} else {
		var keyBytes []byte
		if _, err = os.Stat(keyMaterial); err == nil {
			keyBytes, err = ioutil.ReadFile(keyMaterial)
			if err != nil {
				log.Fatalf("Error reading key material from %s: %s",
					keyMaterial, err)
			}
			block, _ := pem.Decode(keyBytes)
			if block == nil {
				log.Fatalf(
					"Failed to read key material '%s': no key found", keyMaterial)
			}

			if block.Headers["Proc-Type"] == "4,ENCRYPTED" {
				log.Fatalf(
					"Failed to read key '%s': password protected keys are\n"+
						"not currently supported. Please decrypt the key prior to use.", keyMaterial)
			}

		} else {
			keyBytes = []byte(keyMaterial)
		}

		signer, err = authentication.NewPrivateKeySigner(keyID, []byte(keyMaterial), accountName)
		if err != nil {
			log.Fatalf("error creating SSH key signer: %s", err)
		}
	}

	config := &triton.ClientConfig{
		TritonURL:   os.Getenv("SDC_URL"),
		AccountName: accountName,
		Signers:     []authentication.Signer{signer},
	}

	c, err := compute.NewClient(config)
	if err != nil {
		log.Fatalf("couldn't create new compute client: %s", err)
	}

	resp, err := c.Ping(context.Background())
	if err != nil {
		log.Fatalf("couldn't ping endpoint: %s", err)
	}

	fmt.Printf("Ping: %s\n", resp.Ping)
	fmt.Printf("Versions: %v\n", resp.CloudAPI.Versions)
}
