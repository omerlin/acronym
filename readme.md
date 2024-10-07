# Acronym Manager

## Overview

The **Acronym Manager** is a Go-based application that allows users to:
- Search for acronyms and their definitions.
- Add new acronyms and their definitions.

The application supports two modes:
1. **CLI Mode**: Search for or add acronyms via the command line.
2. **HTTP Mode**: Launch a web server with a Bootstrap-based UI for searching and adding acronyms.

## Features

- **Command Line Interface (CLI)** for searching and adding acronyms.
- **Web Interface** built with **Go Fiber** and styled with **Bootstrap**.
- **YAML-based storage** of acronyms and their definitions.

## Directory Structure

```plaintext
acronym-manager/
├── main.go                   # Main Go code
├── acronyms.yaml             # YAML file for storing acronyms and definitions
├── static/                   # Static files (CSS, JS, etc.)
│   ├── css/
│   │   └── bootstrap.min.css  # Bootstrap CSS
│   └── js/
│       └── bootstrap.bundle.min.js  # Bootstrap JS
└── templates/
    └── index.html            # HTML template for web interface
```

## Dependencies

To build and run the application, you need the following dependencies:

- **Go Fiber**: Web framework for Go.
- **Go YAML**: For parsing and writing YAML files.
- **Bootstrap**: For the web UI (local files included).

### Install Dependencies

Install the required Go packages:

```bash
go get github.com/gofiber/fiber/v2
go get github.com/gofiber/template/v2/html
go get github.com/go-yaml/yaml```

## Running the Application

### 1. CLI Mode

You can run the application in CLI mode to search for or add acronyms.

#### Search for an Acronym

```bash
go run main.go -acronym=<ACRONYM>
```

Example:
```bash
go run main.go -acronym=API
```

This will return the definition of the acronym `API`, if it exists in the `acronyms.yaml` file.

### Add a New Acronym

```bash
go run main.go -acronym=<ACRONYM> -definition=<DEFINITION>
```
Example:

```bash
go run main.go -acronym=JSON -definition="JavaScript Object Notation"
```
This will add the acronym `JSON` with its definition to the `acronyms.yaml` file.

### 2. HTTP Mode

To run the application in HTTP mode and launch the web interface:

```bash
go run main.go -http
```

Visit [http://localhost:3000](http://localhost:3000) in your browser to:

- Search for acronyms.
- Add new acronyms via the web interface.

### Command Line Flags

| Flag         | Description                                               | Required        |
| ------------ | --------------------------------------------------------- | --------------- |
| `-acronym`   | The acronym to look up or add.                            | Yes (CLI mode)  |
| `-definition`| The definition to add (required when adding a new acronym).| No (CLI mode)   |
| `-http`      | Runs the application in HTTP mode (launches the web server).| No             |

## YAML Storage
The acronyms and their definitions are stored in an acronyms.yaml file. Below is an example format:

```yaml
acronyms:
  API: Application Programming Interface
  HTTP: HyperText Transfer Protocol
  CLI: Command Line Interface
  JSON: JavaScript Object Notation
```

Example YAML File

```yaml
acronyms:
  API: Application Programming Interface
  HTTP: HyperText Transfer Protocol
  JSON: JavaScript Object Notation
  CLI: Command Line Interface
  ```

## Web Interface

### Pages

- **Home Page** ([http://localhost:3000/](http://localhost:3000/)): Provides a form to search for acronyms and another form to add a new acronym.

### Local Bootstrap Files

- **Bootstrap CSS**: `/static/css/bootstrap.min.css`
- **Bootstrap JS**: `/static/js/bootstrap.bundle.min.js`

## TODO: move to Embed 
[https://zikani.hashnode.dev/go-embed-fs-with-rebed-fiber-web-framework]


---

### Explanation:

1. **Overview**: Describes the application and its two modes (CLI and HTTP).
2. **Directory Structure**: Explains how the files and folders are structured.
3. **Dependencies**: Lists the Go packages needed and the installation commands.
4. **Running the Application**: Details the commands to run the app in CLI or HTTP mode, along with examples.
5. **Web Interface**: Provides an overview of the web interface, Bootstrap file usage, and a sample `index.html` file.
6. **License**: A placeholder section for the project’s license.