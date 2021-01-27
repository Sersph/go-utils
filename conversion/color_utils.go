package conversion

import (
	"fmt"
	"image/color"
	"strconv"
	"strings"
)

// 将 16 位颜色转换为 RGBA 对象
// 接收一个标准的 #开头 16 位颜色值 #FFFFFF(白色) return RGBA{R: 255, G: 255, B: 255, A: 255}
func Color16BitToRGBA(colorCode string) color.RGBA {
	// 头部是否是 # 开始
	flag := strings.Index(colorCode, "#") == 0
	colorLan := len(colorCode) - 1
	// 颜色长度必须符合 6 8 的长度
	flag = flag && (colorLan == 6 || colorLan == 8)

	if !flag {
		panic("请输入标准的16位色值！")
	}

	colorCode = colorCode[1:]
	rgbaColorCode := make([]uint8, 0)
	for len(colorCode) != 0 {
		item := colorCode[0:2]
		if len(colorCode) > 2 {
			colorCode = colorCode[2:]
		} else {
			colorCode = ""
		}
		itemNumber, _ := strconv.ParseInt(item, 16, 32)
		rgbaColorCode = append(rgbaColorCode, uint8(itemNumber))
	}

	if len(rgbaColorCode) == 3 {
		rgbaColorCode = append(rgbaColorCode, 255)
	}

	return color.RGBA{
		R: rgbaColorCode[0],
		G: rgbaColorCode[1],
		B: rgbaColorCode[2],
		A: rgbaColorCode[3],
	}
}

func RGBAToColor16Bit(rgba color.RGBA) string {
	color := fmt.Sprintf("#%X%X%X", rgba.R, rgba.G, rgba.B)
	if rgba.A != 255 {
		color += fmt.Sprintf("%X", rgba.A)
	}
	return color
}
