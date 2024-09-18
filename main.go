package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    // Ginのデフォルトのルーターを作成
    router := gin.Default()

    // ルートハンドラーの設定
    router.GET("/", func(c *gin.Context) {
        c.HTML(200, "index.html", gin.H{
            "title": "Hello, GO!",
        })
    })

    // テンプレートのディレクトリを設定
    router.LoadHTMLGlob("templates/*")

    // サーバーのポートを指定して起動
    router.Run(":8080")
}
