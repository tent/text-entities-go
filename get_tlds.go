// +build ignore

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, err := http.Get("https://data.iana.org/TLD/tlds-alpha-by-domain.txt")
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		panic(err)
	}

	lines := bytes.Split(bytes.ToLower(data), []byte("\n"))
	var idntld, cctld, gtld [][]byte
	for _, line := range lines[1:] {
		if len(line) == 2 {
			cctld = append(cctld, line)
			continue
		}
		if bytes.HasPrefix(line, []byte("xn--")) {
			idntld = append(idntld, line[4:])
			continue
		}
		if len(line) > 0 {
			gtld = append(gtld, line)
		}
	}

	fmt.Printf("regexen[\"GTLD\"] = \"(?:%s)\"\n", bytes.Join(gtld, []byte("|")))
	fmt.Printf("regexen[\"CCTLD\"] = \"(?:%s)\"\n", bytes.Join(cctld, []byte("|")))
}
