# GitHub Trending CLI

<div align="center">

[![Go Version](https://img.shields.io/badge/Go-1.25+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Build Status](https://img.shields.io/badge/build-passing-brightgreen.svg)](https://github.com/letsmakecakes/github-trending-cli)

A command-line interface (CLI) tool that fetches and displays trending repositories from GitHub. Built to help developers discover popular projects and stay updated with the latest trends in the open source community.

[Features](#features) • [Installation](#installation) • [Usage](#usage) • [Development](#development) • [Contributing](#contributing)

</div>

---

## 📋 Overview

GitHub Trending CLI is a lightweight, fast, and easy-to-use command-line tool that interacts with the GitHub API to retrieve and display trending repositories. Whether you're looking for daily hot repositories or exploring monthly trends, this tool provides a clean, formatted output right in your terminal.

### Why GitHub Trending CLI?

- 🚀 **Fast & Lightweight**: Quick execution with minimal dependencies
- 🎯 **Focused Output**: Get only what you need without browser overhead
- ⚙️ **Customizable**: Filter by time range and limit results
- 🎨 **Clean Display**: Beautifully formatted output for easy reading
- 🔄 **CI/CD Friendly**: Perfect for automation and scripts

## ✨ Features

- **Time Range Filtering**: View trending repositories by day, week, month, or year
- **Customizable Limits**: Display between 1-100 repositories
- **Sorted Results**: Repositories sorted by star count
- **Rich Information**: View repository name, description, stars, language, and more
- **Robust Error Handling**: Graceful handling of API errors and invalid input
- **No Authentication Required**: Works out of the box for public repositories
- **Cross-Platform**: Works on Linux, macOS, and Windows

## 🚀 Installation

### Prerequisites

- Go 1.25 or higher

### From Source

1. Clone the repository:
```bash
git clone https://github.com/letsmakecakes/github-trending-cli.git
cd github-trending-cli
```

2. Build the binary:
```bash
make build
```

3. The binary will be available in the `bin/` directory:
```bash
./bin/trending-repos --help
```

### Install Globally

To install the tool globally on your system:

```bash
make install
```

Or manually:

```bash
go install github.com/letsmakecakes/github-trending-cli/cmd/trending-repos@latest
```

After installation, you can run the tool from anywhere:

```bash
trending-repos --duration day --limit 10
```

### Build for Multiple Platforms

To build binaries for all supported platforms:

```bash
make build-all
```

This creates binaries for:
- Linux (amd64)
- macOS (amd64, arm64)
- Windows (amd64)

## 📖 Usage

### Basic Usage

```bash
trending-repos
```

This displays the top 10 trending repositories from the past week (default behavior).

### Command-Line Options

| Flag         | Description                          | Default | Valid Values                   |
| ------------ | ------------------------------------ | ------- | ------------------------------ |
| `--duration` | Time range for trending repositories | `week`  | `day`, `week`, `month`, `year` |
| `--limit`    | Number of repositories to display    | `10`    | `1-100`                        |

### Examples

#### Show today's trending repositories
```bash
trending-repos --duration day
```

#### Display top 20 trending repositories from the past month
```bash
trending-repos --duration month --limit 20
```

#### View top 5 trending repositories from the past year
```bash
trending-repos --duration year --limit 5
```

#### Show 50 trending repositories from the past week
```bash
trending-repos --limit 50
```

### Sample Output

```
================================================================================
                        GitHub Trending Repositories
                            Duration: week | Limit: 10
================================================================================

1. awesome-project
   ⭐ Stars: 15,234 | Language: Python
   📝 A comprehensive collection of awesome resources for developers
   🔗 https://github.com/user/awesome-project

2. next-gen-framework
   ⭐ Stars: 12,891 | Language: TypeScript
   📝 Modern framework for building scalable applications
   🔗 https://github.com/user/next-gen-framework

[... more repositories ...]
```

## 🛠️ Development

### Project Structure

```
github-trending-cli/
├── cmd/
│   └── trending-repos/      # Application entry point
│       └── main.go
├── internal/                # Private application code
│   ├── api/                 # GitHub API client
│   │   ├── github.go
│   │   └── github_test.go
│   ├── config/              # Configuration management
│   │   └── config.go
│   ├── display/             # Output formatting
│   │   ├── formatter.go
│   │   └── formatter_test.go
│   └── models/              # Data models
│       └── repository.go
├── pkg/                     # Public library code
│   └── cli/                 # CLI interface
│       ├── cli.go
│       └── cli_test.go
├── test/                    # Integration tests
│   └── integration/
│       └── integration_test.go
├── bin/                     # Compiled binaries
├── go.mod                   # Go module definition
├── Makefile                 # Build automation
└── README.md               # This file
```

### Setup Development Environment

1. Clone the repository:
```bash
git clone https://github.com/letsmakecakes/github-trending-cli.git
cd github-trending-cli
```

2. Install dependencies:
```bash
go mod download
```

3. Run tests:
```bash
make test
```

### Available Make Commands

```bash
make build              # Build the application
make build-all          # Build for all platforms
make test               # Run all tests with coverage
make test-unit          # Run unit tests only
make test-integration   # Run integration tests only
make clean              # Clean build artifacts
make install            # Install the binary globally
make run                # Build and run the application
make lint               # Run linter
make coverage           # Generate coverage report
make help               # Show all available commands
```

### Running Tests

#### All Tests
```bash
make test
```

#### Unit Tests Only
```bash
make test-unit
```

#### Integration Tests Only
```bash
make test-integration
```

#### Coverage Report
```bash
make coverage
```

### Code Quality

Run the linter to ensure code quality:
```bash
make lint
```

## 🏗️ Architecture

The application follows a clean architecture pattern with clear separation of concerns:

- **cmd/**: Application entry points
- **internal/**: Private application and business logic
  - **api/**: External API integrations (GitHub API)
  - **config/**: Configuration parsing and validation
  - **display/**: Output formatting and presentation
  - **models/**: Data structures and domain models
- **pkg/**: Public reusable packages
  - **cli/**: Command-line interface logic
- **test/**: Integration and end-to-end tests

## 🤝 Contributing

Contributions are welcome! This project is designed to help developers practice building CLI applications and working with APIs.

### How to Contribute

1. **Fork the repository**
2. **Create a feature branch**: `git checkout -b feature/amazing-feature`
3. **Make your changes**
4. **Write/update tests**: Ensure your code is well tested
5. **Run tests**: `make test`
6. **Commit your changes**: `git commit -m 'Add some amazing feature'`
7. **Push to the branch**: `git push origin feature/amazing-feature`
8. **Open a Pull Request**

### Development Guidelines

- Write clear, commented code
- Follow Go best practices and idioms
- Add tests for new features
- Update documentation as needed
- Ensure all tests pass before submitting PR
- Keep commits atomic and well-described

### Ideas for Contributions

- Add language filtering options
- Implement caching to reduce API calls
- Add export functionality (JSON, CSV)
- Create colorized output
- Add interactive mode
- Support for GitHub authentication (higher rate limits)
- Add repository topic filtering
- Implement pagination for large result sets

## 🐛 Troubleshooting

### Common Issues

#### API Rate Limiting
GitHub's API has rate limits for unauthenticated requests (60 requests per hour). If you hit the limit:
- Wait for the rate limit to reset
- Consider implementing authentication for higher limits (5000 requests/hour)

#### Network Issues
If you encounter network errors:
```bash
# Check your internet connection
ping github.com

# Verify GitHub API is accessible
curl https://api.github.com
```

#### Build Errors
If you encounter build errors:
```bash
# Clean and rebuild
make clean
go mod tidy
make build
```

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- [GitHub REST API](https://docs.github.com/en/rest) for providing the data
- The Go community for excellent tools and libraries
- All contributors who help improve this project

## 📬 Contact & Support

- **Issues**: [GitHub Issues](https://github.com/letsmakecakes/github-trending-cli/issues)
- **Discussions**: [GitHub Discussions](https://github.com/letsmakecakes/github-trending-cli/discussions)

## 🌟 Star History

If you find this project useful, please consider giving it a star! It helps others discover the project and motivates continued development.

---

<div align="center">

**Built with ❤️ by developers, for developers**

[⬆ Back to Top](#github-trending-cli)

</div>
