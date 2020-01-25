package synth

import (
	"github.com/hajimehoshi/oto"
	"io"
	"math"
)

const (
	sampleRate      = 48000
	channelNum      = 2
	bitDepthInBytes = 2
)

type SineWave struct {
	freq      float64
	length    int64
	pos       int64
	remaining []byte
}

func (siw *SineWave) Read(buf []byte) (int, error) {
	// Copy the remaining bytes of the sine wave to the io buffer
	if len(siw.remaining) > 0 {
		n := copy(buf, siw.remaining)
		siw.remaining = siw.remaining[n:]
		return n, nil
	}

	// End the read if all of the sine wave is read
	if siw.pos == siw.length {
		return 0, io.EOF
	}

	// Set flag for EOF if this is the final call to Read by calculating the predicted buffer length
	eof := false
	if siw.pos+int64(len(buf)) > siw.length {
		buf = buf[:siw.length-siw.pos]
		eof = true
	}

	// Keep a copy of the original buffer
	// Modify buf
	var origBuf []byte
	if len(buf)%4 > 0 {
		origBuf = buf
		buf = make([]byte, len(origBuf)+4-len(origBuf)%4)
	}

	// Calculate the length of the sine wave and the new position in the buffer
	length := sampleRate / siw.freq
	num := bitDepthInBytes * channelNum
	p := siw.pos / int64(num)

	// TODO: See if there is any benefit to using 1 bit depth, otherwise yeet the switch statement
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

	// increment the buffer position
	siw.pos += int64(len(buf))

	// Copy over the original buffer and update the remaining bytes
	n := len(buf)
	if origBuf != nil {
		n = copy(origBuf, buf)
		siw.remaining = buf[n:]
	}

	// End the read altogether if the EOF flag was set
	if eof {
		return n, io.EOF
	}
	return n, nil
}

// Formula for sine wave: Asin(fx)*A
func NewSineWave(freq float64, duration int64) *SineWave {
	l := channelNum * bitDepthInBytes * sampleRate * duration / 1000
	l = l / 4 * 4
	return &SineWave{
		freq:   freq,
		length: l,
	}
}

func Play(context *oto.Context, freq float64, duration int64) error {
	p := context.NewPlayer()
	s := NewSineWave(freq, duration)
	if _, err := io.Copy(p, s); err != nil {
		return err
	}
	if err := p.Close(); err != nil {
		return err
	}
	return nil
}