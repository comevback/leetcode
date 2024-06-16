package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// 检查命令行参数数量是否正确，如果不是两个参数，则输出使用说明并退出
	// check if the number of command line arguments is correct, if not two arguments, print usage and exit
	if len(os.Args) != 2 {
		fmt.Println("Usage:\nleetcode-setup -n={Question Name}")
		return
	}

	// 使用flag包来解析命令行参数，定义一个名为-n的字符串参数
	// use the flag package to parse command line arguments, define a string parameter named -n
	qsInput := flag.String("n", "", "name of the question")
	flag.Parse() // 解析命令行参数

	var questionName string
	// 如果-n参数被用户设置了，则使用该参数作为问题名称
	// if the -n parameter is set by the user, use it as the question name
	if *qsInput != "" {
		questionName = *qsInput
	} else {
		// 否则，使用第一个非标志参数作为问题名称
		// otherwise, use the first non-flag argument as the question name
		args := flag.Args()
		questionName = args[0]
	}

	// 将问题名称中的点空格（". "）和空格替换成连字符（"-"），用于创建目录和文件名
	// replace dots and spaces in the question name with hyphens for creating directory and file names
	newDirName := strings.ReplaceAll(questionName, ". ", "-")
	newDirName = strings.ReplaceAll(newDirName, " ", "-")

	// 构建Markdown文件和Go文件的文件名
	// build the file names for the Markdown file and the Go file
	mdFileName := newDirName + ".md"
	goFileName := "main.go"

	// 获取当前工作目录的路径
	// get the path of the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Failed to get current working directory:", err)
		return
	}

	// 创建新目录的完整路径
	// create the full path of the new directory
	newDir := filepath.Join(cwd, newDirName)
	newMdFile := filepath.Join(newDir, mdFileName)
	newGoFile := filepath.Join(newDir, goFileName)

	// 输出文件路径（调试用）
	// print file paths (for debugging)
	// fmt.Println(newMdFile)
	// fmt.Println(newGoFile)

	// 创建目录，包括任何必需的父目录
	// create the directory, including any necessary parent directories
	err = os.MkdirAll(newDir, 0755)
	if err != nil {
		fmt.Println("Failed to create directory:", err)
		return
	}

	// 创建并打开Markdown文件，设置权限为0755
	// create and open the Markdown file, set the permission to 0755
	mdFile, err := os.OpenFile(newMdFile, os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		fmt.Println("Failed to open or create Markdown file:", err)
		return
	}
	defer mdFile.Close() // 确保文件最终关闭 (ensure the file is closed eventually)

	// 创建并打开Go源文件，设置权限为0755
	// create and open the Go source file, set the permission to 0755
	goFile, err := os.OpenFile(newGoFile, os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		fmt.Println("Failed to open or create Go file:", err)
		return
	}
	defer goFile.Close() // 确保文件最终关闭 (ensure the file is closed eventually)

	writer := bufio.NewWriter(goFile)                // 创建缓冲写入器以写入Go文件 (create a buffered writer to write to the Go file)
	goContent := "package main\n\nfunc main(){\n\n}" // Go文件的基本内容模板 (basic content template for the Go file)

	// 将Go文件的内容写入文件
	// write the content of the Go file to the file
	_, err = writer.WriteString(goContent)
	if err != nil {
		fmt.Println("Failed to write to Go file:", err)
		return
	}

	// 刷新缓冲写入器，确保所有内容都被写入文件
	// flush the buffered writer to ensure all content is written to the file
	err = writer.Flush()
	if err != nil {
		fmt.Println("Failed to flush writer:", err)
		return
	}

	// 如果一切顺利，打印成功创建文件的消息
	// if everything goes well, print the message of successfully creating files
	fmt.Printf("Successfully created files for question: %s\n", questionName)

	// 提示用户可以使用cd命令进入新目录
	// remind the user that they can use the cd command to enter the new directory
	fmt.Printf("You can use this command to enter the directory:\n\ncd %s\n\n", newDirName)
}
