package keys

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/gob"
	"encoding/pem"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

func GenerateKeys() error {
	reader := rand.Reader
	bitSize := 2048

	key, err := rsa.GenerateKey(reader, bitSize)
	if err != nil {
		log.Error(err)
		return err
	}

	publicKey := key.PublicKey

	err = saveGobKey("private.key", key)
	if err != nil {
		return err
	}

	err = savePEMKey("private.pem", key)
	if err != nil {
		return err
	}

	err = saveGobKey("public.key", publicKey)
	if err != nil {
		return err
	}

	err = savePublicPEMKey("public.pem", publicKey)
	if err != nil {
		return err
	}

	return nil
}

func LoadKeys(privateKeyFile string, publicKeyFile string) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	signBytes, err := ioutil.ReadFile(privateKeyFile)
	if err != nil {
		log.Error("unable to load private key file")
		return nil, nil, err
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		log.Error("unable to parse private key")
		return nil, nil, err
	}

	log.Infof("loading public key from %s\n", publicKeyFile)
	verifyBytes, err := ioutil.ReadFile(publicKeyFile)
	if err != nil {
		log.Error("unable to load public key file")
		return nil, nil, err
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		log.Error("unable to parse public key")
		return nil, nil, err
	}

	return privateKey, publicKey, nil
}

func saveGobKey(fileName string, key interface{}) error {
	outFile, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer outFile.Close()

	encoder := gob.NewEncoder(outFile)
	return encoder.Encode(key)
}

func savePEMKey(fileName string, key *rsa.PrivateKey) error {
	outFile, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer outFile.Close()

	var privateKey = &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}

	return pem.Encode(outFile, privateKey)
}

func savePublicPEMKey(fileName string, pubkey rsa.PublicKey) error {
	asn1Bytes, err := x509.MarshalPKIXPublicKey(&pubkey)
	if err != nil {
		return err
	}

	var pemkey = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: asn1Bytes,
	}

	pemfile, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer pemfile.Close()

	return pem.Encode(pemfile, pemkey)
}

func formatFingerprint(raw string) string {
	var buf bytes.Buffer

	for i, c := range raw {
		buf.WriteByte(byte(c))
		if (i+1)%2 == 0 && i != len(raw)-1 {
			buf.WriteByte(byte(':'))
		}
	}

	return buf.String()
}

func Fingerprint(pubkey rsa.PublicKey) (string, error) {
	asn1Bytes, err := x509.MarshalPKIXPublicKey(&pubkey)
	if err != nil {
		return "", err
	}

	return formatFingerprint(fmt.Sprintf("%x", sha256.Sum256(asn1Bytes))), nil
}
