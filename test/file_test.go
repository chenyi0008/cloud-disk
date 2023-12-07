package test

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"testing"
)

// 分片的大小
// const chunkSize = 100 * 1024 * 1024
const chunkSize = 100 * 1024

// 文件分片
func TestGenerateChunkFile(t *testing.T) {
	fileInfo, err := os.Stat("img/1.pdf")
	if err != nil {
		t.Fatal(err)
	}
	chunkNum := math.Ceil(float64(fileInfo.Size()) / chunkSize)
	//第一个数字 0 表示八进制数。
	//后面的 666 表示实际的权限。
	myFile, err := os.OpenFile("img/1.pdf", os.O_RDONLY, 0666)
	if err != nil {
		t.Fatal(err)
	}
	b := make([]byte, chunkSize)
	for i := 0; i < int(chunkNum); i++ {
		// 指定读取文件的起始位置
		idx := int64(i * chunkSize)
		myFile.Seek(idx, 0)
		// 120M
		// 100M
		if chunkSize > fileInfo.Size()-idx {
			b = make([]byte, fileInfo.Size()-idx)
		}
		myFile.Read(b)

		f, err := os.OpenFile("./"+strconv.Itoa(i)+".chunk", os.O_CREATE|os.O_WRONLY, os.ModePerm)
		if err != nil {
			t.Fatal(err)
		}
		f.Write(b)
		f.Close()

	}
	myFile.Close()

}

// 分片文件的合并
func TestMergeChunkFile(t *testing.T) {
	myFile, err := os.OpenFile("img/2.pdf", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}

	fileInfo, err := os.Stat("img/1.pdf")
	if err != nil {
		t.Fatal(err)
	}
	chunkNum := math.Ceil(float64(fileInfo.Size()) / chunkSize)
	for i := 0; i < int(chunkNum); i++ {
		f, err := os.OpenFile("./"+strconv.Itoa(i)+".chunk", os.O_RDONLY, os.ModePerm)
		if err != nil {
			t.Fatal(err)
		}
		b, err := ioutil.ReadAll(f)
		if err != nil {
			t.Fatal(err)
		}
		myFile.Write(b)
		f.Close()
	}

}

// 文件一致性的校验
func TestCheck(t *testing.T) {
	// 获取第一个文件的信息
	file1, err := os.OpenFile("img/1.pdf", os.O_RDONLY, 0666)
	if err != nil {
		t.Fatal(err)
	}
	b1, err := ioutil.ReadAll(file1)
	if err != nil {
		t.Fatal(err)
	}

	// 获取第二个文件的信息
	file2, err := os.OpenFile("img/2.pdf", os.O_RDONLY, 0666)
	if err != nil {
		t.Fatal(err)
	}
	b2, err := ioutil.ReadAll(file2)
	if err != nil {
		t.Fatal(err)
	}
	s1 := fmt.Sprintf("%x", md5.Sum(b1))
	s2 := fmt.Sprintf("%x", md5.Sum(b2))
	fmt.Println(s1)
	fmt.Println(s2)

}
