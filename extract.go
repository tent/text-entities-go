package text

import (
	"regexp"
	"strings"
)

func init() {
	charGroups := make(map[string][]string)
	regexen := make(map[string]string)

	add := func(p string, r ...rune) {
		if len(r) == 1 {
			charGroups[p] = append(charGroups[p], string(r[0]))
		} else {
			charGroups[p] = append(charGroups[p], string(r[0])+"-"+string(r[1]))
		}
	}

	interpolatePattern := regexp.MustCompile(`#\{\w+\}`)
	interp := func(s string) string {
		replace := func(g string) string {
			g = g[2 : len(g)-1]
			if regex, ok := regexen[g]; ok {
				return regex
			}
			group, ok := charGroups[g]
			if !ok {
				panic(s + ": unknown group or regexen " + g)
			}
			return strings.Join(group, "")
		}
		return interpolatePattern.ReplaceAllStringFunc(s, replace)
	}

	charGroups["spaces"] = []string{
		"\u0020", // White_Space # Zs SPACE
		"\u0085", // White_Space # Cc <control-0085>
		"\u00A0", // White_Space # Zs NO-BREAK SPACE
		"\u1680", // White_Space # Zs OGHAM SPACE MARK
		"\u180E", // White_Space # Zs MONGOLIAN VOWEL SEPARATOR
		"\u2028", // White_Space # Zl LINE SEPARATOR
		"\u2029", // White_Space # Zp PARAGRAPH SEPARATOR
		"\u202F", // White_Space # Zs NARROW NO-BREAK SPACE
		"\u205F", // White_Space # Zs MEDIUM MATHEMATICAL SPACE
		"\u3000", // White_Space # Zs IDEOGRAPHIC SPACE
	}
	add("spaces", '\u0009', '\u000D') // White_Space # Cc [5] <control-0009>..<control-000D>
	add("spaces", '\u2000', '\u200A') // White_Space # Zs [11] EN QUAD..HAIR SPACE

	charGroups["invalid"] = []string{
		"\uFFFE",
		"\uFEFF", // BOM
		"\uFFFF", // Special
	}
	add("invalid", '\u202A', '\u202E') // Directional change

	add("nonLatinHashtag", '\u0400', '\u04ff') // Cyrillic
	add("nonLatinHashtag", '\u0500', '\u0527') // Cyrillic Supplement
	add("nonLatinHashtag", '\u2de0', '\u2dff') // Cyrillic Extended A
	add("nonLatinHashtag", '\ua640', '\ua69f') // Cyrillic Extended B
	add("nonLatinHashtag", '\u0591', '\u05bf') // Hebrew
	add("nonLatinHashtag", '\u05c1', '\u05c2')
	add("nonLatinHashtag", '\u05c4', '\u05c5')
	add("nonLatinHashtag", '\u05c7')
	add("nonLatinHashtag", '\u05d0', '\u05ea')
	add("nonLatinHashtag", '\u05f0', '\u05f4')
	add("nonLatinHashtag", '\ufb12', '\ufb28') // Hebrew Presentation Forms
	add("nonLatinHashtag", '\ufb2a', '\ufb36')
	add("nonLatinHashtag", '\ufb38', '\ufb3c')
	add("nonLatinHashtag", '\ufb3e')
	add("nonLatinHashtag", '\ufb40', '\ufb41')
	add("nonLatinHashtag", '\ufb43', '\ufb44')
	add("nonLatinHashtag", '\ufb46', '\ufb4f')
	add("nonLatinHashtag", '\u0610', '\u061a') // Arabic
	add("nonLatinHashtag", '\u0620', '\u065f')
	add("nonLatinHashtag", '\u066e', '\u06d3')
	add("nonLatinHashtag", '\u06d5', '\u06dc')
	add("nonLatinHashtag", '\u06de', '\u06e8')
	add("nonLatinHashtag", '\u06ea', '\u06ef')
	add("nonLatinHashtag", '\u06fa', '\u06fc')
	add("nonLatinHashtag", '\u06ff')
	add("nonLatinHashtag", '\u0750', '\u077f') // Arabic Supplement
	add("nonLatinHashtag", '\u08a0')           // Arabic Extended A
	add("nonLatinHashtag", '\u08a2', '\u08ac')
	add("nonLatinHashtag", '\u08e4', '\u08fe')
	add("nonLatinHashtag", '\ufb50', '\ufbb1') // Arabic Pres. Forms A
	add("nonLatinHashtag", '\ufbd3', '\ufd3d')
	add("nonLatinHashtag", '\ufd50', '\ufd8f')
	add("nonLatinHashtag", '\ufd92', '\ufdc7')
	add("nonLatinHashtag", '\ufdf0', '\ufdfb')
	add("nonLatinHashtag", '\ufe70', '\ufe74') // Arabic Pres. Forms B
	add("nonLatinHashtag", '\ufe76', '\ufefc')
	add("nonLatinHashtag", '\u200c')           // Zero-Width Non-Joiner
	add("nonLatinHashtag", '\u0e01', '\u0e3a') // Thai
	add("nonLatinHashtag", '\u0e40', '\u0e4e') // Hangul (Korean)
	add("nonLatinHashtag", '\u1100', '\u11ff') // Hangul Jamo
	add("nonLatinHashtag", '\u3130', '\u3185') // Hangul Compatibility Jamo
	add("nonLatinHashtag", '\uA960', '\uA97F') // Hangul Jamo Extended-A
	add("nonLatinHashtag", '\uAC00', '\uD7AF') // Hangul Syllables
	add("nonLatinHashtag", '\uD7B0', '\uD7FF') // Hangul Jamo Extended-B
	add("nonLatinHashtag", '\uFFA1', '\uFFDC') // half-width Hangul
	// Japanese and Chinese
	add("nonLatinHashtag", '\u30A1', '\u30FA')         // Katakana (full-width)
	add("nonLatinHashtag", '\u30FC', '\u30FE')         // Katakana Chouon and iteration marks (full-width)
	add("nonLatinHashtag", '\uFF66', '\uFF9F')         // Katakana (half-width)
	add("nonLatinHashtag", '\uFF70', '\uFF70')         // Katakana Chouon (half-width)
	add("nonLatinHashtag", '\uFF10', '\uFF19')         // \
	add("nonLatinHashtag", '\uFF21', '\uFF3A')         // - Latin (full-width)
	add("nonLatinHashtag", '\uFF41', '\uFF5A')         // /
	add("nonLatinHashtag", '\u3041', '\u3096')         // Hiragana
	add("nonLatinHashtag", '\u3099', '\u309E')         // Hiragana voicing and iteration mark
	add("nonLatinHashtag", '\u3400', '\u4DBF')         // Kanji (CJK Extension A)
	add("nonLatinHashtag", '\u4E00', '\u9FFF')         // Kanji (Unified)
	add("nonLatinHashtag", '\U0002A700', '\U0002B73F') // Kanji (CJK Extension C)
	add("nonLatinHashtag", '\U0002B740', '\U0002B81F') // Kanji (CJK Extension D)
	add("nonLatinHashtag", '\U0002F800', '\U0002FA1F') // Kanji (CJK supplement)
	add("nonLatinHashtag", '\u3003')                   // Kanji iteration mark
	add("nonLatinHashtag", '\u3005')                   // Kanji iteration mark
	add("nonLatinHashtag", '\u303B')                   // Han iteration mark

	add("latinAccent", '\u00c0', '\u00d6')
	add("latinAccent", '\u00d8', '\u00f6')
	add("latinAccent", '\u00f8', '\u00ff')
	// Latin Extended A and B
	add("latinAccent", '\u0100', '\u024f')
	// assorted IPA Extensions
	add("latinAccent", '\u0253', '\u0254')
	add("latinAccent", '\u0256', '\u0257')
	add("latinAccent", '\u0259')
	add("latinAccent", '\u025b')
	add("latinAccent", '\u0263')
	add("latinAccent", '\u0268')
	add("latinAccent", '\u026f')
	add("latinAccent", '\u0272')
	add("latinAccent", '\u0289')
	add("latinAccent", '\u028b')
	// Okina for Hawaiian (it *is* a letter character)
	add("latinAccent", '\u02bb')
	// Combining diacritics
	add("latinAccent", '\u0300', '\u036f')
	// Latin Extended Additional
	add("latinAccent", '\u1e00', '\u1eff')

	// Hashtags
	regexen["punct"] = `\!'#%&'\(\)*\+,\\\-\./:;<=>\?@\[\]\^_{|}~\$`
	regexen["hashSigns"] = "[#＃]"
	regexen["hashtagAlpha"] = interp("[a-z_#{latinAccent}#{nonLatinHashtag}]")
	regexen["hashtagAlphaNumeric"] = interp("[a-z0-9_#{latinAccent}#{nonLatinHashtag}]")
	regexen["hashtagBoundary"] = interp("(?:^|$|[^&a-z0-9_#{latinAccent}#{nonLatinHashtag}])")

	invalidHashtagEnd = regexp.MustCompile(interp(`^(?:#{hashSigns}|://)`))
	hashtagPattern = regexp.MustCompile(interp("(?i)(?:#{hashtagBoundary})((?:#{hashSigns})(#{hashtagAlphaNumeric}*#{hashtagAlpha}#{hashtagAlphaNumeric}*))"))

	// URLs
	regexen["urlPreceding"] = interp("[^A-Za-z0-9@＠$#＃#{invalid}]|^")
	regexen["inUrlWithoutProtocolPreceding"] = `[-_./]$`
	regexen["domainChars"] = interp("[^#{punct}#{spaces}#{invalid}]")
	regexen["subdomain"] = interp(`(?:(?:#{domainChars}(?:[_-]|#{domainChars})*)?#{domainChars}\.)`)
	regexen["domainName"] = interp(`(?:(?:#{domainChars}(?:-|#{domainChars})*)?#{domainChars}\.)`)
	regexen["GTLD"] = "(?:aero|asia|biz|cat|com|coop|edu|gov|info|int|jobs|mil|mobi|museum|name|net|org|pro|tel|travel|xxx)"
	regexen["CCTLD"] = "(?:ac|ad|ae|af|ag|ai|al|am|an|ao|aq|ar|as|at|au|aw|ax|az|ba|bb|bd|be|bf|bg|bh|bi|bj|bm|bn|bo|br|bs|bt|bv|bw|by|bz|" +
		"ca|cc|cd|cf|cg|ch|ci|ck|cl|cm|cn|co|cr|cs|cu|cv|cx|cy|cz|dd|de|dj|dk|dm|do|dz|ec|ee|eg|eh|er|es|et|eu|fi|fj|fk|fm|fo|fr|" +
		"ga|gb|gd|ge|gf|gg|gh|gi|gl|gm|gn|gp|gq|gr|gs|gt|gu|gw|gy|hk|hm|hn|hr|ht|hu|id|ie|il|im|in|io|iq|ir|is|it|je|jm|jo|jp|" +
		"ke|kg|kh|ki|km|kn|kp|kr|kw|ky|kz|la|lb|lc|li|lk|lr|ls|lt|lu|lv|ly|ma|mc|md|me|mg|mh|mk|ml|mm|mn|mo|mp|mq|mr|ms|mt|mu|mv|mw|mx|my|mz|" +
		"na|nc|ne|nf|ng|ni|nl|no|np|nr|nu|nz|om|pa|pe|pf|pg|ph|pk|pl|pm|pn|pr|ps|pt|pw|py|qa|re|ro|rs|ru|rw|" +
		"sa|sb|sc|sd|se|sg|sh|si|sj|sk|sl|sm|sn|so|sr|ss|st|su|sv|sx|sy|sz|tc|td|tf|tg|th|tj|tk|tl|tm|tn|to|tp|tr|tt|tv|tw|tz|" +
		"ua|ug|uk|us|uy|uz|va|vc|ve|vg|vi|vn|vu|wf|ws|ye|yt|za|zm|zw)"
	regexen["punycode"] = "(?:xn--[0-9a-z]+)"
	regexen["domain"] = interp("(?:#{subdomain}*#{domainName}(?:#{GTLD}|#{CCTLD}|#{punycode}))")

	asciiDomain = regexp.MustCompile(interp(`(?i)(?:(?:[\-a-z0-9#{latinAccent}]+)\.)+(?:#{GTLD}|#{CCTLD}|#{punycode})`))
	invalidBeforeDomain = regexp.MustCompile(`[-_./]$`)
	invalidWithoutPath = regexp.MustCompile(interp("(?i)^#{domainName}#{CCTLD}$"))

	regexen["portNumber"] = "[0-9]+"
	regexen["generalUrlPath"] = interp(`[a-z0-9!\*';:=\+,\.\$\/%#\[\]\-_~@|&#{latinAccent}]`)
	regexen["urlBalancedParens"] = interp(`\(#{generalUrlPath}+\)`)
	regexen["urlPathEnding"] = interp(`[\+\-a-z0-9=_#\/#{latinAccent}]|(?:#{urlBalancedParens})`)
	regexen["urlPath"] = interp(`(?:(?:#{generalUrlPath}*(?:#{urlBalancedParens}#{generalUrlPath}*)*#{urlPathEnding})|(?:@#{generalUrlPath}+/))`)
	regexen["urlQuery"] = `[a-z0-9!?\*'@\(\);:&=\+\$/%#\[\]\-_\.,~|]`
	regexen["urlQueryEnding"] = `[a-z0-9_&=#/]`

	urlPattern = regexp.MustCompile(interp(`(?im)(#{urlPreceding})` + // [1] Preceding character
		`(` + // [2] URL
		`(https?://)?` + // [3] Protocol
		`(#{domain})` + // [4] Domain
		`(?::#{portNumber})?` + // Port number
		`(/#{urlPath}*)?` + // [4] URL Path
		`(?:\?#{urlQuery}*#{urlQueryEnding})?)`)) // Query String
}

var hashtagPattern, invalidHashtagEnd *regexp.Regexp
var urlPattern, asciiDomain, invalidBeforeDomain, invalidWithoutPath *regexp.Regexp

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
