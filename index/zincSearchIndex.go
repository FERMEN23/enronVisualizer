package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"runtime"
	"runtime/pprof"

	"log"
	"net/http"
	"os"
)

func index_data() {
	user := "admin"
	password := "Complexpass#123"
	auth := user + ":" + password
	bas64encoded_creds := base64.StdEncoding.EncodeToString([]byte(auth))

	zinc_host := "http://localhost:4080"

	zinc_url := zinc_host + "/api/_bulkv2"

	file, err := os.Open("jSonFinal.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	reader := io.Reader(file)

	req, err := http.NewRequest("POST", zinc_url, reader)

	if err != nil {
		log.Fatal("Error reading request. ", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic "+bas64encoded_creds)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("resp", resp)
	defer resp.Body.Close()
}

func main() {

	cpu, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(cpu)
	defer pprof.StopCPUProfile()

	fmt.Println("Indexing...")
	index_data()
	fmt.Println("Indexing finished!!!!")

	runtime.GC()
	mem, err := os.Create("memory.prof")
	if err != nil {
		log.Fatal(err)
	}
	defer mem.Close()
	if err := pprof.WriteHeapProfile(mem); err != nil {
		log.Fatal(err)
	}

}

//go tool pprof -http=:8020 cpu.prof
