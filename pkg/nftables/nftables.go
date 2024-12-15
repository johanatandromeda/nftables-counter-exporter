package nftables

import "github.com/google/nftables"

func Test() {
	_, err := nftables.New()
	if err != nil {
		panic(err)
	}
}
