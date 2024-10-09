# Password Generator

A simple command-line interface (CLI) password generator written in Golang. This application allows users to generate random passwords with customizable complexity and length.

## Features

- Generate passwords with various levels of complexity.
- Simple CLI for easy use.
- Configurable password length.

## Complexity Levels

The password complexity can be adjusted based on user needs. Here are the complexity levels:

1. **Numeric Only**: Passwords that contain numbers only.
2. **Uppercase Alphanumeric**: Passwords that contain uppercase letters and numbers.
3. **Lowercase Alphanumeric**: Passwords that contain lowercase letters and numbers.
4. **Mixed Alphanumeric**: Passwords that contain both uppercase and lowercase letters along with numbers.
5. **Mixed with Special Characters**: Passwords that contain a mix of letters, numbers, and special characters.

## Installation

To get started with the password generator:

1. Clone the repository:
    ```bash
    git clone https://github.com/urxfa/password-generator.git
    ```

2. Navigate to the project directory:
    ```bash
    cd password-generator
    ```

3. Build the application:
    ```bash
    go build
    ```

## Usage

To generate a password, run the compiled binary and specify your desired options. Example usage might look like this:

```bash
./password-generator 5 32 // It generates a password with a complexity of 5 and a length of 32
```
