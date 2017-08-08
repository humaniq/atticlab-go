package build

import (
	"github.com/humaniq/go/amount"
	"github.com/humaniq/go/support/errors"
	"github.com/humaniq/go/xdr"
)

// SetFee groups the creation of a new SetFeeBuilder with a call to Mutate.
func SetFee(muts ...interface{}) (result SetFeeBuilder) {
	result.Mutate(muts...)
	return
}

// SetFeeMutator is a interface that wraps the
// MutateSetFee operation.  types may implement this interface to
// specify how they modify an xdr.SetFeeOp object
type SetFeeMutator interface {
	MutateSetFee(*xdr.SetFeeOp) (error)
}

// SetFeeBuilder represents a transaction that is being built.
type SetFeeBuilder struct {
	O           xdr.Operation
	E           xdr.SetFeeOp
	Err         error
}

// Mutate applies the provided mutators to this builder's setfee or operation.
func (b *SetFeeBuilder) Mutate(muts ...interface{}) {
	for _, m := range muts {
		var err error

		switch mut := m.(type) {
		case SetFeeMutator:
			err = mut.MutateSetFee(&b.E)
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

// MutateSetFee for NativeAmount sets the SetFeeOp's
// Amount field
func (m NativeAmount) MutateSetFee(o *xdr.SetFeeOp) (err error) {
	baseFee, err := amount.Parse(m.Amount)
	if err == nil{
		o.BaseFee = xdr.Int32(baseFee)
	}
	return
}
