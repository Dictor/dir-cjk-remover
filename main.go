package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
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
	"common-extension": { //rare common characters
		{0x20000, 0x2fa1f},
		{0x30000, 0x3134f},
	},
}

func main() {
	app := &cli.App{Name: "CJK character Remover in directory (with files) name"}
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
			Name:    "print",
			Aliases: []string{"v", "p"},
			Usage:   "Print every tasking directory",
		},
	}

	app.Action = func(ctx *cli.Context) error {
		var (
			c = ctx.Bool("chinese")
			j = ctx.Bool("japanese")
			k = ctx.Bool("korean")
			v = ctx.Bool("print")
		)
		log.Printf("removing option: c=%t j=%t k=%t", c, j, k)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func remove(s string) {
	for _, r := range s {

	}
}
