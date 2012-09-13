//package hashring
package main

import (
    "fmt";
    "crypto/md5";
    "hash";
)

// Creates an md5 based int value for the given string.
func generateKey(key string) []uint8 {
    var h hash.Hash = md5.New();
    h.Write([]byte(key));
    return h.Sum(nil);
}

func main() {
    key := "key_name";
    hex := generateKey(key);
    fmt.Printf("%s: %x\n", key, hex)
}
