package util

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func ReadLocalFile(filename string) ([]byte, error) {
	bytes, err := os.ReadFile(filename)
	return bytes, err
}
func ReadRemoteFile(fileUrl string) ([]byte, error) {
	imageResp, err := http.Get(fileUrl)
	if err != nil {
		return nil, err
	}

	bytes, err := io.ReadAll(imageResp.Body)

	return bytes, err
}

func SaveWaveFile(filename string, pcmData []byte, channels int, rate int, sampleWidth int) error {
	// 创建输出文件
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("创建文件失败: %v", err)
	}
	defer file.Close()

	// 计算音频数据长度（字节）
	dataLength := len(pcmData)
	// 计算文件总长度
	totalLength := 36 + dataLength

	// 写入 WAV 文件头
	err = writeWaveHeader(file, int32(totalLength), int16(channels), int32(rate), int16(sampleWidth))
	if err != nil {
		return fmt.Errorf("写入文件头失败: %v", err)
	}

	// 写入 PCM 数据
	_, err = file.Write(pcmData)
	if err != nil {
		return fmt.Errorf("写入音频数据失败: %v", err)
	}

	return nil
}

func writeWaveHeader(w io.Writer, totalLength int32, channels int16, sampleRate int32, sampleWidth int16) error {
	// 计算每秒字节数
	byteRate := sampleRate * int32(channels) * int32(sampleWidth)
	// 计算块对齐
	blockAlign := channels * sampleWidth
	// 计算位深度
	bitsPerSample := sampleWidth * 8

	// 定义 WAV 文件头结构
	header := []struct {
		Data   []byte
		Format string
		Value  interface{}
	}{
		{[]byte("RIFF"), "s4", nil},
		{nil, "i4", totalLength},
		{[]byte("WAVE"), "s4", nil},
		{[]byte("fmt "), "s4", nil},
		{nil, "i4", int32(16)},     // fmt 块大小
		{nil, "i2", int16(1)},      // PCM 格式
		{nil, "i2", channels},      // 通道数
		{nil, "i4", sampleRate},    // 采样率
		{nil, "i4", byteRate},      // 字节率
		{nil, "i2", blockAlign},    // 块对齐
		{nil, "i2", bitsPerSample}, // 位深度
		{[]byte("data"), "s4", nil},
		{nil, "i4", int32(len([]byte{}))}, // 数据长度（暂时为0，后续填充）
	}
	// 写入文件头
	for _, h := range header {
		if h.Data != nil {
			_, err := w.Write(h.Data)
			if err != nil {
				return err
			}
		} else {
			switch h.Format {
			case "i2":
				err := writeInt16(w, h.Value.(int16))
				if err != nil {
					return err
				}
			case "i4":
				err := writeInt32(w, h.Value.(int32))
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// writeInt16 以小端格式写入 int16
func writeInt16(w io.Writer, v int16) error {
	buf := []byte{
		byte(v & 0xff),
		byte((v >> 8) & 0xff),
	}
	_, err := w.Write(buf)
	return err
}

// writeInt32 以小端格式写入 int32
func writeInt32(w io.Writer, v int32) error {
	buf := []byte{
		byte(v & 0xff),
		byte((v >> 8) & 0xff),
		byte((v >> 16) & 0xff),
		byte((v >> 24) & 0xff),
	}
	_, err := w.Write(buf)
	return err
}
