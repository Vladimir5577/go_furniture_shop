package utils

import (
	"fmt"
	"os"
)

func DeleteImage(imageName string) error {
	if len(imageName) > 0 {
		err := os.Remove("./uploads/" + imageName)
		if err != nil {
			return err
		}
		fmt.Println("image deleted")
	}
	return nil
}
