package config

import (
	"fmt"

	"github.com/Masterminds/semver"
)

func BuildSafeVersion(version string) string {
	safe_version, err := semver.NewVersion(version)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%d.%d.%d", safe_version.Major(), safe_version.Minor(), safe_version.Patch())
}