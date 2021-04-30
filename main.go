package main

import (
	"adsense-ban-check/core"
	"flag"
)

func main() {
	target := flag.String("target", "", "Target is a URL or domain string.")
	flag.Parse()
	result := *target + " is Not banned by Google Adsense."
	defer core.WriteResult(target, &result)

	if core.IsBanned(target, &result) {
		result = *target + " is banned by Google Adsense."
	}
}
