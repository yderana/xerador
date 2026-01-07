package helpers

func Red(s string) string    { return "\x1b[31m" + s + "\x1b[0m" }
func Green(s string) string  { return "\x1b[32m" + s + "\x1b[0m" }
func Yellow(s string) string { return "\x1b[33m" + s + "\x1b[0m" }
func Gray(s string) string   { return "\033[90m" + s + "\033[0m" }
func Cyan(s string) string   { return "\033[36m" + s + "\033[0m" }
