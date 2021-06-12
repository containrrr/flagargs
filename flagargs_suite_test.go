package flagargs_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFlagargs(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Flagargs Suite")
}
