# Mixue POS Application

A command-line Point of Sale (POS) application for Mixue drink orders, demonstrating key Go concepts including pointers, interfaces, goroutines, and concurrent processing. Features cart management, automatic invoice generation, transaction history, and real-time loading animations during checkout.

## Tech Stack

- **Go** 1.22+ - Programming language
- **bufio** - Buffered I/O for user input handling
- **text/tabwriter** - Aligned text/tabular output formatting
- **sync** - Synchronization primitives (WaitGroup, Mutex)
- **time** - Time functions for timestamps and delays
- **math/rand** - Random number generation for invoice IDs

## Prerequisites

Before running this application, make sure you have:

- **Go 1.22 or higher** installed on your system
- **Git** (for cloning the repository)

## How to Run

### 1. Clone the Repository

```bash
git clone https://github.com/ItsnaMaulanaHasan/koda-b4-golang-weekly.git
```

### 2. Initialize Go Module

```bash
go mod init golang-weekly
go mod tidy
```

### 3. Run the Application

```bash
go run main.go
```

## How to Contribute

### 1. Fork the Repository

Click the **Fork** button at the top right of this page.

### 2. Clone Your Fork

```bash
git clone https://github.com/yourusername/koda-b4-golang-weekly.git
```

### 3. Create a Feature Branch

```bash
git checkout -b feature/your-feature-name
```

### 4. Make Your Changes

- Follow Go best practices and conventions
- Keep code clean and well-documented
- Test your changes thoroughly

### 5. Commit Your Changes

```bash
git add .
git commit -m "Add: description of your changes"
```

**Commit Message Convention:**

- `Add:` for new features
- `Fix:` for bug fixes
- `Update:` for improvements
- `Docs:` for documentation changes

### 6. Push to Your Fork

```bash
git push origin feature/your-feature-name
```

### 7. Create a Pull Request

1. Go to the original repository
2. Click **New Pull Request**
3. Select your feature branch
4. Describe your changes clearly
5. Submit the PR

## Features

- **Menu Selection** - Browse 14 different Mixue drinks
- **Cart Management** - Add, edit, and remove items
- **Concurrent Checkout** - Parallel invoice generation and saving
- **Transaction History** - View all past orders
- **Goroutines** - Concurrent processing for better performance
- **Error Handling** - Panic recovery and input validation
