func repeatedCharacter(s string) byte {
    i, seen := 0, make(map[byte]int);

    for i < len(s) {
        c := s[i]
        _, found := seen[c]
        if found {
            return c
        } else {
            seen[c] = i
        }
        i++
    }

    return s[0]
}