package util

import (
	"bytes"
	"encoding/binary"
)

//int转[]byte
func IntToBytes(num int64)([]byte,error){
	//bytes 缓冲区
	buff := new (bytes.Buffer)
	//大端位序排列：binary.BigEndian
	//小端序列排列：binary.LittleEndian
	err :=binary.Write(buff,binary.BigEndian,num)
	if err != nil{
		return  nil,err
	}
	return buff.Bytes(),nil
}
//string转[]byte
func StringToBytes(st string)[]byte{
	return []byte(st)
}