package main

import (
	"log"
	"flag"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
)

// temp1は1つのテンプレート
type templateHandler struct {
	once sync.Once
	filename string
	temp1 *template.Template
}

// ServeHttpはHTTPリクエストを処理します
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func ()  {
		t.temp1 =
			template.Must(template.ParseFiles(filepath.Join("templates",
				t.filename)))
	})
	t.temp1.Execute(w, r)
}

func main() {
		var addr = flag.String("addr", ":8080", "アプリケーションのアドレス")
		flag.Parse() //フラグを解釈する
		r := newRoom()
		http.Handle("/", &templateHandler{filename: "chat.html"})
		http.Handle("/room", r)
		// チャットルームを開始します
		go r.run()
		// Webサーバーを起動します
		log.Println("webサーバーを開始します。ポート：", *addr)
		if err := http.ListenAndServe(*addr, nil); err != nil {
			log.Fatal("ListenAndServe:", err)
		}
		// ルートアクセスハンドラー
		http.Handle("/", &templateHandler{filename: "chat.html"}) 

		// Webサーバを開始
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatal("ListenAndServe:", err)
		}
}