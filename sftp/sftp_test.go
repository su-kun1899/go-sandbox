package sftp_test

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"testing"
)

func Test(t *testing.T) {
	host := os.Getenv("SFTP_HOST")
	port := os.Getenv("SFTP_PORT")
	username := os.Getenv("SFTP_USERNAME")

	cmd := exec.Command("sftp",
		"-o PubkeyAuthentication=no",
		fmt.Sprintf("-P %s", port),
		fmt.Sprintf("%s@%s", username, host))

	stdin, _ := cmd.StdinPipe()
	_, err := io.WriteString(stdin,
		`
		ls
		cd archive
		ls
		`)

	if err != nil {
		log.Fatal(err)
	}
	err = stdin.Close()
	if err != nil {
		log.Fatal(err)
	}

	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("結果: %s", out)
}
