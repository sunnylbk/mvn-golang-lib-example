package mvngolib

import "fmt"
import "github.com/google/uuid"

var Buildstamp string
var svnRevision string

func SomeTestMethod() {
	fmt.Printf(uuid.New().String())
}
