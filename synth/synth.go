package synth

import (
	"github.com/hajimehoshi/oto"
	"io"
	"math"
)

const (
	sampleRate      = 44100
	channelNum      = 2
	bitDepthInBytes = 2
)

type SineWave struct {
	freq      float64
	length    int64
	pos       int64
	remaining []byte
}

func Play(context *oto.Context, freq float64, duration int64) error {
	p := context.NewPlayer()
	s := newSineWave(freq, duration)
	if _, err := io.Copy(p, s); err != nil {
		return err
	}
	if err := p.Close(); err != nil {
		return err
	}
	return nil
}

func newSineWave(freq float64, duration int64) *SineWave {
	l := channelNum * bitDepthInBytes * sampleRate * duration / 1000
	l = l / 4 * 4
	return &SineWave{
		freq:   freq,
		length: l,
	}
}

func (s *SineWave) Read(buf []byte) (int, error) {
	if len(s.remaining) > 0 {
		n := copy(buf, s.remaining)
		s.remaining = s.remaining[n:]
		return n, nil
	}

	if s.pos == s.length {
		return 0, io.EOF
	}

	eof := false
	if s.pos+int64(len(buf)) > s.length {
		buf = buf[:s.length-s.pos]
		eof = true
	}

	var origBuf []byte
	if len(buf)%4 > 0 {
		origBuf = buf
		buf = make([]byte, len(origBuf)+4-len(origBuf)%4)
	}

	length := sampleRate / s.freq

	num := bitDepthInBytes * channelNum
	p := s.pos / int64(num)
	switch bitDepthInBytes {
	case 1:
		for i := 0; i < len(buf)/num; i++ {
			const max = 127
			b := int(math.Sin(2*math.Pi*float64(p)/length) * 0.3 * max)
			for ch := 0; ch < channelNum; ch++ {
				buf[num*i+ch] = byte(b + 128)
			}
			p++
		}
	case 2:
		for i := 0; i < len(buf)/num; i++ {
			const max = 32767
			b := int16(math.Sin(2*math.Pi*float64(p)/length) * 0.3 * max)
			for ch := 0; ch < channelNum; ch++ {
				buf[num*i+2*ch] = byte(b)
				buf[num*i+1+2*ch] = byte(b >> 8)
			}
			p++
		}
	}

	s.pos += int64(len(buf))

	n := len(buf)
	if origBuf != nil {
		n = copy(origBuf, buf)
		s.remaining = buf[n:]
	}

	if eof {
		return n, io.EOF
	}
	return n, nil
}
