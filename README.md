# Code Extractor (ex) - Comprehensive Documentation

[![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://golang.org/)

This tool generates a Markdown documentation file containing the source code from a specified project directory. It's designed to be simple, fast, configurable, and cross-platform (macOS, Linux, Windows).  The tool excludes images, videos, fonts, and archive files by default, focusing on extracting source code.

## Table of Contents

- [Code Extractor (ex) - Comprehensive Documentation](#code-extractor-ex---comprehensive-documentation)
  - [Table of Contents](#table-of-contents)
  - [Installation](#installation)
    - [macOS / Linux](#macos--linux)
    - [Windows](#windows)
    - [Verification](#verification)
    - [Verification](#verification-1)
  - [Usage](#usage)
    - [Options](#options)
    - [Examples](#examples)
  - [Configuration](#configuration)
    - [Excluded File Types](#excluded-file-types)
    - [Language Highlighting](#language-highlighting)
  - [Troubleshooting](#troubleshooting)

## Installation

These instructions explain how to install the Code Extractor and make it accessible via the command `ex`.

### macOS / Linux

1.  **Get the Source Code:**

    You'll need the Go source code (`main.go`).  You can:

    *   **Copy and Paste:** Copy the Go code provided in the previous responses (the complete, improved Go code) and paste it into a new file named `main.go`.
    *   **Download:** If you have a link to a repository (e.g., on GitHub), you can clone it:

        ```bash
        git clone <repository_url>
        cd <repository_directory>
        ```

2.  **Compile:**

    Open a terminal and navigate to the directory containing `main.go`:

    ```bash
    go build
    ```

    This creates an executable file (likely named `main`).

3.  **Install Globally (as `ex`):**

    ```bash
    sudo mv main /usr/local/bin/ex
    sudo chmod +x /usr/local/bin/ex
    ```

    *   **`sudo mv main /usr/local/bin/ex`**:  Renames the executable to `ex` and moves it to `/usr/local/bin`. This directory is usually in your system's `PATH`, making `ex` available everywhere.  `sudo` requires your administrator password.
    *   **`sudo chmod +x /usr/local/bin/ex`**: Makes the `ex` file executable.

### Windows

1.  **Get the Source Code:**  (Same as macOS/Linux - copy/paste or download the `main.go` file).

2.  **Compile:**

    Open a Command Prompt or PowerShell *as a regular user* (not as administrator) and navigate to where you saved `main.go`:

    ```bash
    go build
    ```

    This creates an executable file (likely named `main.exe`).

3.  **Install Globally (as `ex`):**

    *   **Rename:** Rename the executable to `ex.exe`.

    *   **Move to a PATH Directory:** You have two good options:

        *   **Option 1: (Recommended) Create a `bin` Directory:**

            1.  Create a folder named `bin` in your user directory (e.g., `C:\Users\YourName\bin`).  This keeps your personal tools separate.
            2.  Move `ex.exe` into this `bin` folder.
            3.  Add `bin` to your User `PATH` (NOT the System `PATH`):
                *   Search Windows for "environment variables" and select "Edit the system environment variables."
                *   Click "Environment Variables...".
                *   Under **"User variables for YourName"** (IMPORTANT: Use *User* variables), select `Path` and click "Edit...".
                *   Click "New" and add the *full* path: `C:\Users\YourName\bin` (replace `YourName` with your actual username).
                *   Click "OK" on all dialogs.
                *   **Restart your terminal** (or even your computer) for the changes to take effect.

        *   **Option 2: Use an Existing PATH Directory:**

            1.  Find a directory *already* in your User `PATH`.  Open a command prompt and type:

                ```bash
                echo %PATH%
                ```

                Look for directories that seem appropriate for user-installed tools.  *Avoid* system directories like `C:\Windows\System32`. If you have a directory specifically for development tools, that's a good choice.
            2.  Move `ex.exe` into that directory.

### Verification

Open a *new* terminal (or command prompt) and type:
```bash
ex -h
```

### Verification

Open a *new* terminal (or command prompt) and type:

```bash
ex -h
```

You should see the help message, confirming successful installation.  If you see an error (like `command not found` or `'ex' is not recognized...`), proceed to the [Troubleshooting](#troubleshooting) section below.

---

## Usage

The basic command structure is:

```bash
ex [options]
```

### Options

The following table lists the available command-line options:

| Option             | Description                                                                                                        | Default Value                              |
| ------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------ |
| `-src <directory>`  | Specifies the source directory to process. All files and subdirectories within this directory will be scanned.     | `.` (the current working directory)          |
| `-name <name>`     | Sets the project name, which will be used in the generated documentation's title and header.                       | The name of the source directory.           |
| `-exclude-dirs <dirs>` | Provides a comma-separated list of directory names to exclude.  These are matched against the *full path*.      | `node_modules,.git,vendor,dist,build`    |
| `-exclude-exts <exts>` | Provides a comma-separated list of file extensions to exclude (e.g., `.exe,.dll`).                               | (None, but see [Excluded File Types](#excluded-file-types)) |
| `-include-hidden`   | Includes hidden files and directories (those starting with a dot `.`).                                            | `false`                                    |
| `-output <directory>` | Specifies the directory where the generated Markdown file will be created.                                       | `Desktop/code-extracts/<project_name>`    |
| `-verbose`          | Enables verbose output, printing detailed information about each file being processed. Helpful for debugging.     | `false`                                    |
| `-h`               | Displays the help message, showing all available options and their descriptions.                               | (N/A)                                      |

### Examples

*   **Basic Usage:** Process the current directory and save the output to the default location (Desktop/code-extracts/yourProjectName):

    ```bash
    ex
    ```

*   **Specify Source Directory:** Process a specific project located at `/path/to/your/project`:

    ```bash
    ex -src /path/to/your/project
    ```

*   **Exclude Additional Directories:** Exclude the `tests` and `.vscode` directories, in addition to the default exclusions:

    ```bash
    ex -exclude-dirs "node_modules,.git,vendor,dist,build,tests,.vscode"
    ```

*   **Exclude Specific File Extensions:** Exclude `.log` and `.bak` files:

    ```bash
    ex -exclude-exts ".log,.bak"
    ```

*   **Include Hidden Files and Directories:** Include all hidden files and directories in the output:

    ```bash
    ex -include-hidden
    ```

*   **Specify a Custom Output Directory:** Save the generated Markdown file to `/my/documentation/folder`:

    ```bash
    ex -output /my/documentation/folder
    ```

*   **Combined Example (Realistic Scenario):**

    ```bash
    ex -src ~/my-project -name "My Awesome Project" -exclude-dirs "docs,tmp" -exclude-exts ".txt" -output ~/project-docs -verbose
    ```

    This command does the following:
    1.  Processes the code within the `~/my-project` directory.
    2.  Uses "My Awesome Project" as the project name in the output.
    3.  Excludes the `docs` and `tmp` directories.
    4.  Excludes all files with the `.txt` extension.
    5.  Saves the generated Markdown file to the `~/project-docs` directory.
    6.  Enables verbose output, showing detailed processing information in the terminal.

---

## Configuration

### Excluded File Types

The tool *automatically* excludes several common file types that are typically not relevant to code documentation:

*   **Images:** `.jpg`, `.jpeg`, `.png`, `.gif`, `.bmp`, `.svg`, `.ico`, `.webp`, `.tiff`
*   **Videos:** `.mp4`, `.avi`, `.mov`, `.wmv`, `.mkv`, `.flv`, `.webm`
*   **Fonts:** `.woff`, `.woff2`, `.ttf`, `.otf`, `.eot`
*   **Archives:** `.zip`, `.tar`, `.gz`, `.rar`, `.7z`

This built-in exclusion list helps keep the output focused on source code.  You can *add* to this list using the `-exclude-exts` option (e.g., to exclude `.log` files), but you *cannot* override or remove these default exclusions.

### Language Highlighting

The tool automatically detects the programming language of each file based on its extension.  It then uses the correct Markdown syntax for code block highlighting. The table below shows the supported languages and their corresponding extensions:

| Extension | Language     | Extension | Language      |
| --------- | ------------ | --------- | ------------- |
| `.js`     | javascript   | `.rb`     | ruby          |
| `.jsx`    | javascript   | `.php`    | php           |
| `.ts`     | typescript   | `.go`     | go            |
| `.tsx`    | typescript   | `.rs`     | rust          |
| `.css`    | css          | `.swift`  | swift         |
| `.scss`   | scss         | `.kt`     | kotlin        |
| `.html`   | html         | `.sql`    | sql           |
| `.json`   | json         | `.yaml`   | yaml          |
| `.py`     | python       | `.yml`    | yaml          |
| `.java`   | java         | `.xml`    | xml           |
| `.cpp`    | cpp          | `.sh`     | bash          |
| `.c`      | c            | `.bat`    | batch         |
| `.h`      | c            | `.ps1`    | powershell    |
| `.hpp`    | cpp          | `.md`     | markdown      |
|           |              | `.lua`    | lua           |
|           |              | `.r`      | r             |
|           |              | `.dart`   | dart          |
|           |              | `.txt`    | plaintext     |
If a file extension is not recognized, the code block will still be included in the output, but without any language-specific syntax highlighting.

---
## Troubleshooting

This section provides solutions to common problems.

*   **`ex: command not found` (macOS/Linux):**

    *   **Permissions:** Ensure you used `sudo` correctly during installation: `sudo mv ...` and `sudo chmod +x ...`.
    *   **PATH:** Verify that `/usr/local/bin` is in your `$PATH` environment variable.  Run `echo $PATH` in your terminal. If it's *not* listed, you might need to add it to your shell's configuration file (like `~/.bashrc`, `~/.zshrc`, or `~/.profile`). This is less common, but it can happen. *After modifying your shell configuration, open a new terminal window.*
    *   **New Terminal:** Always open a *new* terminal window after installation or making changes to your `PATH`.
    *   **Executable Bit:** Confirm the file has execute permissions: `ls -l /usr/local/bin/ex`. The output should start with `-rwxr-xr-x` (or similar, indicating execute permissions).

*   **`'ex' is not recognized as an internal or external command...` (Windows):**

    *   **PATH:** Double-check that you added the *correct* directory to your *User* `PATH` environment variable (not the System `PATH`, unless you specifically intend for all users on the system to have access). Verify the spelling carefully.
    *   **New Terminal:** Always open a *new* Command Prompt or PowerShell window after modifying the `PATH`.
    *   **Restart:** Sometimes, Windows requires a full system restart for changes to the `PATH` environment variable to take full effect.  Try restarting your computer.
    *   **Executable Name:** Make sure the executable file is named `ex.exe` (no spaces or other characters).

*   **Incorrect or Missing Output:**

    *   **Verbose Mode:** Run the tool with the `-verbose` flag (e.g., `ex -verbose`). This will print detailed information about each file being processed, including any files that are skipped due to exclusions. This is very helpful for debugging.
    *   **Hidden Files:** Remember that hidden files and directories (those starting with a dot `.`) are *excluded* by default. Use the `-include-hidden` flag to include them.
    *   **File Extensions:** Carefully check your `-exclude-exts` option for typos or unintended exclusions.
    * **Output Directory:** Verify the output file path using `-verbose`.

*   **Go Build Errors:** If you encounter errors during the `go build` step, ensure that you have Go installed correctly and that your Go environment (GOPATH, etc.) is set up properly. Consult the official Go documentation for installation and setup instructions.

---
