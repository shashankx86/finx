# Finx

`finx` is an advanced find utility written in Go. 

## Features

- **Type Filtering**
- **Max Depth**
- **more comming soon...**
  
## Installation

To install `finx`, follow the steps below:

### 1. Clone the repository

```bash
git clone https://github.com/shashankx86/finx.git
```

### 2. Build the project

Navigate to the project folder and build the executable:

```bash
cd finx
go build -o finx ./cmd
```

### 3. Run the program

After building the binary, you can run it using:

```bash
./finx <directory_path> <pattern>
```

## Usage

### Basic Syntax

```bash
./finx <directory_path> <pattern> [flags]
```

### Positional Arguments

- **`<directory_path>`**: The path where you want to start your search. (e.g., `./`, `/path/to/search`).
- **`<pattern>`**: The file pattern to search for. Use wildcard characters like `*` or `?` (e.g., `*.go`, `test*`).

### Flags

- **`-type`**: Filter by file type.
  - `f` - Search for files only.
  - `d` - Search for directories only.

- **`-maxdepth`**: Set the maximum recursion depth.
  - Example: `-maxdepth 2` limits the search to two levels deep.

- **`-v`**: Enable verbose output to show the progress of the search.

### Example Usage

```bash
# Search for all Go files in a directory
./finx /path/to/search "*.go"

# Search for all text files in the current directory, limit to 2 levels of depth
./finx /path/to/search "*.txt" -maxdepth 2

# Search for all directories named "test" in a specific path, verbose output enabled
./finx /path/to/search "test*" -type d -v

# Search for files only, no directories, verbose output
./finx /path/to/search "*.md" -type f -v
```

### Example Commands

- **Search for all `.go` files in the current directory:**
  ```bash
  ./finx . "*.go"
  ```

- **Search for `.txt` files, limiting the depth to 3 levels:**
  ```bash
  ./finx /path/to/directory "*.txt" -maxdepth 3
  ```

- **Search for directories only (`-type d`):**
  ```bash
  ./finx /path/to/search "test*" -type d
  ```

- **Enable verbose output to see which files are being processed:**
  ```bash
  ./finx /path/to/search "*.log" -v
  ```

## Flags Explained

- **`-type f`**: Search only for files, ignoring directories.
  
  Example:
  ```bash
  ./finx /path/to/search "*.txt" -type f
  ```

- **`-type d`**: Search only for directories, ignoring files.
  
  Example:
  ```bash
  ./finx /path/to/search "docs*" -type d
  ```

- **`-maxdepth <n>`**: Limits the search depth to `n` levels.
  
  Example:
  ```bash
  ./finx /path/to/search "*.go" -maxdepth 2
  ```

- **`-v`**: Verbose output that displays the files being checked during the search.
  
  Example:
  ```bash
  ./finx /path/to/search "*.log" -v
  ```

## License

This project is licensed under the **GPL-2.0 License**. See the LICENSE file for more details.

## Contributing

Feel free to submit issues or pull requests if you'd like to contribute. Contributions are welcome!

---
