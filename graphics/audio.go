package graphics

// typedef unsigned char Uint8;
// void SineWave(void *userdata, Uint8 *stream, int len);
import "C"
import (
	"math"
	"unsafe"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	toneHz     = 440
	sampleRate = 44100
	dPhase     = 2 * math.Pi * toneHz / sampleRate
)

//export SineWave
func SineWave(userdata unsafe.Pointer, stream *C.Uint8, length C.int) {
	n := int(length)
	hdr := unsafe.Slice(stream, length)
	buf := *(*[]C.Uint8)(unsafe.Pointer(&hdr))

	var phase float64
	for i := 0; i < n; i += 2 {
		phase += dPhase
		sample := C.Uint8((math.Sin(phase) + 0.999999) * 128)
		buf[i] = sample
		buf[i+1] = sample
	}
}

// create audio callback
var AudioCallback = sdl.AudioCallback(C.SineWave)
