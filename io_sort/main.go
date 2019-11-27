package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

/**
class_15
demand: 新增/league端点，返回玩家清单。
1. GET /players/{name} 应该返回一个表示获胜总数的数字
2. POST /players/{name} 应该为玩家赢得游戏记录一次得分，并随每次POST递增
3. GET /league返回玩家清单，按赢的次数排序返回玩家列表
4. 使用标准库io测试接口函数

内容：
. Seeker接口以及它与Reader 和 Writer 的关系
. 处理文件读写
. 为测试创建辅助函数，隐藏文件中杂乱内容
. 使用sort.Slice对切片排序
. 利用编辑器安全对应用程序结构更改
*/
const dbFileName = "game.db.json"

func http_svr() {
	fmt.Println("hi http_svr")
	// server := NewPlayerServer(NewInMemoryPlayerStore())
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("problem opening %s, err: %v", dbFileName, err)
	}
	// store := &FileSystemStore{db}
	store, err := NewFileSystemStore(db)
	if err != nil {
		log.Fatalf("problem creating file system player store, %v", err)
	}

	server := NewPlayerServer(store)

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("listen on 5000 is failed, %v", err)
	}
}

func main() {
	http_svr()
}
