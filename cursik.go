package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
	"math"

	"github.com/go-vgo/robotgo"
)

func moveCursorInCircle(duration time.Duration) {
	startTime := time.Now()

	// The radius of the circle along which the cursor moves.
	radius := 100.0

	// Center position of the current screen. Change this if you want the circle centered differently.
	cx, cy := robotgo.GetScreenSize()
	cx, cy = cx/2, cy/2

	// Loop until the specified duration has passed.
	for time.Since(startTime) < duration {
		// Calculate the amount of time passed and use it to determine the angle.
		elapsed := time.Since(startTime).Seconds()
		angle := 2 * math.Pi * elapsed // Complete circle for each second

		// Calculate the new position of the cursor.
		x := int(radius*math.Cos(angle)) + cx
		y := int(radius*math.Sin(angle)) + cy

		// Move the cursor to the new position.
		robotgo.MoveMouse(x, y)

		// Sleep for a short duration to make the movement smoother.
		time.Sleep(10 * time.Millisecond)
	}
}

func main() {
	fmt.Println("Enter the number of minutes to move the cursor in circles:")

	// Read input from the user.
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		return
	}

	// Parse the input, assuming it's in minutes.
	var minutes float64
	_, err = fmt.Sscanf(input, "%f", &minutes)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Invalid input. Please enter a number:", err)
		return
	}

	// Convert minutes to duration.
	duration := time.Duration(minutes) * time.Minute

	// Inform the user to stop the program by pressing Enter.
	fmt.Println("The cursor will move in circles. To stop, switch back to this window and press 'Enter'.")

	// Start a goroutine to move the cursor in circles.
	go moveCursorInCircle(duration)

	// Wait for the user to hit 'Enter' to stop the program.
	_, _ = reader.ReadString('\n')
	fmt.Println("Exiting program.")
}

