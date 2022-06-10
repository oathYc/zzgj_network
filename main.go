package main

import (
	"zzgj_network/boot"
)

func main() {
	if err := boot.Boot(); nil != err {
		panic(err)
	}
}
