package build

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/humaniq/atticlab-go/xdr"
)

var _ = Describe("SettlementBuilder Mutators", func() {

	var (
		subject SettlementBuilder
		mut     interface{}
	)

	JustBeforeEach(func() {
		subject = SettlementBuilder{}
		subject.Mutate(mut)
	})

	Describe("NativeAmount", func() {
		Context("amount valid", func() {
			BeforeEach(func() {
				mut = NativeAmount{"50.0"}
			})
			It("sets the amount properly", func() {

				Expect(subject.S.Amount).To(Equal(xdr.Int64(500000000)))
			})
			It("succeeds", func() {
				Expect(subject.Err).NotTo(HaveOccurred())
			})
		})


		Context("amount invalid", func() {
			BeforeEach(func() {
				mut = NativeAmount{"test"}
			})

			It("failed", func() {
				Expect(subject.Err).To(HaveOccurred())
			})
		})

	})
})