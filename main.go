package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/sha3"
)

func calculate(id string, pw string, site string) string {
	hash := sha3.New256()
	hash.Write([]byte("salt, pepper, birds, and the thought police"))
	hash.Write([]byte(site))
	hash.Write([]byte(id))
	hash.Write([]byte(pw))

	hash_result := hash.Sum(nil)

	var fixlen_result [12]byte
	offset := 0
	for rem_hash_res := hash_result; len(rem_hash_res) != 0; {
		fixlen_result[offset] = fixlen_result[offset] ^ rem_hash_res[0]

		rem_hash_res = rem_hash_res[1:]
		offset += 1
		if offset == 12 {
			offset = 0
		}
	}

	result := base58.Encode(fixlen_result[:])
	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("root id: ")
	id, _ := reader.ReadString('\n')

	fmt.Print("root password: ")
	root_pw, _ := reader.ReadString('\n')

	for {
		fmt.Print("site: ")
		site, _ := reader.ReadString('\n')

		fmt.Println(calculate(id, root_pw, site))
	}
}
