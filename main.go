package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/fatih/color"
)

type Progress struct {
	Goal       int    `json:"goal"`
	Current    int    `json:"current"`
	LastUpdate string `json:"last_update"`
}

func getProgressFilePath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting user directory:", err)
		os.Exit(1)
	}
	return filepath.Join(homeDir, ".bottle_watter_progress.json")
}

func loadProgress() Progress {
	filePath := getProgressFilePath()
	file, err := os.Open(filePath)
	if os.IsNotExist(err) {
		return Progress{}
	} else if err != nil {
		fmt.Println("Error opening progress file:", err)
		os.Exit(1)
	}
	defer file.Close()

	var progress Progress
	if err := json.NewDecoder(file).Decode(&progress); err != nil {
		fmt.Println("Error reading progress:", err)
		os.Exit(1)
	}
	return progress
}

func saveProgress(progress Progress) {
	filePath := getProgressFilePath()
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error saving progress:", err)
		os.Exit(1)
	}
	defer file.Close()

	if err := json.NewEncoder(file).Encode(progress); err != nil {
		fmt.Println("Error writing progress:", err)
		os.Exit(1)
	}
}

func resetBottleForTheDay(progress *Progress) {
	today := time.Now().Format("2006-01-02")
	if progress.LastUpdate != today {
		progress.Current = 0
		progress.LastUpdate = today
		saveProgress(*progress)
		fmt.Println("Progress reset for the new day!")
	}
}

func displayBottle(goal, current int) {
	fullColor := color.New(color.FgBlue).SprintFunc()
	emptyColor := color.New(color.FgWhite).SprintFunc()

	bottle := ""
	for i := 1; i <= goal; i++ {
		if i <= current {
			bottle += fullColor("█")
		} else {
			bottle += emptyColor("░")
		}
	}

	fmt.Printf("Bottle Progress: [%s]\n", bottle)
}

func displayHelp() {
	fmt.Println("Usage: bottle_watter [command] [options]")
	fmt.Println("\nCommands:")
	fmt.Println("  --help                     Show this help message")
	fmt.Println("  set_goal <quantity>        Set a goal for daily watter intake (in number of bottles)")
	fmt.Println("  drink                      Increment the progress by one bottle")
	fmt.Println("\nExamples:")
	fmt.Println("  bottle_watter set_goal 5    Set a goal of drinking 5 bottles of watter per day")
	fmt.Println("  bottle_watter drink         Increment the progress by one bottle")
	fmt.Println("\nNote: The goal and progress are saved locally and reset at the start of each day.")
}

func main() {
	progress := loadProgress()

	resetBottleForTheDay(&progress)

	if len(os.Args) == 1 {
		displayHelp()
		return
	}

	if len(os.Args) > 1 && os.Args[1] == "--help" {
		displayHelp()
		return
	}

	if len(os.Args) > 1 && os.Args[1] == "set_goal" {
		if len(os.Args) != 3 {
			fmt.Println("Usage: bottle_watter set_goal <quantity>")
			os.Exit(1)
		}
		goal, err := strconv.Atoi(os.Args[2])
		if err != nil || goal <= 0 {
			fmt.Println("Invalid goal. Please enter a number greater than 0.")
			os.Exit(1)
		}
		progress.Goal = goal
		progress.Current = 0
		progress.LastUpdate = time.Now().Format("2006-01-02")
		saveProgress(progress)
		fmt.Printf("Goal set: %d bottles\n", goal)
		displayBottle(progress.Goal, progress.Current)
		return
	}

	if progress.Goal == 0 {
		fmt.Println("No goal set. Use 'bottle_watter set_goal <quantity>' to set one.")
		return
	}

	if progress.Current >= progress.Goal {
		fmt.Println("You have already reached your goal for today! Congratulations!")
		displayBottle(progress.Goal, progress.Current)
		return
	}

	if len(os.Args) > 1 && os.Args[1] == "drink" {
		progress.Current++
		saveProgress(progress)
		displayBottle(progress.Goal, progress.Current)
	}

	if progress.Current >= progress.Goal {
		fmt.Println("Congratulations, you reached your watter intake goal!")
	}
}
