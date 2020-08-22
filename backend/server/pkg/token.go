package pkg

import (
	"fmt"
	"os"
	"strings"
)

func verifyToken(token string) bool {
	SECRET := os.Getenv("SECRET")
	return strings.HasPrefix(token, SECRET)
}

func generateToken(username string) string {
	SECRET := os.Getenv("SECRET")
	return fmt.Sprintf("%s|%s", SECRET, username)
}
