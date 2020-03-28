# dir-cjk-remover
CJK character Remover in directory (with files) name

## How to use
```bash
USAGE:
   dir-cjk-remover [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --chinese, -c              Remove all of 'hanga' characters (default: false)
   --japanese, -j             Remove all 'katakana' and 'hiragana' characters (default: false)
   --korean, -k               Remove all 'hangul' characters (default: false)
   --common, -o               Remove all 'common CJK' characters (default: false)
   --commone, -e              Remove all 'common extension CJK' characters (default: false)
   --silence, -s              Disable detail tasking log (default: false)
   --path value, -p value     Directory for processing
   --replace value, -r value  Set replacing character (default: "_")
   --help, -h                 show help (default: false)
   --version, -v              print the version (default: false)
```
## Example
Below tree is original directory structure.
```
t
├── わら
├── ㅋㅋㅋ
├── 哈哈哈
└── 오잉
    ├── 傷心
    ├── 驚くべき
    ├── 졸리다
    └── 헐
        ├── ㅋ?
        └── 哈!
```
After launch dir-cjk-remover with proper option,
```
./dir-cjk-remover -cj -p t
```
it modify directory like below.
```
t
├── __
├── ___
├── ㅋㅋㅋ
└── 오잉
    ├── __
    ├── ____
    ├── 졸리다
    └── 헐
        ├── _!
        └── ㅋ?
```
