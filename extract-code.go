// package main

// import (
//     "bufio"
//     "fmt"
//     "io/ioutil"
//     "os"
//     "path/filepath"
//     "strings"
// )

// // Map file extensions to markdown code block languages
// var languageMap = map[string]string{
//     ".js":   "javascript",
//     ".jsx":  "javascript",
//     ".ts":   "typescript",
//     ".tsx":  "typescript",
//     ".css":  "css",
//     ".scss": "scss",
//     ".html": "html",
//     ".json": "json",
//     // ".md":   "markdown",
//     ".go":   "go",
//     // Add more mappings as needed
// }

// func getLanguage(ext string) string {
//     if lang, exists := languageMap[ext]; exists {
//         return lang
//     }
//     return ""
// }

// func main() {
//     // Get the project root directory from command line arguments or default to current working directory
//     rootDir := "."
//     if len(os.Args) > 1 {
//         rootDir = os.Args[1]
//     }

//     outputFile := "project_code.md"

//     // Open the output file
//     outFile, err := os.Create(outputFile)
//     if err != nil {
//         fmt.Printf("Error creating output file: %v\n", err)
//         os.Exit(1)
//     }
//     defer outFile.Close()

//     writer := bufio.NewWriter(outFile)

//     // Traverse the directory
//     err = filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
//         if err != nil {
//             return err
//         }

//         // Exclude directories
//         if info.IsDir() {
//             // Exclude certain directories
//             if info.Name() == "node_modules" || info.Name() == ".git" {
//                 return filepath.SkipDir
//             }
//             return nil
//         }

//         // Read files
//         if info.Mode().IsRegular() {
//             relativePath, err := filepath.Rel(rootDir, path)
//             if err != nil {
//                 return err
//             }

//             ext := filepath.Ext(path)
//             language := getLanguage(ext)

//             // Read the file content
//             content, err := ioutil.ReadFile(path)
//             if err != nil {
//                 return err
//             }

//             // Escape backticks in code
//             codeContent := strings.ReplaceAll(string(content), "```", "&#96;&#96;&#96;")

//             // Write to the markdown file
//             fmt.Fprintf(writer, "### %s\n\n```%s\n%s\n```\n\n", relativePath, language, codeContent)
//         }

//         return nil
//     })

//     if err != nil {
//         fmt.Printf("Error walking the path %q: %v\n", rootDir, err)
//         os.Exit(1)
//     }

//     // Flush the writer buffer
//     if err := writer.Flush(); err != nil {
//         fmt.Printf("Error flushing to output file: %v\n", err)
//         os.Exit(1)
//     }

//     fmt.Printf("Code extraction complete! Output written to %s\n", outputFile)
// }
