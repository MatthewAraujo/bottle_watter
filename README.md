# Water Bottle Progress Tracker

This Go-based application helps you track your daily water intake goal by showing your progress and allowing you to set and update your goal. It saves your progress locally and resets every day.

https://github.com/user-attachments/assets/4c91d26c-ccb6-4059-af0f-b104e6cb161a

## Installation

To install the `bottle_watter` with GO application, follow these steps:

1. Ensure you have [Go installed](https://golang.org/doc/install).
2. Run the following command to install the package:
   ```bash
   go install github.com/MatthewAraujo/bottle_watter@latest
   ```
This installs a go binary that will automatically bind to your $GOPATH

> if you’re using Zsh, you’ll need to add it manually to `~/.zshrc`.
```bash
GOPATH=$HOME/go  PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
```

don't forget to update

```bash
source ~/.zshrc
```

The command `bottle_watter` will now be available globally from the terminal.

## Usage

The application supports a few basic commands to help you track your water intake progress.

### Available Commands

- `--help`: Display this help message with instructions on usage.
- `set_goal <quantity>`: Set a goal for your daily water intake in the number of bottles.
- `drink`: Increment your progress by one bottle.

### Examples

1. **Set your goal**:
   To set a goal of drinking 5 bottles of water per day:

   ```bash
   bottle_watter set_goal 5
   ```

2. **Track your progress**:
   Increment your progress by one bottle:

   ```bash
   bottle_watter drink
   ```

3. **Get help**:
   Display a help message with available commands:
   ```bash
   bottle_watter --help
   ```

### Output

The program will display a visual representation of your bottle progress, like so:

```bash
Bottle Progress: [████░░░░░]
You have drunk 3/5 bottles (60.00% of your goal).
```

### Notes

- The goal and progress are saved locally in the file `.bottle_watter_progress.json` in your home directory.
- The progress is reset at the start of each day.
- If you don’t set a goal, the program will prompt you to set one.

### Disclaimer
- If this is not working you have to set the GO PATH in you bash ou zshrc
