package main

import (
	"encoding/base64"
	"fmt"
	"net/url"
)

func main() {
	encrypt := "zMq8HzYFNqoGQASttHIRqZqaLol5r8Q1%2FeeWeBvpGhPGVWvbOqykrkct67p2xJXVp7uoXypbkmxw2W6I3Rfy6Iz46pOQiD8rDfqEFBOhTcwcsZSLyV6goryv%2B0aIquZIeiUejQnFo8U86DLG%2FRbAYn4Pl%2BZxRCyGmKoWioKw4rz9zkxwTC9%2B9S3%2FM4lAwO%2BXrOugqOzy%2BV3F8WVlIJ68AEy5otPERvyn9J6TdzmyvhpDLsaEUNwvh0%2BjI3GP8SMFonv6QcvcARY%3D"

	b, err := base64.StdEncoding.DecodeString(encrypt)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))

	fmt.Println(url.QueryUnescape(encrypt))
}