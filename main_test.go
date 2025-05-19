package main

import "testing"

// All keys in loadStatus should be 1:1 with storeStatus
func TestLoadStatusConversion(t *testing.T) {
	for str, val := range loadStatusMap {
		_, ok := storeStatusMap[val]
		if !ok {
			t.Error("[loadStatusMap] Case not covered in storeStatusMap:", val, str)
		}
	}
}

// All keys in storeStatus should be 1:1 with loadStatus
func TestStoreStatusConversion(t *testing.T) {
	for statusCode, val := range storeStatusMap {
		_, ok := loadStatusMap[val]
		if !ok {
			t.Error("[storeStatusMap] Case not covered in loadStatusMap: ", val, statusCode)
		}
	}
}
