package config

import (
	"fmt"
	"log"

	"github.com/Masterminds/semver"
)

func BuildSafeVersion(version string) string {
	safe_version, err := semver.NewVersion(version)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%d.%d.%d", safe_version.Major(), safe_version.Minor(), safe_version.Patch())
}