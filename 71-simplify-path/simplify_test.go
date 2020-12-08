package _1_simplify_path

import (
	"fmt"
	"path/filepath"
	"testing"
)

func Test(t *testing.T) {
	fmt.Println(simplifyPath("//a//b////c/d//././/.."))
	fmt.Println(filepath.Clean("/a//b////c/d//././/.."))
}


