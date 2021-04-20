# 外部コマンドをGoプログラム内で実行

## 実行したいコマンドを `exec.Command` で準備
定義 ＆ 注意: 再利用は不可能
```
// Cmd represents an external command being prepared or run.
//
// A Cmd cannot be reused after calling its Run, Output or CombinedOutput methods.
```

### `$ ls`　のとき
```
package main

import "os/exec"

func main() {
    cmd := exec.Command("ls")
    ...
}
```


### `$ ls -a -G`　のとき
```
package main

import "os/exec"

var lsOption := []string{"-a", "-G"}

func main() {
    cmd := exec.Command("ls", lsOption...)
    ...
}
```

## 準備されたコマンドを実行
exec.Cmdにはいろいろ用意されている
- Run()
- Start()
- String()
- Wait()
- CombinedOutput()
- Output()
- StderrPipe()
- StdinPipe()
- StdoutPipe()

### String()
人が読める形でコマンドを返す。実行とかは特にしない。
```
string := cmd.String()
```

### Run()
実行したコマンドの終了を待つ。コマンドが出力した結果などは受け取ることはできない。コマンドが正常に動作するかをチェックできる。
```
err := cmd.Run()
```

### Output()
コマンドの標準出力を受け取ることができる。ただし、コマンドの終了を待つ。コマンドの終了を待ち、その結果(標準出力)を使って処理をかける。
```
[]byte, error := cmd.Output()
// コマンドの結果をstringに
result := string([]byte)
```

### CombinedOutput()
コマンドの標準出力＆標準エラー出力を受け取ることができる。ただし、コマンドの終了を待つ。コマンドの終了を待ち、その結果(標準出力&標準エラー出力)を使って処理をかける。
```
[]byte, error := cmd.CombinedOutput()
// コマンドの結果をstringに
result := string([]byte)
```

### Start()
コマンドの終了を待たない。単体で使うよりも `Cmd.Wait()` と組み合わせて使うことが多いと思う。`Cmd.Wait()`でコマンドの終了を待ってくれるのでリアルタイムで標準出力などを検知することができる。（使い勝手がいい子）
```
package main

import (
	"bufio"
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("./a.sh")

	// standard output
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		return
	}
	// standard error output
	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Println(err)
		return
	}

	// start running command
	if err := cmd.Start(); err != nil {
		fmt.Println(err)
		return
	}

	// output
	scan := bufio.NewScanner(stdout)
	scanErr := bufio.NewScanner(stderr)
	go print(scan)
	go print(scanErr)

    // Wait for the command to finish
	if err = cmd.Wait(); err != nil {
		fmt.Println(err)
	}
}

func print(r *bufio.Scanner) {
	for r.Scan() {
		fmt.Println(r.Text())
	}
}
```