package build

import (
	"bitbucket.attic.pw/hum/go/amount"
	"bitbucket.attic.pw/hum/go/support/errors"
	"bitbucket.attic.pw/hum/go/xdr"
)

// Settlement groups the creation of a new SettlementBuilder with a call to Mutate.
func Settlement(muts ...interface{}) (result SettlementBuilder) {
	result.Mutate(muts...)
	return
}

// SettlementMutator is a interface that wraps the
// MutateSettlement operation.  types may implement this interface to
// specify how they modify an xdr.SettlementOp object
type SettlementMutator interface {
	MutateSettlement(*xdr.SettlementOp) error
}

// SettlementBuilder represents a transaction that is being built.
type SettlementBuilder struct {
	O           xdr.Operation
	S           xdr.SettlementOp
	Err         error
}

// Mutate applies the provided mutators to this builder's Settlement or operation.
func (b *SettlementBuilder) Mutate(muts ...interface{}) {
	for _, m := range muts {
		var err error
		switch mut := m.(type) {
		case SettlementMutator:
			err = mut.MutateSettlement(&b.S)
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

// MutateSettlement for NativeAmount sets the SettlementOp's
// Amount field
func (m NativeAmount) MutateSettlement(o *xdr.SettlementOp) (err error) {
	o.Amount, err = amount.Parse(m.Amount)
	return
}
