package main

import (
	"encoding/binary"
	"fmt"
)

func pack2uhex(size int, data interface{}) []byte {
	switch size {
	case 1:
		return []byte{uint8(data.(uint64))}
	case 2:
		buf := make([]byte, 2)
		binary.BigEndian.PutUint16(buf, uint16(data.(uint64)))
		return buf
	case 3:
		buf := make([]byte, 3)
		values := data.([]uint64)
		buf[0] = uint8(values[0] >> 8)
		buf[1] = uint8(values[0])
		buf[2] = uint8(values[1])
		return buf
	case 4:
		buf := make([]byte, 4)
		binary.BigEndian.PutUint32(buf, uint32(data.(uint64)))
		return buf
	case 5:
		buf := make([]byte, 5)
		values := data.([]uint64)
		binary.BigEndian.PutUint32(buf, uint32(values[0]))
		buf[4] = uint8(values[1])
		return buf
	case 6:
		buf := make([]byte, 6)
		values := data.([]uint64)
		binary.BigEndian.PutUint16(buf[2:], uint16(values[0]))
		buf[5] = byte(values[1])
		return buf
	case 7:
		buf := make([]byte, 7)
		values := data.([]uint64)
		binary.BigEndian.PutUint32(buf, uint32(values[0]))
		buf[5] = byte(values[1])
		buf[6] = byte(values[2])
		return buf
	case 8:
		buf := make([]byte, 8)
		binary.BigEndian.PutUint64(buf, data.(uint64))
		return buf
	case 9:
		buf := make([]byte, 9)
		values := data.([]uint64)
		binary.BigEndian.PutUint64(buf, values[0])
		buf[8] = uint8(values[1])
		return buf
	case 10:
		buf := make([]byte, 10)
		values := data.([]uint64)
		binary.BigEndian.PutUint64(buf, values[0])
		binary.BigEndian.PutUint16(buf[8:], uint16(values[1]))
		return buf
	case 11:
		buf := make([]byte, 11)
		values := data.([]uint64)
		binary.BigEndian.PutUint64(buf, values[0])
		binary.BigEndian.PutUint16(buf[8:], uint16(values[1]))
		buf[10] = uint8(values[2])
		return buf
	default:
		panic("please pack it yourself")
	}
}

func main() {
	//data1 := uint64(42)
	//data2 := uint64(1024)
	//data3 := []uint64{255, 42}
	//data4 := uint64(123456789)
	//data5 := []uint64{123456789, 255}
	//data6 := []uint64{65536, 42}
	//data7 := []uint64{65536, 42, 255}
	//data8 := uint64(12345678901234567890)
	//data9 := []uint64{12345678901234567890, 255}
	//data10 := []uint64{12345678901234567890, 65535}
	//data11 := []uint64{12345678901234567890, 65535, 255}

	data1 := uint64(99)
	data2 := uint64(2048)
	data3 := []uint64{128, 99}
	data4 := uint64(987654321)
	data5 := []uint64{987654321, 128}
	data6 := []uint64{8192, 99}
	data7 := []uint64{8192, 99, 128}
	data8 := uint64(9876543210123456789)
	data9 := []uint64{9876543210123456789, 128}
	data10 := []uint64{9876543210123456789, 32767}
	data11 := []uint64{9876543210123456789, 32767, 128}

	testData := map[int]interface{}{
		1:  data1,
		2:  data2,
		3:  data3,
		4:  data4,
		5:  data5,
		6:  data6,
		7:  data7,
		8:  data8,
		9:  data9,
		10: data10,
		11: data11,
	}

	for size := 1; size <= 11; size++ {
		data := testData[size]
		result := pack2uhex(size, data)
		fmt.Printf("size=%d, data=%v, result=%v\n", size, data, result)
	}
}
