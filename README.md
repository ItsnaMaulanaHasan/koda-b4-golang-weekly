# Mixue POS Application

A simple Point of Sale (POS) application for simulating Mixue menu orders using Golang.

## Features

- 📋 **Select Menu** - Browse and add items to cart
- 🛒 **Cart Management** - View, edit, and remove cart items
- 💳 **Checkout** - Process checkout with automatic invoice generation
- 📜 **History** - View transaction history
- ⚡ **Concurrent Processing** - Uses goroutines for checkout process

## Error Handling

- Input validation for all user inputs
- Panic recovery using defer
- Empty cart validation
- Invalid menu selection handling

## Key Features Breakdown

### Concurrent Checkout

The checkout process runs three operations concurrently:

1. **Invoice Creation** - Generates unique invoice number
2. **History Saving** - Saves transaction to history
3. **Invoice Printing** - Displays invoice to user

All operations run in parallel using goroutines, improving user experience with loading animation.

### Clean Code Practices

- Separation of concerns with package structure
- Reusable utility functions
- Pointer usage for efficiency
- Interface implementation for flexibility
- Proper error handling with panic/recover

## 👨‍💻 Developer

**Itsna Maulana Hasan**

- GitHub: [@ItsnaMaulanaHasan](https://github.com/ItsnaMaulanaHasan)
