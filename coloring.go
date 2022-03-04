package log4fox

import (
	"fmt"
)

// color values
const (
	FONT_BLACK       = "\033[0;30m"
	FONT_RED         = "\033[0;31m"
	FONT_GREEN       = "\033[0;32m"
	FONT_YELLOW      = "\033[0;33m"
	FONT_BLUE        = "\033[0;34m"
	FONT_PURPLE      = "\033[0;35m"
	FONT_CYAN        = "\033[0;36m"
	FONT_WHITE       = "\033[0;37m"
	FONT_BLACK_BOLD  = "\033[1;30m"
	FONT_RED_BOLD    = "\033[1;31m"
	FONT_GREEN_BOLD  = "\033[1;32m"
	FONT_YELLOW_BOLD = "\033[1;33m"
	FONT_BLUE_BOLD   = "\033[1;34m"
	FONT_PURPLE_BOLD = "\033[1;35m"
	FONT_CYAN_BOLD   = "\033[1;36m"
	FONT_WHITE_BOLD  = "\033[1;37m"
	FONT_BLACK_UL    = "\033[4;30m"
	FONT_RED_UL      = "\033[4;31m"
	FONT_GREEN_UL    = "\033[4;32m"
	FONT_YELLOW_UL   = "\033[4;33m"
	FONT_BLUE_UL     = "\033[4;34m"
	FONT_PURPLE_UL   = "\033[4;35m"
	FONT_CYAN_UL     = "\033[4;36m"
	FONT_WHITE_UL    = "\033[4;37m"

	BG_BLACK     = "\033[40m"
	BG_RED       = "\033[41m"
	BG_GREEN     = "\033[42m"
	BG_YELLOW    = "\033[43m"
	BG_BLUE      = "\033[44m"
	BG_PURPLE    = "\033[45m"
	BG_CYAN      = "\033[46m"
	BG_WHITE     = "\033[47m"
	BG_BLACK_HI  = "\033[0;100m"
	BG_RED_HI    = "\033[0;101m"
	BG_GREEN_HI  = "\033[0;102m"
	BG_YELLOW_HI = "\033[0;103m"
	BG_BLUE_HI   = "\033[0;104m"
	BG_PURPLE_HI = "\033[0;105m"
	BG_CYAN_HI   = "\033[0;106m"
	BG_WHITE_HI  = "\033[0;107m"

	RESET = "\033[m"
)

func Color(c string, s string) string {
	return c + s + RESET
}

func Colorf(c string, s string, k ...interface{}) string {
	return c + fmt.Sprintf(s, k...) + RESET
}

func Colorln(c string, s ...interface{}) string {
	return c + fmt.Sprintln(s...) + RESET
}
