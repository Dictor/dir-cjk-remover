package main_test

import (
	d "github.com/dictor/dir-cjk-remover"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCharacterRemove(t *testing.T) {
	flags := map[int]map[string]bool{
		1: {"hangul": true, "japanese": false, "hanja": false},
		2: {"hangul": false, "japanese": true, "hanja": true},
		3: {"hangul": false, "japanese": false, "hanja": true},
	}
	set := map[int][][]string{
		1: {
			{"안녕하세요", "_____"},
			{"ㅋㅋㅋzzz", "___zzz"},
			{"hello", "hello"},
		},
		2: {
			{"ハムスター", "_____"},
			{"こころ가 뭔뜻임?", "___가 뭔뜻임?"},
			{"正しい言葉www", "_____www"},
		},
		3: {
			{"金黄地鼠", "____"},
			{"哈哈哈哈ㅋㅋㅋ", "____ㅋㅋㅋ"},
			{"where is 武汉?", "where is __?"},
		},
	}
	for i, nowflag := range flags {
		for _, s := range set[i] {
			res := d.RemoveCharacter(nowflag, s[0])
			assert.Equal(t, s[1], res)
		}
	}
}

func TestFileRemove(t *testing.T) {
	give := []string{
		"/ㅋㅋㅋ",
		"/哈哈哈",
		"/わら",
		"/오잉/",
		"/오잉/헐/",
		"/오잉/졸리다",
		"/오잉/驚くべき",
		"/오잉/傷心",
		"/오잉/헐/ㅋ?",
		"/오잉/헐/哈!",
	}
	want := []string{
		"/___",
		"/___ (1)",
		"/__",
		"/__ (1)/",
		"/__ (1)/_/",
		"/__ (1)/___",
		"/__ (1)/____",
		"/__ (1)/__",
		"/__ (1)/_/_?",
		"/__ (1)/_/_!",
	}

	dir, err := ioutil.TempDir("./", "test")
	assert.NoError(t, err)
	defer os.RemoveAll(dir)

	assert.NoError(t, makeRecursivePaths(dir, give))
	giveres, err := getRecursivePaths(dir)
	assert.NoError(t, err)
	assert.ElementsMatch(t, give, giveres)

	d.Process(dir, map[string]bool{"hangul": true, "japanese": true, "hanja": true})

	wantres, err := getRecursivePaths(dir)
	assert.NoError(t, err)
	assert.ElementsMatch(t, want, wantres)
}

func getRecursivePaths(dir string) ([]string, error) {
	res := []string{}
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			path += "/"
		}
		path = strings.Replace(path, dir, "", -1)
		if path != "/" {
			res = append(res, path)
		}
		return nil
	})
	return res, err
}

func makeRecursivePaths(dir string, paths []string) error {
	for _, p := range paths {
		if p[len(p)-1] != '/' {
			if _, err := os.Create(dir + p); err != nil {
				return err
			}
		} else {
			if err := os.Mkdir(dir+p, 0777); err != nil {
				return err
			}
		}
	}
	return nil
}
