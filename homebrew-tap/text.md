# Homebrew でコマンドラインツールを配布する手順

## 手順
1. 作ったツールのバイナリファイルを github の `release` に登録
2. Formulaファイル `.rb` をおく、リポジトリを作成。
3. `homebrew-tap` にFormulaファイルを作成。

### 1. 作ったツールのバイナリファイルを github の `release` に登録
`go build` などでバイナリーファイルを作成して、github の `release` タグに登録

### 2. Formulaファイル `.rb` をおく、リポジトリを作成。
Formulaファイル `.rb` を置くための、リポジトリ(`homebrew-tap`と命名する人が多い)を作成。

### 3. `homebrew-tap` にFormulaファイルを作成。
Formulaファイルを作成する。
```
brew create <install URL>
```
`<install URL>`は配布したいツールのバイナリファイルのURL。

↓ は一例([ctest](https://github.com/hmarf/homebrew-tap/blob/master/ctest.rb))だが、こんな感じにする。作成されたファイルのコメント消して、ちょっと変更するだけ。

```
class Ctest < Formula
  desc ""
  homepage ""
  url "https://github.com/hmarf/cTest/releases/download/v1.0.1/ctest"
  sha256 "5478c31029a2b0fa6ed3cbfccbe0b23ae641169f258171436d578aee7fd9b6f8"
  license ""

  def install
    bin.install "<package name>"
  end

  test do
    system "true"
  end
end
```

こいつをgithubにあげたらいける！

### install方法

```
$ brew tap <user name>/tap
$ brew install <package name>
```