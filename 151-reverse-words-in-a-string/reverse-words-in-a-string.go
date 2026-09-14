func reverseWords(s string) string {
    	slow, fast, end := 0, 0, len(s)

	words := make([]string, 0)
	for end > fast {
		if s[fast] == ' ' {
			words = append(words, s[slow:fast])
			slow = fast + 1
		}
		fast++
	}

	words = append(words, s[slow:fast])
	filtered := make([]string, 0)

	for _, word := range words {
		if word != "" {
			fmt.Println("Word is : " + word)
			filtered = append(filtered, word)
		}
	}

	slow, fast = 0, len(filtered)-1

	for slow < fast {
		filtered[slow], filtered[fast] = filtered[fast], filtered[slow]
		slow++
		fast--
	}

	return strings.Join(filtered, " ")
}