package main

import (
	"fmt"
	"log"
	"net/http"
)

/**
class_13
TDD: 测试驱动开发，是敏捷开发中的一项核心实践和技术，也是一种设计方法。其原理是在开发代码之前，
先编写单元测试用例代码，由测试代码确定需要编写什么产品代码。
TDD原则：（独立测试、测试列表、测试驱动、先写断言、可测试性、及时重构、小步前进）
demand: 要求创建一个用户服务器，用户可以在其中跟踪玩家赢了多少场游戏。
1. GET /players/{name} 应该返回一个表示获胜总数的数字
2. POST /players/{name} 应该为玩家赢得游戏记录一次得分，并随每次POST递增
*/

func http_svr() {
	fmt.Println("hi http_svr")
	server := &PlayerServer{NewInMemoryPlayerStore()}
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("listen on 5000 is failed, %v", err)
	}
}

func main() {
	http_svr()
}
