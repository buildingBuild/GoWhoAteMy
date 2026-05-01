package monitoring

import "os/exec"

func sendNotification(title string, message string) {
	cmd := exec.Command(
		"osascript",
		"-e",
		`display notification "`+message+`" with title "`+title+`" sound name "Default"`,
	)

	err := cmd.Run()
	if err != nil {
		println("Error:", err)
	}
}
