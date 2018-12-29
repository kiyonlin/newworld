package poker

import (
	"io/ioutil"
	"os"
	"testing"
)

func Test文件存储(t *testing.T) {
	t.Run("联盟按分数降序排序", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
            {"Name": "zs", "Wins": 10},
            {"Name": "ls", "Wins": 33}]`)
		defer cleanDatabase()

		want := []Player{
			{"ls", 33},
			{"zs", 10},
		}

		store, err := NewFileSystemPlayerStore(database)
		assertNoError(t, err)

		got := store.GetLeague()
		assertLeague(t, got, want)

		// read again
		got = store.GetLeague()
		assertLeague(t, got, want)
	})

	t.Run("获取玩家分数", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
            {"Name": "zs", "Wins": 10},
            {"Name": "ls", "Wins": 33}]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)

		assertNoError(t, err)

		assertScoreEquals(t, store.GetPlayerScore("ls"), 33)
	})

	t.Run("为存在的玩家加分", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
            {"Name": "zs", "Wins": 10},
            {"Name": "ls", "Wins": 33}]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)

		assertNoError(t, err)

		store.RecordWin("ls")

		assertScoreEquals(t, store.GetPlayerScore("ls"), 34)
	})

	t.Run("为不存在的玩家加分", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
            {"Name": "zs", "Wins": 10},
            {"Name": "ls", "Wins": 33}]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)

		assertNoError(t, err)

		store.RecordWin("ww")

		assertScoreEquals(t, store.GetPlayerScore("ww"), 1)
	})

	t.Run("处理空文件", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, "")
		defer cleanDatabase()

		_, err := NewFileSystemPlayerStore(database)

		assertNoError(t, err)
	})
}

func createTempFile(t *testing.T, initialData string) (*os.File, func()) {
	t.Helper()
	tmpfile, err := ioutil.TempFile("", "db")
	if err != nil {
		t.Fatalf("无法创建临时文件: %v", err)
	}
	tmpfile.Write([]byte(initialData))
	removeFile := func() {
		os.Remove(tmpfile.Name())
	}
	return tmpfile, removeFile
}
