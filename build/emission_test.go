package build

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"bitbucket.attic.pw/hum/go/xdr"
)

var _ = Describe("EmissionBuilder Mutators", func() {

	var (
		subject EmissionBuilder
		mut     interface{}

		address = "GAXEMCEXBERNSRXOEKD4JAIKVECIXQCENHEBRVSPX2TTYZPMNEDSQCNQ"
		bad     = "foo"
	)

	JustBeforeEach(func() {
		subject = EmissionBuilder{}
		subject.Mutate(mut)
	})

	Describe("SourceAccount", func() {
		Context("using a valid stellar address", func() {
			BeforeEach(func() { mut = SourceAccount{address} })

			It("succeeds", func() {
				Expect(subject.Err).NotTo(HaveOccurred())
			})

			It("sets the destination to the correct xdr.AccountId", func() {
				var aid xdr.AccountId
				aid.SetAddress(address)
				Expect(subject.O.SourceAccount.MustEd25519()).To(Equal(aid.MustEd25519()))
			})
		})

		Context("using an invalid stellar address", func() {
			BeforeEach(func() { mut = SourceAccount{bad} })
			It("failed", func() { Expect(subject.Err).To(HaveOccurred()) })
		})
	})

	Describe("NativeAmount", func() {
		Context("amount valid", func() {
			BeforeEach(func() {
				mut = NativeAmount{"50.0"}
			})
			It("sets the amount properly", func() {

				Expect(subject.E.Amount).To(Equal(xdr.Int64(500000000)))
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