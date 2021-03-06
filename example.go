package example

import (
	"fmt"

	example_4 "github.com/elioengcomp/go-module-example-4"
	example_v3 "github.com/elioengcomp/go-module-example/v3"
)

const version = "v1.0.0"

func Exec() string {

	return fmt.Sprintf("This is go-module-example-3 %s consuming go-module-example/v3: %s and go-module-example-4: %s",
		version, example_v3.Exec(), example_4.Exec())
}
