package main

import (
	"encoding/json"
	"fmt"
)

// 定义一个结构体类型 Person
type Person struct {
	Name string
	Age  int
}

// 为 Person 结构体定义一个方法 Greet
func (p Person) Greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

// 继承
type Student struct {
	Person // 嵌入 Person 结构体，实现继承
	School string
	ID     string
}

func (s Student) PrintInfo() {
	fmt.Printf(`Student Info:
	Name: %s
	Age: %d
	School: %s
	ID: %s`,
		s.Name, s.Age, s.School, s.ID)
	fmt.Println()
}

// 结构体指针 -> 值传递和引用传递
func (s Student) SetName(name string) {
	s.Name = name // 这里修改的是副本，不会影响原结构体
}

func (s *Student) SetNamePtr(name string) {
	s.Name = name // 这里修改的是原结构体
}

// 结构体 tag
// Tag 是一种元数据，可以为结构体字段添加额外的信息，通常用于序列化、数据库映射等场景
// 主要用途是解耦，实现代码与外部系统的解耦，并赋予工具灵活的处理规则。
type Product struct {
	Name  string  `json:"name" db:"product_name"`
	Price float64 `json:"price" db:"product_price"`
}

type User struct {
	UserName string `json:"username" db:"user_name"`
	Email    string `json:"email,omitempty" db:"user_email"` // omitempty: 如果字段值为空，则在 JSON 输出中省略该字段
	Age      int    `json:"age" db:"user_age"`
	password string `json:"-" db:"user_password"` // 私有字段，不会被导出到 JSON
} // Go 语言的访问控制：“小写私有、大写公有”

func main() {
	p1 := Person{Name: "Alice", Age: 30}
	p2 := Person{Name: "Bob", Age: 25}

	p1.Greet()
	p2.Greet()

	p1.Name = "Charlie" // 修改 p1 的 Name 字段
	p1.Greet()

	// 使用指针接收者修改结构体字段
	// 对于Go语言，结构指针调用字段或者方法时，编译器会自动解引用，不需要使用箭头来访问字段或方法
	p2Ptr := &p2
	p2Ptr.Age = 26
	p2Ptr.Greet()

	// 继承
	p := Person{Name: "David", Age: 20}
	s := Student{
		Person: p, // 直接使用 Person 结构体初始化嵌入的 Person 字段
		School: "XYZ University",
		ID:     "S12345",
	}
	s.PrintInfo()

	s1 := Student{
		Person: Person{Name: "Eve", Age: 22}, // 直接在此处初始化嵌入的 Person 结构体
		School: "ABC College",
		ID:     "S67890",
	}
	s1.PrintInfo()
	s1.Greet()                 // 也可以使用父结构体的方法
	fmt.Println(s1.Name)       // 如果标识符不重复，可以直接访问嵌入结构体的字段
	fmt.Println(s1.Person.Age) // 也可以通过嵌入结构体的名字访问字段

	// 结构体指针 -> 值传递和引用传递
	fmt.Println("---------------------")
	student := Student{
		Person: Person{Name: "Frank", Age: 23},
		School: "DEF Institute",
		ID:     "S54321",
	}
	fmt.Println("Before SetName:", student.Name)
	student.SetName("George") // 值传递
	fmt.Println("After SetName:", student.Name)
	student.SetNamePtr("Henry") // 引用传递
	fmt.Println("After SetNamePtr:", student.Name)

	// 结构体 tag
	fmt.Println("---------------------")
	product := Product{
		Name:  "Laptop",
		Price: 999.99,
	}
	fmt.Printf("Product Name: %s, Price: %.2f\n",
		product.Name, product.Price)

	bytedata, _ := json.Marshal(product) // 返回: []byte, error
	fmt.Println("JSON:", string(bytedata))

	user := User{
		UserName: "john_doe",
		Email:    "john@example.com",
		Age:      28,
		password: "secret123",
	}
	userData, _ := json.Marshal(user)
	fmt.Println("User JSON:", string(userData))
}
