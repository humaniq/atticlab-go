package build

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"bitbucket.attic.pw/hum/go/xdr"
)

var _ = Describe("Transaction Mutators:", func() {

	var (
		subject *TransactionBuilder
		mut     TransactionMutator
	)

	BeforeEach(func() { subject = &TransactionBuilder{} })
	JustBeforeEach(func() { subject.Mutate(mut) })

	Describe("Defaults", func() {
		BeforeEach(func() {
			subject.Mutate(Payment())
			mut = Defaults{}
		})
		It("sets the fee", func() { Expect(subject.TX.Fee).To(BeEquivalentTo(100)) })
		It("sets the network passphrase", func() { Expect(subject.NetworkPassphrase).To(Equal(DefaultNetwork.Passphrase)) })

		Context("on a transaction with 2 operations", func() {
			BeforeEach(func() { subject.Mutate(Payment()) })
			It("sets the fee to 200", func() { Expect(subject.TX.Fee).To(BeEquivalentTo(200)) })
		})
	})

	Describe("MemoHash", func() {
		BeforeEach(func() { mut = MemoHash{[32]byte{0x01}} })
		It("sets a Hash memo on the transaction", func() {
			Expect(subject.TX.Memo.Type).To(Equal(xdr.MemoTypeMemoHash))
			Expect(subject.TX.Memo.MustHash()).To(Equal(xdr.Hash([32]byte{0x01})))
		})
	})

	Describe("MemoID", func() {
		BeforeEach(func() { mut = MemoID{123} })
		It("sets an ID memo on the transaction", func() {
			Expect(subject.TX.Memo.Type).To(Equal(xdr.MemoTypeMemoId))
			Expect(subject.TX.Memo.MustId()).To(Equal(xdr.Uint64(123)))
		})
	})

	Describe("MemoReturn", func() {
		BeforeEach(func() { mut = MemoReturn{[32]byte{0x01}} })
		It("sets a Hash memo on the transaction", func() {
			Expect(subject.TX.Memo.Type).To(Equal(xdr.MemoTypeMemoReturn))
			Expect(subject.TX.Memo.MustRetHash()).To(Equal(xdr.Hash([32]byte{0x01})))
		})
	})

	Describe("MemoText", func() {
		BeforeEach(func() { mut = MemoText{"hello"} })
		It("sets a TEXT memo on the transaction", func() {
			Expect(subject.TX.Memo.Type).To(Equal(xdr.MemoTypeMemoText))
			Expect(subject.TX.Memo.MustText()).To(Equal("hello"))
		})

		Context("a string longer than 255 bytes", func() {
			BeforeEach(func() { mut = MemoText{"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456"} })
			It("sets an error", func() {
				Expect(subject.Err).ToNot(BeNil())
			})
		})
	})

	Describe("AllowTrustBuilder", func() {
		BeforeEach(func() { mut = AllowTrust() })
		It("adds itself to the tx's operations", func() {
			Expect(subject.TX.Operations).To(HaveLen(1))
		})
	})

	Describe("PaymentBuilder", func() {
		BeforeEach(func() { mut = Payment() })
		It("adds itself to the tx's operations", func() {
			Expect(subject.TX.Operations).To(HaveLen(1))
		})
	})

	Describe("SourceAccount", func() {
		Context("with a valid address", func() {
			address := "GAXEMCEXBERNSRXOEKD4JAIKVECIXQCENHEBRVSPX2TTYZPMNEDSQCNQ"
			BeforeEach(func() { mut = SourceAccount{address} })
			It("sets the AccountId correctly", func() {
				var aid xdr.AccountId
				aid.SetAddress(address)
				Expect(subject.TX.SourceAccount.MustEd25519()).To(Equal(aid.MustEd25519()))
			})
		})

		Context("with bad address", func() {
			BeforeEach(func() { mut = SourceAccount{"foo"} })
			It("fails", func() { Expect(subject.Err).To(HaveOccurred()) })
		})
	})

	Describe("Sequence", func() {
		BeforeEach(func() { mut = Sequence{12345} })
		It("succeeds", func() { Expect(subject.Err).NotTo(HaveOccurred()) })
		It("sets the sequence", func() { Expect(subject.TX.SeqNum).To(BeEquivalentTo(12345)) })
	})

	Describe("AutoSequence", func() {
		BeforeEach(func() {
			mock := &MockSequenceProvider{
				Data: map[string]xdr.SequenceNumber{
					"GAXEMCEXBERNSRXOEKD4JAIKVECIXQCENHEBRVSPX2TTYZPMNEDSQCNQ": 2,
				},
			}

			mut = AutoSequence{mock}
		})

		Context("with no source account set", func() {
			It("fails", func() { Expect(subject.Err).To(HaveOccurred()) })
		})

		Context("with a source account set", func() {
			BeforeEach(func() {
				subject.Mutate(SourceAccount{
					"GAXEMCEXBERNSRXOEKD4JAIKVECIXQCENHEBRVSPX2TTYZPMNEDSQCNQ",
				})
			})

			It("succeeds", func() { Expect(subject.Err).NotTo(HaveOccurred()) })
			It("sets the sequence", func() { Expect(subject.TX.SeqNum).To(BeEquivalentTo(3)) })
		})
	})
})
