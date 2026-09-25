func myAtoi(s string) int {
		i, l, r, negative, started := 0, len(s), 0, false, false
	toInt := func(c byte) int {
		d := "0"
		return int(c) - int(d[0])
	}

	for i < l {
		if !started && (s[i] == ' ') {
			i++
			continue
		}

		if !started && (s[i] == '-' || s[i] == '+') {
			negative = s[i] == '-'
			started = true
			i++
			continue
		}

		if s[i] < 48 || s[i] > 57 {
			break
		}

		intVal := toInt(s[i])

		if s[i] <= 57 && s[i] >= 48 {
			started = true
			r = r*10 + intVal
		}

		if !negative && r > math.MaxInt32 {
			return math.MaxInt32
		}

		if negative && -r < math.MinInt32 {
			return math.MinInt32
		}

		i++
	}

	if negative {
		r = -r
	}

	return r
}