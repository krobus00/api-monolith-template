package cachekey

import "fmt"

func NewUserByIdentifierCacheKey(identifier string) string {
	return fmt.Sprintf("user:identifier:%s", identifier)
}

func NewUserByIDCacheKey(id string) string {
	return fmt.Sprintf("user:id:%s", id)
}

func NewUserNonPrimaryKeyCacheKeysPatterns() []string {
	return []string{NewUserByIdentifierCacheKey("*")}
}
