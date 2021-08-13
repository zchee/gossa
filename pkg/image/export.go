// export by github.com/goplus/interp/cmd/qexp

package image

import (
	"image"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("image", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*image.Alpha).AlphaAt":             (*image.Alpha).AlphaAt,
	"(*image.Alpha).At":                  (*image.Alpha).At,
	"(*image.Alpha).Bounds":              (*image.Alpha).Bounds,
	"(*image.Alpha).ColorModel":          (*image.Alpha).ColorModel,
	"(*image.Alpha).Opaque":              (*image.Alpha).Opaque,
	"(*image.Alpha).PixOffset":           (*image.Alpha).PixOffset,
	"(*image.Alpha).Set":                 (*image.Alpha).Set,
	"(*image.Alpha).SetAlpha":            (*image.Alpha).SetAlpha,
	"(*image.Alpha).SubImage":            (*image.Alpha).SubImage,
	"(*image.Alpha16).Alpha16At":         (*image.Alpha16).Alpha16At,
	"(*image.Alpha16).At":                (*image.Alpha16).At,
	"(*image.Alpha16).Bounds":            (*image.Alpha16).Bounds,
	"(*image.Alpha16).ColorModel":        (*image.Alpha16).ColorModel,
	"(*image.Alpha16).Opaque":            (*image.Alpha16).Opaque,
	"(*image.Alpha16).PixOffset":         (*image.Alpha16).PixOffset,
	"(*image.Alpha16).Set":               (*image.Alpha16).Set,
	"(*image.Alpha16).SetAlpha16":        (*image.Alpha16).SetAlpha16,
	"(*image.Alpha16).SubImage":          (*image.Alpha16).SubImage,
	"(*image.CMYK).At":                   (*image.CMYK).At,
	"(*image.CMYK).Bounds":               (*image.CMYK).Bounds,
	"(*image.CMYK).CMYKAt":               (*image.CMYK).CMYKAt,
	"(*image.CMYK).ColorModel":           (*image.CMYK).ColorModel,
	"(*image.CMYK).Opaque":               (*image.CMYK).Opaque,
	"(*image.CMYK).PixOffset":            (*image.CMYK).PixOffset,
	"(*image.CMYK).Set":                  (*image.CMYK).Set,
	"(*image.CMYK).SetCMYK":              (*image.CMYK).SetCMYK,
	"(*image.CMYK).SubImage":             (*image.CMYK).SubImage,
	"(*image.Gray).At":                   (*image.Gray).At,
	"(*image.Gray).Bounds":               (*image.Gray).Bounds,
	"(*image.Gray).ColorModel":           (*image.Gray).ColorModel,
	"(*image.Gray).GrayAt":               (*image.Gray).GrayAt,
	"(*image.Gray).Opaque":               (*image.Gray).Opaque,
	"(*image.Gray).PixOffset":            (*image.Gray).PixOffset,
	"(*image.Gray).Set":                  (*image.Gray).Set,
	"(*image.Gray).SetGray":              (*image.Gray).SetGray,
	"(*image.Gray).SubImage":             (*image.Gray).SubImage,
	"(*image.Gray16).At":                 (*image.Gray16).At,
	"(*image.Gray16).Bounds":             (*image.Gray16).Bounds,
	"(*image.Gray16).ColorModel":         (*image.Gray16).ColorModel,
	"(*image.Gray16).Gray16At":           (*image.Gray16).Gray16At,
	"(*image.Gray16).Opaque":             (*image.Gray16).Opaque,
	"(*image.Gray16).PixOffset":          (*image.Gray16).PixOffset,
	"(*image.Gray16).Set":                (*image.Gray16).Set,
	"(*image.Gray16).SetGray16":          (*image.Gray16).SetGray16,
	"(*image.Gray16).SubImage":           (*image.Gray16).SubImage,
	"(*image.NRGBA).At":                  (*image.NRGBA).At,
	"(*image.NRGBA).Bounds":              (*image.NRGBA).Bounds,
	"(*image.NRGBA).ColorModel":          (*image.NRGBA).ColorModel,
	"(*image.NRGBA).NRGBAAt":             (*image.NRGBA).NRGBAAt,
	"(*image.NRGBA).Opaque":              (*image.NRGBA).Opaque,
	"(*image.NRGBA).PixOffset":           (*image.NRGBA).PixOffset,
	"(*image.NRGBA).Set":                 (*image.NRGBA).Set,
	"(*image.NRGBA).SetNRGBA":            (*image.NRGBA).SetNRGBA,
	"(*image.NRGBA).SubImage":            (*image.NRGBA).SubImage,
	"(*image.NRGBA64).At":                (*image.NRGBA64).At,
	"(*image.NRGBA64).Bounds":            (*image.NRGBA64).Bounds,
	"(*image.NRGBA64).ColorModel":        (*image.NRGBA64).ColorModel,
	"(*image.NRGBA64).NRGBA64At":         (*image.NRGBA64).NRGBA64At,
	"(*image.NRGBA64).Opaque":            (*image.NRGBA64).Opaque,
	"(*image.NRGBA64).PixOffset":         (*image.NRGBA64).PixOffset,
	"(*image.NRGBA64).Set":               (*image.NRGBA64).Set,
	"(*image.NRGBA64).SetNRGBA64":        (*image.NRGBA64).SetNRGBA64,
	"(*image.NRGBA64).SubImage":          (*image.NRGBA64).SubImage,
	"(*image.NYCbCrA).AOffset":           (*image.NYCbCrA).AOffset,
	"(*image.NYCbCrA).At":                (*image.NYCbCrA).At,
	"(*image.NYCbCrA).Bounds":            (*image.NYCbCrA).Bounds,
	"(*image.NYCbCrA).COffset":           (*image.NYCbCrA).COffset,
	"(*image.NYCbCrA).ColorModel":        (*image.NYCbCrA).ColorModel,
	"(*image.NYCbCrA).NYCbCrAAt":         (*image.NYCbCrA).NYCbCrAAt,
	"(*image.NYCbCrA).Opaque":            (*image.NYCbCrA).Opaque,
	"(*image.NYCbCrA).SubImage":          (*image.NYCbCrA).SubImage,
	"(*image.NYCbCrA).YCbCrAt":           (*image.NYCbCrA).YCbCrAt,
	"(*image.NYCbCrA).YOffset":           (*image.NYCbCrA).YOffset,
	"(*image.Paletted).At":               (*image.Paletted).At,
	"(*image.Paletted).Bounds":           (*image.Paletted).Bounds,
	"(*image.Paletted).ColorIndexAt":     (*image.Paletted).ColorIndexAt,
	"(*image.Paletted).ColorModel":       (*image.Paletted).ColorModel,
	"(*image.Paletted).Opaque":           (*image.Paletted).Opaque,
	"(*image.Paletted).PixOffset":        (*image.Paletted).PixOffset,
	"(*image.Paletted).Set":              (*image.Paletted).Set,
	"(*image.Paletted).SetColorIndex":    (*image.Paletted).SetColorIndex,
	"(*image.Paletted).SubImage":         (*image.Paletted).SubImage,
	"(*image.RGBA).At":                   (*image.RGBA).At,
	"(*image.RGBA).Bounds":               (*image.RGBA).Bounds,
	"(*image.RGBA).ColorModel":           (*image.RGBA).ColorModel,
	"(*image.RGBA).Opaque":               (*image.RGBA).Opaque,
	"(*image.RGBA).PixOffset":            (*image.RGBA).PixOffset,
	"(*image.RGBA).RGBAAt":               (*image.RGBA).RGBAAt,
	"(*image.RGBA).Set":                  (*image.RGBA).Set,
	"(*image.RGBA).SetRGBA":              (*image.RGBA).SetRGBA,
	"(*image.RGBA).SubImage":             (*image.RGBA).SubImage,
	"(*image.RGBA64).At":                 (*image.RGBA64).At,
	"(*image.RGBA64).Bounds":             (*image.RGBA64).Bounds,
	"(*image.RGBA64).ColorModel":         (*image.RGBA64).ColorModel,
	"(*image.RGBA64).Opaque":             (*image.RGBA64).Opaque,
	"(*image.RGBA64).PixOffset":          (*image.RGBA64).PixOffset,
	"(*image.RGBA64).RGBA64At":           (*image.RGBA64).RGBA64At,
	"(*image.RGBA64).Set":                (*image.RGBA64).Set,
	"(*image.RGBA64).SetRGBA64":          (*image.RGBA64).SetRGBA64,
	"(*image.RGBA64).SubImage":           (*image.RGBA64).SubImage,
	"(*image.Uniform).At":                (*image.Uniform).At,
	"(*image.Uniform).Bounds":            (*image.Uniform).Bounds,
	"(*image.Uniform).ColorModel":        (*image.Uniform).ColorModel,
	"(*image.Uniform).Convert":           (*image.Uniform).Convert,
	"(*image.Uniform).Opaque":            (*image.Uniform).Opaque,
	"(*image.Uniform).RGBA":              (*image.Uniform).RGBA,
	"(*image.YCbCr).At":                  (*image.YCbCr).At,
	"(*image.YCbCr).Bounds":              (*image.YCbCr).Bounds,
	"(*image.YCbCr).COffset":             (*image.YCbCr).COffset,
	"(*image.YCbCr).ColorModel":          (*image.YCbCr).ColorModel,
	"(*image.YCbCr).Opaque":              (*image.YCbCr).Opaque,
	"(*image.YCbCr).SubImage":            (*image.YCbCr).SubImage,
	"(*image.YCbCr).YCbCrAt":             (*image.YCbCr).YCbCrAt,
	"(*image.YCbCr).YOffset":             (*image.YCbCr).YOffset,
	"(image.Image).At":                   (image.Image).At,
	"(image.Image).Bounds":               (image.Image).Bounds,
	"(image.Image).ColorModel":           (image.Image).ColorModel,
	"(image.PalettedImage).At":           (image.PalettedImage).At,
	"(image.PalettedImage).Bounds":       (image.PalettedImage).Bounds,
	"(image.PalettedImage).ColorIndexAt": (image.PalettedImage).ColorIndexAt,
	"(image.PalettedImage).ColorModel":   (image.PalettedImage).ColorModel,
	"(image.Point).Add":                  (image.Point).Add,
	"(image.Point).Div":                  (image.Point).Div,
	"(image.Point).Eq":                   (image.Point).Eq,
	"(image.Point).In":                   (image.Point).In,
	"(image.Point).Mod":                  (image.Point).Mod,
	"(image.Point).Mul":                  (image.Point).Mul,
	"(image.Point).String":               (image.Point).String,
	"(image.Point).Sub":                  (image.Point).Sub,
	"(image.Rectangle).Add":              (image.Rectangle).Add,
	"(image.Rectangle).At":               (image.Rectangle).At,
	"(image.Rectangle).Bounds":           (image.Rectangle).Bounds,
	"(image.Rectangle).Canon":            (image.Rectangle).Canon,
	"(image.Rectangle).ColorModel":       (image.Rectangle).ColorModel,
	"(image.Rectangle).Dx":               (image.Rectangle).Dx,
	"(image.Rectangle).Dy":               (image.Rectangle).Dy,
	"(image.Rectangle).Empty":            (image.Rectangle).Empty,
	"(image.Rectangle).Eq":               (image.Rectangle).Eq,
	"(image.Rectangle).In":               (image.Rectangle).In,
	"(image.Rectangle).Inset":            (image.Rectangle).Inset,
	"(image.Rectangle).Intersect":        (image.Rectangle).Intersect,
	"(image.Rectangle).Overlaps":         (image.Rectangle).Overlaps,
	"(image.Rectangle).Size":             (image.Rectangle).Size,
	"(image.Rectangle).String":           (image.Rectangle).String,
	"(image.Rectangle).Sub":              (image.Rectangle).Sub,
	"(image.Rectangle).Union":            (image.Rectangle).Union,
	"(image.YCbCrSubsampleRatio).String": (image.YCbCrSubsampleRatio).String,
	"image.Black":                        &image.Black,
	"image.Decode":                       image.Decode,
	"image.DecodeConfig":                 image.DecodeConfig,
	"image.ErrFormat":                    &image.ErrFormat,
	"image.NewAlpha":                     image.NewAlpha,
	"image.NewAlpha16":                   image.NewAlpha16,
	"image.NewCMYK":                      image.NewCMYK,
	"image.NewGray":                      image.NewGray,
	"image.NewGray16":                    image.NewGray16,
	"image.NewNRGBA":                     image.NewNRGBA,
	"image.NewNRGBA64":                   image.NewNRGBA64,
	"image.NewNYCbCrA":                   image.NewNYCbCrA,
	"image.NewPaletted":                  image.NewPaletted,
	"image.NewRGBA":                      image.NewRGBA,
	"image.NewRGBA64":                    image.NewRGBA64,
	"image.NewUniform":                   image.NewUniform,
	"image.NewYCbCr":                     image.NewYCbCr,
	"image.Opaque":                       &image.Opaque,
	"image.Pt":                           image.Pt,
	"image.Rect":                         image.Rect,
	"image.RegisterFormat":               image.RegisterFormat,
	"image.Transparent":                  &image.Transparent,
	"image.White":                        &image.White,
	"image.ZP":                           &image.ZP,
	"image.ZR":                           &image.ZR,
}

var typList = []interface{}{
	(*image.Alpha)(nil),
	(*image.Alpha16)(nil),
	(*image.CMYK)(nil),
	(*image.Config)(nil),
	(*image.Gray)(nil),
	(*image.Gray16)(nil),
	(*image.Image)(nil),
	(*image.NRGBA)(nil),
	(*image.NRGBA64)(nil),
	(*image.NYCbCrA)(nil),
	(*image.Paletted)(nil),
	(*image.PalettedImage)(nil),
	(*image.Point)(nil),
	(*image.RGBA)(nil),
	(*image.RGBA64)(nil),
	(*image.Rectangle)(nil),
	(*image.Uniform)(nil),
	(*image.YCbCr)(nil),
	(*image.YCbCrSubsampleRatio)(nil),
}