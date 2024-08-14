//go:generate service-export ./$GOFILE
//go:generate mockgen -source $GOFILE -destination ./service_mock.go -package $GOPACKAGE
package auth

import "time"

type Service interface {
	GenerateAccessToken(username string, userID int32, expirationTime time.Time, secret []byte) (string, error)
}
