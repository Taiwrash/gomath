Solving Data Structure and Algorigthm question with NEETCODE and we need to get a maximum number of an array in golang but we coudn't get a package for it. They I created this.

## Max and Min

### Installation
```bash
go get github.com/Taiwrash/goarrayminmax
```

### Usage


```go
package main

import (
	"fmt"

	"github.com/Taiwrash/gomath"
)

func main() {
	min := gomath.Min([]int{3, 5, 9, 7})
	max := gomath.Max([]int{3, 5, 9, 7})
	fmt.Println("Min: ", min)
	fmt.Println("Max: ", max)
}
```

This will return Min: 3 and Max: 9