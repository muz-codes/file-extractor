package constant

var Html = []string{"html"}
var BlockLevelElements = map[string]struct{}{
	"p":          {},
	"div":        {},
	"article":    {},
	"section":    {},
	"header":     {},
	"footer":     {},
	"main":       {},
	"aside":      {},
	"h1":         {},
	"h2":         {},
	"h3":         {},
	"h4":         {},
	"h5":         {},
	"h6":         {},
	"ol":         {},
	"ul":         {},
	"li":         {},
	"dl":         {},
	"dt":         {},
	"dd":         {},
	"blockquote": {},
	"figure":     {},
	"figcaption": {},
	//"pre":        {}, 	//added to inline
	"table":    {},
	"caption":  {},
	"thead":    {},
	"tbody":    {},
	"tfoot":    {},
	"colgroup": {},
	"col":      {},
	"tr":       {},
	"td":       {},
	"th":       {},
	"form":     {},
	"fieldset": {},
	"legend":   {},
	"button":   {},
	"address":  {},
	"details":  {},
	"hr":       {},
	"menu":     {},
	"object":   {},
	"canvas":   {},
	"video":    {},
	"audio":    {},
	"picture":  {},
	"select":   {},
	"textarea": {},
	"portal":   {},
	"datalist": {},
	"source":   {},
	"option":   {},
	"optgroup": {},
	"summary":  {},
	"dialog":   {},
	"nav":      {},
	"slot":     {},
	"template": {},
	"ins":      {},
	"del":      {},
	// Following can wrap block elements, that's why included here
	"label": {},
}

var DeprecatedBlockLevelElements = map[string]struct{}{
	"center":   {},
	"dir":      {},
	"frame":    {},
	"frameset": {},
	"noframes": {},
	"applet":   {},
	"isindex":  {},
	"spacer":   {},
	"content":  {},
	"marquee":  {},
	"menuitem": {},
	"nobr":     {},
	"noembed":  {},
}

// pre, plaintext and xmp are included in the inline so that if any of these tags contain html element, it is not be traversed
const InlineElements = "a, progress, embed, pre,  abbr, b, bdi, bdo, br, cite, code, data, dfn, em, i, input, kbd, label, mark, meter, output, q, rp, rt, ruby, s, samp, small, span, strong, sub, sup, time, u, var, wbr"
const DeprecatedInlineElements = "acronym, big, basefont, tt, font, strike, blink, image, param, plaintext, rb, rtc, shadow, xmp"

var IgnoredElements = map[string]struct{}{
	"svg":      {},
	"math":     {},
	"style":    {},
	"script":   {},
	"noscript": {},
	// following elements don't have content
	"iframe": {},
	"img":    {},
	"area":   {},
	"map":    {},
	"track":  {},
}

const (
	Gt          = "&gt;"
	Lt          = "&lt;"
	Ampersand   = "&#38;"
	SingleQuote = "&#39;"
	DoubleQuote = "&#34;"
)
