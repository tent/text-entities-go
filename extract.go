package text

type Match struct {
	Text    string
	Indices [2]int
}

func ExtractHashtags(s string) []string {
	hashtags := ExtractHashtagMatches(s)
	res := make([]string, len(hashtags))
	for i, tag := range hashtags {
		res[i] = tag.Text
	}
	return res
}

func ExtractHashtagMatches(s string) []Match {
	matches := hashtagPattern.FindAllStringSubmatchIndex(s, -1)
	res := make([]Match, 0, len(matches))
	for _, m := range matches {
		if invalidHashtagEnd.MatchString(s[m[3]:]) {
			continue
		}
		res = append(res, Match{s[m[4]:m[5]], [2]int{m[2], m[3]}})
	}
	return res
}

func ExtractURLs(s string) []string {
	urls := ExtractURLMatches(s)
	res := make([]string, len(urls))
	for i, url := range urls {
		res[i] = url.Text
	}
	return res
}

func ExtractURLMatches(s string) []Match {
	matches := urlPattern.FindAllStringSubmatchIndex(s, -1)
	res := make([]Match, 0, len(matches))
	for _, m := range matches {
		var before, protocol, path string
		if m[2] != -1 {
			before = s[m[2]:m[3]]
		}
		if m[6] != -1 {
			protocol = s[m[6]:m[7]]
		}
		domain := s[m[8]:m[9]]
		afterDomain := s[m[9]:]
		if m[10] != -1 {
			path = s[m[10]:m[11]]
		}

		// We don't have lookahead, so manually check that the tld doesn't end in the middle of a word
		if afterDomain != "" && ((afterDomain[0] >= '0' && afterDomain[0] <= '9') ||
			(afterDomain[0] >= 'A' && afterDomain[0] <= 'Z') ||
			(afterDomain[0] >= 'a' && afterDomain[0] <= 'z')) {
			continue
		}

		if protocol == "" {
			if invalidBeforeDomain.MatchString(before) {
				continue
			}
			ad := asciiDomain.FindString(domain)
			if len(domain) != len(ad) || invalidWithoutPath.MatchString(ad) && path == "" {
				continue
			}
		}
		res = append(res, Match{s[m[4]:m[5]], [2]int{m[4], m[5]}})
	}
	return res
}
