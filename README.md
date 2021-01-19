## Fax image processing

* ### ToHalftone

 Through the error diffusion algorithm, more or less similar images are generated in order to
Try to restore visually matched two-color images。


* ### Converted to 1720, 2280 size fax bitmap。
   
 ```
func ToFaxImg(AImage image.Image) *image.Gray {
	m0 := Resize(AImage, 1720, 2280)
	m1 := ToGray(m0)
	m2 := ToHalftone(m1)
	return m2
}
```