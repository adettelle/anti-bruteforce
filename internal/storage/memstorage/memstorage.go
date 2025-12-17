package memstorage

import (
	"log"
	"time"

	"github.com/adettelle/anti-bruteforce/pkg/bucket"
)

type MemStorage struct {
	// cfg       *config.Config
	Logins           map[string]bucket.LeakyBucket
	Pwds             map[string]bucket.LeakyBucket
	IPs              map[string]bucket.LeakyBucket // ["ip1":LeakyBucket{}, "ip2":LeakyBucket{}]
	capacityForLogin int                           // 10
	capacityForPwd   int                           // 100
	capacityForIP    int                           // 1000
}

type WBStorage struct { // TODO не используется
	// cfg       *config.Config
	WhiteList map[string]bool // ["ip1":true, "ip2":true]
	BlackList map[string]bool // ["ip3":true, "ip4":true]
}

func New(capacityForLogin, capacityForPwd, capacityForIP int) *MemStorage {
	logins := map[string]bucket.LeakyBucket{}
	pwds := map[string]bucket.LeakyBucket{}
	ips := map[string]bucket.LeakyBucket{}

	return &MemStorage{
		Logins:           logins,
		Pwds:             pwds,
		IPs:              ips,
		capacityForLogin: capacityForLogin,
		capacityForPwd:   capacityForPwd,
		capacityForIP:    capacityForIP,
	}
}

func NewWB() *WBStorage {
	wl := map[string]bool{}
	return &WBStorage{WhiteList: wl}
}

func (wb *WBStorage) CheckInWhiteList(ip string) bool {
	_, ok := wb.WhiteList[ip]
	return ok
}

func (ms *MemStorage) Register(login string, pwd string, ip string) bool {
	added := time.Now()

	if existingBucketForLogin, ok := ms.Logins[login]; ok {
		ok := existingBucketForLogin.AddRequest()
		if !ok {
			log.Println("failed to add request by login")
		}
	} else {
		bucketForLogin := bucket.NewLeakyBucket(ms.capacityForLogin, added)
		ms.Logins[login] = *bucketForLogin
		ok := existingBucketForLogin.AddRequest()
		if !ok {
			log.Println("failed to add request by login in existing bucket")
		}
	}
	log.Printf("added by login: %s\n", login)

	if existingBucketForPwd, ok := ms.Pwds[pwd]; !ok {
		ok := existingBucketForPwd.AddRequest()
		if !ok {
			log.Println("failed to add request by password")
		}
	} else {
		bucketForLogin := bucket.NewLeakyBucket(ms.capacityForLogin, added)
		ms.Logins[login] = *bucketForLogin
		ok := existingBucketForPwd.AddRequest()
		if !ok {
			log.Println("failed to add request by password in existing bucket")
		}
	}
	log.Printf("added by password: %s\n", pwd)

	if existingBucketForIP, ok := ms.IPs[ip]; !ok {
		ok := existingBucketForIP.AddRequest()
		if !ok {
			log.Println("failed to add request by ip")
		}
	} else {
		bucketForLogin := bucket.NewLeakyBucket(ms.capacityForLogin, added)
		ms.Logins[login] = *bucketForLogin
		ok := existingBucketForIP.AddRequest()
		if !ok {
			log.Println("failed to add request by ip in existing bucket")
		}
	}
	log.Printf("added by ip: %s\n", ip)

	return true
}
