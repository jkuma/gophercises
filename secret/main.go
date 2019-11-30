package main

import (
	"fmt"
	"secret/secret"
)

func main() {
	vault := &secret.FileVault{
		Key:      []byte("6368616e676520746869732070617373"),
		Filename: "./test.secret",
	}

	err := vault.Set("twitter_aPi", "jesaispas")
	err = vault.Set("facebook_api", "facebookpassword")
	err = vault.Set("discogs_api", "supermotdepasse")
	err = vault.Set("root", "root")

	val, err := vault.Get("discogs_api")
	fmt.Println(val)

	err = vault.Set("kea_api", "supercocaine man")

	fmt.Println(err)
}
