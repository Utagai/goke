package internal

import "io"

// NewPrefixWriter creates a PrefixWriter.
func NewPrefixWriter(w io.Writer) *PrefixWriter {
	return &PrefixWriter{
		w:  w,
		nl: true,
	}
}

// PrefixWriter wraps an io.Writer to automatically add a prefix at the beginning
// of every line.
type PrefixWriter struct {
	w      io.Writer
	prefix []byte
	nl     bool
}

// SetPrefix sets the prefix.
func (w *PrefixWriter) SetPrefix(prefix []byte) {
	w.prefix = prefix
}

func (w *PrefixWriter) Write(p []byte) (n int, err error) {
	for _, c := range p {
		if w.nl {
			_, err = w.w.Write(w.prefix)
			if err != nil {
				return n, err
			}
			w.nl = false
		}

		_, err = w.w.Write([]byte{c})
		if err != nil {
			return n, err
		}

		n++
		w.nl = c == '\n'
	}

	return n, nil
}
