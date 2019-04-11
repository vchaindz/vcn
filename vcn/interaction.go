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

package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/ethereum/go-ethereum/common"
	"github.com/fatih/color"
	"github.com/pkg/browser"
	"github.com/sirupsen/logrus"
)

var displayProgress = true

func dashboard() {
	// open dashboard
	// we intentionally do not read the customer's token from disk
	// and GET the dashboard => this would be insecure as tokens would
	// be visible in server logs. in case the anyhow long-running web session
	// has expired the customer will have to log in
	url := DashboardURL()
	fmt.Println(fmt.Sprintf("Taking you to <%s>", url))
	browser.OpenURL(url)
}

func login(in *os.File) {
	if in == nil {
		in = os.Stdin
	}

	// file system: token exists && api: token is valid
	// no => enter email
	//        api: publisher exists
	//        yes => enter pw
	//               authenticate()
	//               fails => retry pw entry up to 3 times
	//        no  => hint at registration
	// filesystem: keystore exists
	// no => createkeystore
	// synckeys

	token, _ := LoadToken()
	tokenValid, _ := CheckToken(token)

	if tokenValid == false {
		email, err := ProvidePlatformUsername()
		if err != nil {
			log.Fatal(err)
		}
		publisherExists := CheckPublisherExists(email)

		if publisherExists {

			LOG.WithFields(logrus.Fields{
				"email": email,
			}).Debug("Publisher exists")

			password, err := ProvidePlatformPassword()
			if err != nil {
				log.Fatal(err)
			}
			_, returnCode := Authenticate(email, password)
			if returnCode == 401 {
				log.Fatal("Invalid password")
			} else if returnCode == 400 {
				log.Fatal("Your email address was not confirmed.\n" +
					"Please confirm it by clicking on the link we sent to " + email + ".\n" +
					"If you did not receive the email, please go to dashboard.codenotary.io and click on the link \"Resend email\"")
			}
		} else {
			fmt.Println("It looks like you have not yet registered.")
			color.Set(StyleAffordance())
			fmt.Printf("Please create an account first at %s", DashboardURL())
			color.Unset()
			fmt.Println()
			dashboard()
			os.Exit(1)
		}

	}

	_ = TrackPublisher("VCN_LOGIN")

	hasKeystore, err := HasKeystore()
	if err != nil {
		LOG.WithFields(logrus.Fields{
			"error": err,
		}).Fatal("Could not access keystore directory")
	}
	if hasKeystore == false {

		fmt.Println("You have no keystore set up yet.")
		fmt.Println("<vcn> will now do this for you and upload the public key to the platform.")

		color.Set(StyleAffordance())
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
				PrintErrorURLCustom("password", 404)
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

		pubKey, wallet := CreateKeystore(keystorePassphrase)

		fmt.Println("Keystore successfully created. We are updating your user profile.\n" +
			"You will be able to sign your first asset in one minute")
		fmt.Println("Public key:\t", pubKey)
		fmt.Println("Keystore:\t", wallet)

	}

	//
	SyncKeys()

	fmt.Println("Login successful.")

	WG.Wait()

}

// Commit => "sign"
func Sign(filename string, state Status, visibility Visibility, quit bool, acknowledge bool) {

	// check for token
	token, _ := LoadToken()
	checkOk, _ := CheckToken(token)
	if !checkOk {
		fmt.Println("You need to be logged in to sign.")
		fmt.Println("Proceed by authenticating yourself using <vcn login>")
		// PrintErrorURLCustom("token", 428)
		os.Exit(1)
	}

	// keystore
	hasKeystore, _ := HasKeystore()
	if hasKeystore == false {
		fmt.Printf("You need a keystore to sign.\n")
		fmt.Println("Proceed by authenticating yourself using <vcn auth>")
		// PrintErrorURLCustom("keystore", 428)
		os.Exit(1)
	}

	var err error
	var artifactHash string
	var fileSize int64 = 0

	if strings.HasPrefix(filename, "docker:") {
		artifactHash, err = GetDockerHash(filename)
		if err != nil {
			log.Fatal("failed to get hash for docker image", err)
		}
		fileSize, err = GetDockerSize(filename)
		if err != nil {
			log.Fatal("failed to get size for docker image", err)
		}
	} else {
		// file mode
		artifactHash = hash(filename)
		fi, err := os.Stat(filename);
		if err != nil {
			log.Fatal(err)
		}
		fileSize = fi.Size()
	}

	reader := bufio.NewReader(os.Stdin)

	if !acknowledge {
		fmt.Println("CodeNotary - code signing in 1 simple step:")
		fmt.Println()
		fmt.Println("Attention, by signing this asset with CodeNotary you implicitly claim its ownership.")
		fmt.Println("Doing this can potentially infringe other publisher's intellectual property under the laws of your country of residence.")
		fmt.Println("vChain and the Zero Trust Consortium cannot be held responsible for legal ramifications.")
		color.Set(color.FgGreen)
		fmt.Println()
		fmt.Println("If you are the owner of the asset (e.g. author, creator, publisher) you can continue")
		color.Unset()
		fmt.Println()
		fmt.Print("I understand and want to continue. (y/n)")
		question, _ := reader.ReadString('\n')
		if strings.ToLower(strings.TrimSpace(question)) != "y" {
			os.Exit(1)
		}
	}

	passphrase, err := ProvideKeystorePassword()
	if err != nil {
		log.Fatal(err)
	}

	go displayLatency()

	_ = TrackPublisher("VCN_SIGN")
	_ = TrackSign(artifactHash, filepath.Base(filename), state)

	// TODO: return and display: block #, trx #
	_, _ = commitHash(artifactHash, passphrase, filepath.Base(filename), fileSize, state, visibility)
	fmt.Println("")
	fmt.Println("Asset:\t", filename)
	fmt.Println("Hash:\t", artifactHash)
	// fmt.Println("Date:\t\t", time.Now())
	// fmt.Println("Signer:\t", "<pubKey>")

	WG.Wait()
	displayProgress = false
	if !quit {
		if _, err := fmt.Scanln(); err != nil {
			log.Fatal(err)
		}
	}
}

func VerifyAll(files []string, quit bool) {
	_ = TrackPublisher("VCN_VERIFY")
	var success = true
	for _, file := range files {
		success = success && verify(file)
	}
	if !quit {
		if _, err := fmt.Scanln(); err != nil {
			log.Fatal(err)
		}
	}
	if success {
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}

func verify(filename string) (success bool) {
	var artifactHash string
	var err error

	if strings.HasPrefix(filename, "docker:") {
		artifactHash, err = GetDockerHash(filename)
		if err != nil {
			log.Fatal("failed to get hash for docker image", err)
		}
	} else {
		artifactHash = strings.TrimSpace(hash(filename))
	}
	_ = TrackVerify(artifactHash, filepath.Base(filename))
	verification, err := BlockChainVerify(artifactHash)
	if err != nil {
		log.Fatal("unable to verify hash", err)
	}

	var artifact *ArtifactResponse
	if verification.Owner != common.BigToAddress(big.NewInt(0)) {
		metaHash := hashAsset(verification)
		if metaHash != "" {
			artifact, _ = LoadArtifactForHash(artifactHash, metaHash)
		}
	}
	if artifact != nil {
		if artifact.Visibility == "PUBLIC" {
			fmt.Println("Asset:\t", artifact.Filename)
			fmt.Println("Hash:\t", artifactHash)
			fmt.Println("Date:\t", verification.Timestamp)
			fmt.Println("Signer:\t", artifact.Publisher)
			fmt.Println("Name:\t", artifact.Name)
			fmt.Println("Size:\t", humanize.Bytes(artifact.FileSize))
			fmt.Println("Level:\t", LevelName(verification.Level))
		}
		if artifact.Visibility == "PRIVATE" {
			fmt.Println("Asset:\t", filepath.Base(filename))
			fmt.Println("Hash:\t", artifactHash)
			fmt.Println("Date:\t", verification.Timestamp)
			fmt.Println("Signer:\t", verification.Owner.Hex())
			fmt.Println("Name:\t", "NA")
			fmt.Println("Size:\t", "NA")
			fmt.Println("Level:\t", LevelName(verification.Level))
		}
	} else {
		fmt.Println("Asset:\t", filepath.Base(filename))
		fmt.Println("Hash:\t", artifactHash)
		fmt.Println("Signer:\t", "NA")
		if verification.Timestamp != time.Unix(0, 0) {
			fmt.Println("Date:\t", verification.Timestamp)
		} else {
			fmt.Println("Date:\t", "NA")
		}
		fmt.Println("Level:\t", "NA")
		fmt.Println("Name:\t", "NA")
		fmt.Println("Size:\t", "NA")
	}
	fmt.Print("Status:\t ")
	if verification.Status == StatusTrusted {
		color.Set(StyleSuccess())
		success = true
	} else if verification.Status == StatusUnknown {
		color.Set(StyleWarning())
		success = false
	} else {
		color.Set(StyleError())
		success = false
	}
	fmt.Print(StatusName(verification.Status))
	color.Unset()
	fmt.Println()
	return success
}

func displayLatency() {
	i := 0
	for displayProgress {
		i++
		fmt.Printf("\033[2K\rIn progress %02dsec", i)
		// fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
}
