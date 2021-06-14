package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

const size = 4 //分割の数

var url = "https://www.robots.ox.ac.uk/~vgg/data/vgg_face/vgg_face_dataset.tar.gz"
var fileph = "./vgg_face_dataset.tar.gz"
var muex sync.Mutex
var wg = sync.WaitGroup{}

func main() {
	fmt.Println("ダウンロードスタート...")
	datetime := time.Now()
	fmt.Println(datetime)
	filelen := Getfilelen()
	fmt.Println("ファイルサイズ：", calculatelen(filelen))
	//sigChan := make(chan os.Signal)
	//signal.Notify(sigChan)
	//ctx, cancel := context.WithCancel(context.Background())
	file, err := os.Create(fileph)
	if err != nil {
		panic(err)
	}
	part := filelen/size + 1 //startとendを指定する
	wg.Add(size)
	go func() {
		i := 1
		Download(url, file, filelen, i, part)
		fmt.Println("channel 1 over")
		wg.Done()
	}()
	go func() {
		i := 2
		Download(url, file, filelen, i, part)
		fmt.Println("channel 2 over")
		wg.Done()
	}()
	go func() {
		i := 3
		Download(url, file, filelen, i, part)
		fmt.Println("channel 3 over")
		wg.Done()
	}()
	go func() {
		i := 4
		Download(url, file, filelen, i, part)
		fmt.Println("channel 4 over")
		wg.Done()
	}()
	/*select {
	case sig := <-sigChan:
		defer file.Close()
		wg.Wait()
		os.Remove(fileph)
		fmt.Println("ダウンロード停止", sig)
	default:*/
	//cancel()
	defer file.Close()
	wg.Wait()
	fmt.Println("ダウンロード終了。")
	datetime2 := time.Now()
	fmt.Println(datetime2)
	//}
}

//ダウンロード機能
func Download(url string, file *os.File, filelen int, i int, part int) {
	var start, end int
	fmt.Println(strconv.Itoa(i) + "番目の部分をダウンロードスタート")
	muex.Lock()
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if i != 0 { //rangeのstartとendを指定
		start = i*part + 1
	} else {
		start = 0
	}
	if i+1 == size {
		end = filelen
	} else {
		end = (i + 1) * part
	}

	str := strconv.FormatInt(int64(start), 10) + "-" + strconv.FormatInt(int64(end), 10)
	fmt.Println("Range"+strconv.Itoa(i), "byte="+str)
	req.Header.Set("Range", "bytes="+str)
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("client.Do err:%v", err)
	}

	fmt.Println(strconv.Itoa(i)+"番目のダウンロードの長さ:", resp.Header.Get("Content-Length"))
	op, err := file.Seek(int64(start), 0)
	fmt.Println("今までダウンロードの長さ:", op)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("read file error: %v", err)
	}
	if len(body) == 0 {
		fmt.Println("no file")
	}
	len, err := file.Write(body)
	fmt.Println(strconv.Itoa(i)+"番目のファイルを書き込み、長さ：", len)
	resp.Body.Close()
	muex.Unlock()
}

//ファイルの長さを取得
func Getfilelen() (filelen int) {
	resp, err := http.Get(string(url))
	if err != nil {
		panic(err)
	}

	filelenStr, ext := resp.Header["Content-Length"]
	if !ext || len(filelenStr) == 0 {
		filelenStr = []string{"0"}
	}

	filelen, fileerr := strconv.Atoi(filelenStr[0])
	fmt.Println("ファイルの長さ：", filelen)
	if fileerr != nil {
		fmt.Println(fileerr)
		filelen = 0
		fmt.Println("ファイルサイズ未知")
	}
	return filelen
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
