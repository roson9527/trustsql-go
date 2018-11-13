package util

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ripemd160"
)

const (
	version            = byte(0x00)
	addressChecksumLen = 4
	privateKeyLen      = 32
)

func NewKeyPair() ([]byte, []byte, error) {
	curve := elliptic.P256()
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		return nil, nil, err
	}

	publicKey, err := GeneratePubKeyByPrvKey(privateKey.D.Bytes())
	if err != nil {
		return nil, nil, err
	}

	return privateKey.D.Bytes(), publicKey, nil
}

func GeneratePubKeyByPrvKey(p []byte) ([]byte, error) {
	curve := btcec.S256()
	_, publicKey := btcec.PrivKeyFromBytes(curve, p)
	return publicKey.SerializeCompressed(), nil
}

func GenerateAddrByPubKey(publicKey []byte) (string, error) {
	publicKeyHash, err := HashPublicKey(publicKey)
	if err != nil {
		return "", err
	}

	versionPayload := append([]byte{version}, publicKeyHash...)
	checksum := checksum(versionPayload)

	fullPayload := append(versionPayload, checksum...)
	address := base58.Encode(fullPayload)

	return address, nil
}

func HashPublicKey(publicKey []byte) ([]byte, error) {
	publicKeySHA256 := sha256.Sum256(publicKey)

	RIPEMD160Hash := ripemd160.New()
	_, err := RIPEMD160Hash.Write(publicKeySHA256[:])
	if err != nil {
		return nil, err
	}
	publicKeyRIPEMD160 := RIPEMD160Hash.Sum(nil)

	return publicKeyRIPEMD160, nil
}

func checksum(payload []byte) []byte {
	firstSHA := sha256.Sum256(payload)
	secondSHA := sha256.Sum256(firstSHA[:])

	return secondSHA[:addressChecksumLen]
}
