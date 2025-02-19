package version

import (
	"github.com/Masterminds/semver/v3"
)

func ValidateVersion(version string) error {
	_, err := semver.NewVersion(version)

	if err == nil {
		return nil
	} else {
		return NewInvalidVersionError(version)
	}
}
