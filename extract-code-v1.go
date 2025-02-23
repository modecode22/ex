// package main

// import (
//     "bufio"
//     "flag"
//     "fmt"
//     "io/ioutil"
//     "os"
//     "path/filepath"
//     "strings"
//     "time"
//     "sort"
// )

// type FileEntry struct {
//     Path     string
//     Language string
//     Size     int64
//     ModTime  time.Time
//     Content  string
// }

// type Config struct {
//     SourceDir     string
//     OutputDir     string
//     ExcludeDirs   []string
//     ExcludeExts   []string
//     IncludeHidden bool
//     ProjectName   string
// }

// // Map file extensions to markdown code block languages
// var languageMap = map[string]string{
//     ".js":    "javascript",
//     ".jsx":   "javascript",
//     ".ts":    "typescript",
//     ".tsx":   "typescript",
//     ".css":   "css",
//     ".scss":  "scss",
//     ".html":  "html",
//     ".json":  "json",
//     ".py":    "python",
//     ".java":  "java",
//     ".cpp":   "cpp",
//     ".c":     "c",
//     ".h":     "c",
//     ".hpp":   "cpp",
//     ".rb":    "ruby",
//     ".php":   "php",
//     ".go":    "go",
//     ".rs":    "rust",
//     ".swift": "swift",
//     ".kt":    "kotlin",
//     ".sql":   "sql",
//     ".yaml":  "yaml",
//     ".yml":   "yaml",
//     ".xml":   "xml",
//     ".sh":    "bash",
//     ".bat":   "batch",
//     ".ps1":   "powershell",
//     ".md":    "markdown",
//     ".lua":   "lua",
//     ".r":     "r",
//     ".dart":  "dart",
// }

// func getLanguage(ext string) string {
//     if lang, exists := languageMap[ext]; exists {
//         return lang
//     }
//     return ""
// }

// func formatSize(size int64) string {
//     const unit = 1024
//     if size < unit {
//         return fmt.Sprintf("%d B", size)
//     }
//     div, exp := int64(unit), 0
//     for n := size / unit; n >= unit; n /= unit {
//         div *= unit
//         exp++
//     }
//     return fmt.Sprintf("%.1f %cB", float64(size)/float64(div), "KMGTPE"[exp])
// }

// func shouldExclude(path string, config Config) bool {
//     // Check if it's a hidden file/directory
//     if !config.IncludeHidden && strings.HasPrefix(filepath.Base(path), ".") {
//         return true
//     }

//     // Check excluded directories
//     for _, dir := range config.ExcludeDirs {
//         if dir != "" && strings.Contains(path, dir) {
//             return true
//         }
//     }

//     // Check excluded extensions
//     ext := filepath.Ext(path)
//     for _, excludeExt := range config.ExcludeExts {
//         if excludeExt != "" && ext == excludeExt {
//             return true
//         }
//     }

//     return false
// }

// func ensureDirectory(path string) error {
//     return os.MkdirAll(path, 0755)
// }

// func main() {
//     // Parse command line flags
//     var config Config
    
//     flag.StringVar(&config.SourceDir, "src", ".", "Source directory to extract code from")
//     flag.StringVar(&config.ProjectName, "name", "", "Project name (default: directory name)")
//     excludeDirsStr := flag.String("exclude-dirs", "node_modules,.git,vendor,dist,build", "Comma-separated list of directories to exclude")
//     excludeExtsStr := flag.String("exclude-exts", "", "Comma-separated list of file extensions to exclude (e.g., .exe,.dll)")
//     flag.BoolVar(&config.IncludeHidden, "include-hidden", false, "Include hidden files and directories")
//     outputPath := flag.String("output", "", "Output directory (default: Desktop/code-extracts)")
//     verbose := flag.Bool("verbose", false, "Show verbose output")
    
//     flag.Parse()

//     // Process excluded directories and extensions
//     config.ExcludeDirs = strings.Split(*excludeDirsStr, ",")
//     config.ExcludeExts = strings.Split(*excludeExtsStr, ",")

//     // Get absolute path of source directory
//     absSourceDir, err := filepath.Abs(config.SourceDir)
//     if err != nil {
//         fmt.Printf("Error resolving source directory: %v\n", err)
//         os.Exit(1)
//     }
//     config.SourceDir = absSourceDir

//     if *verbose {
//         fmt.Printf("Source directory: %s\n", config.SourceDir)
//         fmt.Printf("Excluded directories: %v\n", config.ExcludeDirs)
//         fmt.Printf("Excluded extensions: %v\n", config.ExcludeExts)
//     }

//     // Set project name if not provided
//     if config.ProjectName == "" {
//         config.ProjectName = filepath.Base(absSourceDir)
//     }

//     // Setup output directory
//     if *outputPath == "" {
//         homeDir, err := os.UserHomeDir()
//         if err != nil {
//             fmt.Printf("Error getting home directory: %v\n", err)
//             os.Exit(1)
//         }
//         config.OutputDir = filepath.Join(homeDir, "Desktop", "code-extracts", config.ProjectName)
//     } else {
//         config.OutputDir = filepath.Join(*outputPath, config.ProjectName)
//     }

//     // Ensure output directory exists
//     if err := ensureDirectory(config.OutputDir); err != nil {
//         fmt.Printf("Error creating output directory: %v\n", err)
//         os.Exit(1)
//     }

//     // Collect all files
//     var files []FileEntry
//     err = filepath.Walk(config.SourceDir, func(path string, info os.FileInfo, err error) error {
//         if err != nil {
//             return err
//         }

//         if *verbose {
//             fmt.Printf("Checking file: %s\n", path)
//         }

//         if info.IsDir() {
//             if shouldExclude(path, config) {
//                 if *verbose {
//                     fmt.Printf("Skipping directory: %s\n", path)
//                 }
//                 return filepath.SkipDir
//             }
//             return nil
//         }

//         if info.Mode().IsRegular() {
//             if shouldExclude(path, config) {
//                 if *verbose {
//                     fmt.Printf("Skipping file: %s\n", path)
//                 }
//                 return nil
//             }

//             relativePath, err := filepath.Rel(config.SourceDir, path)
//             if err != nil {
//                 return err
//             }

//             content, err := ioutil.ReadFile(path)
//             if err != nil {
//                 return err
//             }

//             if *verbose {
//                 fmt.Printf("Adding file: %s\n", relativePath)
//             }

//             files = append(files, FileEntry{
//                 Path:     relativePath,
//                 Language: getLanguage(filepath.Ext(path)),
//                 Size:     info.Size(),
//                 ModTime:  info.ModTime(),
//                 Content:  string(content),
//             })
//         }
//         return nil
//     })

//     if err != nil {
//         fmt.Printf("Error walking the path %q: %v\n", config.SourceDir, err)
//         os.Exit(1)
//     }

//     // Sort files by path
//     sort.Slice(files, func(i, j int) bool {
//         return files[i].Path < files[j].Path
//     })

//     // Create output files
//     timestamp := time.Now().Format("2006-01-02-15-04-05")
//     outputFile := filepath.Join(config.OutputDir, fmt.Sprintf("%s-%s.md", config.ProjectName, timestamp))
    
//     outFile, err := os.Create(outputFile)
//     if err != nil {
//         fmt.Printf("Error creating output file: %v\n", err)
//         os.Exit(1)
//     }
//     defer outFile.Close()

//     writer := bufio.NewWriter(outFile)

//     // Write header
//     fmt.Fprintf(writer, "# %s - Code Documentation\n\n", config.ProjectName)
//     fmt.Fprintf(writer, "## Project Information\n\n")
//     fmt.Fprintf(writer, "- **Project Name:** %s\n", config.ProjectName)
//     fmt.Fprintf(writer, "- **Source Directory:** `%s`\n", config.SourceDir)
//     fmt.Fprintf(writer, "- **Generated On:** %s\n", time.Now().Format("2006-01-02 15:04:05"))
//     fmt.Fprintf(writer, "- **Total Files:** %d\n\n", len(files))

//     // Write table of contents
//     fmt.Fprintf(writer, "## Table of Contents\n\n")
//     for _, file := range files {
//         fmt.Fprintf(writer, "- [%s](#%s)\n", file.Path, strings.ReplaceAll(file.Path, "/", "-"))
//     }
//     fmt.Fprintf(writer, "\n---\n\n")

//     // Write file contents
//     for _, file := range files {
//         fmt.Fprintf(writer, "## %s\n\n", file.Path)
//         fmt.Fprintf(writer, "**File Information:**\n")
//         fmt.Fprintf(writer, "- Size: %s\n", formatSize(file.Size))
//         fmt.Fprintf(writer, "- Last Modified: %s\n", file.ModTime.Format("2006-01-02 15:04:05"))
//         if file.Language != "" {
//             fmt.Fprintf(writer, "- Language: %s\n", file.Language)
//         }
//         fmt.Fprintf(writer, "\n")

//         // Escape backticks in code
//         codeContent := strings.ReplaceAll(file.Content, "```", "&#96;&#96;&#96;")

//         // Write the code block
//         fmt.Fprintf(writer, "```%s\n%s\n```\n\n", file.Language, codeContent)
//         fmt.Fprintf(writer, "---\n\n")
//     }

//     if err := writer.Flush(); err != nil {
//         fmt.Printf("Error flushing to output file: %v\n", err)
//         os.Exit(1)
//     }

//     fmt.Printf("Documentation generated successfully!\n")
//     fmt.Printf("Output file: %s\n", outputFile)
//     fmt.Printf("Total files processed: %d\n", len(files))
// }
