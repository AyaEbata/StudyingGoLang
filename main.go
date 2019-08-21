package main

import (
    "fmt"
    "math"
    "runtime"
    "time"
)

var i, j int = 1, 2  // 初期化子(initializer)

func main() {
    // Go プログラムの基本的なコンポーネントを学びます。
    basic()

    // 条件文とループ、switch、defer を使ってコードの流れをコントロールする方法を学びます。
    grammar()
}

func basic() {
    fmt.Println(math.Pi)  // importしたやつ(math)の中のやつ(Pi)を指定するときは大文字から始める
    fmt.Println(add(42, 13))

    // a, b := swap("hello", "world")  // :=で暗黙的な型宣言ができる
    var a, b string = swap("hello", "world")  // varでも書ける
    fmt.Println(a, b)

    fmt.Println(split(17))

    var c, python, java = true, false, "no!"
    fmt.Println(i, j, c, python, java)

    // 初期値
    var in int
    var f float64
    var bo bool
    var s string
    fmt.Printf("%v %v %v %q\n", in, f, bo, s)

    // 型変換
    var ii int = 42
    var ff float64 = float64(ii)
    var uu uint = uint(ff)
    fmt.Println(ii, ff, uu)

    const Pi = 3.14  // 定数
}

// func add(x int, y int) int {  // 変数名の後ろに型を書く
func add(x, y int) int {  // 型を1つに省略できる
    return x + y
}

func swap(x, y string) (string, string) {  // 一番後ろの括弧が戻り値の型
    return y, x
}

func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return  // 最初に戻り値を指定しているので、戻り値に何も書かないでいい
}

func grammar() {
    defer fmt.Println("end")  // 関数の終わりに呼び出される

    for i := 0; i < 5; i++ {
        defer fmt.Println(i)  // 4から出力される
    }

    // for
    sum := 0
    for i := 0; i < 10; i++ {  // カッコいらない
        sum += i
    }
    fmt.Println(sum)

    // while
    sum2 := 1
    for sum2 < 1000 {  // ifを使う
        sum2 += sum2
    }
    fmt.Println(sum2)

    // if
    x := 6
    if x < 5 {  // カッコいらない&ステートメント書ける（変数書いて条件書ける）
        fmt.Println(x)
    } else {
        fmt.Println("elseだよ")
    }

    // switch
    fmt.Print("Go runs on ")
    switch os := runtime.GOOS; os {  // ここにも変数置ける（書かなくてもいい）
        case "darwin":
            fmt.Println("OS X.")
        case "linux":
            fmt.Println("Linux.")
        default:
            // freebsd, openbsd,
            // plan9, windows...
            fmt.Printf("%s.", os)
    }

    t := time.Now()
    switch {  // switch trueを意味していて、caseで条件を書く
       case t.Hour() < 12:
            fmt.Println("Good morning!")
        case t.Hour() < 17:
            fmt.Println("Good afternoon.")
        default:
            fmt.Println("Good evening.")
    }
}