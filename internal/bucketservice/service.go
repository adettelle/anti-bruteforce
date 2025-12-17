package bucketservice

import (
	"log"

	"github.com/adettelle/anti-bruteforce/pkg/bucket"
)

func XXX() {
	logins := map[string]bucket.LeakyBucket{}
	pwds := map[string]bucket.LeakyBucket{}
	ips := map[string]bucket.LeakyBucket{}

	log.Println(logins, pwds, ips)
}
