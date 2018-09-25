// 生成gif图像
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

// color.Color类型的数组
var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

// lissajous ...
func lissajous(out io.Writer)  {
	const (
		cycles = 5				// 
		res = 0.001
		size = 100
		nframes = 64
		delay = 8
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes} // 复合声明，GIF是个结构，LoopCount是它的成员，其余字段会是各自默认的零值
	phase := 0.0
	for i := 0; i< nframes; i++ {
		rect := image.Rect(0, 0, 2*size + 1, 2*size+1)
		// 构建指定大小的图像。默认各像素都是零值，即palette的第一个元素的值(color.White)
		img := image.NewPaletted(rect, palette) 
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			// 设置曲线上的点为黑色
			img.SetColorIndex(size + int(x*size+0.5),
				size + int(y*size +0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image,img)
	}
	gif.EncodeAll(out, &anim)
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout)
}
