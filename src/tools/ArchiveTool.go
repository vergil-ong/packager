package tools

import (
	"os"
	"fmt"
	"io"
	"path"
	"errors"
	"compress/gzip"
	"archive/tar"
	"io/ioutil"
	"path/filepath"
	"log"
	"strings"
)

// 判断档案是否存在
func Exists(name string) bool {
	_, err := os.Stat(name)
	return err == nil || os.IsExist(err)
}

// 判断文件是否存在
func FileExists(filename string) bool {
	fi, err := os.Stat(filename)
	return (err == nil || os.IsExist(err)) && !fi.IsDir()
}

// 判断目录是否存在
func DirExists(dirname string) bool {
	fi, err := os.Stat(dirname)
	return (err == nil || os.IsExist(err)) && fi.IsDir()
}

func CreateDir(filePath string)  error  {
	if !FileExists(filePath) {
		err := os.MkdirAll(filePath,os.ModePerm)
		return err
	}
	return nil
}

func getCurrentDirectory(path string) string {
	dir, err := filepath.Abs(filepath.Dir(path))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func CopyLocalFileToTarget(localPath string, targetPath string, serial string)  {
	localFileInfo, err := os.Stat(localPath)
	if err != nil {
		return
	}
	if !localFileInfo.Mode().IsRegular(){
		return
	}

	localFile, err := os.Open(localPath)
	if err != nil {
		return
	}

	directory := getCurrentDirectory(targetPath)
	CreateDir(directory)
	out, err := os.OpenFile(targetPath,
		os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Printf("cannot open file %s",targetPath)
		return
	}
	io.Copy(out, localFile)
	defer out.Close()
}

func TarGZFiles(src string, dstTar string, failIfExist bool) (err error) {
	src = path.Clean(src)
	if !Exists(src) {
		return errors.New("要打包的文件或目录不存在：" + src)
	}

	if FileExists(dstTar) {
		if failIfExist { // 不覆盖已存在的文件
			return errors.New("目标文件已经存在：" + dstTar)
		} else { // 覆盖已存在的文件
			if er := os.Remove(dstTar); er != nil {
				return er
			}
		}
	}

	fw, er := os.Create(dstTar)
	if er != nil {
		return er
	}
	defer fw.Close()

	gw := gzip.NewWriter(fw)
	defer gw.Close()

	tw := tar.NewWriter(gw)
	defer func() {
		if er := tw.Close(); er != nil {
			err = er
		}
	}()

	fi, er := os.Stat(src)
	if er != nil {
		return er
	}
	srcBase, srcRelative := path.Split(path.Clean(src))

	if fi.IsDir() {
		tarDir(srcBase, srcRelative, tw, fi)
	} else {
		tarFile(srcBase, srcRelative, tw, fi)
	}

	return nil
}

// 因为要执行遍历操作，所以要单独创建一个函数
func tarDir(srcBase, srcRelative string, tw *tar.Writer, fi os.FileInfo) (err error) {
	// 获取完整路径
	srcFull := srcBase + srcRelative

	// 在结尾添加 "/"
	last := len(srcRelative) - 1
	if srcRelative[last] != os.PathSeparator {
		srcRelative += string(os.PathSeparator)
	}

	// 获取 srcFull 下的文件或子目录列表
	fis, er := ioutil.ReadDir(srcFull)
	if er != nil {
		return er
	}

	// 开始遍历
	for _, fi := range fis {
		if fi.IsDir() {
			tarDir(srcBase, srcRelative+fi.Name(), tw, fi)
		} else {
			tarFile(srcBase, srcRelative+fi.Name(), tw, fi)
		}
	}

	// 写入目录信息
	if len(srcRelative) > 0 {
		hdr, er := tar.FileInfoHeader(fi, "")
		if er != nil {
			return er
		}
		hdr.Name = srcRelative

		if er = tw.WriteHeader(hdr); er != nil {
			return er
		}
	}

	return nil
}

func tarFile(srcBase, srcRelative string, tw *tar.Writer, fi os.FileInfo) (err error) {
	// 获取完整路径
	srcFull := srcBase + srcRelative

	// 写入文件信息
	hdr, er := tar.FileInfoHeader(fi, "")
	if er != nil {
		return er
	}
	hdr.Name = srcRelative

	if er = tw.WriteHeader(hdr); er != nil {
		return er
	}

	// 打开要打包的文件，准备读取
	fr, er := os.Open(srcFull)
	if er != nil {
		return er
	}
	defer fr.Close()

	// 将文件数据写入 tw 中
	if _, er = io.Copy(tw, fr); er != nil {
		return er
	}
	return nil
}