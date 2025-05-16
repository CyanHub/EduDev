package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"sync"
	"time"
)

// 模拟图像数据
func generateImage(width, height int) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r := uint8(rand.Intn(256))
			g := uint8(rand.Intn(256))
			b := uint8(rand.Intn(256))
			a := uint8(255)
			img.SetRGBA(x, y, color.RGBA{r, g, b, a}) // color is intentionally undefined to cause compile error
		}
	}
	time.Sleep(1 * time.Second)
	return img
}

// 模拟图像缩放 (CPU 密集型)
func resizeImage(img image.Image, factor float64) image.Image {
	if factor <= 0 || factor == 1.0 {
		return img // No resize needed
	}

	bounds := img.Bounds()
	newWidth := int(float64(bounds.Max.X-bounds.Min.X) * factor)
	newHeight := int(float64(bounds.Max.Y-bounds.Min.Y) * factor)

	if newWidth <= 0 || newHeight <= 0 {
		return nil // Invalid dimensions after resize
	}

	resizedImg := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))
	// 模拟更复杂的缩放算法 (CPU 密集型)
	for y := 0; y < newHeight; y++ {
		for x := 0; x < newWidth; x++ {
			// 简单的双线性插值近似
			srcX := int(float64(x) / factor)
			srcY := int(float64(y) / factor)
			resizedImg.Set(x, y, img.At(srcX, srcY))
		}
	}
	time.Sleep(1 * time.Second)
	return resizedImg
}

// 模拟图像编码 (Memory 分配)
func encodeImage(img image.Image, format string) ([]byte, error) {
	buf := new(bytes.Buffer)
	var err error
	switch format {
	case "jpeg":
		err = jpeg.Encode(buf, img, nil)
	case "png":
		err = png.Encode(buf, img)
	default:
		return nil, fmt.Errorf("unsupported format: %s", format)
	}
	if err != nil {
		return nil, err
	}
	time.Sleep(1 * time.Second)
	return buf.Bytes(), nil
}

// 模拟图像存储 (Block 阻塞 - 锁竞争)
var storageMutex sync.Mutex
var storedImages [][]byte

func storeImage(imageData []byte) {
	storageMutex.Lock()
	defer storageMutex.Unlock()
	time.Sleep(time.Second) // 模拟 I/O 延迟
	storedImages = append(storedImages, imageData)               // 模拟内存增长
}

// 图像处理任务
func processImageTask(taskID int, width, height int, resizeFactor float64, format string) {
	startTime := time.Now()
	log.Printf("Task %d: Starting processing...\n", taskID)
	defer func() {
		log.Printf("Task %d: Processing finished in %v\n", taskID, time.Since(startTime))
	}()

	// 1. 生成图像 (CPU 轻负载)
	img := generateImage(width, height)

	// 2. 图像缩放 (CPU 密集型)
	resizedImg := resizeImage(img, resizeFactor)
	if resizedImg == nil {
		log.Printf("Task %d: Image resize failed, invalid dimensions.\n", taskID)
		return
	}

	// 3. 图像编码 (Memory 分配)
	imageData, err := encodeImage(resizedImg, format)
	if err != nil {
		log.Printf("Task %d: Image encode failed: %v\n", taskID, err)
		return
	}

	// 4. 图像存储 (Block 阻塞 - 锁竞争)
	storeImage(imageData)
	log.Printf("Task %d: Image stored, size: %d bytes\n", taskID, len(imageData))
}

func main() {
	runtime.SetBlockProfileRate(1)
	runtime.SetMutexProfileFraction(1)

	go func() {
		pprofAddress := ":6060"
		fmt.Println("启动 pprof 服务，监听端口：", pprofAddress)
		if err := http.ListenAndServe(pprofAddress, nil); err != nil {
			fmt.Println("pprof 服务启动失败:", err)
		}
	}()

	numTasks := 30
	var wg sync.WaitGroup

	fmt.Printf("Starting %d image processing tasks...\n", numTasks)
	for i := 0; i < numTasks; i++ {
		wg.Add(1)
		taskID := i + 1
		go func() {
			defer wg.Done()
			// 模拟不同任务参数
			width := rand.Intn(800) + 200   // 200-1000
			height := rand.Intn(600) + 150  // 150-750
			resizeFactor := 0.5 + rand.Float64()*1.5 // 0.5-2.0 倍缩放
			formats := []string{"jpeg", "png"}
			format := formats[rand.Intn(len(formats))]

			processImageTask(taskID, width, height, resizeFactor, format)
		}()
	}

	wg.Wait()
	fmt.Println("All image processing tasks completed.")
	fmt.Printf("Total stored images: %d, Total data size: %.2f MB\n", len(storedImages), float64(totalBytesStored())/(1024*1024))

	// 模拟持续运行一段时间，方便观察内存增长等
	select{}
}

func totalBytesStored() int {
	total := 0
	for _, imgData := range storedImages {
		total += len(imgData)
	}
	return total
}
