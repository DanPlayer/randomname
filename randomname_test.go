package randomname

import (
	"fmt"
	"testing"
)

func TestGenerateName(t *testing.T) {
	name := GenerateName()
	fmt.Println(name)

	name = GenerateName()
	fmt.Println(name)

	name = GenerateName()
	fmt.Println(name)
}
