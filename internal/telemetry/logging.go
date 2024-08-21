// Copyright Project Harbor Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package telemetry

import (
	"context"
	"errors"
	"io"
	"time"

	"go.opentelemetry.io/otel/log"
)

// Logger returns a logger with the given name.
func Logger(name string) log.Logger {
	return loggerProvider.Logger(name) // TODO more instrumentation attrs
}

// SpanStdio returns a pair of io.WriteClosers which will send log records with
// stdio.stream=1 for stdout and stdio.stream=2 for stderr. Closing either of
// them will send a log record for that stream with an empty body and
// stdio.eof=true.
//
// SpanStdio should be used when a span represents a process that writes to
// stdout/stderr and terminates them with an EOF, to confirm that all data has
// been received. It should not be used for general-purpose logging.
//
// Both streamsm must be closed to ensure that draining completes.
func SpanStdio(ctx context.Context, name string, attrs ...log.KeyValue) SpanStreams {
	logger := Logger(name)
	return SpanStreams{
		Stdout: &spanStream{
			Writer: &Writer{
				ctx:    ctx,
				logger: logger,
				attrs:  append([]log.KeyValue{log.Int(StdioStreamAttr, 1)}, attrs...),
			},
		},
		Stderr: &spanStream{
			Writer: &Writer{
				ctx:    ctx,
				logger: logger,
				attrs:  append([]log.KeyValue{log.Int(StdioStreamAttr, 2)}, attrs...),
			},
		},
	}
}

// Writer is an io.Writer that emits log records.
type Writer struct {
	ctx    context.Context
	logger log.Logger
	attrs  []log.KeyValue
}

// NewWriter returns a new Writer that emits log records with the given logger
// name and attributes.
func NewWriter(ctx context.Context, name string, attrs ...log.KeyValue) io.Writer {
	return &Writer{
		ctx:    ctx,
		logger: Logger(name),
		attrs:  attrs,
	}
}

// Write emits a log record with the given payload as a string body.
func (w *Writer) Write(p []byte) (int, error) {
	w.Emit(log.StringValue(string(p)))
	return len(p), nil
}

// Emit sends a log record with the given body and additional attributes.
func (w *Writer) Emit(body log.Value, attrs ...log.KeyValue) {
	rec := log.Record{}
	rec.SetTimestamp(time.Now())
	rec.SetBody(body)
	rec.AddAttributes(w.attrs...)
	rec.AddAttributes(attrs...)
	w.logger.Emit(w.ctx, rec)
}

// SpanStreams contains the stdout and stderr for a span.
type SpanStreams struct {
	Stdout io.WriteCloser
	Stderr io.WriteCloser
}

// Calling Close closes both streams.
func (sl SpanStreams) Close() error {
	return errors.Join(
		sl.Stdout.Close(),
		sl.Stderr.Close(),
	)
}

type spanStream struct {
	*Writer
}

// Close emits an EOF log record.
func (w *spanStream) Close() error {
	w.Writer.Emit(log.StringValue(""), log.Bool(StdioEOFAttr, true))
	return nil
}
