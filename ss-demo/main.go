package main

import (
	"fmt"
)
const secretKey = "AC123456789012345678901234567890ab" 

func main() {
	fmt.Println("Twilio API Token:", "AC123456789012345678901234567890ab")
	fmt.Println("AWS Session Token:", "FwoGZXIvYXdzEMX//////////EXAMPLETOKEN1234567890ABCDEFGHIJKLMNO")
}


// Go Example
// secretKey = "AC123456789012345678901234567890ab" 

package main

import (
    "fmt"
    "os"
)

func main() {
    secretKey := os.Getenv("SECRET_KEY")
    fmt.Println(secretKey)
}
