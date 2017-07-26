package build

import (
	"bitbucket.attic.pw/hum/go/support/errors"
	"bitbucket.attic.pw/hum/go/xdr"
)

// RestrictAccount groups the creation of a new RestrictAccount with a call to Mutate.
func RestrictAccount(muts ...interface{}) (result RestrictAccountBuilder) {
	result.Mutate(muts...)
	return
}

// RestrictAccountMutator is a interface that wraps the
// MutateRestrictAccount operation.  types may implement this interface to
// specify how they modify an xdr.RestrictAccountOp object
type RestrictAccountMutator interface {
	MutateRestrictAccount(*xdr.RestrictAccountOp) error
}

// RestrictAccountBuilder represents a transaction that is being built.
type RestrictAccountBuilder struct {
	O   xdr.Operation
	SO  xdr.RestrictAccountOp
	Err error
}

// Mutate applies the provided mutators to this builder's payment or operation.
func (b *RestrictAccountBuilder) Mutate(muts ...interface{}) {
	for _, m := range muts {
		var err error
		switch mut := m.(type) {
		case RestrictAccountMutator:
			err = mut.MutateRestrictAccount(&b.SO)
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

// SetBlockIncoming sets BlockIncomingFlag on RestrictAccount operation
func SetBlockIncoming() SetFlag {
	return SetFlag(xdr.AccountFlagsBlockIncoming)
}

// SetBlockOutgoing sets BlockOutgoingFlag on RestrictAccount operation
func SetBlockOutgoing() SetFlag {
	return SetFlag(xdr.AccountFlagsBlockOutgoing)
}

// MutateRestrictAccount for SetFlag sets the RestrictAccountOp's SetFlags field
func (m SetFlag) MutateRestrictAccount(o *xdr.RestrictAccountOp) (err error) {
	if !isFlagValidForRestrict(xdr.AccountFlags(m)) {
		return errors.New("Unknown flag in SetFlag mutator")
	}

	var val xdr.Uint32
	if o.SetFlags == nil {
		val = xdr.Uint32(m)
	} else {
		val = xdr.Uint32(m) | *o.SetFlags
	}
	o.SetFlags = &val
	return
}

// ClearBlockIncoming clears AuthRequiredFlag on RestrictAccount operation
func ClearBlockIncoming() ClearFlag {
	return ClearFlag(xdr.AccountFlagsBlockIncoming)
}

// ClearBlockOutgoing clears AuthRevocableFlag on RestrictAccount operation
func ClearBlockOutgoing() ClearFlag {
	return ClearFlag(xdr.AccountFlagsBlockOutgoing)
}

// MutateRestrictAccount for ClearFlag sets the RestrictAccountOp's ClearFlags field
func (m ClearFlag) MutateRestrictAccount(o *xdr.RestrictAccountOp) (err error) {
	if !isFlagValidForRestrict(xdr.AccountFlags(m)) {
		return errors.New("Unknown flag in SetFlag mutator")
	}

	var val xdr.Uint32
	if o.ClearFlags == nil {
		val = xdr.Uint32(m)
	} else {
		val = xdr.Uint32(m) | *o.ClearFlags
	}
	o.ClearFlags = &val
	return
}

func isFlagValidForRestrict(flag xdr.AccountFlags) bool {
	if flag != xdr.AccountFlagsBlockIncoming &&
		flag != xdr.AccountFlagsBlockOutgoing {
		return false
	}
	return true
}

// MutateSpendFee for Destination sets the RestrictAccountOp's Account
// field
func (m Destination) MutateRestrictAccount(o *xdr.RestrictAccountOp) error {
	return setAccountId(m.AddressOrSeed, &o.Account)
}

