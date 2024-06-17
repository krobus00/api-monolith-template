package cachekey

import "fmt"

func NewRefreshTokenCacheKey(userID, tokenID string) string {
	return fmt.Sprintf("refresh-token:userId:%s:tokenId:%s", userID, tokenID)
}
