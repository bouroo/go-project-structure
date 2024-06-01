# Go Project Structure

Welcome to the **Go Project Structure** repository! This project is dedicated to learning for organizing a Go web application. Whether you are new to Go or looking to improve your project architecture, this guide will inspire you create a maintainable and scalable Go project.

## Table of Contents

- [Introduction](#introduction)
- [Getting Started](#getting-started)
- [Project Structure](#project-structure)
- [Dependencies](#dependencies)
- [Building and Running](#building-and-running)
- [Testing](#testing)
- [Contributing](#contributing)
- [License](#license)

## Introduction

In this repository, you will find a sample Go web application with a commonly used folder structure. This structure helps in maintaining clear separation of concerns, improving code readability, and making it easier to manage large projects.

## Getting Started

To get started with this project, you'll need to have Go installed on your machine. Follow the instructions on the [official Go website](https://go.dev/doc/install) to set up your Go environment.

### Clone the Repository

```bash
git clone https://github.com/bouroo/go-project-structure.git
cd go-project-structure
```

### Install Dependencies
```bash
go mod tidy
```

## Project Structure

Here is an overview of the project structure:

```bash
go-project-structure/
├── cmd/                 # Main applications of the project
│   └── app/             # Specific application folder
│       └── main.go      # Entry point of the application
├── internal/            # Private application and library code
│   └── .../             # Specific application folder
│       ├── handler/     # HTTP handlers
│       ├── service/     # Business logic
│       └── repository/  # Data access
├── pkg/                 # Library code that's ok to use by external applications
├── web/                 # Web related files (static assets, templates)
├── configs/             # Configuration files
├── infrastructure/      # Infrastructure code and configuration
├── scripts/             # Scripts to perform various build, install, analysis, etc.
├── .air.toml            # Configuration file for Air live reload
├── .gitignore           # Git ignore file
├── go.mod               # Go module file
└── README.md            # Project README
```

### Folder Descriptions

- cmd/: Contains the main applications of the project. Each application has its own subdirectory.
- internal/: Contains private application and library code. This directory is not meant to be imported by other projects.
- pkg/: Contains library code that can be used by other projects. This is where you can place reusable code.
- web/: Contains web-related files such as HTML templates, CSS, JavaScript, and other static assets.
- configs/: Contains configuration files for different environments.
- infrastructure/: Contains infrastructure code and configuration files.
- scripts/: Contains scripts for various tasks like building, testing, and deploying the application.

## Dependencies

This project uses several external libraries to enhance the functionality of the Go web application:

[Fiber](https://github.com/gofiber/fiber): An Express-inspired web framework for Go.
[Echo](https://github.com/labstack/echo): A high performance, minimalist Go web framework.
[Gorm](https://gorm.io/): An ORM library for Go.
[Viper](https://github.com/spf13/viper): A complete configuration solution for Go applications.
[Air](https://github.com/cosmtrek/air): Live reload for Go applications.

You can find all dependencies listed in the `go.mod` file.

## Building and Running

To build and run the application, use the following commands:

```bash
go build -o bin/app ./cmd/app
./bin/app
```

Alternatively, you can run the application directly using:

```bash
go run ./cmd/app
```

#### Using Air for Live Reload

To use Air for live reloading during development, first install Air:

```bash
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh
```

Then, start Air:

```bash
air
```

This will automatically rebuild and reload the application whenever you make changes to the source code.

## Testing

This project will includes unit tests to ensure the correctness of the application. To run the tests, use the following command:

```bash
go test ./...
```

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request. Follow these steps to contribute:

Fork the repository
- Create a new branch (`git checkout -b feature/your-feature`)
- Commit your changes (`git commit -m 'Add some feature'`)
- Push to the branch (`git push origin feature/your-feature`)
- Open a pull request

## License

This project is licensed under the MIT License. See the `LICENSE` file for more details.