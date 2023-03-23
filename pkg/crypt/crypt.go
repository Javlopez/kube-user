package crypt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"log"
	"os"

	"github.com/Javlopez/kube-user/pkg/models"
)

const (
	keySize = 2048
)

type Core struct {
}

func New() *Core {
	return &Core{}
}

func (c Core) CreateRSAPrivateKey(user string) error {
	pk, err := c.GenerateRSAPrivateKey()
	if err != nil {
		return err
	}
	_, err = c.WriteKeyFile(pk, user)
	if err != nil {
		return err
	}

	return nil
}

func (c Core) GenerateRSAPrivateKey() (*rsa.PrivateKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		fmt.Println("Error generating private key:", err)
		return nil, err
	}

	return privateKey, nil
}

func (c Core) WriteKeyFile(privateKey *rsa.PrivateKey, user string) (string, error) {

	// Encode private key in PEM format
	privateKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}

	// Write private key to file
	filename := fmt.Sprintf("%s.key", user)
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return "", err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal("unable to close the file")
		}
	}(file)

	err = pem.Encode(file, privateKeyPEM)
	if err != nil {
		fmt.Println("Error encoding private key:", err)
		return "", err
	}

	fmt.Printf("Private key saved to %s\n", filename)
	return filename, nil
}

func (c Core) CSR(privateKey *rsa.PrivateKey, username string) error {
	// f, err := os.Open("/tmp/dat")
	// check(err)
	// Create a CSR template
	csrTemplate := x509.CertificateRequest{
		Subject: pkix.Name{
			CommonName: username,
		},
		SignatureAlgorithm: x509.SHA256WithRSA,
	}
	// Create the CSR using the private key and the template
	csrBytes, err := x509.CreateCertificateRequest(rand.Reader, &csrTemplate, privateKey)
	if err != nil {
		return (err)
	}

	// Write the CSR to a file
	csrFile, err := os.Create(username + ".csr")
	if err != nil {
		return (err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal("unable to close the file")
		}
	}(csrFile)

	err = pem.Encode(csrFile, &pem.Block{Type: "CERTIFICATE REQUEST", Bytes: csrBytes})
	if err != nil {
		return (err)
	}

	return nil

}

// key
func (c Core) Build(opts models.Options) error {
	pk, err := c.GenerateRSAPrivateKey()
	if err != nil {
		return err
	}
	_, err = c.WriteKeyFile(pk, opts.User)
	if err != nil {
		return err
	}

	err = c.CSR(pk, opts.User)
	if err != nil {
		return err
	}
	return nil
}

/*
// Create a CSR template
    csrTemplate := x509.CertificateRequest{
        Subject: pkix.Name{
            CommonName: username,
        },
        SignatureAlgorithm: x509.SHA256WithRSA,
    }

    // Create the CSR using the private key and the template
    csrBytes, err := x509.CreateCertificateRequest(rand.Reader, &csrTemplate, privateKey)
    if err != nil {
        panic(err)
    }

    // Write the CSR to a file
    csrFile, err := os.Create(username + ".csr")
    if err != nil {
        panic(err)
    }
    pem.Encode(csrFile, &pem.Block{Type: "CERTIFICATE REQUEST", Bytes: csrBytes})
    csrFile.Close()

    // Write the private key to a file
    privateKeyFile, err := os.Create(username + ".key")
    if err != nil {
        panic(err)
    }
    //pem.Encode(privateKeyFile, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)})
    //privateKeyFile.Close()

    fmt.Println("CSR and private key generated successfully.")
*/

/*
openssl genrsa -out [username].key 2048
*/

/*
In this code, we first define the username and key size as variables. Then, we use the rsa.GenerateKey function to generate a new RSA private key with the specified key size. We encode the private key in PEM format using the x509.MarshalPKCS1PrivateKey function and create a pem.Block with the appropriate type and bytes.

Finally, we write the private key to a file with the specified filename using os.Create and pem.Encode. The defer statement ensures that the file is closed when the function returns.
*/
/*
package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func main() {
	username := "your-username"
	keySize := 2048

	// Generate RSA private key
	privateKey, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		fmt.Println("Error generating private key:", err)
		return
	}

	// Encode private key in PEM format
	privateKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}

	// Write private key to file
	filename := fmt.Sprintf("%s.key", username)
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	err = pem.Encode(file, privateKeyPEM)
	if err != nil {
		fmt.Println("Error encoding private key:", err)
		return
	}

	fmt.Printf("Private key saved to %s\n", filename)
}

*/
