package interfaces

type PasswordHasher interface {
	GenerateHash(password string) string
	CheckPassword(password string, hashedPassword string) error
}
