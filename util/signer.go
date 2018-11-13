package util

import (
	"crypto/sha256"
	"encoding/base64"
	"github.com/btcsuite/btcd/btcec"
)

func BytesToString(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func StringToBytes(str string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(str)
}

type Signer struct {
	PublicKeyBytes  []byte
	PrivateKeyBytes []byte
	PrivateKey      *btcec.PrivateKey
	PublicKey       *btcec.PublicKey
}

func NewSigner(prvKey []byte) (*Signer, error) {
	var err error
	s := new(Signer)
	s.PrivateKeyBytes = prvKey
	s.PublicKeyBytes, err = GeneratePubKeyByPrvKey(s.PrivateKeyBytes)
	if err != nil {
		return nil, err
	}

	s.PrivateKey, s.PublicKey = btcec.PrivKeyFromBytes(btcec.S256(), prvKey)

	return s, nil
}

func (s *Signer) PublicKeyStr() string {
	return base64.StdEncoding.EncodeToString(s.PublicKeyBytes)
}

func (s *Signer) PrivateKeyStr() string {
	return base64.StdEncoding.EncodeToString(s.PrivateKeyBytes)
}

func (s *Signer) Address() string {
	addr, _ := GenerateAddrByPubKey(s.PublicKeyBytes)
	return addr
}

// 签名
func (s *Signer) Signature(data []byte, isHash256 bool) string {
	if !isHash256 {
		dataHash := sha256.Sum256(data)
		data = dataHash[:]
	}

	signature, _ := s.PrivateKey.Sign(data)
	return BytesToString(signature.Serialize())
}

// 验证
func (s *Signer) Verify(sign, data []byte, isHash256 bool) bool {
	signature, err := btcec.ParseDERSignature(sign, btcec.S256())
	if err != nil {
		return false
	}

	if !isHash256 {
		dataHash := sha256.Sum256(data)
		data = dataHash[:]
	}

	return signature.Verify(data, s.PublicKey)
}

// 解密
func (s *Signer) Decrypt(data []byte) ([]byte, error) {
	return btcec.Decrypt(s.PrivateKey, data)
}

// 加密
func (s *Signer) Encrypt(data []byte) ([]byte, error) {
	return btcec.Encrypt(s.PublicKey, data)
}
