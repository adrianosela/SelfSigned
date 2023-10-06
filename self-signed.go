package main

import (
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"log"
	"math/big"
	"math/rand"
	"os"
	"time"
)

var randReader = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	generateSelfSignedCert(os.Args[1])
}
func generateSelfSignedCert(cn string) *tls.Certificate {
	template := x509.Certificate{
		SerialNumber: big.NewInt(rand.Int63()), //time seeded serial number
		Subject: pkix.Name{
			CommonName:   cn,
			Country:      []string{"CA"},
			Organization: []string{"AdrianoSela Inc."},
		},
		NotBefore:             time.Now().Add(time.Minute * -5),                         // valid from 5 minutes ago (allow for clock skews)
		NotAfter:              time.Now().Add(time.Duration(time.Hour * 24 * 365 * 10)), // valid until 10 years from now
		IsCA:                  true,
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	//generate a key pair
	key, err := rsa.GenerateKey(randReader, 2048)
	if err != nil {
		log.Fatal(err)
	}
	keyBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	})

	certData, err := x509.CreateCertificate(randReader, &template, &template, key.Public(), key)
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

	if err = os.WriteFile("key.pem", keyBytes, 0644); err != nil {
		log.Fatal(err)
	}

	if err = os.WriteFile("cert.pem", certBytes, 0644); err != nil {
		log.Fatal(err)
	}

	return &cert
}
