package helpers

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/briandowns/spinner"
)

type Task struct {
	Title string
	Run   func() error
}

type SkipError struct {
	Reason string
}

func (e SkipError) Error() string { return e.Reason }

func Skip(reason string) error { return SkipError{Reason: reason} }

func isSkip(err error) (string, bool) {
	var se SkipError
	if errors.As(err, &se) {
		return se.Reason, true
	}
	return "", false
}

// RunTasks: satu baris spinner saat running, setelah selesai jadi ✔ Title
func RunTasks(tasks []Task) error {
	for _, t := range tasks {
		// Spinner dots (kamu bisa ganti charset lain)
		s := spinner.New(spinner.CharSets[14], 120*time.Millisecond) // 9 = dot-ish
		s.Color("yellow")
		s.Suffix = "  " + t.Title
		s.Writer = os.Stderr // biar gak campur sama stdout task (kalau ada)
		s.Start()

		err := t.Run()

		// stop + clear line spinner supaya gak ninggalin teks spinner
		s.Stop()
		// clear current line (ANSI): carriage return + clear line
		fmt.Fprint(os.Stderr, "\r\033[2K")

		if err == nil {
			fmt.Println("✔  " + t.Title)
			continue
		}

		if reason, ok := isSkip(err); ok {
			gray := "\033[90m"
			reset := "\033[0m"
			// kamu bisa pilih mau tampil seperti apa kalau skip
			// fmt.Println("↷  " + t.Title + " (" + reason + ")")
			// fmt.Printf("\033[90m↷\033[0m  %s \033[90m[skipped]\033[0m\n \033[90m→ (%s)\033[0m\n", t.Title, reason)
			fmt.Printf("%s↷%s  %s %s[skipped]%s\n", gray, reset, t.Title, gray, reset)
			if reason != "" {
				fmt.Printf("%s   → %s%s\n", gray, reason, reset)
			}

			continue
		}

		fmt.Println("✖  " + t.Title)
		return fmt.Errorf("%s: %w", t.Title, err)
	}
	return nil
}
