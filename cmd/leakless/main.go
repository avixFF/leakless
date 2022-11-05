package main

import (
	"encoding/json"
	"log"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"

	"github.com/ysmood/leakless/pkg/shared"
)

func main() {
	ignoreSignals()

	// Strip binary path from the argument list
	os.Args = os.Args[1:]

	options := shared.ParseOptionsString(strings.Join(os.Args, " "))

	// Iterate over all arguments that start with "--", process them
	// and remove from the os.Args list
	for len(os.Args) > 0 && strings.HasPrefix(os.Args[0], "--") {
		if os.Args[0] == "--version" {
			_, _ = os.Stdout.WriteString(shared.Version + "\n")
			return
		}

		os.Args = os.Args[1:]
	}

	if len(os.Args) < 3 {
		panic("wrong args, usage: leakless [--flags...] <uid> <tcp-addr> <cmd> [args...]")
	}

	uid := os.Args[0]
	addr := os.Args[1]

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatalln(err)
	}

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = osSetupCmd(cmd, options)
	if err != nil {
		_ = send(conn, 0, err.Error())
		log.Fatalln(err)
	}

	err = cmd.Start()
	if err != nil {
		_ = send(conn, 0, err.Error())
		log.Fatalln(err)
	}

	err = send(conn, cmd.Process.Pid, "")
	if err != nil {
		kill(cmd.Process)
		log.Fatalln(err)
	}

	go guard(conn, uid, cmd.Process)

	err = cmd.Wait()
	if err != nil {
		exitErr, ok := err.(*exec.ExitError)
		if ok {
			os.Exit(exitErr.ExitCode())
			return
		}
		err = send(conn, 0, err.Error())
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func guard(conn net.Conn, uid string, p *os.Process) {
	defer kill(p)

	dec := json.NewDecoder(conn)

	var msg shared.Message
	err := dec.Decode(&msg)
	if err != nil {
		return
	}
	if msg.UID != uid {
		return
	}

	_ = dec.Decode(&msg)
}

func send(conn net.Conn, pid int, errMessage string) error {
	enc := json.NewEncoder(conn)
	return enc.Encode(shared.Message{PID: pid, Error: errMessage})
}

// OS may send signals to interrupt processes in the same group, as a guard process leakless shouldn't be stopped by them.
func ignoreSignals() {
	signal.Ignore(syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
}
