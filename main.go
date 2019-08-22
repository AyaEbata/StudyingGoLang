package main

import (
    "fmt"
    "math"
    "runtime"
    "time"
)

type Vertex struct {
    // X int
    // Y int
    X, Y int  // 書き方省略できる
}

var i, j int = 1, 2  // 初期化子(initializer)

func main() {
    // Go プログラムの基本的なコンポーネントを学びます。
    basic()

    // 条件文とループ、switch、defer を使ってコードの流れをコントロールする方法を学びます。
    grammar()

    // 既存の型に基づいて新しい型を定義する方法を学ぶ: このレッスンでは構造体、配列、slice、map について説明します。
    types()
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

func types() {
    // ポインタ
    i := 42
    p := &i         // point to i
    fmt.Println(*p) // read i through the pointer

    // 構造体
    v := Vertex{1, 2}
    fmt.Println(v)
    fmt.Println(v.X)

    // 片っぽだけも指定できる
    var v2 = Vertex{X: 1}
    fmt.Println(v2)

    // 配列（長さ固定）
    var a [2]string
    a[0] = "Hello"
    a[1] = "World"
    fmt.Println(a[0], a[1])
    fmt.Println(a)

    primes := [6]int{2, 3, 5, 7, 11, 13}
    fmt.Println(primes)

    // slices（配列の長さ変更できるやつ）
    var s []int = []int{2, 3, 5, 7, 11, 13}  // []intはnil
    fmt.Println(s)

    ss := s[1:3]
    ss[1] = 99  // 元のs配列の値も変わる
    fmt.Println(s)
    fmt.Println(ss)

    // slice指定方法
    fmt.Println(s[1:4], len(s), cap(s))  // 長さ(length)と容量(capacity)を持ってる
    fmt.Println(s[:4], len(s), cap(s))  // make([]int, <lenの値>, <capの値>)で割り当てられる
    fmt.Println(s[1:])
    fmt.Println(s[:])

    s = append(s, 100, 111)  // 追加できる
    fmt.Println(s)

    // 配列を回す
    for i, v := range s {  // indexとvalueを持ってる
    // for _, v := range s {  // valueだけを取得
        fmt.Printf("i = %d, v = %d\n", i, v)
    } 

    // map
    map1 := make(map[string]Vertex)
    map1["Bell Labs"] = Vertex{
        40, -74,
    }
    fmt.Println(map1["Bell Labs"])

    // Map literals
    var map2 = map[string]Vertex{
        "Bell Labs": Vertex{
            40, -74,
        },
        "Google": {  // Vertex省略できる
            37, -122,
        },
    }
    fmt.Println(map2)

    // mapの操作
    map3 := map[string]string{"Question":"aaa", "Answer":"bbb"}
    delete(map3, "Answer")
    fmt.Println("The value:", map3["Answer"])

    value, ok := map3["Answer"]
    fmt.Println("The value:", value, "Present?", ok)

    // 関数も変数
    hypot := func(x, y float64) float64 {
        return math.Sqrt(x*x + y*y)
    }
    fmt.Println(hypot(5, 12))

    // Function closures
    add := adder()
    for i := 0; i < 10; i++ {
        fmt.Println(add(i))  // adder()の中のfuncに値を渡せる
    }

    // フィボナッチ
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f(i))
    }
}

func adder() func(int) int {
    sum := 0
    return func(x int) int {
        sum += x
        return sum
    }
}

func fibonacci() func(int) int {
    sum := 0
    x, y := 0, 1
    return func(v int) int {
        if v == 0 {
            return 0
        }
        if v == 1 {
            return 1
        }
        sum = x + y
        x = y
        y = sum
        return sum
    }
}