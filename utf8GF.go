package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func contieneCaracteresNoASCII(nombreArchivo string) bool {
	for _, char := range nombreArchivo {
		if int(char) > 127 {
			return true
		}
	}
	return false
}

func buscarArchivosNoASCII(directorio, archivoLog string) {
	fmt.Println("\nArchivos con caracteres no ASCII en el nombre:")

	logFile, err := os.Create(archivoLog)
	if err != nil {
		fmt.Println("Error al crear el archivo de registro:", err)
		return
	}
	defer logFile.Close()

	err = filepath.Walk(directorio, func(rutaCompleta string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		nombreArchivo := info.Name()
		if contieneCaracteresNoASCII(nombreArchivo) {
			fmt.Println(rutaCompleta)
			logFile.WriteString(fmt.Sprintf("%s\n", rutaCompleta))
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error al buscar archivos:", err)
	}
}

func mostrarArchivosLargos(directorio string, longitudMinima int) {
	fmt.Printf("\nArchivos con nombres más largos que %d caracteres en %s:\n", longitudMinima, directorio)

	err := filepath.Walk(directorio, func(ruta string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error al acceder a %s: %v\n", ruta, err)
			return nil
		}

		if !info.IsDir() && len(info.Name()) > longitudMinima {
			fmt.Println(ruta)
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error al caminar por el directorio:", err)
	}
}

func mostrarArchivosConEspacios(directorio string) {
	fmt.Println("\nArchivos que contienen espacios en su nombre:")

	err := filepath.Walk(directorio, func(ruta string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error al acceder a %s: %v\n", ruta, err)
			return nil
		}

		if !info.IsDir() && strings.Contains(info.Name(), " ") {
			fmt.Println(ruta)
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error al caminar por el directorio:", err)
	}
}

func mostrarDirectoriosConEspacios(directorio string) {
	fmt.Println("\nDirectorios que contienen espacios en su nombre:")

	err := filepath.Walk(directorio, func(rutaCompleta string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && strings.Contains(info.Name(), " ") {
			fmt.Println(rutaCompleta)
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error al buscar directorios:", err)
	}
}

func contarArchivosPorExtension(directorio string) {
	fmt.Println("\nCantidad de archivos por extensión:")

	extensiones := make(map[string]int)

	err := filepath.Walk(directorio, func(rutaCompleta string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			ext := filepath.Ext(info.Name())
			extensiones[ext]++
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error al contar archivos por extensión:", err)
		return
	}

	for ext, cantidad := range extensiones {
		fmt.Printf("%s: %d\n", ext, cantidad)
	}
}

func copiarArchivosRAW(directorioOrigen, directorioDestino string) {
	fmt.Println("\nCopiando archivos RAW a un directorio específico:")

	extensionesRAW := map[string]bool{".rw2": true, ".dng": true}

	err := filepath.Walk(directorioOrigen, func(rutaCompleta string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			ext := filepath.Ext(info.Name())
			if extensionesRAW[ext] {
				rutaDestino := filepath.Join(directorioDestino, info.Name())
				err := copiarArchivo(rutaCompleta, rutaDestino)
				if err != nil {
					fmt.Printf("Error al copiar %s: %v\n", info.Name(), err)
				} else {
					fmt.Printf("Copiado %s a %s\n", rutaCompleta, rutaDestino)
				}
			}
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error al copiar archivos RAW:", err)
	}
}

func copiarArchivo(origen, destino string) error {
	contenido, err := os.ReadFile(origen)
	if err != nil {
		return err
	}

	err = os.WriteFile(destino, contenido, 0644)
	if err != nil {
		return err
	}

	return nil
}

func menu() {
	fmt.Println("Seleccione una opción:")
	fmt.Println("1. Buscar archivos con problemas de codificación UTF-8")
	fmt.Println("2. Mostrar archivos con nombres largos")
	fmt.Println("3. Mostrar archivos con espacios en el nombre")
	fmt.Println("4. Mostrar directorios con espacios en el nombre")
	fmt.Println("5. Contar archivos por extensión")
	fmt.Println("6. Copiar archivos RAW a un directorio específico")
	fmt.Println("0. Salir")
}

func main() {
	directorioABuscar := "D:\\Pictures"
	directorioRAWDestino := "D:\\Pictures\\RAW"
	archivoLogResultados := "D:\\Pictures\\log_resultados.txt"

	for {
		menu()
		fmt.Print("Ingrese el número de la opción deseada (0-6): ")
		var opcion string
		fmt.Scanln(&opcion)

		switch opcion {
		case "0":
			fmt.Println("Saliendo del programa.")
			return
		case "1":
			buscarArchivosNoASCII(directorioABuscar, archivoLogResultados)
		case "2":
			mostrarArchivosLargos(directorioABuscar, 50)
		case "3":
			mostrarArchivosConEspacios(directorioABuscar)
		case "4":
			mostrarDirectoriosConEspacios(directorioABuscar)
		case "5":
			contarArchivosPorExtension(directorioABuscar)
		case "6":
			copiarArchivosRAW(directorioABuscar, directorioRAWDestino)
			fmt.Println("Archivos RAW copiados.")
		default:
			fmt.Println("Opción no válida. Ingrese un número del 0 al 6.")
		}
	}
}
