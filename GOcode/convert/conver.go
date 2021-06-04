package convert

import (
	"fmt"
	"image"
	"image/gif"
	_ "image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"path/filepath"
)

//再帰的にファイルを読み込む
func GetAllFile(pathname string, ipt []string) ([]string, error) {
	rd, err := ioutil.ReadDir(pathname)
	if err != nil {
		fmt.Println("Failed to read dir:", err)
		return ipt, err
	}
	for _, fi := range rd {
		if fi.IsDir() {
			fullDir := pathname + "/" + fi.Name()
			ipt, err = GetAllFile(fullDir, ipt)
			if err != nil {
				fmt.Println("Failed to read dir:", err)
				return ipt, err
			}
		} else {
			fullName := pathname + "/" + fi.Name()
			ipt = append(ipt, fullName)
		}
	}
	return ipt, nil
}

//画像の形式を変換する
func Conv(ipt, opt string) {
	file, err := os.Open(ipt)
	assert(err, "Invalid image file path ")
	defer file.Close()

	img, _, err := image.Decode(file)
	assert(err, "Failed to convert file to image.")

	out, err := os.Create(opt)
	assert(err, "Failed to create output path.")
	defer out.Close()

	switch filepath.Ext(opt) {
	case ".png":
		png.Encode(out, img)
	case ".gif":
		gif.Encode(out, img, nil)
	}
}

func assert(err error, msg string) {
	if err != nil {
		panic(err.Error() + ":" + msg)
	}
}
