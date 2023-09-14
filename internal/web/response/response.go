package response

import (
	"encoding/json"
	"net/http"

	"github.com/umalmyha/authy/internal/web/header"
)

func JSON(w http.ResponseWriter, status int, data any) error {
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.WriteHeader(status)
	header.SetHeader(w, header.HeaderContentType, header.MIMEApplicationJSONCharsetUTF8)
	if _, err = w.Write(body); err != nil {
		return err
	}

	return nil
}

func Status(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
}
