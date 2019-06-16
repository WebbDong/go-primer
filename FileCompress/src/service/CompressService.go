package service

import (
	zip2 "archive/zip"
	"fmt"
	"github.com/gpmgo/gopm/modules/cae/zip"
	"io"
	"os"
	"strings"
	"utils"
)

type CompressServicer interface {
	// 压缩
	Compress(decompressPath string, decompressSavePath string) bool
	// 解压缩
	Decompress(compressPath string, compressSavePath string) bool
}

type CompressService struct {
}

func (c *CompressService) Decompress(decompressPath string, decompressSavePath string) bool {
	archive, err := zip.OpenFile(decompressPath, os.O_RDONLY, 4)
	defer archive.Close()
	if err != nil {
		fmt.Println(err)
		return false
	}

	files := archive.File
	for _, f := range files {
		if f.FileInfo().IsDir() {
			continue
		}
		srcFile, err := f.Open()
		defer srcFile.Close()
		if err != nil {
			fmt.Println(err)
			continue
		}

		if strings.Contains(f.Name, "/") {
			path := decompressSavePath + f.Name[0:strings.LastIndex(f.Name, "/")]
			if path != "" {
				// GO中汉字采用的是UTF-8编码，而Windows中采用的是GBK，需要进行转换。
				// 处理中文乱码
				newPath, err := utils.UTF8ToGBK(path)
				if err != nil {
					fmt.Println(err)
					continue
				}
				err = os.MkdirAll(newPath, os.ModePerm)
				if err != nil {
					fmt.Println(err)
					continue
				}
			}
		}

		newFilePath := decompressSavePath + "\\" + f.Name
		newFilePath, err = utils.UTF8ToGBK(newFilePath)
		if err != nil {
			fmt.Println(err)
			continue
		}
		newFile, err := os.Create(newFilePath)
		defer newFile.Close()
		if err != nil {
			fmt.Println(err)
			continue
		}
		/*		buffer := make([]byte, 2 * 1024)
				for {
					_, err := srcFile.Read(buffer)
					if err == io.EOF {
						break
					}
					newFile.Write(buffer)
				}*/

		/*		bytes, err := ioutil.ReadAll(srcFile)
				if err != nil {
					fmt.Println(err)
					continue
				}
				fmt.Println(len(bytes))
				newFile.Write(bytes)
		*/
		_, err2 := io.Copy(newFile, srcFile)
		if err2 != nil {
			fmt.Println(err2)
			continue
		}
	}
	return true
}

func (c *CompressService) Compress(compressPath string, compressSavePath string) bool {
	compressFile, err := os.Create(compressSavePath)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer compressFile.Close()
	file, err := os.Open(compressPath)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer file.Close()
	info, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return false
	}
	fileHeader, err := zip2.FileInfoHeader(info)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
