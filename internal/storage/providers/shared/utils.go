package shared

import (
	"fmt"
	config "terrapak/internal/config/mid"
)

func BuldPathValues(mid config.MID) (prefix, filename string) {
	prefix = fmt.Sprintf("%s/%s/%s/%s", mid.Namespace, mid.Provider, mid.Name, mid.Version)
	filename = fmt.Sprintf("%s.zip", mid.Name)
	return
}