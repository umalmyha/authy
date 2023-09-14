package header

import "net/http"

const (
	charsetUTF8 = "charset=UTF-8"

	HeaderContentType = "Content-Type"

	MIMEApplicationJSON            = "application/json"
	MIMEApplicationJSONCharsetUTF8 = MIMEApplicationJSON + "; " + charsetUTF8
)

func Header(r *http.Request, header string) string {
	return r.Header.Get(header)
}

func SetHeader(w http.ResponseWriter, header string, value string) {
	w.Header().Set(header, value)
}
