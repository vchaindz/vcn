/*
 * Copyright (c) 2018-2020 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package api

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	sdk "github.com/vchain-us/ledger-compliance-go/grpcclient"
	"github.com/vchain-us/vcn/pkg/store"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
	"io/ioutil"
	"strconv"
	"time"
)

func NewLcClient(lcApiKey, host, port, lcCertPath string) (*sdk.LcClient, error) {
	p, _ := strconv.Atoi(port)

	dialOptions := []grpc.DialOption{
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                20 * time.Second,
			Timeout:             10 * time.Second,
			PermitWithoutStream: true,
		}),
	}
	if lcCertPath != "" {
		tlsCredentials, err := loadTLSCertificate(lcCertPath)
		if err != nil {
			return nil, fmt.Errorf("cannot load TLS credentials: %s", err)
		}
		dialOptions = append(dialOptions, grpc.WithTransportCredentials(tlsCredentials))
	} else {
		dialOptions = append(dialOptions, grpc.WithInsecure())
	}

	return sdk.NewLcClient(sdk.ApiKey(lcApiKey), sdk.Host(host), sdk.Port(p), sdk.Dir(store.CurrentConfigFilePath()), sdk.DialOptions(dialOptions)), nil
}

func loadTLSCertificate(certPath string) (credentials.TransportCredentials, error) {
	cert, err := ioutil.ReadFile(certPath)
	if err != nil {
		return nil, err
	}
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(cert) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}
	config := &tls.Config{
		RootCAs: certPool,
	}
	return credentials.NewTLS(config), nil
}
