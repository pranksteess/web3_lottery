package main

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/liangyaopei/bloom"
	"log"
	"sort"
	"strings"
	"time"
	"web3_lottery/config"
)

func genNewAddr() (string, string) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyStr := hexutil.Encode(privateKeyBytes)[2:]

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	return address[2:], privateKeyStr
}

func bloomSearch(l []string)  {
	filter := bloom.New(200000, 14, false)
	for i := 0; i < len(l); i++ {
		filter.AddString(l[i])
	}
	for count := 0;; count ++ {
		addr, pri := genNewAddr()
		ok := filter.TestString(strings.ToLower(addr))
		if ok {
			fmt.Println("CONGRATS! addr: ", addr, " pri: ", pri, " count: ", count)
			break
		} else {
			if count % 10000000 == 0 {
				fmt.Println("count: ", count, " time: ", time.Now().Unix())
			}
		}
	}
}

func binarySearch(l []string)  {
	length := len(l)
	for count := 0;; count ++ {
		addr, pri := genNewAddr()
		index := sort.Search(length, func(i int) bool {
			return l[i] >= strings.ToLower(addr)
		})
		if index < length && l[index] == addr {
			fmt.Println("SUCCESS FOUND! addr: ", addr, " pri: ", pri, " count: ", count)
			break
		} else {
			if count % 10000000 == 0 {
				fmt.Println("count: ", count, " time: ", time.Now().Unix())
			}
		}
	}
}

func main() {
	if err := config.InitCfg(""); err != nil {
		fmt.Println("init failed")
	}
	if config.Cfg.SearchMethod == "binary" {
		sort.Sort(sort.StringSlice(config.Cfg.WhaleETHAddr))
		binarySearch(config.Cfg.WhaleETHAddr)
	} else if config.Cfg.SearchMethod == "bloom" {
		bloomSearch(config.Cfg.WhaleETHAddr)
	}
}
