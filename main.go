package main

import file "github.com/migueel15/TextEditor/File"

func main() {
	// file, err := file.OpenFile("/home/miguel/main.y")
	// if err != nil {
	// 	fmt.Printf("Error al abrir el archivo.", err)
	// }
	//
	// file.Buffer.Delete(0, 0, 1)
	//
	// fmt.Printf("file.Buffer: %v\n", file.Buffer)
	// file.Save()
	file := file.NewFile("/home/miguel/prueba.txt")
	file.Buffer.Append("Primera linea")
	file.Buffer.Append("Segunda linea")
	file.Save()

}
