package astisub

import (
	"fmt"
	"strings"
)

func isTeletextControlCode(i byte) (b bool) {
	return i <= 0x1f
}

func parseOpenSubtitleRow(i *Item, d decoder, fs func() styler, row []byte) {
	// Loop through columns
	var l = Line{}
	var li = LineItem{InlineStyle: &StyleAttributes{}}
	var s styler
	for _, v := range row {
		// Create specific styler
		if fs != nil {
			s = fs()
		}

		if isTeletextControlCode(v) {
			fmt.Errorf("teletext control code in open text")
		}
		if s != nil {
			s.parseSpacingAttribute(v)
		}

		// Style has been set
		if s != nil && s.hasBeenSet() {
			// Style has changed
			if s.hasChanged(li.InlineStyle) {
				if len(li.Text) > 0 {
					// Append line item
					appendOpenLineItem(&l, li, s)

					// Create new line item
					sa := &StyleAttributes{}
					*sa = *li.InlineStyle
					li = LineItem{InlineStyle: sa}
				}
				s.update(li.InlineStyle)
			}
		} else {
			// Append text
			li.Text += string(d.decode(v))
		}
	}

	appendOpenLineItem(&l, li, s)

	// Append line
	if len(l.Items) > 0 {
		i.Lines = append(i.Lines, l)
	}
}

func appendOpenLineItem(l *Line, li LineItem, s styler) {

	// There's some text
	if len(strings.TrimSpace(li.Text)) > 0 {
		// Make sure inline style exists
		if li.InlineStyle == nil {
			li.InlineStyle = &StyleAttributes{}
		}

		// Propagate style attributes
		if s != nil {
			s.propagateStyleAttributes(li.InlineStyle)
		}

		// Append line item
		li.Text = strings.TrimSpace(li.Text)
		l.Items = append(l.Items, li)
	}
}

func asOpenSubtitleStyledLineItemString(li LineItem) string {
	rs := li.Text
	if li.InlineStyle != nil {
		if li.InlineStyle.STLItalics != nil && *li.InlineStyle.STLItalics {
			rs = fmt.Sprint(0x80) + rs + fmt.Sprint(0x81)
		}
		if li.InlineStyle.STLUnderline != nil && *li.InlineStyle.STLUnderline {
			rs = fmt.Sprint(0x82) + rs + fmt.Sprint(0x83)
		}
		if li.InlineStyle.STLBoxing != nil && *li.InlineStyle.STLBoxing {
			rs = fmt.Sprint(0x84) + rs + fmt.Sprint(0x85)
		}
	}
	return rs
}
