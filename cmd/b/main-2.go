package main

var Version = "dev"  // 会被编译时的 -ldflags 替换

func main() {
    println("Version:", Version)
    // 你的其他代码
}