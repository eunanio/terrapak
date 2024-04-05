package helpers

import (
	"archive/zip"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func Pack(localPath, modulePath, moduleName string)(string,string,error){
	path := fmt.Sprintf("%s/%s",localPath,modulePath)
	filepath := fmt.Sprintf("%s/%s.zip",localPath,moduleName)

	err := ZipDir(path,filepath); if err != nil {
		fmt.Println(err)
		return "","",err
	}
	hash, err := HashZip(filepath); if err != nil {
        return "","", err
	}

	return filepath,hash,nil
}

func HashZip(filepath string) (string,error){
	file, err := os.Open(filepath); if err != nil {
		return "",err
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "",err
	}

	return hex.EncodeToString(hash.Sum(nil)),nil
}

func HashFiles(dirpath string) (string, error) {
    var fileHashes []string
    err := filepath.Walk(dirpath, func(path string, info os.FileInfo, err error) error {
        if err != nil {
			fmt.Println("walk file error")
            return err
        }

        if info.IsDir() {
            return nil
        }

        data, err := os.ReadFile(path)
        if err != nil {
            return err
        }
        
        hash := sha256.Sum256(data)
        fileHashes = append(fileHashes, hex.EncodeToString(hash[:]))

        return nil
    })

    if err != nil {
        return "", err
    }

    sort.Strings(fileHashes)
    hash := sha256.Sum256([]byte(strings.Join(fileHashes, "")))

    return hex.EncodeToString(hash[:]), nil
}

func ZipDir(source, target string) error {

	f, err := os.Create(target)
	if err != nil {
		return err
	}
	defer f.Close()

	writer := zip.NewWriter(f)
	defer writer.Close()

	return filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		header.Method = zip.Deflate
		header.Name, err = filepath.Rel(source, path)
		if err != nil {
			return err
		}
		if info.IsDir() {
			header.Name += "/"
		}

		headerWriter, err := writer.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = io.Copy(headerWriter, f)
		return err
	})
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}