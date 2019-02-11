package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"syscall"
	"time"

	"github.com/dghubble/sling"
	"golang.org/x/crypto/ssh/terminal"
)

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokenResponse struct {
	Token string `token:"token"`
}

type Error struct {
	Status    int    `json:"status"`
	Message   string `json:"message"`
	Path      string `json:"path"`
	Timestamp string `json:"timestamp"`
	Error     string `json:"error"`
}

func Authenticate() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter your email address: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSuffix(email, "\n")

	token := new(TokenResponse)
	authError := new(Error)
	fmt.Print("Password:")
	password, err := terminal.ReadPassword(int(syscall.Stdin))
	fmt.Println(".")

	if err != nil {
		log.Fatal(err)
	}
	r, err := sling.New().
		Post(PublisherEndpoint()+"/auth").
		BodyJSON(AuthRequest{Email: email, Password: string(password)}).
		Receive(token, authError)
	if err != nil {
		log.Fatal(err)
	}
	if r.StatusCode != 200 {
		log.Fatalf("request failed: %s (%d)", authError.Message,
			authError.Status)
	}
	err = ioutil.WriteFile(TokenFile(), []byte(token.Token),
		os.FileMode(0600))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Authentication successful.")

	// check for a keystore right now and hint at creating one
	createKeystoreForDashboardUsers()

}

func createKeystoreForDashboardUsers() {
	// this is for the case of a newly registered customer
	// coming in from the dashboard
	// having authenticatedat vcn and no keystore is yet present
	hasKeystore, err := HasKeystore()
	if err != nil {
		log.Fatal(err)
	}
	if hasKeystore == false {
		fmt.Println("You have no keystore set up yet. <vcn> will now do this for you and upload the public key to the platform.")

		keystorePassphrase, err := readPassword("Keystore passphrase:")
		if err != nil {
			log.Fatal(err)
		}

		CreateKeystore(keystorePassphrase)
		SyncKeys()

	}

}

func Register() {
	var keystorePassphrase string
	authError := new(Error)
	hasKeystore, err := HasKeystore()
	if err != nil {
		log.Fatal(err)
	}

	// get email
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("User registration for vChain.us: New account")
	fmt.Println("If you already have an account pls abort and simply authenticate yourself using <vcn auth>")
	fmt.Print("Please enter your email address: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSuffix(email, "\n")

	accountPassword, err := readPassword("Account password:")
	if err != nil {
		log.Fatal(err)
	}
	if !hasKeystore {
		keystorePassphrase, err = readPassword("Keystore passphrase:")
		if err != nil {
			log.Fatal(err)
		}
	}
	r, err := sling.New().
		Post(PublisherEndpoint()).
		BodyJSON(AuthRequest{Email: email, Password: accountPassword}).
		Receive(nil, authError)
	if err != nil {
		log.Fatal(err)
	}
	if r.StatusCode != 200 {
		log.Fatalf("request failed: %s (%d)", authError.Message,
			authError.Status)
	}
	log.Println("We've sent you an email to: ", email,
		"Click the link and you will be automatically logged in")
	if hasKeystore {
		SyncKeys()
		return
	} else {
		CreateKeystore(keystorePassphrase)
		err = WaitForConfirmation(email, accountPassword,
			60, 2*time.Second)
		if err != nil {
			log.Fatal(err)
		}
		SyncKeys()
	}
}

func WaitForConfirmation(email string, password string, maxRounds uint64,
	pollInterval time.Duration) error {
	token := new(TokenResponse)
	authError := new(Error)
	for i := uint64(0); i < maxRounds; i++ {
		r, err := sling.New().
			Post(PublisherEndpoint()+"/auth").
			BodyJSON(AuthRequest{Email: email, Password: password}).
			Receive(token, authError)
		if err != nil {
			return err
		}
		if r.StatusCode == 400 {
			time.Sleep(pollInterval)
			continue
		}
		if r.StatusCode == 200 {
			err = ioutil.WriteFile(TokenFile(), []byte(token.Token),
				os.FileMode(0600))
			if err != nil {
				return err
			}
			return nil
		}
		return fmt.Errorf("wait for confirmation failed: %s (%d)",
			authError.Message, authError.Status)
	}
	return fmt.Errorf("confirmation timed out")
}

func LoadToken() (jwtToken string, err error) {
	contents, err := ioutil.ReadFile(TokenFile())
	if err != nil {
		return "", err
	}
	return string(contents), nil
}