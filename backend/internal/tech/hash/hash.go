package hash

import "github.com/rs/zerolog/log"

var (
	SecretKey []byte
)

func NewSecretKey(secretKey string) {
	if secretKey == "" {
		log.Fatal().Msg("error: no secret key provided")
	}

	SecretKey = []byte(secretKey)
}
