package request

import (
	"encoding/json"
	"net/http"

	"github.com/umalmyha/authy/internal/web"
	"github.com/umalmyha/authy/internal/web/header"
)

func BindJSON(r *http.Request, target any) error {
	if ctyp := header.Header(r, header.HeaderContentType); ctyp != header.MIMEApplicationJSON {
		return web.ErrUnsupportedMediaType
	}
	return json.NewDecoder(r.Body).Decode(target)
}
