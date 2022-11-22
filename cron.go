package main

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"

	"github.com/robfig/cron"
	"golang.org/x/text/encoding/simplifiedchinese"
)

type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)

func ConvertByte2String(byte []byte, charset Charset) string {

	var str string
	switch charset {
	case GB18030:
		decodeBytes, _ := simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}

	return str
}

func Base64Encode(input string) string {
	return base64.StdEncoding.EncodeToString([]byte(input))
}

func Base64Decode(input string) string {
	str, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return input
	}

	return string(str)
}

func RunCmd(path string, cmd string) (msg string, err error) {
	if runtime.GOOS == "windows" {
		return RunProc(path, "cmd", "/c", cmd)
	} else if runtime.GOOS == "linux" {
		return RunProc(path, "bash", cmd)
	} else if runtime.GOOS == "darwin" {
		return RunProc(path, "zsh", cmd)
	} else {
		return RunProc(path, cmd)
	}

}

func RunProc(path, name string, arg ...string) (msg string, err error) {
	cmd := exec.Command(name, arg...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	cmd.Dir = path
	err = cmd.Run()
	if err != nil {
		msg_str := ConvertByte2String(stderr.Bytes(), GB18030)
		msg = fmt.Sprint(err) + ": " + msg_str
		err = errors.New(msg)
	}
	out_str := ConvertByte2String(out.Bytes(), GB18030)
	fmt.Println("run:", cmd.Args, "\n", out_str)
	return out_str, err
}

func main() {
	if len(os.Args) != 2 {
		return
	}
	c := cron.New()
	zp := regexp.MustCompile(`(?mi)(^[^#\r\n]+)`)
	lines := zp.FindAllString(Base64Decode(os.Args[1]), -1)
	for _, line := range lines {
		cmd := strings.Split(line, "!!")
		if len(cmd) != 2 {
			return
		}
		fmt.Println(cmd[0], cmd[1])
		c.AddFunc(cmd[0], func() {
			rst, err := RunCmd(".", cmd[1])
			if len(rst) > 0 {
				fmt.Println(rst)
			}
			if err != nil {
				fmt.Println(err)
			}
		})
	}
	c.Start()
	select {}
}
