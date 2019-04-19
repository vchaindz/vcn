/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 *
 * User Interaction
 *
 * This part of the vcn code handles the concern of interaction (the *V*iew)
 *
 */

package cli

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
	"github.com/vchain-us/vcn/internal/errors"
	"github.com/vchain-us/vcn/pkg/api"
	"github.com/vchain-us/vcn/pkg/logs"
	"github.com/vchain-us/vcn/pkg/meta"
)

func Login() {
	token, _ := api.LoadToken()
	tokenValid, err := api.CheckToken(token)
	if err != nil {
		log.Fatal(err)
	}
	if !tokenValid {
		email, err := ProvidePlatformUsername()
		if err != nil {
			log.Fatal(err)
		}
		publisherExists, err := api.CheckPublisherExists(email)
		if err != nil {
			log.Fatal(err)
		}
		if publisherExists {
			password, err := ProvidePlatformPassword()
			if err != nil {
				log.Fatal(err)
			}
			err = api.Authenticate(email, password)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatal("no such user. please create an account at: ", meta.DashboardURL())
		}
	}

	_ = api.TrackPublisher(meta.VcnLoginEvent)

	hasKeystore, err := api.HasKeystore()
	if err != nil {
		logs.LOG.WithFields(logrus.Fields{
			"error": err,
		}).Fatal("Could not access keystore directory")
	}
	if hasKeystore == false {

		fmt.Println("You have no keystore set up yet.")
		fmt.Println("<vcn> will now do this for you and upload the public key to the platform.")

		color.Set(meta.StyleAffordance())
		fmt.Print("Attention: Please pick a strong passphrase. There is no recovery possible.")
		color.Unset()
		fmt.Println()

		var keystorePassphrase string
		var keystorePassphrase2 string

		match := false
		counter := 0
		for match == false {

			counter++

			if counter == 4 {
				fmt.Println("Too many attempts failed.")
				errors.PrintErrorURLCustom("password", 404)
				os.Exit(1)

			}

			// TODO: solution for reading from file inputs whose compilation does not fail on windows
			// if terminal.IsTerminal(syscall.Stdin) {

			keystorePassphrase, _ = readPassword("Keystore passphrase: ")
			keystorePassphrase2, _ = readPassword("Keystore passphrase (reenter): ")
			fmt.Println("")
			/*} else {

				keystorePassphrase, _ = reader.ReadString('\n')
				keystorePassphrase = strings.TrimSuffix(keystorePassphrase, "\n")

				keystorePassphrase2, _ = reader.ReadString('\n')
				keystorePassphrase2 = strings.TrimSuffix(keystorePassphrase2, "\n")
			}*/

			if keystorePassphrase == "" {
				fmt.Println("Your passphrase must not be empty.")
			} else if keystorePassphrase != keystorePassphrase2 {
				fmt.Println("Your two inputs did not match. Please try again.")
			} else {
				match = true
			}

		}

		pubKey, wallet := api.CreateKeystore(keystorePassphrase)

		fmt.Println("Keystore successfully created. We are updating your user profile.\n" +
			"You will be able to sign your first asset in one minute")
		fmt.Println("Public key:\t", pubKey)
		fmt.Println("Keystore:\t", wallet)

	}

	//
	api.SyncKeys()

	fmt.Println("Login successful.")
}
