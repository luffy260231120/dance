package patch

import (
	"bytes"
	"encoding/json"
	"fmt"
	l "github.com/sirupsen/logrus"
	"time"
)

type LogFormatter struct {
	*l.JSONFormatter
}

// Format renders a single log entry
func (f *LogFormatter) Format(entry *l.Entry) ([]byte, error) {
	data := make(l.Fields, len(entry.Data)+4)
	for k, v := range entry.Data {
		switch v := v.(type) {
		case error:
			// Otherwise errors are ignored by `encoding/json`
			// https://github.com/sirupsen/logrus/issues/137
			data[k] = v.Error()
		default:
			data[k] = v
		}
	}

	if f.DataKey != "" {
		newData := make(l.Fields, 4)
		newData[f.DataKey] = data
		data = newData
	}

	timestampFormat := f.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = time.RFC3339
	}

	//if entry.err != "" {
	//	data[f.FieldMap.resolve(l.FieldKeyLogrusError)] = entry.err
	//}
	if !f.DisableTimestamp {
		data[l.FieldKeyTime] = entry.Time.Format(timestampFormat)
	}
	data[l.FieldKeyMsg] = entry.Message
	data[l.FieldKeyLevel] = entry.Level.String()
	if entry.HasCaller() {
		data[l.FieldKeyFunc] = entry.Caller.Function
		data[l.FieldKeyFile] = fmt.Sprintf("%s:%d", entry.Caller.File, entry.Caller.Line)
	}

	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	encoder := json.NewEncoder(b)
	if f.PrettyPrint {
		encoder.SetIndent("", "  ")
	}
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(data); err != nil {
		return nil, fmt.Errorf("Failed to marshal fields to JSON, %v", err)
	}
	return b.Bytes(), nil
}
