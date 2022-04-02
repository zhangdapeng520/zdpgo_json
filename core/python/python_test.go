package python

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

type account struct {
	Email    string  `json:"email"`
	password string  `json:"password"` // 不会处理私有变量
	Money    float64 `json:"money"`
}

type user struct {
	Name    string
	Age     int
	Roles   []string
	Skill   map[string]float64
	account account
}

// 使用Golang的Marshal方法
func TestGoMarshal(t *testing.T) {
	a := account{
		Email:    "张大鹏",
		password: "123456",
		Money:    100.5,
	}

	rs, err := json.Marshal(a)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(rs)
	fmt.Println(string(rs))
}

// 使用类似Python的dumps方法
func TestPythonDumps(t *testing.T) {
	a := account{
		Email:    "张大鹏",
		password: "123456",
		Money:    100.5,
	}
	result, _ := Dumps(a)
	fmt.Println(result)
}

func TestGoUnmarshal(t *testing.T) {
	str := "{\"access_token\":\"uAUS6o5g-9rFWjYt39LYa7TKqiMVsIfCGPEN4IZzdAk5-T-ryVhL7xb8kYciuU_m\",\"expires_in\":7200}"
	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(str), &dat); err == nil {
		fmt.Println(dat)
		fmt.Println(dat["expires_in"])
	} else {
		fmt.Println(err)
	}
}

func TestPythonLoads(t *testing.T) {
	str := "{\"access_token\":\"uAUS6o5g-9rFWjYt39LYa7TKqiMVsIfCGPEN4IZzdAk5-T-ryVhL7xb8kYciuU_m\",\"expires_in\":7200}"
	var dat map[string]interface{}
	_ = Loads(str, &dat)
	fmt.Println(dat)

	s := `{"email":"张大鹏","money":100.5}`
	var a account
	_ = Loads(s, &a)
	fmt.Println(a, a.Email, a.Money)
}

func TestMap1(t *testing.T) {

	// 处理map
	skill := make(map[string]float64)
	skill["python"] = 99.5
	skill["elixir"] = 90
	skill["ruby"] = 80.0

	user := user{
		Name:  "张大鹏",
		Age:   27,
		Roles: []string{"Owner", "Master"}, // 处理切片
		Skill: skill,
	}

	rs, err := Dumps(user)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(rs)
}

func TestPythonDump(t *testing.T) {
	a := account{
		Email:    "张大鹏",
		password: "123456",
		Money:    100.5,
	}

	// 处理map
	skill := make(map[string]float64)
	skill["python"] = 99.5
	skill["elixir"] = 90
	skill["ruby"] = 80.0

	u := user{
		Name:    "张大鹏",
		Age:     27,
		Roles:   []string{"Owner", "Master"},
		Skill:   skill,
		account: a,
	}

	// 写json文件
	// 创建文件
	filePath := "user.json"
	_ = Dump(filePath, u)
}

func TestPythonLoad(t *testing.T) {
	filePath := "user.json"
	var u user
	_ = Load(filePath, &u)
	fmt.Println(u)
}
