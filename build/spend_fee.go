package build

import (
	"github.com/humaniq/go/amount"
	"github.com/humaniq/go/support/errors"
	"github.com/humaniq/go/xdr"
)

// SpendFee groups the creation of a new SpendFeeBuilder with a call to Mutate.
func SpendFee(muts ...interface{}) (result SpendFeeBuilder) {
	result.Mutate(muts...)
	return
}

// SpendFeeMutator is a interface that wraps the
// MutateSpendFee operation.  types may implement this interface to
// specify how they modify an xdr.SpendFeeOp object
type SpendFeeMutator interface {
	MutateSpendFee(*xdr.SpendFeeOp) (error)
}

// SpendFeeBuilder represents a transaction that is being built.
type SpendFeeBuilder struct {
	O           xdr.Operation
	E           xdr.SpendFeeOp
	Err         error
}

// Mutate applies the provided mutators to this builder's spendFee or operation.
func (b *SpendFeeBuilder) Mutate(muts ...interface{}) {
	for _, m := range muts {
		var err error

		switch mut := m.(type) {
		case SpendFeeMutator:
			err = mut.MutateSpendFee(&b.E)
		case OperationMutator:
			err = mut.MutateOperation(&b.O)
		default:
			err = errors.New("Mutator type not allowed")
		}

		if err != nil {
			b.Err = err
			return
		}
	}
}

// MutateSpendFee for Destination sets the SpendFeeOp's Destination
// field
func (m Destination) MutateSpendFee(o *xdr.SpendFeeOp) error {
	return setAccountId(m.AddressOrSeed, &o.Destination)
}

// MutateSpendFee for NativeAmount sets the SpendFeeOp's
// Amount field
func (m NativeAmount) MutateSpendFee(o *xdr.SpendFeeOp) (err error) {
	o.Amount, err = amount.Parse(m.Amount)
	return
}
