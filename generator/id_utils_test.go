package generator

import (
	"testing"
)

func TestIdUtils(t *testing.T) {
	//println("id: ", GetSonyId())

	for i := 0; i < 10; i++ {
		println("id: ", SteamID())
	}
}
