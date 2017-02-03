package main

import (
	"math"
	"image"
    "image/color"
    "image/draw"
    "image/jpeg"
    "os"
)


var (
	white color.Color = color.RGBA{255, 255, 255, 255}
	black color.Color = color.RGBA{0, 0, 0, 255}
	blue  color.Color = color.RGBA{0, 0, 255, 255}
)



func julia (x complex128) complex128{

	return x*x-complex(0.221, 0.713)
}

func toinfinityandbeyond (x complex128) uint16 {

	exitloop := 0

	var iterations uint16 = 0
	complexiteration := x

	for exitloop < 1 {


		complexiteration = julia(complexiteration)

		if math.Abs(real(complexiteration)) >= 2 || iterations > 9999 {
			exitloop = 1
		}
		iterations = iterations + 1	
	}

	return iterations

}

func main() {
	
	const h = 3000
	const w = 3000

	var indexjdummy float64 = -1
	var indexidummy float64 = -1

 	var twoD [h][w]complex128
 	var twoDmagnitude [h][w]uint16



 	for i := 0; i < h; i++ {
		
		if i == 0{
			indexidummy = indexidummy + 0
		} else{
			indexidummy = indexidummy + .001
		}	

		indexjdummy = -1
     
        for j := 0; j < w; j++ {
	
			if j == 0{
				indexjdummy = indexjdummy + 0
			} else{
				indexjdummy = indexjdummy + .001
			}        
		
			twoD[i][j] = complex(indexidummy, indexjdummy)
		
		}
	}


 	for i := 0; i < h; i++ {

        for j := 0; j < w; j++ {
	
			twoDmagnitude[i][j] = 100*toinfinityandbeyond(twoD[i][j])
	
			}


		}


	img := image.NewRGBA(image.Rect(0,0,h, w))
	
	draw.Draw(img, img.Bounds(), &image.Uniform{white}, image.ZP, draw.Src)

    for i := 0; i < h; i++ {

        for j := 0; j < w; j++ {

        	img.Set(i, j, color.Gray16{Y: twoDmagnitude[i][j]})
        	
		}
	}

	
	picture, _ := os.Create("julia.jpg")
	defer picture.Close()
	jpeg.Encode(picture, img, &jpeg.Options{jpeg.DefaultQuality}) 


}