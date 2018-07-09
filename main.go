package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"time"

	"github.com/jkuri/party-face-detection/mtcnn"
	"gocv.io/x/gocv"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	videoFile = kingpin.Flag("file", "Input video file.").Default("./data/videos/ag.mp4").Short('f').String()
	modelFile = kingpin.Flag("model", "MTCNN TF model file").Default("./data/models/mtcnn.pb").Short('m').String()
)

func main() {
	kingpin.Parse()

	det, err := mtcnn.NewMtcnnDetector(*modelFile)
	if err != nil {
		log.Fatal(err)
	}
	defer det.Close()

	det.Config(0, 0, []float32{0.7, 0.7, 0.95})

	webcam, err := gocv.OpenVideoCapture(*videoFile)
	if err != nil {
		log.Fatal(err)
	}
	defer webcam.Close()

	window := gocv.NewWindow("gocv-mtcnn")
	defer window.Close()

	frame := gocv.NewMat()
	defer frame.Close()

	red := color.RGBA{255, 0, 0, 0}
	frameCounter := 0
	tick := 0
	timeBegin := time.Now()
	var fps int

	for {
		if ok := webcam.Read(&frame); !ok {
			log.Fatalf("Error: cannot read frame from file\n")
		}
		if frame.Empty() {
			continue
		}

		img, _ := frame.ToImage()
		imageBytes := new(bytes.Buffer)
		err := jpeg.Encode(imageBytes, img, nil)
		if err != nil {
			log.Fatal(err)
		}

		tensorImg, err := mtcnn.TensorFromJpeg(imageBytes.Bytes())
		if err != nil {
			log.Fatal(err)
		}

		bbox, err := det.DetectFaces(tensorImg)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("%d faces found\n", len(bbox))

		for _, b := range bbox {
			rect := image.Rect(int(b[0]), int(b[1]), int(b[2]), int(b[3]))
			gocv.Rectangle(&frame, rect, red, 1)
		}

		frameCounter++
		timeNow := int(time.Now().Sub(timeBegin).Seconds())
		if timeNow-tick >= 1 {
			tick++
			fps = frameCounter
			frameCounter = 0
		}

		gocv.PutText(&frame, fmt.Sprintf("FPS: %d", fps), image.Pt(10, 20), gocv.FontHersheyPlain, 1.2, color.RGBA{0, 255, 0, 0}, 2)

		window.IMShow(frame)
		if window.WaitKey(1) == 27 {
			break
		}
	}

}
