package text

import (
	"regexp"
	"strings"
)

func add(p string, r ...rune) {
	if len(r) == 1 {
		charGroups[p] = append(charGroups[p], string(r[0]))
	} else {
		charGroups[p] = append(charGroups[p], string(r[0])+"-"+string(r[1]))
	}
}

var charGroups = make(map[string][]string)
var regexen = make(map[string]string)
var validHashtagPattern, invalidHashtagEnd *regexp.Regexp

func init() {
	charGroups["unicodeSpaces"] = []string{
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
	add("unicodeSpaces", '\u0009', '\u000D') // White_Space # Cc [5] <control-0009>..<control-000D>
	add("unicodeSpaces", '\u2000', '\u200A') // White_Space # Zs [11] EN QUAD..HAIR SPACE

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

	regexen["punct"] = `\!'#%&'\(\)*\+,\\\-\.\/:;<=>\?@\[\]\^_{|}~\$`
	regexen["hashSigns"] = "[#＃]"
	regexen["hashtagAlpha"] = regexpInterpolate("[a-z_#{latinAccent}#{nonLatinHashtag}]")
	regexen["hashtagAlphaNumeric"] = regexpInterpolate("[a-z0-9_#{latinAccent}#{nonLatinHashtag}]")
	regexen["hashtagBoundary"] = regexpInterpolate("(?:^|$|[^&a-z0-9_#{latinAccent}#{nonLatinHashtag}])")
	invalidHashtagEnd = regexp.MustCompile(regexpInterpolate(`^(?:#{hashSigns}|:\/\/)`))
	validHashtagPattern = regexp.MustCompile(regexpInterpolate("(?i)(?:#{hashtagBoundary})(?:#{hashSigns})(#{hashtagAlphaNumeric}*#{hashtagAlpha}#{hashtagAlphaNumeric}*)"))
}

func ExtractHashtags(s string) []string {
	indices := ExtractHashtagIndices(s)
	res := make([]string, len(indices))
	for i, idx := range indices {
		res[i] = s[idx[0]:idx[1]]
	}
	return res
}

func ExtractHashtagIndices(s string) [][2]int {
	matches := validHashtagPattern.FindAllStringSubmatchIndex(s, -1)
	res := make([][2]int, 0, len(matches))
	for _, match := range matches {
		if invalidHashtagEnd.MatchString(s[match[3]:]) {
			continue
		}
		res = append(res, [2]int{match[2], match[3]})
	}
	return res
}

func regexpInterpolate(s string) string {
	replace := func(g string) string {
		g = g[2 : len(g)-1]
		if regex, ok := regexen[g]; ok {
			return regex
		}
		group, ok := charGroups[g]
		if !ok {
			panic("unknown group or regexen " + g)
		}
		return strings.Join(group, "")
	}
	return regexp.MustCompile(`#\{\w+\}`).ReplaceAllStringFunc(s, replace)
}
