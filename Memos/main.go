package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
)

func main() {

	var filename string
	fmt.Print("file name: ")
	fmt.Scan(&filename)

	format(filename)
}

func format(filename string) {
	file, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// add # to the start
	var lines []string
	lines = append(lines, "# ")

	// iterate over the file line by line, and append the line to the lines slice
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		if line == "Solved\n" || line == "Topics\n" || line == "Companies\n" {
			line = "\n"
		}

		re := regexp.MustCompile(`Example \d:\n`)
		if re.Match([]byte(line)) {
			line = re.ReplaceAllString(line, "\n$0 ")
		}
		lines = append(lines, line)

	}

	// 清空文件并准备写入
	if _, err = file.Seek(0, 0); err != nil {
		log.Fatal(err)
	}
	if err = file.Truncate(0); err != nil {
		log.Fatal(err)
	}

	// 将更改后的内容写回文件
	writer := bufio.NewWriter(file)
	for _, line := range lines {
		if _, err = writer.WriteString(line); err != nil {
			log.Fatal(err)
		}
	}
	if err = writer.Flush(); err != nil {
		log.Fatal(err)
	}
}
