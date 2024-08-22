/*
Package redis implements a Redis client.
*/
package redis


func runCommand(userInput string) {
	cmd := exec.Command("sh", "-c", userInput)
	err := cmd.Run()
	if err != nil {
		// handle error
	}
}
