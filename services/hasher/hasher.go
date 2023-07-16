package hasher

type Hasher interface {
	HashItem() (string, error)
	VerifyItem(hash string) bool
}
