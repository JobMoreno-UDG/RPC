package main

import (
	//"errors"
	"fmt"
	"net"
	"net/rpc"
	"strings"
	"strconv"
)
	//// mapa de materias
	var materias = make(map[string]map[string]float64)
	////  mapa alumnos
	var alumnos = make(map[string]map[string]float64)
type Server struct{}
type Error struct{
	msg string
}
func NewErrGOT(mensaje string) *Error {
    return &Error{
        msg: mensaje,
    }
}

func (e *Error) Error() string {
    return fmt.Sprintf("%s ", e.msg)
}
func (this *Server) Agregar(datos string, reply *int64)error{
	//// alumno - calificaion
	alumno := make(map[string]float64)
	//// materia - calificacion
	materia := make(map[string]float64)
	datosF := strings.Split(datos,"/")
	cont :=0
	nom := datosF[0]
	mat := datosF[1]
	al := 0
	cal, err := strconv.ParseFloat(datosF[2], 64)
	alumno[nom] = cal
	materia[mat] = cal
	if err != nil {
		fmt.Println(err)
	}
	for i,k := range alumnos{
		if i == nom{
			al = 1
			for i,_ := range k{
				if mat == i{
					return NewErrGOT("Error calificaion ya registrada!")
				}
			}
		}
	}
	for i,_ := range materias{
		if i == mat{
			cont = 1
			materias[mat][nom]=cal
		}
	}
	if cont == 0{		
		materias[mat] = alumno
	}
	if al == 0{
		alumnos[nom]= materia
	}else{
		alumnos[nom][mat]= cal
	} 
	
	fmt.Println(alumnos,materias)
	return nil
}
func (this *Server)Promedio_Alumno(nombre string, reply * int64) error {
	var prom,num float64
	num = 0.0
	prom = 0.0
	fmt.Println(nombre)
	for i,mat:=range alumnos{
		fmt.Println(nombre,i)
		if nombre == i{
			for _,cal:= range mat{
				prom += cal
				num +=1
			}
		}
	}
	return NewErrGOT(fmt.Sprintf("%f", prom/num))
}
func (this *Server)Promedio_General(x string, reply *int64) error  {
	var prom,num float64
	num = 0.0
	prom = 0.0
	for _,mat:=range alumnos{
		for _,cal:= range mat{
			prom += cal
			num +=1
		}
	}
	fmt.Println(fmt.Sprintf("%f", prom/num))
	return NewErrGOT(fmt.Sprintf("%f", prom/num))
}
func (this *Server)Promedio_Materia(materia string, reply *int64)  error{
	var prom,num float64
	num = 0.0
	prom = 0.0
	for i,al:=range materias{
		if i == materia{
			for _,cal:= range al{
				prom += cal
				num +=1
			}
		}
	}
	fmt.Println(fmt.Sprintf("%f", prom/num))
	return NewErrGOT(fmt.Sprintf("%f", prom/num))
}
func server() {

	rpc.Register(new(Server))
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go rpc.ServeConn(c)
	}
}

func main() {
	go server()

	var input string
	fmt.Scanln(&input)
}
