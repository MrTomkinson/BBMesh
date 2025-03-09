package main

import (
        "fmt"
        "net"
        "time"

        "google.golang.org/protobuf/proto"
        pb "github.com/MrTomkinson/BBMesh/pb"
)

func main() {
        conn, err := net.Dial("tcp", "127.0.0.1:8080")
        if err != nil {
                fmt.Println("Error connecting:", err.Error())
                return
        }
        defer conn.Close()

        ping := &pb.BBMeshMessage{
                Content: &pb.BBMeshMessage_Ping{
                        Ping: &pb.Ping{Timestamp: uint32(time.Now().Unix())},
                },
        }

        pingBytes, err := proto.Marshal(ping)
        if err != nil {
                fmt.Println("Error marshalling:", err.Error())
                return
        }

        conn.Write(pingBytes)

        buffer := make([]byte, 1024)
        n, err := conn.Read(buffer)
        if err != nil {
                fmt.Println("Error reading:", err.Error())
                return
        }

        msg := &pb.BBMeshMessage{}
        err = proto.Unmarshal(buffer[:n], msg)
        if err != nil {
                fmt.Println("Error unmarshalling:", err.Error())
                return
        }

        if pong := msg.GetPong(); pong != nil {
                fmt.Println("Pong received:", pong.Timestamp)
        }

}
