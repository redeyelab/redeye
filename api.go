/*
RedEye the Smart Camera Package.
*/
package redeye

// Blob is either a JPEG image as an array of bytes
// or a *gocv.Mat, we wrap this structure such that modules
// can be built without relying on building GOCV.
type VideoFrame interface {
	Bytes() []bytes
	Blob() interface{}
}

// VideoSource Play channel delivers either []byte or *gocv.Mat
type VideoSource interface {
	Name() string
	Play(stop chan interface{}) chan *Blob
}

// VidePlayer provides an MJPEG video stream over a specific
// IP address and port.
type VideoPlayer interface {
	Name() string
	PlayVideo(VideoSource)
}

// VideoPipe accepts an JPEG image converts it to
// type VideoReader interface {
// }
