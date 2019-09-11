package main

import (
	"fmt"
	"github.com/gookit/color"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		color.Error.Println("Error: a unix timestamp is required.")
		os.Exit(1)
	}
	ts_str := os.Args[1]
	ts, err := strconv.ParseInt(ts_str, 10, 64)
	if err == nil {
		switch len(ts_str) {
		case 16:
			fmt.Println(time.Unix(0, ts*int64(time.Microsecond)))
		case 13:
			fmt.Println(time.Unix(0, ts*int64(time.Millisecond)))
		case 10:
			fmt.Println(time.Unix(0, ts*int64(time.Second)))
		default:
			color.Error.Println("Error: timestamps must be 10, 13, or 16 integers.")
			os.Exit(1)
		}
	} else {
		color.Error.Println("Error: timestamps must be 10, 13, or 16 integers.")
		os.Exit(1)
	}

}
