package listing12

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	checkMark = "\u2713"
	ballotx   = "\u2717"
)

//feed 模仿了我们期望接收的XML文档
var feed = `<?xml version="1.0" encoding="UTF-8"?>
<rss>
	<channel>
		<title>Going GO Programming</title>
		<description>Golang : https://github.com/goinggo</description>
		<link>http://www.goinggo.net/</link>
		<item>
			<pubDate>Sun, 15 Mar 2015 15:04:00 +0000</pubDate>
			<title>Object Oriented Programming Mechanics</title>
			<description>Go is an object oriented language.</description>
			<link>http://www.goinggo.net/2015/03/object-oriented</link>
		</item>
	</channel>
</rss>`

//mockserver返回用来处理请求的服务器的指针
func mockServer() *httptest.Server {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/xml")

		fmt.Println(w, feed)
	}
	return httptest.NewServer(http.HandlerFunc(f))
}

// TestDownload 确认http 包的 Get 函数可以下载内容
// 并且内容可以被正确地反序列化并关闭
func TestDownload(t *testing.T) {
	statusCode := http.StatusOK

	server := mockServer()
	defer server.Close() //此时没有关闭server，但是在return之前一定会关闭

	t.Log("Given the need to test downloading content.")
	{
		t.Logf("\tWhen checking \"%s\" for status code \"%d\"",
			server.URL, statusCode)
		{
			resp, err := http.Get(server.URL)
			if err != nil {
				t.Fatal("\t\tShould be able to make the Get call.",
					ballotx, err)
			}
			t.Log("\t\tShould be able to make the Get call.",
				checkMark)

			defer resp.Body.Close() //此时不一定关闭了Body，但是在return之前一定会关闭

			if resp.StatusCode != statusCode {
				t.Fatalf("\t\tShould receive a \"%d\" status. %v %v",
					statusCode, ballotx, resp.StatusCode)
			}
			t.Logf("\t\tShould receive a \"%d\" status. %v",
				statusCode, checkMark)
		}
	}
}
