# cgo-tool-sample
cgoツールによるコード生成のサンプル

## cgoツールの実行コマンド

```cmd
go tool cgo main.go
```

## 生成されるファイル

- `_obj/`
  - `_cgo_export.c`, `_cgo_export.h`:
    - Goで`export`された関数をCから呼ぶための定義など
  - `_cgo_gotypes.go`:
    - Cの型や関数をGoで扱うための型定義や関数呼び出しのスタブなど
  - `_cgo_main.c`:
    - Cリンカが適切にリンクするための一時的なCソース
  - `main.cgo1.go`:
    - `main.go`のC参照部分を`_cgo_gotypes.go`の定義に置換したGoコード
  - `main.cgo2.c`:
    - `main.go`内のCコードと、Goからそれを呼び出すためのスタブコードなど
- `_cgo_2.o`:
  - 生成されたCコードをコンパイルしたオブジェクト
