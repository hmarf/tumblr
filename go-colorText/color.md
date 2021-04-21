# Go言語で標準出力にカラーをつけるライブラリーのベンチマーク比較

- `fmt`
- `github.com/fatih/color`
- `github.com/gookit/color`

比較したコード：

```
import (
	"fmt"
	"testing"

	"github.com/fatih/color"
	gcolor "github.com/gookit/color"
)

// https://github.com/fatih/color
func Benchmark_fatih_color(t *testing.B) {
	color.Cyan("Prints text in cyan.")
}

// https://github.com/gookit/color
func Benchmark_gookit_color(t *testing.B) {
	gcolor.Cyan.Println("Prints text in cyan.")
}

// https://groups.google.com/g/golang-nuts/c/nluStAtr8NA?pli=1
func Benchmark_fmt_color(t *testing.B) {
	fmt.Printf("\x1b[36m%s\x1b[0m", "Prints text in cyan.")
}
```

結果：
```
cpu: Intel(R) Core(TM) i7-5775R CPU @ 3.30GHz
github.com/fatih/color:  1000000000           0.0000173 ns/op           0 B/op           0 allocs/op
github.com/gookit/color: 1000000000           0.0000118 ns/op           0 B/op           0 allocs/op
fmt:                     1000000000           0.0000114 ns/op           0 B/op           0 allocs/op
```
*結果の関数名は見やすように変更*

`fmt.Printf()`が最も高速だが、他も早い。  
ただ直感的に描けるし、早いのでカラーつけるときは `"github.com/gookit/color"` 使おうかな。