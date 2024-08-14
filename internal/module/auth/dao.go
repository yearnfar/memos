//go:generate mockgen -source $GOFILE -destination ./dao_mock.go -package $GOPACKAGE
package auth

type DAO interface {
}
