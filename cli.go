package kgo

import (
    "fmt"
    "strings"
)

type Func func(string,...interface{})string

// CliColorRender 命令行输出带颜色字体.
func CliColorRender(str string, color int, weight int, extraArgs ...interface{}) string {
    //闪烁效果
    isBlink := 0
    //if len(extraArgs) > 0 {
    //	isBlink, _ = toInt(extraArgs[0])
    //}
    //下划线效果
    isUnderLine := 0
    //if len(extraArgs) > 1 {
    //	isUnderLine, _ = extraArgs[1].ToInt()
    //}
    var mo []string
    if isBlink > 0 {
        mo = append(mo, "05")
    }
    if isUnderLine > 0 {
        mo = append(mo, "04")
    }
    if weight > 0 {
        mo = append(mo, fmt.Sprintf("%d", weight))
    }
    if len(mo) <= 0 {
        mo = append(mo, "0")
    }
    return fmt.Sprintf("\033[%s;%dm"+str+"\033[0m", strings.Join(mo, ";"), color)
}

// PrettyFileSize 把bit转化为更大的单位显示
func (kc *LKKCli) PrettyFileSize(fileSize int64) (size string) {
    if fileSize < 1024 {
        //return strconv.FormatInt(fileSize, 10) + "B"
        return fmt.Sprintf("%.2fB ", float64(fileSize)/float64(1))
    } else if fileSize < (1024 * 1024) {
        return fmt.Sprintf("%.2fKB ", float64(fileSize)/float64(1024))
    } else if fileSize < (1024 * 1024 * 1024) {
        return fmt.Sprintf("%.2fMB ", float64(fileSize)/float64(1024*1024))
    } else if fileSize < (1024 * 1024 * 1024 * 1024) {
        return fmt.Sprintf("%.2fGB ", float64(fileSize)/float64(1024*1024*1024))
    } else if fileSize < (1024 * 1024 * 1024 * 1024 * 1024) {
        return fmt.Sprintf("%.2fTB ", float64(fileSize)/float64(1024*1024*1024*1024))
    } else { //if fileSize < (1024 * 1024 * 1024 * 1024 * 1024 * 1024)
        return fmt.Sprintf("%.2fEB ", float64(fileSize)/float64(1024*1024*1024*1024*1024))
    }
}

func (kc *LKKCli) ColorWrapStr(str string,color CliColor) string{
    funcMap := make(map[CliColor]Func)
    funcMap[CLI_GREEN] = kc.Green
    funcMap[CLI_LIGHT_GREEN] = kc.LightGreen
    funcMap[CLI_CYAN] = kc.Cyan
    funcMap[CLI_LIGHT_CYAN] = kc.LightCyan
    funcMap[CLI_RED] = kc.Red
    funcMap[CLI_LIGHT_RED] = kc.LightRed
    funcMap[CLI_BLUE] = kc.Blue
    funcMap[CLI_LIGHT_BLUE] = kc.LightBlue
    funcMap[CLI_PURPLE] = kc.Purple
    funcMap[CLI_LIGHT_PURPLE] = kc.LightPurple
    funcMap[CLI_GRAY] = kc.Gray
    funcMap[CLI_LIGHT_GRAY] = kc.LightGray
    funcMap[CLI_YELLOW] = kc.Yellow
    funcMap[CLI_BLACK] = kc.Black
    funcMap[CLI_WHITE] = kc.White
    return funcMap[color](str)
}

func (kc *LKKCli) FormatStrWithColor(str string,args ...interface{}) string{
    colorArr := args[len(args)-1]
    cArr,ok:=colorArr.([]CliColor)
    if !ok{
        return fmt.Sprintf(str,args)
    }
    lc:=len(cArr)
    for i:=0;i<len(args)-1;i++{
        color:=cArr[lc-1]
        if i<=lc-1{
            color=cArr[i]
        }
        args[i]=kc.ColorWrapStr(str,color)
    }
    return fmt.Sprintf(str,args[0:len(args)-1])
}

// Green 绿色字体，modifier里，第一个控制闪烁，第二个控制下划线
func (kc *LKKCli) Green(str string, modifier ...interface{}) string {
    return CliColorRender(str, 32, 0, modifier...)
}

// LightGreen 淡绿
func (kc *LKKCli) LightGreen(str string, modifier ...interface{}) string {
    return CliColorRender(str, 32, 1, modifier...)
}

// Cyan 青色/蓝绿色
func (kc *LKKCli) Cyan(str string, modifier ...interface{}) string {
    return CliColorRender(str, 36, 0, modifier...)
}

// LightCyan 淡青色
func (kc *LKKCli) LightCyan(str string, modifier ...interface{}) string {
    return CliColorRender(str, 36, 1, modifier...)
}
// Red 红字体
func (kc *LKKCli) Red(str string, modifier ...interface{}) string {
    return CliColorRender(str, 31, 0, modifier...)
}

// LightRed 淡红色
func (kc *LKKCli) LightRed(str string, modifier ...interface{}) string {
    return CliColorRender(str, 31, 1, modifier...)
}

// Yellow 黄色字体
func (kc *LKKCli) Yellow(str string, modifier ...interface{}) string {
    return CliColorRender(str, 33, 0, modifier...)
}

// Black 黑色
func (kc *LKKCli) Black(str string, modifier ...interface{}) string {
    return CliColorRender(str, 30, 0, modifier...)
}

// Gray 深灰色
func (kc *LKKCli) Gray(str string, modifier ...interface{}) string {
    return CliColorRender(str, 30, 1, modifier...)
}

// LightGray 浅灰色
func (kc *LKKCli) LightGray(str string, modifier ...interface{}) string {
    return CliColorRender(str, 37, 0, modifier...)
}

// White 白色
func (kc *LKKCli) White(str string, modifier ...interface{}) string {
    return CliColorRender(str, 37, 1, modifier...)
}

// Blue 蓝色
func (kc *LKKCli) Blue(str string, modifier ...interface{}) string {
    return CliColorRender(str, 34, 0, modifier...)
}

// LightBlue 淡蓝
func (kc *LKKCli) LightBlue(str string, modifier ...interface{}) string {
    return CliColorRender(str, 34, 1, modifier...)
}

// Purple 紫色
func (kc *LKKCli) Purple(str string, modifier ...interface{}) string {
    return CliColorRender(str, 35, 0, modifier...)
}

// LightPurple 淡紫色
func (kc *LKKCli) LightPurple(str string, modifier ...interface{}) string {
    return CliColorRender(str, 35, 1, modifier...)
}

// Brown 棕色
func (kc *LKKCli) Brown(str string, modifier ...interface{}) string {
    return CliColorRender(str, 33, 0, modifier...)
}
