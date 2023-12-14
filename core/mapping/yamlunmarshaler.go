package mapping

import (
	"io"

	"github.com/wuntsong-org/go-zero-plus/internal/encoding"
)

// UnmarshalYamlBytes unmarshals content into v.
func UnmarshalYamlBytes(content []byte, v any, opts ...UnmarshalOption) error {
	b, err := encoding.YamlToJson(content)
	if err != nil {
		return err
	}

	return UnmarshalJsonBytes(b, v, opts...)
}

// UnmarshalYamlReader unmarshals content from reader into v.
func UnmarshalYamlReader(reader io.Reader, v any, opts ...UnmarshalOption) error {
	b, err := io.ReadAll(reader)
	if err != nil {
		return err
	}

	return UnmarshalYamlBytes(b, v, opts...)
}
