package drivers

import "strings"

func GetApplicationName(processName, title string) string {
	if processName != "" {
		return strings.Title(strings.ReplaceAll(processName, ".exe", ""))
	}
	titles := strings.Split(title, " - ")
	if len(titles) > 1 {
		return titles[len(titles)-1]
	}
	return title
}
