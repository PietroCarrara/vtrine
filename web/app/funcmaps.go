package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"github.com/PietroCarrara/vtrine/pkg/torrent"
	"github.com/dustin/go-humanize"
)

var funcmaps = template.FuncMap{
	"env":         os.Getenv,
	"bestTorrent": bestTorrent,
	"url": func(s string) template.URL {
		return template.URL(s)
	},
	"byte": func(s uint64) string {
		return humanize.Bytes(s)
	},
	"rand": func() string {
		return strconv.Itoa(rand.Int())
	},
	"cat": func(s ...string) string {
		return strings.Join(s, "")
	},
	"json": func(i interface{}) template.JS {
		obj, _ := json.Marshal(i)
		return template.JS(obj)
	},
	"dict": func(values ...interface{}) (map[string]interface{}, error) {
		if len(values)%2 != 0 {
			return nil, errors.New("invalid dict call")
		}
		dict := make(map[string]interface{}, len(values)/2)
		for i := 0; i < len(values); i += 2 {
			key, ok := values[i].(string)
			if !ok {
				return nil, errors.New("dict keys must be strings")
			}
			dict[key] = values[i+1]
		}
		return dict, nil
	},
	"percent": func(f float32) string {
		num := fmt.Sprintf("%.2f", f*100)

		var discard, decimal int
		fmt.Sscanf(num, "%d.%d", &discard, &decimal)

		if decimal == 0 {
			return fmt.Sprintf("%.0f", f*100)
		}

		return num
	},
	"complete": func(arr []torrent.ClientData) []torrent.ClientData {
		res := make([]torrent.ClientData, 0)

		for _, v := range arr {
			if v.Complete {
				res = append(res, v)
			}
		}

		return res
	},
	"nonComplete": func(arr []torrent.ClientData) []torrent.ClientData {
		res := make([]torrent.ClientData, 0)

		for _, v := range arr {
			if !v.Complete {
				res = append(res, v)
			}
		}

		return res
	},
}
