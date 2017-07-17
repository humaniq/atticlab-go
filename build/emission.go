package build

import (
	"bitbucket.attic.pw/hum/go/amount"
	"bitbucket.attic.pw/hum/go/support/errors"
	"bitbucket.attic.pw/hum/go/xdr"
)

// Emission groups the creation of a new EmissionBuilder with a call to Mutate.
func Emission(muts ...interface{}) (result EmissionBuilder) {
	result.Mutate(muts...)
	return
}

// EmissionMutator is a interface that wraps the
// MutateEmission operation.  types may implement this interface to
// specify how they modify an xdr.EmissionOp object
type EmissionMutator interface {
	MutateEmission(*xdr.EmissionOp) (error)
}

// EmissionBuilder represents a transaction that is being built.
type EmissionBuilder struct {
	O           xdr.Operation
	E           xdr.EmissionOp
	Err         error
}

// Mutate applies the provided mutators to this builder's emission or operation.
func (b *EmissionBuilder) Mutate(muts ...interface{}) {
	for _, m := range muts {
		var err error

		switch mut := m.(type) {
		case EmissionMutator:
			err = mut.MutateEmission(&b.E)
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

// MutateEmission for Destination sets the EmissionOp's Destination
// field
func (m Destination) MutateEmission(o *xdr.EmissionOp) error {
	return setAccountId(m.AddressOrSeed, &o.Destination)
}

// MutateEmission for NativeAmount sets the EmissionOp's
// Amount field
func (m NativeAmount) MutateEmission(o *xdr.EmissionOp) (err error) {
	o.Amount, err = amount.Parse(m.Amount)
	return
}
