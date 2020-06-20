package app

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"image/color"
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
	"f32toui64": func(f float32) uint64 {
		return uint64(f)
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
	"mix": func(s1, s2 string, weight float32) (string, error) {
		if strings.HasPrefix(s1, "#") {
			s1 = s1[1:]
		}
		if strings.HasPrefix(s2, "#") {
			s2 = s2[1:]
		}

		b1, err := hex.DecodeString(s1)
		if err != nil {
			return "", err
		}
		b2, err := hex.DecodeString(s2)
		if err != nil {
			return "", err
		}

		c1 := color.RGBA{b1[0], b1[1], b1[2], 0}
		c2 := color.RGBA{b2[0], b2[1], b2[2], 0}

		w1 := 1 - weight
		w2 := weight

		// Great type system... /s
		mix := func(c1, c2 byte) byte {
			return byte(float32(c1)*w1) + byte(float32(c2)*w2)
		}

		res := color.RGBA{
			R: mix(c1.R, c2.R),
			G: mix(c1.G, c2.G),
			B: mix(c1.B, c2.B),
			A: 0,
		}

		return "#" + hex.EncodeToString([]byte{res.R, res.G, res.B}), nil
	},

	// For lisp-style math with floats
	"add": func(args ...float32) float32 {
		res := args[0]
		for _, v := range args[1:] {
			res += v
		}

		return res
	},
	"sub": func(args ...float32) float32 {
		res := args[0]
		for _, v := range args[1:] {
			res -= v
		}

		return res
	},
	"div": func(args ...float32) float32 {
		res := args[0]
		for _, v := range args[1:] {
			res /= v
		}

		return res
	},
	"mul": func(args ...float32) float32 {
		res := args[0]
		for _, v := range args {
			res *= v
		}

		return res
	},
}
