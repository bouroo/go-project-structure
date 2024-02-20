package helper

import (
	"bytes"
	"fmt"
	"os"
	"time"
)

type LogWriter struct {
	DirPath      string
	LokiURL      string
	ConsolePrint bool
}

func (l LogWriter) Write(p []byte) (n int, err error) {
	buffer := bytes.NewBuffer(p)

	if l.ConsolePrint {
		go os.Stdout.Write(buffer.Bytes())
	}

	if len(l.DirPath) != 0 {
		err = os.MkdirAll(l.DirPath, 0755)
		if err != nil {
			return n, err
		}
		filePath := fmt.Sprintf("%s/%s", l.DirPath, time.Now().Format("20060102")+".log")
		f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return n, err
		}
		defer f.Close()
		n, err = f.Write(buffer.Bytes())
		if err != nil {
			return n, err
		}
	}

	// if len(l.LokiURL) != 0 {
	// 	go l.sendToLoki(buffer.Bytes())
	// }

	return n, err
}
