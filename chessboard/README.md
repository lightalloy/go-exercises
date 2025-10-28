# Chessboard

Use app:

```bash
go run main.go
```

```
width: 30
length: 12
Chessboard for 30x12:

# # # # # # # # # # # # # # # 
 # # # # # # # # # # # # # # #
# # # # # # # # # # # # # # # 
 # # # # # # # # # # # # # # #
# # # # # # # # # # # # # # # 
 # # # # # # # # # # # # # # #
# # # # # # # # # # # # # # # 
 # # # # # # # # # # # # # # #
# # # # # # # # # # # # # # # 
 # # # # # # # # # # # # # # #
# # # # # # # # # # # # # # # 
 # # # # # # # # # # # # # # #
```

Run tests:

```bash
cd chessboard
go test
```


Use package in your code:

```go
import (
	"github.com/lightalloy/go-exercises/chessboard/chessboard"
)
func main() {
    // pass any width and length <= 1000
	board, err := chessboard.Get(100, 100)
    // your code
}
```
