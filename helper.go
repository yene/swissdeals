package main

import "strings"

func removeNewline(s string) string {
	s = strings.Replace(s, "\n", " ", -1)
	s = strings.TrimSpace(s)
	return s
}

func removeShitFromURL(url string) string {
	que := strings.Index(url, "?")
	if que == -1 {
		return url
	}
	return url[0:que]
}

func cutoffafterprice(p string) string {
	pos := strings.LastIndex(p, ".–")
	if pos == -1 {
		return p
	}
	pos += len(".–") // The dash has length of 3
	return p[:pos]
}

func extractSrc(s string) string {
	start := strings.Index(s, "src=\"")
	if start == -1 {
		return ""
	}
	start += 5
	substring := s[start:len(s)]

	end := strings.Index(substring, "\"")
	if start == -1 {
		return ""
	}
	return substring[0:end]
}
