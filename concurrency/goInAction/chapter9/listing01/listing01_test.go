package listing01

import "testing"
import "net/http"

const (
	checkMark = "\u2713"
	ballotx   = "\u2717"
)

//testDownload 确认htpp包的get函数可以下载内容
func TestDownload(t *testing.T) {
	url := "http://wwww.goinggo.net/feeds/posts/default?alt=rss"
	statusCode := 200

	t.Log("Given the need to test downloding content.")
	{
		t.Logf("\tWhen checking \"%s\" for status code \"%d\"", url, statusCode)
		{
			resp, err := http.Get(url)
			if err != nil {
				//如果是执行到此，就会停止执行
				t.Fatal("\t\t Should be able to make the Get Call.", ballotx, err)
			}
			t.Log("\t\tShould be able to make the Get call.", checkMark)
			defer resp.Body.Close()

			if resp.StatusCode == statusCode {
				t.Logf("\t\tShould receive a \"%s\" status. %v", statusCode, checkMark)
			} else {
				//如果是遇见这个，会显示错误，但还会接着运行
				t.Errorf("\t\tShould receive a \"%d\" status. %v %v", statusCode, ballotx, resp.StatusCode)
			}

		}
	}
}
