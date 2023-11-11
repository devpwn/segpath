# URL Manipulation Tool

This tool is a command-line utility written in Go, designed for processing a list of URLs. It allows users to modify URLs by controlling the number of path segments and optionally keeping or removing query parameters.

## Features

- **Path Segment Control**: Specify the number of path segments to keep in each URL.
- **Query Parameter Handling**: Option to remove or keep the query parameters in URLs.
- **Flexible Input Options**: Read URLs from a provided file or standard input.

## Installation

Ensure Go is installed on your system. Clone or download this repository to use the tool.

## Usage

The tool is executed using the `go run` command, with several flags for customization:

- `-p`: Number of path segments to keep in the URL (default is 0).
- `-f`: File path for reading URLs. If not provided, the tool reads from standard input.
- `-keep-query`: Keep the query parameters in the URL (default is false).

### Reading from a File

To process URLs from a file, use the `-f` flag with the file path. Example:

```bash
go run main.go -f urls.txt -p 2
```

This command processes URLs in urls.txt, keeping the first two path segments and removing query parameters.

### Reading from Standard Input

Alternatively, you can pipe URLs into the tool from standard input. For example:

```bash
cat urls.txt | segpath -p 2 -keep-query=true
```

This command reads URLs from urls.txt, keeps the first two path segments, and retains the query parameters.

1. Removing All Path Segments and Query Parameters

Command: `go run main.go -f urls.txt`

Input: `https://example.com/path1/path2?query=1`

Output: `https://example.com`

2. Keeping One Path Segment, Removing Query Parameters

Command: `go run main.go -f urls.txt -p 1`

Input: `https://example.com/path1/path2?query=1`

Output: `https://example.com/path1`

3. Keeping Path Segments and Query Parameters

Command: `cat urls.txt | go run main.go -p 2 -keep-query`

Input: `https://example.com/path1/path2/path3?query=1`

Output: `https://example.com/path1/path2?query=1`

- Proof

https://github.com/devpwn/segpath/raw/main/segpath.mov

