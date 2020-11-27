package main

import (
	"fmt"
	"net/rpc"
	"bufio"
	"os"
)
func cliente()  {
	c,err:=rpc.Dial("tcp","127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	var op, result int64
	var nombre, materia,calificacion string
	scanner := bufio.NewScanner(os.Stdin)
	for{
		fmt.Println("\t- Menu -")
		fmt.Println("1.- Agregar Calificación")
		fmt.Println("2.- Mostrar Promedio Alumno")
		fmt.Println("3.- Promedio General")
		fmt.Println("4.- Promedio Materia")
		fmt.Println("5.- Salir")
		fmt.Scan(&op)
		if op == 5{
			break
		}else if op ==1{
			fmt.Println("Agregar Nombre Alumno")
			scanner.Scan()
			scanner.Scan()
			nombre = scanner.Text()
			fmt.Println("Agregar Nombre Materia")
			scanner.Scan()
			materia = scanner.Text()
			fmt.Println("Agregar Calificación Alumno")
			fmt.Scan(&calificacion)
			final := nombre+"/"+materia+"/"+calificacion
			err = c.Call("Server.Agregar", final,&result)
			if err != nil {
				fmt.Println(err)
			}
		}else if op ==2{
			fmt.Println("Agregar Nombre Alumno")
			scanner.Scan()
			scanner.Scan()
			nombre = scanner.Text()
			err = c.Call("Server.Promedio_Alumno", nombre,&result)
			if err != nil {
				fmt.Println("El Promedio es: ",err)
			}
			
		}else if op ==3{
			err = c.Call("Server.Promedio_General", "0",&result)
			if err != nil {
				fmt.Println("El Promedio Genereal es: ",err)
			}
			
		}else if op ==4{
			fmt.Println("Agregar Nombre Materia")
			scanner.Scan()
			scanner.Scan()
			materia = scanner.Text()
			err = c.Call("Server.Promedio_Materia", materia,&result)
			if err != nil {
				fmt.Println("El Promedio por la materia",materia," es: ",err)
			}
			
		}
	}
}
func main()  {
	cliente()
}