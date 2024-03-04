package logger

import (
	"fmt"
	"io"
	"os"

	"github.com/fatih/color"
)

type Logger struct {
	w io.Writer
}

func New(w io.Writer) Logger {
	return Logger{w: w}
}

func (l Logger) Print(message string) {
	_, _ = fmt.Fprintln(l.w, message)
}

func grayString(format string, a ...any) string {
	if color.NoColor {
		return fmt.Sprintf(format, a...)
	} else {
		return fmt.Sprintf("\x1B[38;5;239m"+format+"\x1B[0m", a...)
	}
}

func (l Logger) log(level string, message string) {
	_, _ = fmt.Fprintf(l.w, "%s%s%s\n", level, grayString(" ┃ "), message)
}

func (l Logger) Trace(message string) {
	l.log(color.HiGreenString("TRACE"), message)
}

func (l Logger) Debug(message string) {
	l.log(color.GreenString("DEBUG"), message)
}

func (l Logger) Info(message string) {
	l.log(color.HiWhiteString("INFO "), message)
}

func (l Logger) Warning(message string) {
	l.log(color.YellowString("WARN "), message)
}

func (l Logger) Error(message string) {
	l.log(color.RedString("ERROR"), message)
}

func (l Logger) Fatal(message string) {
	l.log(color.HiRedString("FATAL"), message)
	os.Exit(1)
}
