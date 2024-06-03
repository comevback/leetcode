package main

func main() {
    // 解析命令行参数
    url := flag.String("url", "", "The URL of the LeetCode problem")
    flag.Parse()

    if *url == "" {
        fmt.Println("Please provide a LeetCode problem URL using the -url flag.")
        return
    }

    problem, err := fetchProblem(*url)
    if err != nil {
        panic(err)
    }

    tmpl := `# {{.ID}}. {{.Title}}
{{.Status}}
{{.Difficulty}}
Topics
Companies
{{.Description}}

{{range .Examples}}Example:
> Input: {{.Input}}
Output: {{.Output}}
Explanation: {{.Explanation}}

{{end}}{{.Constraints}}

---

# Code
\```go
{{.Code}}
```
`

    t := template.Must(template.New("problem").Parse(tmpl))
    f, err := os.Create(fmt.Sprintf("%s_%s.md", problem.ID, strings.ReplaceAll(problem.Title, " ", "_")))
    if err != nil {
        panic(err)
    }
    defer f.Close()

    if err := t.Execute(f, problem); err != nil {
        panic(err)
    }
}