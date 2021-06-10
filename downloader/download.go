package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func main() {
	//args := os.Args
	//url, fileph := args[1], args[2]
	url, fileph := "https://github.com/yuan-h-1994/Gocode.git", "./GOcode.zip"
	if err := Download(url, fileph); err != nil {
		panic(err)
	}

}

func Download(url string, fileph string) error {
	resp, err := http.Get(string(url))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	filelenStr, ext := resp.Header["Content-Length"]
	if !ext || len(filelenStr) == 0 {
		filelenStr = []string{"0"}
	}

	filelen, fileerr := strconv.Atoi(filelenStr[0])
	if fileerr != nil {
		fmt.Println(fileerr)
		filelen = 0
		fmt.Println("ファイルサイズ未知")
	}
	fmt.Println("ファイルサイズ：", calculatelen(filelen))

	//resp.Header.Set("Range",fmt.Sprintf("bytes=%d-%d", 0, t.file_size))

	file, err := os.Create(fileph)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	return err
}

//ファイルサイズを計算する
func calculatelen(len int) string {
	if len < 1024 {
		return fmt.Sprintf("%d Btye", len)
	}
	kb := float32(len) / 1024
	if kb < 1024 {
		return fmt.Sprintf("%f Kb", kb)
	}
	mb := kb / 1024
	if mb < 1024 {
		return fmt.Sprintf("%f Mb", mb)
	}
	gb := mb / 1024
	if mb < 1024 {
		return fmt.Sprintf("%f GB", gb)
	}
	return fmt.Sprintf("%f PB", gb/1024)
}
