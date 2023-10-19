# Circle Cursor Movement for Teams Status

## Description
This script is specifically designed for individuals who need to prevent their status from appearing as "away" or "offline" on applications like Microsoft Teams during periods of inactivity. Written in Go, it utilizes the `robotgo` package to move the mouse cursor in a circular pattern for a user-defined duration. While the script is running, peripheral activity is simulated, thus avoiding automatic status changes based on inactivity.

## Installation

Before running the script, ensure that Go is installed on your machine. Confirm the installation by running:

```bash
go version
```

This script also requires the `robotgo` package. Install it using the following command:

```bash
go get -u github.com/go-vgo/robotgo
```

**Note:** The `robotgo` package has specific system dependencies. Please refer to the robotgo's documentation for detailed installation instructions.

## Usage

### Building yourself

Follow these steps to run the script:

1. Clone the repository or download the script from the repository.
2. Navigate to the directory containing the script.
3. Execute the script with the command:

```bash
go run main.go
```

4. Upon execution, the program requests the user to specify the duration (in minutes) for the cursor movement.
5. After inputting a valid duration, you can switch to any window or screen. The cursor will move in circles, preventing your status from changing to "away" or "offline." To stop the program, return to the terminal and press 'Enter'.

### Using executable

For your convenience there is an already prebuilt executable in the repo which you
can easily run as follows:

1. Ensure it has rights to be executed:
```bash
chmod +x cursik
```
2. Execute the program:
```bash
./cursik
```

## Contributions

Contributions, issues, and feature requests are welcome. For those looking to contribute, please check the issues page.

## License

This script is distributed under the MIT License. See `LICENSE` for more information.

## Contact

Your Name - nikita.vostrosablin@gmail.com
Project Link: https://github.com/pupersosition/cursik
