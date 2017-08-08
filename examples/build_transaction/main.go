package main

import (
	"fmt"

	b "github.com/humaniq/go/build"
	"github.com/humaniq/go/clients/horizon"
)

func main() {
	// address: GB6S3XHQVL6ZBAF6FIK62OCK3XTUI4L5Z5YUVYNBZUXZ4AZMVBQZNSAU
	from := "SCRUYGFG76UPX3EIUWGPIQPQDPD24XPR3RII5BD53DYPKZJGG43FL5HI"

	// seed: SDLJZXOSOMKPWAK4OCWNNVOYUEYEESPGCWK53PT7QMG4J4KGDAUIL5LG
	to := "GA3A7AD7ZR4PIYW6A52SP6IK7UISESICPMMZVJGNUTVIZ5OUYOPBTK6X"

	tx := b.Transaction(
		b.SourceAccount{from},
		b.AutoSequence{horizon.DefaultTestNetClient},
		b.Payment(
			b.Destination{to},
			b.NativeAmount{"0.1"},
		),
	)

	txe := tx.Sign(from)
	txeB64, err := txe.Base64()

	if err != nil {
		panic(err)
	}

	fmt.Printf("tx base64: %s", txeB64)
}
