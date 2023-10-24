package worker

import "sync"

var configMutex sync.RWMutex

func GetPayingWorkerCount() int {
	configMutex.RLock() // Read lock
	defer configMutex.RUnlock()
	return payingWorkerCount
}

func GetFreeWorkerCount() int {
	configMutex.RLock() // Read lock
	defer configMutex.RUnlock()
	return freeWorkerCount
}

func GetRateLimit() int {
	configMutex.RLock() // Read lock
	defer configMutex.RUnlock()
	return rateLimit
}
