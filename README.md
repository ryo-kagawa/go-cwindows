# go-cwindows

GoでWindows APIをC言語ライクな型定義と関数呼び出しを行う

## 機能

- Windows APIのGoラッパー

## ディレクトリ構成

```
.
├── cwindows/                    # Windows SDKのGoラッパーに独自の定数・関数・型定義を追加
│   ├── custom_const_*.go        # Windowsに存在しない定数定義
│   ├── custom_func_*.go         # Windowsに存在しない関数
│   ├── custom_type_*.go         # Windowsに存在しない型定義
│   ├── windows_const_*.go       # Windows SDK定数定義
│   ├── windows_func_*.go        # Windows SDK関数
│   └── windows_type_*.go        # Windows SDK型定義
├── cwindowsnative/              # Windows SDKの純粋なGoラッパー
│   ├── windows_const_*.go       # Windows SDK定数定義
│   ├── windows_func_*.go        # Windows SDK関数
│   └── windows_type_*.go        # Windows SDK型定義
├── dll/                         # Windows SDKの純粋なGoラッパー
│   └── DLL名/                   # Windows SDKで使用するDLL
│       ├── Windows API名/       # Windows SDKで使用するDLLに含まれる関数
│       │   └── Windows API名.go # Windows SDKで使用するDLLに含まれる関数
│       └── DLL名.go             # DLLのLazyLoadを行う
├── go.mod                       # Goモジュール定義
└── go.sum                       # Goモジュールチェックサム
```

## cwindows と cwindowsnative の違い

- `cwindows`: Windows SDKの定義に独自定義
- `cwindowsnative`: Windows SDKの定義に忠実にC言語風に扱えるラッパー

## 使い方

1. init関数内で各関数の初期化処理を呼び出す（`Init${DLL名}${関数名}()`）
```
package main

import "github.com/ryo-kagawa/go-cwindows/cwindows"

func init() {
	cwindows.InitKernel32CloseHandle()
}
```
2. 対象の関数を使用する（`${関数名}()`）
```
cwindows.CloseHandle(handle)
```

## 例外

Windows APIのコールバック関数をGoで構造体に渡すuintptrとする必要があるため各種構造体の型定義はuintptrに変更されているので各`Set${フィールド名(${コールバック関数})}`を定義してください
```
options := cwindows.ETW_OPEN_TRACE_OPTIONS{}
options.SetEventCallback(
  func(EventRecord cwindows.PEVENT_RECORD) {
  },
)
```

`windows_const_*.go`に実際には定数として扱いたいが、Goでは設定できないためvarで定義しているものが存在する

## インストール

```bash
go get github.com/ryo-kagawa/go-cwindows@latest
```
