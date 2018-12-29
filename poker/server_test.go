package poker

import (
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func Test获取玩家分数(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"zs": 20,
			"ls": 10,
		},
		nil,
		nil,
	}
	server := NewPlayerServer(&store)

	t.Run("返回张三的分数", func(t *testing.T) {
		request := newGetScoreRequest("zs")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("返回李四的分数", func(t *testing.T) {
		request := newGetScoreRequest("ls")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "10")
	})

	t.Run("不存在的人则返回404", func(t *testing.T) {
		request := newGetScoreRequest("ww")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusNotFound)
	})
}

func Test保存胜利者(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
		[]string{},
		nil,
	}
	server := NewPlayerServer(&store)

	t.Run("POST 记录胜利者", func(t *testing.T) {
		player := "ll"

		request := newPostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusAccepted)

		AssertPlayerWin(t, &store, player)
	})
}

func Test联盟成员(t *testing.T) {
	t.Run("获取JSON格式的联盟成员表", func(t *testing.T) {
		wantedLeague := []Player{
			{"zs", 32},
			{"ls", 20},
			{"ww", 14},
		}

		store := StubPlayerStore{nil, nil, wantedLeague}
		server := NewPlayerServer(&store)

		request := newLeagueRequest()
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		got := getLeagueFromResponse(t, response.Body)

		assertStatus(t, response.Code, http.StatusOK)

		assertContentType(t, response, jsonContentType)

		assertLeague(t, got, wantedLeague)
	})
}

func newGetScoreRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, "/players/"+name, nil)
	return request
}

func newPostWinRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodPost, "/players/"+name, nil)
	return request
}

func newLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/League", nil)
	return req
}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("状态码结果 %d 期望 %d", got, want)
	}
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("请求响应错误, 结果 '%s', 期望 '%s'", got, want)
	}
}

func getLeagueFromResponse(t *testing.T, body io.Reader) []Player {
	t.Helper()
	league, err := NewLeague(body)

	if err != nil {
		t.Fatalf("不能解析服务端响应 '%s' 为 Player 切片, '%v'", body, err)
	}

	return league
}

func assertLeague(t *testing.T, got, want []Player) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("比较 JSON 结果 %v 期望 %v", got, want)
	}
}

func assertContentType(t *testing.T, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	if response.Header().Get("content-type") != want {
		t.Errorf("响应头 content-type 期望 %s, 得到的是 %v", want, response.HeaderMap)
	}
}

func assertScoreEquals(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("结果 %d 期望 %d", got, want)
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("发生错误 %v", err)
	}
}
