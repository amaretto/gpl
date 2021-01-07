package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os/exec"
)

type GoListResult struct {
	Deps []string `json:"Deps"`
}

func main() {
	flag.Parse()
	if len(flag.Args()) != 1 {
	}

	args := []string{"list", "-json"}
	args = append(args, flag.Arg(0))
	out, err := exec.Command("go", args...).Output()
	if err != nil {
		log.Fatal(err)
	}

	var result GoListResult
	err = json.NewDecoder(bytes.NewReader(out)).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	args = []string{"list"}
	args = append(args, result.Deps...)
	out, err = exec.Command("go", args...).Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", string(out))

}
