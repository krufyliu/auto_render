package auto_render

import (
  "io"
  "bytes"
  "encoding/binary"
  "auto_render/protocol"
  pb "github.com/golang/protobuf"
)

//以固定4字节长度(大端)分割的消息
func Fixed4Encode(message []byte) []byte, error {
  length := len(message)
  buffer := bytes.NewBuffer(make([]byte, 0, (length + 4)))
  if err := binary.Write(buffer, binary.BigEndian, length); err != nil {
    return nil, err
  }
  if _, err := buffer.Write(message); err != nil {
    return nil, err
  }
  return buffer, nil
}

//解码以固定4字节(大端)分割的message
func Fixed4Decode(r io.Reader) []byte, error {
  buffer := make([]byte, 0, 4)
  if _, err := r.Read(buffer); err != nil {
    return nil, err
  }
  var length int
  if err := binary.Read(bytes.NewReader(buffer), binary.BigEndian, &length); err != nil {
    return nil, err
  }
  buffer = make([]byte, 0, length)
  if _, err := r.Read(buffer); err != nil {
    return nil, err
  }
  return buffer, nil
}

func Fixed4EncodePb(message *protocol.Directive) []byte, error {
  buffer := pb.Marshal(message)
  return Fixed4Decode(buffer)
}

func Fixed4DecodePb(r io.Reader) *protocol.Directive, error {
  buffer := Fixed4Decode(r)
  directive := &protocol.Directive{}
  if err := pb.Unmarshal(buffer); err != nil {
    return nil, err
  }
  return directive, nil
} 
