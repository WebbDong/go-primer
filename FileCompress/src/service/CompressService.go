package service

import (
	"fmt"
	"github.com/gpmgo/gopm/modules/cae/zip"
	"io"
	"os"
	"strings"
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
				err := os.MkdirAll(path, os.ModePerm)
				if err != nil {
					fmt.Println(err)
					continue
				}
			}
		}

		newFilePath := decompressSavePath + "\\" + f.Name
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
	return false
}
