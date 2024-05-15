# Go Packages

---

# Strings
### 把字符拼接成字符串 strings.Builder
```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	// 定义一个strings.Builder对象为builder
	var builder strings.Builder
	// 遍历3到1，用fmt.Fprintf()方法将数字加到builder中，也可以用builder.WriteByte()的方法写入
	for i := 3; i >= 1; i-- {
		fmt.Fprintf(&builder, "%d...", i)
	}
	// 将字符串"ignition"加到builder中
	builder.WriteString("ignition")
	// 将builder中的内容由字符转换为字符串，并打印输出
	fmt.Println(builder.String())

}
```
```terminal
Output:
3...2...1...ignition
```

### 字符串分割 strings.Split
有几种方法：
1. strings.Split(s, sep)：按照 sep 分割字符串 s，返回一个字符串切片。
2. strings.Fields(s)：按照空格分割字符串 s，返回一个字符串切片。
3. strings.FieldsFunc(s, f)：按照自定义函数 f 分割字符串 s，返回一个字符串切片。
4. strings.SplitAfter(s, sep)：按照 sep 分割字符串 s，返回一个字符串切片，每个元素包含 sep。

> - 一个或多个空格分割字符串 ==> strings.Fields(s)
> - 单个空格或其他字符分割字符串 ==> strings.Split(s, sep)
> - 自定义函数分割字符串 ==> strings.FieldsFunc(s, f)
> - 每个元素包含分割的字符 ==> strings.SplitAfter(s, sep)

- FieldsFunc 函数接受一个自定义的函数 f，该函数接受一个 rune 类型的参数，并返回一个 bool 类型的值。FieldsFunc 会根据 f 的返回值来决定是否在该位置分割字符串。如果 f 返回 true，则在该位置分割字符串；如果 f 返回 false，则不分割。
- 如果提供的分割函数 f 返回 true 时连续出现几次（即字符串中有连续的分割符），FieldsFunc 会将它们视为单一的分隔点，并且在这些位置分割字符串。所有连续返回 true 的分隔符都会被忽略，这意味着不会有空字符串被包括在返回的切片中。

example:
```go
func main() {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	fmt.Printf("Fields are: %q", strings.FieldsFunc("  foo1;bar2,baz3...", f))
}
```
```terminal
Output:
Fields are: ["foo1" "bar2" "baz3"]
```

### 字符串连接 strings.Join
```go
func main() {
    // 定义一个字符串切片为slice
    slice := []string{"foo", "bar", "baz"}
    // 将slice中的元素用"|"连接起来，赋值给str
    str := strings.Join(slice, "|")
    // 打印输出str
    fmt.Println(str)
}
```
```terminal
Output:
foo|bar|baz
```

### 字符串替换 strings.Replace
```go
func main() {
    // 定义一个字符串s
    s := "foo foo foo"
    // 将s中的"foo"替换为"bar"，替换次数为2，赋值给str
    str := strings.Replace(s, "foo", "bar", 2)
    // 打印输出str
    fmt.Println(str)
}
```
```terminal
Output:
bar bar foo
```

### 字符串重复 strings.Repeat
```go
func main() {
    // 定义一个字符串s
    s := "foo"
    // 将s重复3次，赋值给str
    str := strings.Repeat(s, 3)
    // 打印输出str
    fmt.Println(str)
}
```
```terminal
Output:
foofoofoo
```

### 字符串修剪 strings.Trim
```go
func main() {
    // 定义一个字符串s
    s := "  foo bar baz  "
    // 将s两侧的空格去掉，赋值给str
    str := strings.Trim(s, " ")
    // 打印输出str
    fmt.Println(str)
}
```
```terminal

Output:
foo bar baz
```

### 字符串搜索 strings.Contains 和 strings.Index
```go
func main() {
    // 定义一个字符串s
    s := "foo bar baz"
    // 判断s中是否包含"bar"，赋值给b
    b := strings.Contains(s, "bar")
    // 打印输出b
    fmt.Println(b)

    // 获取s中"bar"的索引，赋值给i
    i := strings.Index(s, "bar")
    // 打印输出i
    fmt.Println(i)
}
```
```terminal
Output:
true
4
```

### 字符串比较 strings.Compare
用于比较两个字符串，如果字符串相等，返回0；如果字符串s1小于s2，返回-1；如果字符串s1大于s2，返回1。
字符串的比较是按照字典顺序进行的，即按照字符的 Unicode 编码值进行比较，从字符串的第一个字符开始逐个比较，直到找到不同的字符为止，如果这个不同的字符在 s1 中的 Unicode 编码值小于 s2 中的 Unicode 编码值，则 s1 小于 s2，返回 -1；如果大于，则 s1 大于 s2，返回 1；如果两个字符串完全相同，则返回 0。
```go
func main() {
    // 定义两个字符串s1和s2
    s1 := "foo"
    s2 := "bar"
    // 比较s1和s2，赋值给i
    i := strings.Compare(s1, s2)
    // 打印输出i
    fmt.Println(i)
}
```
```terminal
Output:
1
```






# os
<https://pkg.go.dev/os>

### 创建并打开文件: **os.Create**
```go

file, err := os.Create("file.txt") // 创建文件
if err != nil {
    log.Fatal(err)
}
defer file.Close()
```

### 打开已有文件：**os.Open** (只读，不能写入) 和 **os.OpenFile** 可以自定义权限
如果简单情况，只需要读，用os.Open即可
```go
file, err := os.Open("File.txt")
if err != nil {
    log.Fatal(err)
}
defer file.Close()

// 读取文件内容
reader := bufio.NewReader(file)
line, err := reader.ReadString('\n')
if err != nil {
    log.Fatal(err)
}
fmt.Println(line)
```

如果要自定义控制的场景，用os.OpenFile

```go
func OpenFile(name string, flag int, perm os.FileMode) (*os.File, error)
// - name：文件的路径和名称。
// - flag：指定打开文件的模式，如 os.O_RDONLY（只读）、os.O_WRONLY（只写）、os.O_RDWR（读写）、os.O_CREATE（不存在时创建）、os.O_APPEND（写操作时追加数据）等。
// - perm：设置文件权限，在创建文件时使用。
// 返回值：
// - *os.File：返回一个文件对象指针，可以用于后续的文件 I/O 操作。
// - error：如果操作出错，返回一个错误值。
```

```go
// 打开一个文件，如果不存在则创建，只写模式，权限为644
file, err := os.OpenFile("filename.txt", os.O_CREATE|os.O_WRONLY, 0644)
if err != nil {
    log.Fatal(err)
}
defer file.Close()
// 写入文件内容...
```

### 读取整个文件内容：**os.ReadFile** 和 **ioutil.ReadAll**

用os.ReadFile方法（推荐）
```go
// 指定文件路径
filename := "example.txt"

// 使用 os.ReadFile 读取文件内容
content, err := os.OpenFile("filename.txt", os.O_CREATE|os.O_WRONLY, 0644)
if err != nil {
    log.Fatalf("Failed to read file: %v", err)
}

// 将内容（字节切片）转换为字符串并打印
fmt.Println(string(content))
```
或用ioutil
```go
file, err := os.Open("File.txt")
if err != nil {
    log.Fatal(err)
}
defer file.Close()

content, err := ioutil.ReadAll(file)
if err != nil {
    log.Fatal(err)
}

fmt.Println(string(content))
```


### 读写文件：file.WriteString
```go
file, err := os.Open("File.txt")
if err != nil {
    log.Fatal(err)
}
defer file.Close()

_, err = file.WriteString("Hello, world!") // 写入字符串到文件
if err != nil {
    log.Fatal(err)
}
```

### 删除文件：os.Remove
```go
err = os.Remove("file.txt") // 删除文件
if err != nil {
    log.Fatal(err)
}
```

### 重命名或移动文件：os.Rename
```go
err = os.Rename("old_filename.txt", "new_filename.txt") // 重命名文件
if err != nil {
    log.Fatal(err)
}
```

### 获取和修改文件属性：os.Stat
```go
info, err := os.Stat("file.txt") // 获取文件状态
if err != nil {
    log.Fatal(err)
}
fmt.Println("File size:", info.Size()) // 打印文件大小
```
---

# bufio
<https://pkg.go.dev/bufio>

- 打开已有文件，并读取文件内容：
```go
file, err := os.Open("File.txt")
if err != nil {
    log.Fatal(err)
}
defer file.Close()

// 读取文件内容
reader := bufio.NewReader(file) // 创建了一个bufio.Reader对象

// 读取一行
line, err := reader.ReadString('\n')  //读取一行的数据，读到'\n'换行符时，就停止
if err != nil {
    log.Fatal(err)
}
fmt.Println(line)

// 读取整个文件
for reader{
    line, err := reader.ReadString('\n')  //读取一行的数据，读到'\n'换行符时，就停止
    // 如果是其他错误，打印即可
    if err != nil {
        log.Fatal(err)
    }
    // 如果是到达文件末尾，也就是遇到io.EOF
    if err == io.EOF {
        log.Fatal(err) // 不一定要打印，可选
        break
    }
    fmt.Println(line)
}
```

一般特别简单的输入，可以用fmt.Scan(&a)，如下：
```go
var input string
_, err := fmt.Scanln(&input)

if err != nil {
    fmt.Println("err:", err)
}

fmt.Println(input)
``` 


**bufio捕获单次输入**
```go
// 创建一个输入检测对象scanner
scanner := bufio.NewScanner(os.Stdin)

// 单次输入====>

// 扫描你输入的值，直到按enter为止
scanner.Scan()
// 获取得到的值
str := scanner.Text()
fmt.Println(str)
```

**bufio捕获持续输入**
```go
// 创建一个输入检测对象scanner
scanner := bufio.NewScanner(os.Stdin)

// 每一个输入，都输出对应的值
for scanner.Scan() {
    // 存储到的值为input
    input := scanner.Text()
    / /如果input是"exit"，退出循环
    if input == "exit" {
        break
    }
    // 如果没有错误，打印出存储的值
    fmt.Print("your input is: ")
    fmt.Println(input)
}

// 如果有错误，打印错误
if err := scanner.Err(); err != nil {
    fmt.Fprintf(os.Stderr, "error reading input: %s", err)
}
```

## bufio.Scanner 使用SplitFunc自定义分隔符

bufio.SplitFunc 是一个函数类型，其定义如下：

```go
type SplitFunc func(data []byte, atEOF bool) (advance int, token []byte, err error)
```
- data: 输入的待处理数据。
- atEOF: 标志表示是否已到达数据的末尾。
- **advance**: 表示切掉多少字节。
- **token**: 表示返回的下一个令牌（如果有的话），返回的不一定等于切掉的全部字节。
- err: 处理中遇到的任何错误。

### 实例：遇到逗号分隔
```go
// 遇到逗号分隔
func ComSplitter(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF {
		return len(data), data, nil
	}

	// 在data里向后查找
	for i := 0; i < len(data); i++ {
		// 如果发现某个字节为逗号
		if data[i] == ',' {
			// advance = i + 1：切掉包括逗号之前的字符串
			// token = data[:i]：返回从头开始不包括逗号的字符串
			// err = nil：不返回错误
			return i + 1, data[:i], nil
		}
	}
	return 0, nil, nil
}
```

### 示例: 每两个字符分割字符串
```go
// doubleCharSplitter 是一个自定义的 bufio.SplitFunc
func doubleCharSplitter(data []byte, atEOF bool) (advance int, token []byte, err error) {
    if atEOF && len(data) == 0 {
        return 0, nil, nil
    }
    if len(data) >= 2 {
        // 返回两个字符作为一个令牌
        return 2, data[:2], nil
    }
    // 如果到了文件末尾但还有未处理的数据（少于两个字符），也返回它们
    return len(data), data, nil
}

func main() {
    const input = "123456789"
    scanner := bufio.NewScanner(strings.NewReader(input))
    scanner.Split(doubleCharSplitter)

    fmt.Println("Tokens:")
    for scanner.Scan() {
        fmt.Printf("%q\n", scanner.Text())
    }
}
```

```terminal
Tokens:
"12"
"34"
"56"
"78"
"9"
```


**为什么可以循环**
- 读取机制：scanner.Scan() 在内部维护一个缓冲区，每次调用时它都会尝试从其底层的 io.Reader（在你的例子中是标准输入）中填充这个缓冲区，并从中提取下一个 token。

- 状态管理：**scanner.Scan() 返回一个布尔值**，如果 scanner.Scan() 能从输入中成功读取数据，它会更新其内部状态并准备下一次读取，返回 true。如果遇到输入结束（EOF）或错误，它会返回 false，这也会导致 for 循环终止。

### bufio.NewScanner(file) 和 bufio.NewReader(file) 有什么不同

`bufio.NewScanner` 和 `bufio.NewReader` 在 Go 语言中都用于读取数据，但它们提供了不同的接口和功能，适用于不同的使用场景。下面是两者之间的主要区别：

### bufio.NewScanner

- **用途**：`bufio.NewScanner` 主要用于逐行读取文本文件。它提供了一个简便的接口来逐行扫描文件。
- **基本方法**：使用 `Scan()` 方法逐行读取，每次调用 `Scan()` 都会移动到下一行。读取的每一行可以通过 `Text()` 或 `Bytes()` 方法获取。
- **内存管理**：默认情况下，`Scanner` 的缓冲区大小为 4096 字节，但可以通过 `Buffer()` 方法自定义缓冲区的大小和容量，这对于处理超长行非常有用。
- **错误处理**：可以通过 `Err()` 方法检查扫描过程中是否发生错误。
- **使用场景**：非常适合处理标准或结构化的文本文件，尤其是需要按行读取的情况。

### bufio.NewReader

- **用途**：`bufio.NewReader` 提供了更通用的缓冲读取功能，可以用来读取任何类型的数据，不限于文本文件。
- **基本方法**：提供多种读取方法，如 `Read()`, `ReadByte()`, `ReadBytes()`, `ReadLine()`, `ReadString()` 等。这使得 `NewReader` 在读取数据时更为灵活。
- **缓冲区**：内置的缓冲机制可以减少对底层 IO 的直接调用次数，提高读取效率。缓冲区大小默认通常为 4096 字节，但这是在底层实现中设置的，通常不需要自定义。
- **错误处理**：错误通常在各个读取方法中直接返回，需要在每次调用时检查。
- **使用场景**：更适合需要更复杂数据处理的场景，例如处理二进制文件，或者需要从数据流中读取特定格式数据（如直到遇到特定字符）。

### 示例比较

如果你只需要逐行读取文本文件，并对每一行进行处理，使用 `Scanner` 更为简单直接：

```go
scanner := bufio.NewScanner(file)
for scanner.Scan() {
    line := scanner.Text()
    // 处理 line
}
if err := scanner.Err(); err != nil {
    log.Fatal(err)
}
```

如果需要更灵活地处理数据，比如读取直到遇到特定字符，或者处理可能不仅仅是按行分割的数据，则 `NewReader` 更合适：

```go
reader := bufio.NewReader(file)
for {
    line, err := reader.ReadString('\n')
    if err != nil {
        if err == io.EOF {
            break
        }
        log.Fatal(err)
    }
    // 处理 line
}
```

总的来说，选择哪个依赖于你的具体需求，`bufio.NewScanner` 提供了更简单的接口对于基于行的读取，而 `bufio.NewReader` 则提供了更多的控制和灵活性对于各种类型的输入处理。

---

# strconv
<https://pkg.go.dev/strconv>

> 规律：字符串转换为别的类型的方法，返回值有一个error， 别的类型转换为字符串，只返回string，不返回error

#### 字符串和整数互相 Atoi Itoa
```go
i, err := strconv.Atoi("-42")   // string to int
s := strconv.Itoa(-42)          // int to string
```

#### 从字符串变成其他类型 Parse*
```go
b, err := strconv.ParseBool("true")         // to bool
f, err := strconv.ParseFloat("3.1415", 64)  // to float
i, err := strconv.ParseInt("-42", 10, 64)   // to int
u, err := strconv.ParseUint("42", 10, 64)   // to unsigned int 无符号的整数，等于正数
```

#### 从其他类型变成字符串 Format*
```go
s := strconv.FormatBool(true)                   // from bool
s := strconv.FormatFloat(3.1415, 'E', -1, 64)   // from float
s := strconv.FormatInt(-42, 16)                 // from int
s := strconv.FormatUint(42, 16)                 // from unsigned int 正数
```
#### 特殊字符转译或不转译

**strconv.Quote**
strconv.Quote 函数用于将输入的字符串转换为 Go 语言的安全引用字符串。这意味着函数会返回一个能够直接用作 Go 语言代码中字符串字面值的字符串，它会自动处理字符串中的特殊字符，例如引号、控制字符等，通过转义它们来确保字符串的正确表示和安全使用。

```go
q := strconv.Quote("Hello, 世界")
```
在这行代码中，q 将被赋值为 "\"Hello, 世界\""，包括字符串两侧的双引号。这使得 q 可以直接用作代码中的一个字符串字面值，包括所有特殊字符都已被适当转义。

**strconv.QuoteToASCII**
strconv.QuoteToASCII 函数的功能类似于 strconv.Quote，但它确保返回的字符串只包含 ASCII 字符。这是通过将所有非 ASCII 字符（比如 Unicode 字符）转换为 \uXXXX 或 \UXXXXXXXX 形式的转义序列来实现的。

```go
q := strconv.QuoteToASCII("Hello, 世界")
```
在这行代码中，q 将被赋值为 "\"Hello, \u4e16\u754c\""。这里的中文字符“世界”被转换为了其对应的 Unicode 转义序列。这使得字符串在只支持 ASCII 字符的环境中也能被正确显示或处理。

---

# 定义一个接收各种类型的函数

### 1. 使用switch语句，
   可以根据传入的值的类型执行不同的操作。这在处理不同类型的输入时非常有用，例如从用户输入中读取值并根据其类型执行不同的操作。类型值用 value.(type)来获取
```go
func printType(value interface{}) {
    switch v := value.(type) {
    case int:
        fmt.Println("int:", v)
    case string:
        fmt.Println("string:", v)
    case bool:
        fmt.Println("bool:", v)
    default:
        fmt.Println("unknown type")
    }
}
```

###  2. 使用 value.('具体类型')来断言具体值
```go
func printType(value interface{}) {
    if v, ok := value.(int); ok {
        fmt.Println("int:", v)
    } else if v, ok := value.(string); ok {
        fmt.Println("string:", v)
    } else if v, ok := value.(bool); ok {
        fmt.Println("bool:", v)
    } else {
        fmt.Println("unknown type")
    }
}
```
或 
```go
func add(a, b interface{}) interface{} {
    // 尝试将 a 和 b 转换为 int 类型
    aInt, aIsInt := a.(int)
    bInt, bIsInt := b.(int)
    // 如果ab都是Int类型，返回int类型的a+b
    if aIsInt && bIsInt{
        return aInt + bInt
    }

    // 同上，float类型
    aFloat, aIsFloat := a.(float64)
    bFloat, bIsFloat := b.(float64)

    if aIsFloat && bIsFloat{
        return aFloat + bFloat
    }
}
```

### 3. 使用 reflect 包来获取值的类型
```go
import "reflect"

func printType(value interface{}) {
    t := reflect.TypeOf(value)
    fmt.Println("Type:", t)
}
```

### 4. 使用 fmt.Sprintf 来获取值的类型
fmt.Sprintf 函数可以将任何值转换为字符串，包括其类型。这使得我们可以使用 fmt.Sprintf("%T", value) 来获取值的类型。
使用方法
fmt.Sprintf 的语法如下：

```go
s := fmt.Sprintf(format string, a ...interface{}) string

// format 是一个格式字符串，它告诉函数如何格式化后续的参数。
// a 是一个可变数量的参数，它们将根据格式字符串进行格式化。
// 函数返回一个新的字符串，该字符串包含按指定格式格式化的参数。
```

- %v：默认格式表示
- %+v：当输出结构体时，会添加字段名
- %#v：输出该值的 Go 语法表示
- %T：输出一个值的类型
- %d：十进制表示
- %s：字符串表示
- %f：浮点表示
- %t：布尔表示
- %p：指针的地址

实例:

```go
func printType(value interface{}) {
    // 使用 fmt.Sprintf 获取值的类型
    t := fmt.Sprintf("%T", value)
    fmt.Println("Type:", t)
}
```

# 泛型 generics

假如有一个方法，可以接收不同的类型参数，返回不同类型值，怎么做

例如：这个方法可以接收各种类型参数，但是返回值都是any / interface{}, 不利于对这个结果的使用。
```go
func add(a, b interface{}) interface{} {
    // 尝试将 a 和 b 转换为 int 类型
    aInt, aIsInt := a.(int)
    bInt, bIsInt := b.(int)
    // 如果ab都是Int类型，返回int类型的a+b
    if aIsInt && bIsInt{
        return aInt + bInt
    }

    // 同上，float类型
    aFloat, aIsFloat := a.(float64)
    bFloat, bIsFloat := b.(float64)

    if aIsFloat && bIsFloat{
        return aFloat + bFloat
    }
}
```

这时候可以这样,在方法名后面，参数前面加一个[]，里面填入类型代称和类型名
```go
func add[T any](a, b T) T {
    // 尝试将 a 和 b 转换为 int 类型
    aInt, aIsInt := a.(int)
    bInt, bIsInt := b.(int)
    // 如果ab都是Int类型，返回int类型的a+b
    if aIsInt && bIsInt{
        return aInt + bInt
    }

    // 同上，float类型
    aFloat, aIsFloat := a.(float64)
    bFloat, bIsFloat := b.(float64)

    if aIsFloat && bIsFloat{
        return aFloat + bFloat
    }
}
```

```go
func add[T int|float64|string](a, b T) T {
    // T的类型是int|float64|string中的一种，确保可以执行+操作
    return a + b
}
```