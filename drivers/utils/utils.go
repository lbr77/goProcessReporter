package utils

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

func ReplaceString(input string, replace []string, replaceTo []string) string {
	for i, r := range replace {
		input = strings.ReplaceAll(input, r, replaceTo[i])
	}
	return input
}

func HideString(input string, keywords []string) string {
	for _, r := range keywords {
		if strings.Contains(input, r) {
			return ""
		}
	}
	return input
}
