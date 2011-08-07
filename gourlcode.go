package gourlcode

var dec2hex = [16]byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}

func Unescape(buf *[]byte, length int) (unescaped_length int) {
	pos := 0
	for c := 0; c < length; c++ {
		if (*buf)[c] == '%' {
			switch {
			case '0' <= (*buf)[c+1] && (*buf)[c+1] <= '9':
				(*buf)[pos] = (*buf)[c+1] - '0'
			case 'a' <= (*buf)[c+1] && (*buf)[c+1] <= 'f':
				(*buf)[pos] = 10 + (*buf)[c+1] - 'a'
			case 'A' <= (*buf)[c+1] && (*buf)[c+1] <= 'F':
				(*buf)[pos] = 10 + (*buf)[c+1] - 'A'
			}
			(*buf)[pos] <<= 4
			switch {
			case '0' <= (*buf)[c+2] && (*buf)[c+2] <= '9':
				(*buf)[pos] += (*buf)[c+2] - '0'
			case 'a' <= (*buf)[c+2] && (*buf)[c+2] <= 'f':
				(*buf)[pos] += 10 + (*buf)[c+2] - 'a'
			case 'A' <= (*buf)[c+2] && (*buf)[c+2] <= 'F':
				(*buf)[pos] += 10 + (*buf)[c+2] - 'A'
			}
			c += 2
		} else {
			(*buf)[pos] = (*buf)[c]
		}
		pos++
	}
	return pos
}

func Escape(raw []byte) []byte {
	output := make([]byte, 3*len(raw)) // TODO: make more dynamic?
	pos := 0
	for _, c := range raw {
		if 'A' <= c && c <= 'Z' || 'a' <= c && c <= 'z' || '0' <= c && c <= '9' {
			output[pos] = c
			pos++
		} else {
			output[pos] = '%'
			output[pos+1] = dec2hex[(c>>4)&0x0f]
			output[pos+2] = dec2hex[c&0x0f]
			pos += 3
		}
	}
	return output[0:pos]
}

