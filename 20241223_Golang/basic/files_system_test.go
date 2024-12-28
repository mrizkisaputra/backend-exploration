package basic

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func TestCreateFile(t *testing.T) {
	_, err := os.Stat("notes.txt")
	if os.IsNotExist(err) {
		file, err := os.Create("notes.txt")
		if err != nil {
			panic(err)
		}
		defer file.Close()
	}
}

func TestWriteToFile(t *testing.T) {
	file, err := os.OpenFile("notes.txt", os.O_RDWR, 0664)
	if err != nil {
		panic(err)
	}

	_, err = file.WriteString("Di hari minggu saya akan:\n1. Berolahraga pagi.\n2. Membersihkan halaman rumah.\n3. Menonton film.\n4. Membaca buku Laskar Pelangi.")
	if err != nil {
		panic(err)
	}

	file.Sync()
}

func TestReadFile(t *testing.T) {
	file, _ := os.OpenFile("notes.txt", os.O_RDONLY, 0664)
	buffer := make([]byte, 10)
	for {
		n, err := file.Read(buffer)
		if err != nil || err == io.EOF {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

}
