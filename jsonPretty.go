package codecs

import (
	"encoding/json"

	"github.com/spartanlogs/spartan/codecs"
	"github.com/spartanlogs/spartan/config"
	"github.com/spartanlogs/spartan/event"
	"github.com/spartanlogs/spartan/utils"
)

type jsonConfig struct {
	indent string
}

var jsonConfigSchema = []config.Setting{
	{
		Name:    "indent",
		Type:    config.String,
		Default: "  ",
	},
}

// The JSONPrettyCodec encodes/decodes an event as formatted, pretty JSON.
type JSONPrettyCodec struct {
	codecs.BaseCodec
	config *jsonConfig
}

func init() {
	codecs.Register("json_pretty", newJSONPrettyCodec)
}

func newJSONPrettyCodec(options utils.InterfaceMap) (codecs.Codec, error) {
	c := &JSONPrettyCodec{
		config: &jsonConfig{},
	}
	return c, c.setConfig(options)
}

func (c *JSONPrettyCodec) setConfig(options utils.InterfaceMap) error {
	var err error
	options, err = config.VerifySettings(options, jsonConfigSchema)
	if err != nil {
		return err
	}

	c.config.indent = options.Get("indent").(string)

	return nil
}

// Encode Event as JSON object.
func (c *JSONPrettyCodec) Encode(e *event.Event) []byte {
	data := e.Data()
	j, _ := json.MarshalIndent(data, "", c.config.indent)
	return j
}
