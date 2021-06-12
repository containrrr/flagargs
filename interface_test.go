package flagargs_test

import (
	"flag"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/pflag"

	fa "containrrr.dev/flagargs"
)

var _ = Describe("Implementation", func() {

	Describe("stdlib flag.FlagSet", func() {
		It("should implement FlagSet interface", func() {
			_, ok := (interface{}(&flag.FlagSet{})).(fa.FlagSet)
			Expect(ok).To(BeTrue(), "casting failed")
		})
	})

	Describe("spf13/pflags.FlagSet", func() {
		It("should implement FlagSet interface", func() {
			_, ok := (interface{}(&pflag.FlagSet{})).(fa.FlagSet)
			Expect(ok).To(BeTrue(), "casting failed")
		})

		It("should implement FlagSetWithDash interface", func() {
			_, ok := (interface{}(&pflag.FlagSet{})).(fa.FlagSetWithDash)
			Expect(ok).To(BeTrue(), "casting failed")
		})
	})

})
