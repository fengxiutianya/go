package listing08

import (
	"net/http"
	"testing"
)

const (
	checkMark = "\u2713"
	ballotx   = "\u2717"
)

//表组测试 产生一组不同的输入并产生不同的输出
//TestDownload 确认http包的get函数可以下载内容
//并正确处理不同的状态
func TestDownload(t *testing.T) {
	var urls = []struct {
		url        string
		statusCode int
	}{
		{
			"http://www.goinggo.net/feeds/posts/default?alt=rss",
			http.StatusOK,
		},
		{
			"http://rss.cnn.com/rss/cnn_topstbadurl.rss",
			http.StatusNotFound,
		},
	}
	t.Log("Gvien the need to test downloding different content.")
	{
		for _, u := range urls {
			t.Logf("\t When checking \"%s\" for status code \"%d\"",
				u.url, u.statusCode)
			{
				resp, err := http.Get(u.url)
				if err != nil {
					t.Fatal("\t\t Should be able to Get the url", ballotx, err)
				}
				t.Log("\t\tShould be able to Get the url", checkMark)
				defer resp.Body.Close()
				if resp.StatusCode == u.statusCode {
					t.Logf("\t\tShould have a \"%s\" status.%v", u.statusCode, checkMark)
				} else {
					t.Errorf("\t\tShould have a \"%d\" status %v %v", u.statusCode, ballotx, resp.StatusCode)
				}
			}
		}
	}

}
