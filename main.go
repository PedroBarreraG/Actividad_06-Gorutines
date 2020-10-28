package main

import (
	"fmt"
	"time"
)

var banderas []int

func Proceso(id uint64, view *int) {
	i := uint64(0)
	for {
		if banderas[id] == 0 {
			return
		}
		if *view == 2 {
			fmt.Printf("id %d: %d\n", id, i)
		}
		i = i + 1
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	op := 0
	cont := 0
	for op != 4 {
		fmt.Print("1) Agregar proceso\n2) Mostrar procesos\n3) Terminar proceso\n4) Salir\n")
		fmt.Scan(&op)
		switch op {
		case 1:
			banderas = append(banderas, 1)
			id := uint64(cont)
			go Proceso(id, &op)
			cont++
		case 2:
			var opTem int
			fmt.Scan(&opTem)
			if opTem == 2 {
				op = 0
			}
		case 3:
			var idTer int
			fmt.Print("id del proceso a terminar: ")
			fmt.Scan(&idTer)
			if idTer < cont {
				banderas[idTer] = 0
				fmt.Printf("Proceso %d Terminado\n", idTer)
			}
		}
	}
}
