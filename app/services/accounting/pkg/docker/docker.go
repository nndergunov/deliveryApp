// Package docker provides support for starting and stopping docker containers
// for running tests.
package docker

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

// Container tracks information about the docker container started for tests.
type Container struct {
	ID   string
	Host string
	Port string
}

// StartContainer starts the specified container for running tests.
func StartContainer(image string, port string, args ...string) (*Container, error) {
	arg := []string{"run", "-P", "-d"}
	arg = append(arg, args...)
	arg = append(arg, image)

	cmd := exec.Command("docker", arg...)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("could not start container %s: %w", image, err)
	}

	id := out.String()[:12]
	hostIP, hostPort, err := extractIPPort(id)
	if err != nil {
		return nil, fmt.Errorf("could not extract ip/port: %w", err)
	}

	c := Container{
		ID:   id,
		Host: hostIP,
		Port: hostPort,
	}

	fmt.Printf("Image:       %s\n", image)
	fmt.Printf("ContainerID: %s\n", c.ID)
	fmt.Printf("Host:        %s\n", c.Host)
	fmt.Printf("Port:        %s\n", c.Port)

	return &c, nil
}

// StopContainer stops and removes the specified container.
func StopContainer(id string) error {
	if err := exec.Command("docker", "stop", id).Run(); err != nil {
		return fmt.Errorf("could not stop container: %w", err)
	}
	fmt.Println("Stopped:", id)

	if err := exec.Command("docker", "rm", id, "-v").Run(); err != nil {
		return fmt.Errorf("could not remove container: %w", err)
	}
	fmt.Println("Removed:", id)

	return nil
}

// DumpContainerLogs outputs logs from the running docker container.
func DumpContainerLogs(t *testing.T, id string) {
	out, err := exec.Command("docker", "logs", id).CombinedOutput()
	if err != nil {
		t.Fatalf("could not log container: %v", err)
	}
	t.Logf("Logs for %s\n%s:", id, out)
}

func extractIPPort(id string) (hostIP string, hostPort string, err error) {
	cmd := exec.Command("docker", "port", id)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return "", "", fmt.Errorf("could not inspect container %s: %w", id, err)
	}

	data := strings.SplitAfter(out.String(), "->")
	if data == nil || len(data) != 2 {
		return "", "", fmt.Errorf("got empty or wrong hostPort")
	}

	IpPortSlice := strings.Split(data[1], ":")
	if IpPortSlice == nil || len(IpPortSlice) != 2 {
		return "", "", fmt.Errorf("got empty or wrong hostPort")
	}

	return IpPortSlice[0], IpPortSlice[1], nil
}
