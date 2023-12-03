package test

import (
	"cloud-disk/core/helper"
	"fmt"
	"testing"
)

func TestGenerateUUID(t *testing.T) {
	v4 := helper.GetUUID()
	fmt.Println(v4)

}
