package codecs

import (
	"encoding/json"
	"io"

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
	if err := config.VerifySettings(options, jsonConfigSchema); err != nil {
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

// EncodeWriter reads events from in and writes them to w
func (c *JSONPrettyCodec) EncodeWriter(w io.Writer, in <-chan *event.Event) {}

// Decode byte slice into an Event. CURRENTLY NOT IMPLEMENTED.
func (c *JSONPrettyCodec) Decode(data []byte) (*event.Event, error) {
	return nil, nil
}

// DecodeReader reads from r and creates an event sent to out
func (c *JSONPrettyCodec) DecodeReader(r io.Reader, out chan<- *event.Event) {}
