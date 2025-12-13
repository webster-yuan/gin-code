package encode

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"google.golang.org/protobuf/proto"
	"gopkg.in/yaml.v3"
)

type Person struct {
	UserId   string `xml:"id"`
	Username string `xml:"name"`
	Age      int    `xml:"age"`
	Address  string `xml:"address"`
}

func sourceDir() string {
	_, file, _, _ := runtime.Caller(0) // 第0层=调用者本身
	return filepath.Dir(file)
}

func XMLMain() {
	person := Person{
		UserId:   "1",
		Username: "webster",
		Age:      23,
		Address:  "usa",
	}
	// 编码为 XML 字符串，以必须带缩进的方式返回
	// GO 中 的 bytes 才相当于 C++ 中的 std::vector<char>
	bytes, err := xml.MarshalIndent(person, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bytes))
	xmlPath := filepath.Join(sourceDir(), "person.xml")
	// 将 bytes 保存进入文件
	if err := os.WriteFile(xmlPath, bytes, 0644); err != nil {
		fmt.Println("写入文件失败")
		return
	}
	// 从文件中读取 XML 字符串
	raw, err := os.ReadFile(xmlPath)
	if err != nil {
		fmt.Println("读取文件失败")
		return
	}
	// 解码成新的 GO 对象
	var loaded Person
	if err := xml.Unmarshal(raw, &loaded); err != nil {
		fmt.Println("解码失败", err)
		return
	}
	// 打印解码后的对象
	fmt.Printf("%+v\n", loaded)
}

func YMLMain() {
	type Config struct {
		Database string `yaml:"database"`
		Url      string `yaml:"url"`
		Port     int    `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	}
	config := Config{
		Database: "oracle",
		Url:      "localhost",
		Port:     3326,
		Username: "root",
		Password: "123456",
	}
	out, err := yaml.Marshal(config)
	if err != nil {
		fmt.Println("编码失败", err)
		return
	}
	fmt.Println(string(out))
	ymlPath := filepath.Join(sourceDir(), "config.yml")
	if err := os.WriteFile(ymlPath, out, 0644); err != nil {
		fmt.Println("写入文件失败")
		return
	}
	raw, err := os.ReadFile(ymlPath)
	if err != nil {
		fmt.Println("读取文件失败")
		return
	}
	var loaded Config
	if err := yaml.Unmarshal(raw, &loaded); err != nil {
		fmt.Println("解码失败", err)
		return
	}
	fmt.Printf("%+v\n", loaded)
}

func JSONMain() {
	type Person struct {
		UserId   string
		Username string
		Age      int
		Address  string
	}
	person := Person{
		UserId:   "1",
		Username: "webster",
		Age:      23,
		Address:  "usa",
	}
	bytes, err := json.Marshal(person)
	if err != nil {
		fmt.Println("编码失败", err)
		return
	}
	// 打印编码后的 JSON 字符串
	fmt.Println(string(bytes))
	jsonPath := filepath.Join(sourceDir(), "person.json")
	if err := os.WriteFile(jsonPath, bytes, 0644); err != nil {
		fmt.Println("写入文件失败", err)
		return
	}
	raw, err := os.ReadFile(jsonPath)
	if err != nil {
		fmt.Println("读取文件失败", err)
		return
	}
	var loaded Person
	if err := json.Unmarshal(raw, &loaded); err != nil {
		fmt.Println("解码失败", err)
		return
	}
	fmt.Printf("%+v\n", loaded)
	// 字段重命名
	// 注意：JSON 标签中使用的字段名必须是导出的（首字母大写）
	type PersonWithJSONTags struct {
		UserId   string `json:"user_id"`
		Username string `json:"username"`
		Age      int    `json:"age"`
		Address  string `json:"address"`
	}
	loadedWithTags := PersonWithJSONTags{
		UserId:   loaded.UserId,
		Username: loaded.Username,
		Age:      loaded.Age,
		Address:  loaded.Address,
	}
	fmt.Printf("%+v\n", loadedWithTags)
	// 写入时带缩进
	// MarshalIndent = Marshal + 自动换行 + 前缀控制(\t)
	bytes, err = json.MarshalIndent(loadedWithTags, "", "\t")
	if err != nil {
		fmt.Println("编码失败", err)
		return
	}
	// 打印编码后的 JSON 字符串
	fmt.Println(string(bytes))
	jsonPath = filepath.Join(sourceDir(), "person_with_tags.json")
	if err := os.WriteFile(jsonPath, bytes, 0644); err != nil {
		fmt.Println("写入文件失败", err)
		return
	}
	raw, err = os.ReadFile(jsonPath)
	if err != nil {
		fmt.Println("读取文件失败", err)
		return
	}
	var loadedWithJSONTags PersonWithJSONTags
	if err := json.Unmarshal(raw, &loadedWithJSONTags); err != nil {
		fmt.Println("解码失败", err)
		return
	}
	fmt.Printf("%+v\n", loadedWithJSONTags)
}

func ProtoMain() {
	// 1. 创建一个符合 proto 的对象
	p := PersonProto{
		Name:   "webster",
		Age:    23,
		Gender: Gender_MALE,
	}
	// 2. 编码为 proto 字节流
	bytes, err := proto.Marshal(&p)
	if err != nil {
		fmt.Println("编码失败", err)
		return
	}
	// 3. 落盘
	protoPath := filepath.Join(sourceDir(), "person.bin")
	if err := os.WriteFile(protoPath, bytes, 0644); err != nil {
		fmt.Println("写入文件失败", err)
		return
	}
	// 4. 从文件中读取
	raw, err := os.ReadFile(protoPath)
	if err != nil {
		fmt.Println("读取文件失败", err)
		return
	}
	// 5. 反序列化
	var loaded PersonProto
	if err := proto.Unmarshal(raw, &loaded); err != nil {
		fmt.Println("反序列化失败", err)
		return
	}
	// 6. 打印反序列化后的对象
	//fmt.Printf("%+v\n", loaded)
	fmt.Printf("name: %s\n", loaded.GetName())
	fmt.Printf("age: %d\n", loaded.GetAge())
	fmt.Printf("gender: %s\n", loaded.GetGender())
}
