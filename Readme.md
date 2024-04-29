# UTF8 Gallery Folder Check
# Menu 

```go
fmt.Println("Seleccione una opción:")
fmt.Println("1. Buscar archivos con problemas de codificación UTF-8")
fmt.Println("2. Mostrar archivos con nombres largos")
fmt.Println("3. Mostrar archivos con espacios en el nombre")
fmt.Println("4. Mostrar directorios con espacios en el nombre")
fmt.Println("5. Contar archivos por extensión")
fmt.Println("6. Copiar archivos RAW a un directorio específico")
```


# Resultados (opcion 1)
<pre>
Archivos con caracteres no ASCII en el nombre:
D:\..\Pictures\Raiz\Imágenes_de_muestra.lnk
D:\..\Pictures\archivo_ñóarchivo_漢字.txt.txt
</pre>


# Resultados (opcion 2)
<pre>
Archivos con nombres más largos que 50 caracteres en D:\..\Pictures:
D:\..\DS_city_rain_tree_cyberpunk_0dfdda6a-73f0-46e8-8824-889fb69ec160.png
D:\..\DS_city_rain_tree_cyberpunk_be37bb73-3b3b-42d8-9474-f5f64c0bd97b.png
D:\..\DS_desolation_in_red_aefc811d-883c-4eb7-886f-1e0391457f34.png
D:\..\DS_futuristic_venice_city_5c1ab218-bcb7-4c0f-96dd-1c593c2266c0.png
D:\..\Screenshot_20230301_214052_com.android.gallery3d.png
D:\..\Screenshot_20230309_180208_com.google.android.googlequicksearchbox.png
</pre>

# Resultados (opcion 5)
<pre>
Cantidad de archivos por extensión:
.lnk: 3
.HTM: 6
.pp3: 30
.mp4: 618
.ini: 338
.THM: 53
.zip: 19
.dng: 720
.3gp: 50
.jfif: 1
.wmv: 23
.avi: 3
.MP4: 50
.png0: 4
.png: 292
.thm: 6
.html: 1
.psd: 7
.xcf: 1
.db: 274
.MOV: 4
.jpg-large: 1
.pdf: 2
.jpg: 10409
.tif: 4
.MPG: 69
.JPG: 10637
.gif: 14
.jpeg: 53
.info: 42
.AVI: 7
.RW2: 120
</pre>
