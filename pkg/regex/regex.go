package regex

import "regexp"

// NormalMatch matchs just a normal match
func NormalMatch(regx string, arg string) bool {
	matched, err := regexp.MatchString(regx, arg)
	if err != nil {
		return false
	}
	return matched
}

// Match matches a regex from begin of a string to the end
func Match(regx string, arg string) bool {
	matched, err := regexp.MatchString("^("+regx+")$", arg)
	if err != nil {
		return false
	}
	return matched
}

// FindMatch finds a specific match and returns a specific part of the match
func FindMatch(toMatch string, regx string, selecter int) string {
	re := regexp.MustCompile(regx)
	out := re.FindStringSubmatch(string(toMatch))
	if len(out) > selecter {
		return out[selecter]
	}
	return ""
}

// FindAllMatch find all matches
func FindAllMatch(toMatch string, regx string, selecter int) []string {
	re := regexp.MustCompile(regx)
	out := re.FindAllStringSubmatch(string(toMatch), -1)
	toReturn := []string{}
	for _, value := range out {
		toReturn = append(toReturn, value[selecter])
	}
	return toReturn
}

// Replace with a regex
func Replace(toReplace string, Replaceval string, regx string) string {
	re := regexp.MustCompile(regx)
	return re.ReplaceAllString(toReplace, Replaceval)
}
