# GO 输入输出（bufio）

```go
// 创建一个输入检测对象scanner
scanner := bufio.NewScanner(os.Stdin)

// 单次输入

// 
scanner.Scan()
str := scanner.Text()
fmt.Println(str)
```






在 Go 语言中，进行输入输出操作主要可以通过 **`fmt`** 包和 **`bufio`** 包实现。以下是两者的主要使用场景和示例，帮助理解何时使用哪一个更为合适。

### **`fmt` 包**

- **场景**:
    - 用于简单的输入输出，例如从标准输入读取数据或向标准输出打印数据。
    - 适合读写不频繁且数据量不大的场合。
- **优点**:
    - 简单直接，易于使用。
    - 支持格式化输入输出。
- **示例**:
    
    ```go
    
    package main
    
    import "fmt"
    
    func main() {
        var age int
        fmt.Print("Enter your age: ")
        fmt.Scan(&age)
        fmt.Printf("Your age is: %d\n", age)
    }
    ```
    
    这个例子中，使用 **`fmt.Scan`** 从标准输入读取一个整数，然后使用 **`fmt.Printf`** 将其打印出来。
    

---

### **`bufio` 包**

**`bufio`** 包在 Go 语言中用于处理缓冲的输入和输出，提高处理大量数据时的效率。它封装了基础的 **`io.Reader`** 和 **`io.Writer`** 接口，添加了缓冲功能。这里是 **`bufio`** 的一些常见用法：

### **1. `bufio.Reader`**

**`bufio.Reader`** 用于从一个 **`io.Reader`** 对象（如文件、网络连接等）读取数据，提供缓冲和一些便利的方法来读取文本数据。

- **读取字节和字符串**:
    - **`Read`** 方法读取数据到给定的字节数组中。
    - **`ReadByte`** 和 **`ReadRune`** 方法用于读取单个字节或 UTF-8 编码的字符。
    
    ```go
    
    reader := bufio.NewReader(os.Stdin)
    byte, err := reader.ReadByte()
    rune, size, err := reader.ReadRune()
    ```
    
- **读取一行数据**:
    - **`ReadBytes`** 方法读取数据直到遇到指定的分隔符。
    - **`ReadString`** 方法类似于 **`ReadBytes`**，但直接返回字符串。
    
    ```go
    
    line, err := reader.ReadString('\n')
    ```
    
- **预读取功能**:
    - **`Peek`** 方法允许在不移动读取指针的情况下查看接下来的数据。
    
    ```go
    
    buffer, err := reader.Peek(5)
    ```
    

### **2. `bufio.Writer`**

**`bufio.Writer`** 用于向 **`io.Writer`** 对象写数据，例如文件或网络连接。它提供了缓冲以及相关的写入方法。

- **写入数据**:
    - **`Write`** 方法接受一个字节切片并将其写入缓冲区。
    - **`WriteByte`** 和 **`WriteRune`** 方法用于写入单个字节或字符。
    - **`WriteString`** 方法直接写入字符串数据。
    
    ```go
    writer := bufio.NewWriter(os.Stdout)
    writer.WriteString("Hello, World!\n")
    writer.Flush() // 确保所有缓冲的数据被写出
    ```
    

### **3. `bufio.Scanner`**

**`bufio.Scanner`** 是用来逐行或逐段（基于自定义的分隔符）读取数据的便捷工具，通常用于读取标准输入或文件的行。

- **基本用法**:
    - 默认按行分隔（**`\n`**），可以通过 **`Split`** 方法自定义分隔函数。
    
    ```go
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        fmt.Println("Read line:", scanner.Text())
    }
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "reading input:", err)
    }
    ```
    

### **4. 自定义分隔函数**

可以通过 **`Split`** 方法为 **`bufio.Scanner`** 设置自定义的分隔函数，例如按空格分隔或其他任何标准。

```go
scanner := bufio.NewScanner(os.Stdin)
scanner.Split(bufio.ScanWords)
for scanner.Scan() {
    fmt.Println("Read word:", scanner.Text())
}
```

这些是 **`bufio`** 包中的一些常见用法，它们提供了比直接使用 **`io`** 包更高效和更方便的数据处理方式。通过这种方式，可以显著提高处理大量数据时的性能，尤其是在读取或写入大文件时。

### **总结**

选择 **`fmt`** 还是 **`bufio`** 主要取决于应用的具体需求：

- 如果需要简单的、偶尔的输入输出，或需要格式化输出，使用 **`fmt`**。
- 如果处理的是大量数据或需要高效率的连续数据处理，尤其是文件操作，使用 **`bufio`** 会更加高效。