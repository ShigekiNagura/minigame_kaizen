package game

import (
	"math/rand"
	"time"
)

type Charactor struct {
	Name      string
	Rarity    string
	Image     string
	SoundFile string
}

// func (c *Charactor) RaritytoStar() string {
// 	var star string
// 	for i := 0; i < c.Rarity; i++ {
// 		star += "★"
// 	}
// 	return star
// }

var dog = Charactor{
	Name:   "イヌ",
	Rarity: "⭐︎3",
	Image: `
__    __
o-''))_____\\
"--__/ * * * )
c_c__/-c____/
`,
	SoundFile: "manuke",
}

var cat = Charactor{
	Name:   "ネコ",
	Rarity: "⭐︎4",
	Image: `
/\___/\
( o   o )
(  =^=  )   ニャー
(        )
(         )
___v__^__v___
/             \
|               |
|               |
\             /
""""""""""""""
`,
	SoundFile: "hakusyu",
}

var shiro = Charactor{
	Name:   "しろ",
	Rarity: "⭐︎5",
	Image: `
　　　　　　_ -‐──-ﾆ二ー-､
　 _／ ﾌ‐´　　　　　´｀ヽﾍ｀ー'
　/　,//´｀　　　　　　 ・　 ',
　ゝ' ,' 　・　　　　　　　　　 l
　　　!　　　　　 ━　 ＿, -'
　　　ヽ＿＿,-＝＝i´
　　　　　　 ,'　　　　.',
　　　／,'., '　　　　　 ',
　　　',_.ﾆl 　 .l l　.l l. ﾉ
　　　　　 ｀￣U─U´
`,
	SoundFile: "level_up",
}

var doraemon = Charactor{
	Name:   "ドラえもん",
	Rarity: "⭐︎6",
	Image: `
　　　　　　　　　　　　　　　　　 ______,,,,,,,,,,,,,,,,______
　　　　　　　　　　　　　,,,,:::::::ﾞﾞﾞﾞ､-‐‐-､::::::::-‐‐-､ﾞﾞﾞﾞ::::::,,,,
　　　　　　　　　　　,,::"::::::::::::::/　　　　ヽ/　　　　ヽ:::::::::::::"::,,
　　　　　　　　　 ／::::::::::::::::;;;;l　　　　●|●　　　　l;;;;::::::::::::::::＼
　　　　　　　　／:::::::::::: ''" 　 ヽ.　　　,.-‐-､　　　ノ　　"'' ::::::::::::＼
　　　　　　　/::::::::::／　 ｰ-､,,,_　 ￣´l::::::::::::l｀￣ 　_,,,､-‐　＼:::::::::ヽ
　　　　　　 i':::::,､-‐-､.　　　　 ｀'''‐-　‐-‐'　-‐'''´　　　　,.-‐-､::::::::i,
　　　　 　 i'::::/　　　　 　──-----　　|　　-----── 　　 　ヽ:::::::i,
　　　　　 i':::::{.　　　　　-----‐‐‐‐‐　 │　 ‐‐‐‐‐-----　　　　}::::::::i
　　　　　.|:::::i ヽ.,　　　　　　　　 _____,,,,,,,,|,,,,,,,_____　　　　　　 　 ,ノ i:::::::|
　　　　　.|::::|　　 ｀'t‐----‐''''''´　　　　　　　　　｀''''''‐---‐t''´　　|::::::i
　　　　　 i::::i　　　　i 　　　　　　　　　　　　　　　　　　　　　i　　　 i:::::i'
　　　　　 .'i:::i　　　　i 　　　　　　　　　　　　　　　　　　　　 i　　　 i::::i'
　　, -‐‐- ､::i,　　　 ヽ. 　　　　　 　　　　　　　　　　　　　 /　　　/::i'
　/　　　　　ヽi,　　　　ヽ　　　／ﾞﾞﾞﾞﾞﾞﾞ"'‐--‐'"ﾞﾞﾞﾞﾞ＼　　／ 　　 /:i'
　{　　　　　　} ヽ　　　　 ＼ /　　　　　　　　　　　　 i／ 　　　./'´
　ヽ 　　　　ノ:::::::＼　　　　 '''‐-､,,,,,,,,,_______,,,,,,,､-‐'´　　　 ／
　　｀'''''''''t":::::::::::::::::＼,,,,__　 　　　　　　　　　　　　　__,,,,,／
　　　　 　 ＼::::::::::::::/;,,,,,,,,"""'''''''''''''ゝ‐-､ ''''''''''''''""",,,,,,,},,,,,,,,____, -‐- ､
　　　　　　　 ＼::::::/:::::::::::"""'''''''''''''{=＝=}'''''''''''''""":::::::::::::::::::: / 　　　 ヽ
　　　　　　　　 ､'''ﾞﾞ￣￣ﾞﾞヽ.／　　　'ｰﾞ‐"　　　 ＼:::::::::::|:::::::::::{　　　　　}
　　　　　　　 /　　　　　　　 ヽ　　　　　　　　　　　 ヽ::::::::|‐‐--ヽ､______ノ
　　　　　　　/　　　　　　　　　|─--､､､,,,,,,,______　　 |::::::::|
　　　　　　　| 　　　　　　　　　|　　　　　　　　　|　　 }::::::::l
　　　　　 　 |　　　　　　　　　 |　　　　　　　　 /　　./::::::/
　　　　　　　ヽ. 　　　　　　　 /ヽ､,,,,________,,,／　 ／::::::/
　　　　　　　　＼ 　　　　　／ｰ----------‐‐''´:::::::::/
　　　　　　　　　 ｀'ー--‐''ﾞ　　　　　｀ﾞヽ::::::::::::::::::::::::::/
　　　　　　　　　　　　　　　　　　　　　／ヽ;;＿＿;;;／､
　　　　　　　 　 :::::::::::::::::::::::::::::::::::::::::::{, 　　　　　　　　}
　　　　　　 ::::::::::::::::::::::::::::::::::::::::::::::::::::::'ー-------‐
`,
	SoundFile: "kansei_2",
}

var uron = Charactor{
	Name:   "ウーロン",
	Rarity: "⭐︎7",
	Image: `
　　　　　　　　　　　　　　　　　　　　＿_,.- ､_　　 　 　 ＿＿
　　　　　　　 　 　 　 　 ＿＿＿,._'^l ／⌒ヽ' ｀＞‐''" ＿＿_｀ヽ、
　　　　　　　　　　　 ／＿＿／'　| / 　 　 　 　 　 ／:.:.:.:.:.:.:.:｀ヽ
.　　　　　_r―- ､ ／／　　/ /´｡　｡　⌒)　　U 　 l:.:.:.:.:
　　　　r'　 ⌒ヽ__l　l　 　 /ｒn＾nl'　 ｰ‐' ,.-‐-､_ 　ヽ､_／r:.
　　　 f　_￣ヽ/_　）　　.:/ {_U_Uj 　 ,.-―‐-､　　　 　 /´.:.
　　　 {　　｀Y＾ー''｀ヽ､_{　 ヽ､＿_／ _／￣｀ l　　　　 ヽ_;:-_,.-‐'"
　　　　ヽ'ｰ' _ノ 　 　 ヽ!　　　lｰ一'" 　 　 　 !　　..:.:,r'⌒ヽ、
　　　　　｀弋　　　　　　ヽ、　 ヽ　r‐ｙ'⌒l⌒/|　:.:.:.l^!-‐　'　}￣
　　　　　 　 ヽ　　 　 　 　 ヽ､　ヽ{ _{＿＿ノノ /:.:._L|ー　'
　　 　 　 　 　 ＼　 　 　 　 ll::l'lヽ'ー‐一'"_ノ‐ｧ'　ｲ:(　' ,
　　　　　 　 　 　 ＼　　　ヽ|l::|　ヽ ￣￣￣　／ ／ |:::ト＜＿_
　 　 　 　 　 　 　 　 ヽ､ 　 ||::|　　 'ー―一′_,,..,,_ l:::l
`,
	SoundFile: "donpahu",
}

var charactorLevel3 = []Charactor{dog}
var charactorLevel4 = []Charactor{cat}
var charactorLevel5 = []Charactor{shiro}
var charactorLevel6 = []Charactor{doraemon}
var charactorLevel7 = []Charactor{uron}

func generateCharactor() Charactor {

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(100) + 1

	var charactor Charactor

	switch {
	case 1 <= n && n <= 60:
		charactor = charactorLevel3[len(charactorLevel3)-1]
	case 61 <= n && n <= 85:
		charactor = charactorLevel4[len(charactorLevel4)-1]
	case 86 <= n && n <= 95:
		charactor = charactorLevel5[len(charactorLevel5)-1]
	case 96 <= n && n <= 99:
		charactor = charactorLevel6[len(charactorLevel6)-1]
	case n == 100:
		charactor = charactorLevel7[len(charactorLevel7)-1]
	}

	return charactor
}
