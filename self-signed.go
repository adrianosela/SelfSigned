package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"io/ioutil"
	"log"
	"math/big"
	random "math/rand"
	"os"
	"time"
)

func main() {
	generateSelfSignedCert(os.Args[1])
}

func generateSelfSignedCert(cn string) *tls.Certificate {
	template := x509.Certificate{
		SerialNumber: big.NewInt(random.New(random.NewSource(time.Now().UnixNano())).Int63()), //time seeded serial number
		Subject: pkix.Name{
			CommonName:   cn,
			Country:      []string{"CA"},
			Organization: []string{"AdrianoSela Inc."},
		},
		NotBefore:             time.Now(),                                               //valid from now on
		NotAfter:              time.Now().Add(time.Duration(time.Hour * 24 * 365 * 10)), //valid for 10 years
		IsCA:                  true,
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	//generate a key pair
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatal(err)
	}
	keyBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	})

	certData, err := x509.CreateCertificate(rand.Reader, &template, &template, key.Public(), key)
	if err != nil {
		log.Fatal(err)
	}
	certBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certData,
	})

	cert, err := tls.X509KeyPair(certBytes, keyBytes)
	if err != nil {
		log.Fatal(err)
	}

	if err = ioutil.WriteFile("key.pem", keyBytes, 0644); err != nil {
		log.Fatal(err)
	}

	if err = ioutil.WriteFile("cert.pem", certBytes, 0644); err != nil {
		log.Fatal(err)
	}

	return &cert
}
