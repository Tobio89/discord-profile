package magiclink

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

const TokenSize = 32

type IssuedToken struct {
	RawToken  string
	TokenHash string
}

func IssueToken(pepper string) (IssuedToken, error) {
	tokenBytes := make([]byte, TokenSize)
	if _, err := rand.Read(tokenBytes); err != nil {
		return IssuedToken{}, fmt.Errorf("generate token: %w", err)
	}

	rawToken := base64.RawURLEncoding.EncodeToString(tokenBytes)
	tokenHash := HashTokenBytes(tokenBytes, pepper)

	return IssuedToken{
		RawToken:  rawToken,
		TokenHash: tokenHash,
	}, nil
}

func HashTokenBytes(tokenBytes []byte, pepper string) string {
	mac := hmac.New(sha256.New, []byte(pepper))
	mac.Write(tokenBytes)
	return hex.EncodeToString(mac.Sum(nil))
}

func DecodeToken(rawToken string) ([]byte, error) {
	if rawToken == "" {
		return nil, fmt.Errorf("missing token")
	}

	tokenBytes, err := base64.RawURLEncoding.DecodeString(rawToken)
	if err != nil {
		return nil, fmt.Errorf("invalid token encoding")
	}

	if len(tokenBytes) != TokenSize {
		return nil, fmt.Errorf("invalid token size")
	}

	return tokenBytes, nil
}
