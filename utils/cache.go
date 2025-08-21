package utils

import (
	"sync"
	"time"

	"github.com/patrickmn/go-cache"
)

var (
	BalanceCache = cache.New(10*time.Second, 15*time.Second)
	WalletLocks  = sync.Map{} // map[string]*sync.Mutex
)

func GetWalletLock(wallet string) *sync.Mutex {
	lock, _ := WalletLocks.LoadOrStore(wallet, &sync.Mutex{})
	return lock.(*sync.Mutex)
}