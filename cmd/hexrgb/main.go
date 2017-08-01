package main

import (
	"flag"
	"fmt"
	"log"
)

import "github.com/thales17/hexrgb"

var (
	hexUsage = ""
	hex      = flag.String("-hex", "", hexUsage)
	h        = flag.String("h", "", hexUsage)
	rgbUsage = ""
	rgb      = flag.String("-rgb", "", rgbUsage)
	r        = flag.String("r", "", rgbUsage)
)

func main() {
	flag.Parse()
	usage := "Usage: hexrbg --hex=#AABB11\nhexrgb -h=#AABB11\nhexrgb --rgb=255,255,12\nhexrbg -r=255,255,12"
	if *hex == "" && *h == "" && *rgb == "" && *r == "" {
		fmt.Println(usage)
	}

	if *hex != "" && *h != "" {
		log.Fatal("Please use either --hex or -h not both")
	}

	if *rgb != "" && *r != "" {
		log.Fatal("Please use either --rgb or -r not both")
	}

	if (*hex != "" || *h != "") && (*rgb != "" || *r != "") {
		log.Fatal("Please supply either hex or rgb not both")
	}

	if *hex != "" {
		output, err := hexrgb.ToRGB(*hex)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(output)
	}

	if *h != "" {
		output, err := hexrgb.ToRGB(*h)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(output)
	}

	if *rgb != "" {
		output, err := hexrgb.ToHex(*rgb)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(output)
	}

	if *r != "" {
		output, err := hexrgb.ToHex(*r)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(output)
	}
}
