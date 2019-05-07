package deepPackage

import "fmt"

func init() { 
	fmt.Println("init 1 from deepPackage")
}
func DeepLevel() {
	fmt.Println("Hello from deep level")
}