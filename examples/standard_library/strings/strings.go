package strings

import (
	"fmt"
	"log"
	"strings"
)

func CloneString() {
	originalString := "hello webster"
	copyString := strings.Clone(originalString)
	fmt.Println(originalString, copyString)
	// 验证 copyString 是否是 originalString 的副本
	fmt.Println(originalString == copyString) // true
	fmt.Println(&originalString, &copyString) // 地址不同
}

func CompareString() {
	// 按照字典顺序进行比较，如果 a < b 返回-1 如果 a == b 返回0 如果 a > b 返回1
	fmt.Println(strings.Compare("abc", "abe"))
	fmt.Println(strings.Compare("abcd", "abe"))
	fmt.Println(strings.Compare("abijk", "abe"))
	fmt.Println(strings.Compare("abe", "abe"))
}

func ContainsString() {
	fmt.Println(strings.Contains("abcdefg", "a"))
	fmt.Println(strings.Contains("abcdefg", "abc"))
	fmt.Println(strings.Contains("abcdefg", "ba"))
}

func ContainsAnyString() {
	fmt.Println(strings.ContainsAny("abcdefg", "ab"))
	fmt.Println(strings.ContainsAny("abcdefg", "ac"))
	fmt.Println(strings.ContainsAny("abcdefg", "z"))
}

func CountSubString() {
	fmt.Println(strings.Count("3.1415926", "1"))       //2
	fmt.Println(strings.Count("there is a girl", "e")) //2
	fmt.Println(strings.Count("there is a girl", ""))  //16
}

func CutString() {
	// 删除在 s 中第一次出现的 prefix 和 suffix
	fmt.Println(strings.Cut("hello webster", "web"))     //hello  ster true
	fmt.Println(strings.Cut("hello webster", "ster"))    //hello web  true
	fmt.Println(strings.Cut("hello webster", "webster")) //hello   true
	fmt.Println(strings.Cut("hello webster", "hello"))   //webster true
	fmt.Println(strings.Cut("hello webster", "x"))       //hello webster  false
}

func EqualString() {
	// 返回字符串 s 和 t 在忽略大小写情况下是否相等
	fmt.Println(strings.EqualFold("abc", "abc")) // true
	fmt.Println(strings.EqualFold("abc", "Abc")) // true
	fmt.Println(strings.EqualFold("abc", "aBc")) // true
	fmt.Println(strings.EqualFold("abc", "AbC")) // true
	fmt.Println(strings.EqualFold("abc", "def")) // false
}

func FieldsString() {
	// 根据空格来分割字符串，将 s 切分成多个子字符串
	fmt.Println(strings.Fields("  hello  webster  ")) // [hello webster]
	// 根据函数的返回值判断是否要进行切割
	fmt.Println(strings.FieldsFunc("  hello  webster  ", func(r rune) bool {
		return r == ' '
	})) // [hello webster]
	fmt.Println(strings.FieldsFunc(" hello webster  ", func(r rune) bool {
		return r == 'e'
	})) // [ h llo w bst r  ]
}

func HasPrefixString() {
	str := "abbc cbba"
	fmt.Println(strings.HasPrefix(str, "abb"))
	fmt.Println(strings.HasPrefix(str, "bba"))
}

func HasSuffixString() {
	str := "abbc cbba"
	fmt.Println(strings.HasSuffix(str, "c"))
	fmt.Println(strings.HasSuffix(str, "bba"))
}

func IndexSubString() {
	str := "abce cbba"
	// 在str中找到bb第一次出现的索引
	fmt.Println(strings.Index(str, "bb"))
	// 在 str 中找到【字符集合】cb 中【任何一个字符】第一次在 str 中出现的字符的索引
	fmt.Println(strings.IndexAny(str, "cb"))
	// 在str中找到单个符文 rune e 第一次出现的索引
	fmt.Println(strings.IndexRune(str, 'e'))

	// 返回最后一次出现的子串的下标
	fmt.Println(strings.LastIndex(str, "b"))
	fmt.Println(strings.LastIndexAny(str, "ea"))
	fmt.Println(strings.LastIndexByte(str, 'e'))
}

func MapString() {
	// Map 返回的都是字符串的副本，不会修改原始字符串
	// 根据映射函数的返回值，来决定是删除（< 0），亦或者保留并转换为其他字符
	fmt.Println(strings.Map(func(r rune) rune {
		if r == 'e' {
			return 'E'
		}
		return r
	}, "abcde")) //abcdE
	fmt.Println(strings.Map(func(r rune) rune {
		if r < 'F' {
			// 如果返回值是 负值（< 0），表示“把这个字符删掉”
			return -1
		} else {
			return r // 保留
		}
	}, "ABCDEFGHIJK")) //FGHIJK
}

func RepeatString() {
	// 重复复制字符串
	fmt.Println(strings.Repeat("hello", 3)) //hellohellohello
	fmt.Println(strings.Repeat("a", 10))
}

func ReplaceString() {
	// 替换字符串
	// 替换字符串中的所有 old 子字符串为 new 子字符串
	// n 指的是替换次数，n 小于 0 时表示不限制替换次数 相当于 ReplaceAll
	fmt.Println(strings.Replace("hello webster", "webster", "world", -1)) //hello world
	fmt.Println(strings.ReplaceAll("Hello this is golang", "o", "c++"))

	fmt.Println(strings.Replace("hello webster", "webster", "world", 1)) //hello world
}

func SplitString() {
	// 分割字符串
	// 根据 sep 来分割字符串 s，返回一个字符串切片，分割次数由n决定

	fmt.Println(strings.Split("hello webster", " ")) // [hello webster]
	fmt.Println(strings.Split("hello webster", ""))  // [h e l l o   w e b s t e r]
	fmt.Println(strings.Split("hello webster", "l")) // [he o webster]

	// 如果 n 为 0，则返回一个空切片
	// 如果 n 为 -1，则表示不限制分割次数
	fmt.Println(strings.SplitN("hello webster", ",", 0)) // []
	fmt.Println(strings.SplitN("hello webster", "", -1)) // [h e l l o   w e b s t e r]
	fmt.Println(strings.SplitN("hello webster", ",", 2)) // [hello webster]
}

func UpperLowerString() {
	// 转换为大写或小写
	fmt.Println(strings.ToUpper("hello webster")) // HELLO WEBSTER
	fmt.Println(strings.ToLower("HELLO WEBSTER")) // hello webster
}

func TrimString() {
	// 修剪字符串两端，把 cutset 当成字符集合（不是子串！），
	// 左右两端只要出现集合里的任意一个字符就砍掉，直到遇到第一个不属于集合的字符为止
	fmt.Println(strings.Trim("aabba", "a")) // bb
	// 修剪字符串左端，把 cutset 当成字符集合（不是子串！），
	// 从左往右只要出现集合里的任意一个字符就砍掉，直到遇到第一个不属于集合的字符为止
	fmt.Println(strings.TrimLeft("aabba", "a")) // bba
	// 修剪字符串左端前缀，把 prefix 当成完整子串，
	// 只有 s 正好以它开头时才一次性削掉，否则原样返回。
	// 查看 s 的前 len(prefix) 的字符是否等于 s[:len(prefix)]
	fmt.Println(strings.TrimPrefix("aabba", "a")) // abba
}

func BuilderString() {
	builder := strings.Builder{}
	// 写入字符串
	builder.WriteString("hello")
	// 写入字节切片
	builder.Write([]byte(" webster"))
	// 写入单个字符
	builder.WriteRune('!')
	// 打印结果
	fmt.Println(builder.String()) // hello webster!
}

func ReplacerString() {
	// 专门用于替换字符串
	replacer := strings.NewReplacer("hello", "hi", "webster", "world")
	fmt.Println(replacer.Replace("hello webster")) // hi world
}

func ReaderString() {
	reader := strings.NewReader("hello world")
	buffer := make([]byte, reader.Len())
	read, err := reader.Read(buffer)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(read)
	fmt.Println(string(buffer[:read]))
}

func StringMain() {
	ReaderString()
}
