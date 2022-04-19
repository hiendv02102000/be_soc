package main

import (
	"be_soc/pkg/share/utils"
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("models.py", os.O_RDWR, 0644)
	fmt.Println(err)
	// err = utils.UploadFile(file, "download")
	// fmt.Println(err)
	url, err := utils.UploadFile(file, "models.py")

	fmt.Println(url)
	fmt.Println(err)
}
