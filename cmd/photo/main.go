package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"time"
)

func main() {
	photos := make(map[string]string)
	filepath.WalkDir("/Users/panda/Documents/photos", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return err
		}
		f, err := d.Info()
		if err != nil {
			fmt.Println(err)
		}
		md5sum, contentType, err := CalcFileMD5(path)
		if err != nil {
			return err
		} else {
			_, ok := photos[md5sum]
			if !ok {
				photos[md5sum] = path
				extType := strings.Split(contentType, "/")[0]
				switch extType {
				case "image":
					rename(md5sum, contentType, path, f)
				case "video":
					rename(md5sum, contentType, path, f)
				default:
					fmt.Println(contentType, f)
				}
			}
			//fmt.Println(path, d, photos)
			//runtime.Breakpoint()
		}
		return err
	})
}

func rename(sum, contentType, path string, f fs.FileInfo) {

	var ext string
	var nf string
	switch contentType {
	case "application/octet-stream":
		ext = strings.ToLower(filepath.Ext(path))
		fmt.Println(contentType, f)
		nf = fmt.Sprintf("/Users/panda/Downloads/photos/%s%s", sum, ext)

	default:
		pre := strings.Split(contentType, ";")[0]
		ext = strings.Split(pre, "/")[1]
		fattr := f.Sys().(*syscall.Stat_t)
		fileDay := time.Unix(fattr.Birthtimespec.Sec, 0).Format("2006-01-02")
		fileDir := fmt.Sprintf("/Users/panda/Downloads/photo/%s", fileDay)
		os.MkdirAll(fileDir, 0755)
		nf = filepath.Join(fileDir, fmt.Sprintf("%s_%s.%s", fileDay, sum, ext))
	}
	fmt.Println(nf)
	err := os.Rename(path, nf)
	if err != nil {
		fmt.Println(err)
	}

}
func CalcFileMD5(filename string) (string, string, error) {
	f, err := os.Open(filename) //打开文件
	if nil != err {
		fmt.Println(err)
		return "", "", err
	}
	defer f.Close()

	md5Handle := md5.New() //创建 md5 句柄

	buffer := make([]byte, 512)
	_, err = f.Read(buffer)
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}
	contentType := http.DetectContentType(buffer)

	md5Handle.Write(buffer)

	_, err = io.Copy(md5Handle, f) //将文件内容拷贝到 md5 句柄中
	if nil != err {
		fmt.Println(err)
		return "", "", err
	}

	md := md5Handle.Sum(nil)        //计算 MD5 值，返回 []byte
	md5str := fmt.Sprintf("%x", md) //将 []byte 转为 string

	return md5str, contentType, nil
}
