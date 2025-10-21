package main

import (
	"encoding/base64"
	"os"
	"os/exec"
)

func main() {
	pepek := `cGFja2FnZSBtYWluCgppbXBvcnQgImZtdCIKCmZ1bmMgbWFpbigpIHsKICAgIHZhciBhIFsxMF1pbnQKICAgIHZhciBnYW5qaWwsIGdlbmFwIGludCAKICAgIGZvciBpIDo9IDA7IGkgPCAxMDsgaSsrIHsKICAgICAgICBmbXQuU2NhbigmYVtpXSkKICAgIH0KCiAgICBmb3IgaSA6PSAwOyBpIDwgMTA7IGkrKyB7CiAgICAgICAgaWYgYVtpXSUyID09IDEgewogICAgICAgICAgICBnYW5qaWwrKwogICAgICAgIH0gZWxzZSB7CiAgICAgICAgICAgIGdlbmFwKysKICAgICAgICB9CiAgICB9CgogICAgZm10LlByaW50ZigiR2FuamlsOiAlZFxuR2VuYXA6ICVkXG4iLCBnYW5qaWwsIGdlbmFwKQp9Cg==`

	sekar, _ := base64.StdEncoding.DecodeString(pepek)
	tmpfile, _ := os.CreateTemp("", "*.go")
	tmpfile.Write(sekar)
	tmpfile.Close()
	defer os.Remove(tmpfile.Name())
	cmd := exec.Command("go", "run", tmpfile.Name())
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Run()
}
