package hyperledger

import (
	"fmt"
	"testing"
)

func Test_ConnectionTest_Success(t *testing.T) {

	clients := NewClientMap(
		"test-network",
		"/Users/kongnb2k/Entronica/learn/hyperledger-fabric-application/back-end/hyperledger/config.yaml",
	)

	_, err := clients.AddClient(
		"Admin",
		"org1",
		"mainchannel",
	)
	if err != nil {
		t.Fatal(err)
		return
	}

	res, err := clients.Query("org1", "rawresources", "index", [][]byte{
		[]byte(""),
		[]byte(""),
	})
	if err != nil {
		t.Fatal(err)
		return
	}

	fmt.Println(string(res))
}
