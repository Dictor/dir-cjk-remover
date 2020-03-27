package main_test

import (
	d "github.com/dictor/dir-cjk-remover"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemove(t *testing.T) {
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
