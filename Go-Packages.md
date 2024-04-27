# Go Packages

---

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
content, err := os.ReadFile(filename)
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


