//+build mage

package main

import "github.com/magefile/mage/sh"

// Build creates binary with Git revision as version string, use with "mage -v build"
func Build() error {
	gitRevision, err := sh.Output("git", "rev-parse", "--short", "HEAD")
	if err != nil {
		return err
	}

	if err := sh.Run("go", "build", "-i", "-ldflags", "-X main.build="+gitRevision, "build_id.go"); err != nil {
		return err
	}

	return nil
}
