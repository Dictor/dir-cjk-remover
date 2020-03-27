package main

import (
	"errors"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

var unicodeAddress map[string]([][]int) = map[string]([][]int){
	"hanja": { //han
		{0x3400, 0x4dbf},
		{0x4e00, 0x9fff},
		{0xf900, 0xfaff},
	},
	"japanese": { //hiragana and katakana
		{0x3040, 0x309f},
		{0x3040, 0x30ff},
		{0x31f0, 0x31ff},
		{0x32d0, 0x32ff},
		{0x3300, 0x3370},
		{0xff66, 0xff9f},
		{0x1b000, 0x1b16f},
	},
	"hangul": { //hangul
		{0x1100, 0x11ff},
		{0x3130, 0x318f},
		{0x3200, 0x321f},
		{0x3260, 0x327f},
		{0xa960, 0xa97f},
		{0xac00, 0xd7af},
		{0xd7b0, 0xd7ff},
		{0xffa0, 0xffdf},
	},
	"common": { //common characters (like symbol and additional)
		{0x2e80, 0x2eff},
		{0x2ff0, 0x2fff},
		{0x3000, 0x303f},
		{0x31c0, 0x31ef},
		{0x3220, 0x325f},
		{0x3280, 0x32cf},
		{0x3371, 0x33ff},
		{0xfe30, 0xfe4f},
		{0xfe10, 0xfe1f},
		{0xfe00, 0xfeff},
		{0xfe50, 0xfe6f},
		{0xffe0, 0xffef},
	},
	"common-ext": { //rare common characters
		{0x20000, 0x2fa1f},
		{0x30000, 0x3134f},
	},
}

var (
	printMode   bool
	replaceChar string
)

func main() {
	app := &cli.App{
		Name:    "dir-cjk-remover",
		Usage:   "CJK character Remover in directory (with files) name",
		Version: "1.0",
		Authors: []*cli.Author{{Name: "Dictor", Email: "kimdictor@gmail.com"}},
	}
	app.UseShortOptionHandling = true
	app.Flags = []cli.Flag{
		&cli.BoolFlag{
			Name:    "chinese",
			Aliases: []string{"c"},
			Usage:   "Remove all of 'hanga' characters",
		},
		&cli.BoolFlag{
			Name:    "japanese",
			Aliases: []string{"j"},
			Usage:   "Remove all 'katakana' and 'hiragana' characters",
		},
		&cli.BoolFlag{
			Name:    "korean",
			Aliases: []string{"k"},
			Usage:   "Remove all 'hangul' characters",
		},
		&cli.BoolFlag{
			Name:    "common",
			Aliases: []string{"o"},
			Usage:   "Remove all 'common CJK' characters",
		},
		&cli.BoolFlag{
			Name:    "commone",
			Aliases: []string{"e"},
			Usage:   "Remove all 'common extension CJK' characters",
		},
		&cli.BoolFlag{
			Name:        "silence",
			Aliases:     []string{"s"},
			Usage:       "Disable detail tasking log",
			Destination: &printMode,
			Value:       false,
		},
		&cli.StringFlag{
			Name:     "path",
			Aliases:  []string{"p"},
			Usage:    "Directory for processing",
			Required: true,
		},
		&cli.StringFlag{
			Name:    "replace",
			Aliases: []string{"r"},
			Usage:   "Set replacing character",
			Value:   "_",
		},
	}

	app.Action = func(ctx *cli.Context) error {
		var (
			langFlags map[string]bool = map[string]bool{
				"hanga":      ctx.Bool("chinese"),
				"japanese":   ctx.Bool("japanese"),
				"hangul":     ctx.Bool("korean"),
				"common":     ctx.Bool("common"),
				"common-ext": ctx.Bool("commone"),
			}
			dir = ctx.String("path")
		)

		var noFlag bool = true
		for _, f := range langFlags {
			if f {
				noFlag = false
			}
		}
		if noFlag {
			log.Printf("No character flag passed! Let's see 'dir-cjk-remover --help'")
			return nil
		}

		Process(dir, langFlags)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func Process(dir string, lang map[string]bool) {
	for {
		var complete = true
		err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				printDetail("Error: %s - %s\n", path, err)
				return err
			}
			newpath := getDuplicatePath(path, RemoveCharacter(lang, path))

			if newpath != path {
				if err := os.Rename(path, newpath); err != nil {
					printDetail("Error: %s - %s\n", path, err)
				} else {
					printDetail("Change: %s → %s", path, newpath)
					if info.IsDir() {
						if path == dir {
							dir = newpath
						}
						complete = false
						return errors.New("Retry")
					}
				}
			}
			return nil
		})
		if complete {
			break
		} else if err != nil && err.Error() != "Retry" {
			break
		}
	}
}

func getDuplicatePath(path, newpath string) string {
	var newpathpost = ""
	for {
		if getPostfixPath(newpath, newpathpost) == path {
			break
		}
		if _, err := os.Stat(getPostfixPath(newpath, newpathpost)); err == nil {
			if newpathpost == "" {
				newpathpost = strconv.Itoa(1)
			} else {
				i, _ := strconv.Atoi(newpathpost)
				newpathpost = strconv.Itoa(i + 1)
			}
		} else {
			break
		}
	}
	return getPostfixPath(newpath, newpathpost)
}

func getPostfixPath(path, post string) string {
	if post == "" {
		return path
	} else {
		return path + " (" + post + ")"
	}
}

func printDetail(format string, param ...interface{}) {
	if !printMode {
		log.Printf(format, param...)
	}
}

func RemoveCharacter(flag map[string]bool, s string) string {
	for idx, b := range flag {
		if b {
			var rs string
			for _, r := range s {
				if !hasCharacter(idx, r) {
					rs += string(r)
				} else {
					rs += "_"
				}
			}
			s = rs
		}
	}
	return s
}

func hasCharacter(index string, s rune) bool {
	for i := 0; i < len(unicodeAddress[index]); i++ {
		var ci []int = unicodeAddress[index][i]
		if int32(s) >= int32(ci[0]) && int32(s) <= int32(ci[1]) {
			return true
		}
	}
	return false
}
