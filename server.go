package main

import (
        "fmt"
        "net"
        "time"

        "google.golang.org/protobuf/proto"
        pb "github.com/MrTomkinson/BBMesh/pb"
)

func main() {
        listener, err := net.Listen("tcp", ":8080")
        if err != nil {
                fmt.Println("Error listening:", err.Error())
                return
        }
        defer listener.Close()

        fmt.Println("Server listening on :8080")

        for {
                conn, err := listener.Accept()
                if err != nil {
                        fmt.Println("Error accepting:", err.Error())
                        continue
                }
                go handleConnection(conn)
        }
}

func handleConnection(conn net.Conn) {
        defer conn.Close()
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

        if ping := msg.GetPing(); ping != nil {
                pong := &pb.BBMeshMessage{
                        Content: &pb.BBMeshMessage_Pong{
                                Pong: &pb.Pong{Timestamp: uint32(time.Now().Unix())},
                        },
                }

                pongBytes, err := proto.Marshal(pong)
                if err != nil {
                        fmt.Println("Error marshalling:", err.Error())
                        return
                }
                conn.Write(pongBytes)
        }

}
