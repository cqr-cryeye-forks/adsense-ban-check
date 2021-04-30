package main

import (
	"flag"
	"github.com/fabelx/adsense-ban-check/core"
)

func main() {
	target := flag.String("target", "", "Target is a URL or domain string.")
	flag.Parse()
	result := core.Result{Target: *target, Data: *target + " is Not banned by Google Adsense."}
	defer core.WriteResult(&result)

	if core.IsBanned(&result) {
		result.Data = *target + " is banned by Google Adsense."
		result.IsBanned = true
	}
}
