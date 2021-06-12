package flagargs_test

import (
	"flag"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/pflag"

	fa "containrrr.dev/flagargs"
)

var _ = Describe("Parser", func() {
	Describe("SplitArgs", func() {
		It("should return one known and no extra when one known and one arg is passed", func() {
			known, extra := fa.SplitArgs([]string{"foo"}, "foo")
			Expect(known).To(ContainElements("foo"))
			Expect(extra).To(BeEmpty())
		})
		It("should return one extra and one known when one known and two args are passed", func() {
			known, extra := fa.SplitArgs([]string{"foo", "bar"}, "foo")
			Expect(known).To(ContainElements("foo"))
			Expect(extra).To(ContainElements("bar"))
		})
		It("should return only extra args when no known args are passed", func() {
			known, extra := fa.SplitArgs([]string{"foo"})
			Expect(known).To(BeEmpty())
			Expect(extra).To(ContainElements("foo"))
		})
		When("two known and two args are passed", func() {

			When("passing double dash after fist argument", func() {
				It("should return one extra and one known", func() {
					known, extra := fa.SplitArgs([]string{"foo", "--", "bar"}, "known1", "known2")
					Expect(known).To(ContainElements("foo"))
					Expect(extra).To(ContainElements("bar"))
				})
			})
			When("passing double dash as fist argument", func() {
				It("should return two extra and no known", func() {
					known, extra := fa.SplitArgs([]string{"--", "foo", "bar"}, "known1", "known2")
					Expect(known).To(BeEmpty())
					Expect(extra).To(ContainElements("foo", "bar"))
				})
			})
		})
	})
	Describe("ParseAndUpdateFlags", func() {
		When("two known arguments are passed", func() {

			When("passing double dash after fist argument", func() {
				It("should return one extra and one knownd", func() {
					knownArgs := []string{"first", "second"}
					parser := fa.NewParser(knownArgs...)
					flags := flag.NewFlagSet("", flag.ContinueOnError)
					var known1, known2 string
					flags.StringVar(&known1, knownArgs[0], "", "")
					flags.StringVar(&known2, knownArgs[1], "", "")
					extra, err := parser.ParseAndUpdateFlags(flags, []string{"foo", "--", "bar"})

					Expect(err).NotTo(HaveOccurred())
					Expect(known1).To(Equal("foo"))
					Expect(known2).To(BeEmpty())
					Expect(extra).To(ContainElements("bar"))
				})
			})
		})

		When("failing to assign a known arg", func() {
			It("should return an error", func() {
				parser := fa.NewParser("known1")
				flags := flag.NewFlagSet("", flag.ContinueOnError)
				_, err := parser.ParseAndUpdateFlags(flags, []string{"value"})
				Expect(err).To(HaveOccurred())
			})
		})

		When("using spf13/pflag.FlagSet", func() {
			When("passing double dash after fist argument", func() {
				It("should return one extra and one known", func() {
					parser := fa.NewParser("known1")
					flags := pflag.NewFlagSet("", pflag.ExitOnError)
					var known1, known2 string
					flags.StringVar(&known1, "known1", "", "")
					flags.StringVar(&known2, "known2", "", "")
					extra, err := parser.ParseAndUpdateFlags(flags, []string{"foo", "--", "bar"})
					Expect(err).NotTo(HaveOccurred())
					Expect(known1).To(Equal("foo"))
					Expect(known2).To(BeEmpty())
					Expect(extra).To(ContainElements("bar"))
				})
			})
		})
	})
})
