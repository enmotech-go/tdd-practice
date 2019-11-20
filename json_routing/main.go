package main

import (
	"fmt"
	"log"
	"net/http"
)

/**
class_14
demand: 新增/league端点，返回玩家清单。
1. GET /players/{name} 应该返回一个表示获胜总数的数字
2. POST /players/{name} 应该为玩家赢得游戏记录一次得分，并随每次POST递增
3. GET /league返回玩家清单
*/

func http_svr() {
	fmt.Println("hi http_svr")
	server := NewPlayerServer(NewInMemoryPlayerStore())
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("listen on 5000 is failed, %v", err)
	}
}

func main() {
	http_svr()
}
