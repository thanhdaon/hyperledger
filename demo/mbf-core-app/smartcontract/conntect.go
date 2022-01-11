package smartcontract

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
	"github.com/pkg/errors"
)

func Connect() (*gateway.Contract, error) {
	wallet, err := prepareWallet()
	if err != nil {
		return nil, errors.Wrap(err, "prepare wallet fail: ")
	}

	ccpPath := filepath.Join(
		"..",
		"network",
		"organizations",
		"peerOrganizations",
		"org1.example.com",
		"connection-org1.yaml",
	)

	gw, err := gateway.Connect(
		gateway.WithConfig(config.FromFile(filepath.Clean(ccpPath))),
		gateway.WithIdentity(wallet, "appUser"),
	)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to connect to gateway: ")
	}
	defer gw.Close()

	network, err := gw.GetNetwork("mychannel")
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get network: ")
	}

	return network.GetContract("phonecard"), nil
}

func prepareWallet() (*gateway.Wallet, error) {
	if err := os.Setenv("DISCOVERY_AS_LOCALHOST", "true"); err != nil {
		return nil, errors.Wrap(err, "Error setting DISCOVERY_AS_LOCALHOST environemnt variable: ")
	}

	wallet, err := gateway.NewFileSystemWallet("wallet")
	if err != nil {
		return nil, errors.Wrap(err, "create wallet: ")
	}

	if !wallet.Exists("appUser") {
		if err := populateWallet(wallet); err != nil {
			return nil, errors.Wrap(err, "Failed to populate wallet contents")
		}
	}

	return wallet, nil
}

func populateWallet(wallet *gateway.Wallet) error {
	credPath := filepath.Join(
		"..",
		"network",
		"organizations",
		"peerOrganizations",
		"org1.example.com",
		"users",
		"User1@org1.example.com",
		"msp",
	)

	certPath := filepath.Join(credPath, "signcerts", "cert.pem")
	// read the certificate pem
	cert, err := ioutil.ReadFile(filepath.Clean(certPath))
	if err != nil {
		return err
	}

	keyDir := filepath.Join(credPath, "keystore")
	// there's a single file in this dir containing the private key
	files, err := ioutil.ReadDir(keyDir)
	if err != nil {
		return err
	}
	if len(files) != 1 {
		return fmt.Errorf("keystore folder should have contain one file")
	}

	keyPath := filepath.Join(keyDir, files[0].Name())
	key, err := ioutil.ReadFile(filepath.Clean(keyPath))
	if err != nil {
		return err
	}

	identity := gateway.NewX509Identity("Org1MSP", string(cert), string(key))

	return wallet.Put("appUser", identity)
}
