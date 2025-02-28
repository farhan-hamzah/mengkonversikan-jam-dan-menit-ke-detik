package main
import "fmt"
func waktu(j int, m int, d int) int {
    return (j*3600) + (m*60) + d
}
func main(){
    var jam, menit, detik, hasil int
    fmt.Scan(&jam, &menit, &detik)
    hasil = waktu(jam, menit, detik)
    fmt.Print(hasil)
}