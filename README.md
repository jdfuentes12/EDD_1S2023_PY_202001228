# EDD_1S2023_PY_202001228
### Primera Fase del Proyecto de Estructura de Datos
### Ing. Luis Espino 
### Aux. Cristian Suy
### Jose Daniel  Fuentes Orozco<br>202001228

# Manual  Técnico
### Informaciond del Programa
El programa esta realizado con el lenguaje de Go en su version "go1.20.1", con la herramienta de "graphviz" en su version "7.1.0".

### Estructuras utilizadas
* Pila
* Cola
* Lista Doble

### Modulos usados
En el proyecto se usarion varios modulos los cuales contendra diferentes metodos, los cuales se usaron para realizar una instrucción especifica. Por ejemplo el de iniciar sesión, la carga masiva, etc. <br>
A demas se uno una carpeta llamada "Estructura" la cual contiene los nodos y constructores de cada pila y lista usada en la elaboracion del proyecto.
<br>
Para poder acceder a la carpetad de estructuras se uso un comando para que se puedan implementar en otodo el proyecto y ademas no de error a la hora de compilarlo. El comando es "go mod init (y el nombre que se le quiera colocar a la estructura)" 
<br>
<img align='center' src="https://user-images.githubusercontent.com/88565998/222049210-170be3aa-bed2-4d10-b0f4-1a1a02b284a3.png" width="230"><br>



### Main
A continucacion se le presentará una imagen con las importaciones necesarias para poder ejecutarse en el main correctamente. <br>
<img align='center' src="https://user-images.githubusercontent.com/88565998/222045149-f0c93349-8911-42c2-8797-9892f8f97766.png" width="230"><br>
Se inicializara una variable con nombre de "listaCola" con la cual se podra acceder a la cola de todos los estudiantes. Para que no se tenga ningun inconveniente en la manipulacion de datos. <br>
Por medio de la importacion de "encoding/csv" se pudo realizar la lectura del archivo con extension de "csv" el cual fue establecido para el proyecto. Despues de cada una de las lecturas del archivo se pocede a agregarlos a la lista de espera, la cual despues se le dara el acceso para que se puedan registrar en el sistema de "EDD - GoDrive"<br>
Pequeño fracmento de codigo para la lectura del csv, por estetica no se agregara todo el codigo.<br>

<img align='center' src="https://user-images.githubusercontent.com/88565998/222049794-90043419-f28e-4c1f-b0f3-84e0892a7e1c.png" width="230"><br>



### Agregar estudiante al sistema
Por medio de una cola, se establecen todos los usuarios que son candidatos para agregarlos al sistema de "EDD -GoDrive" con el cual se aceptara o rechazara el estudiante, si se rechaza automaticamente se eliminara del sistema de cola y no se agregara al sistema, pero si se agrega tambien se eliminara de la cola pero se agregara al sistema por medio de una lista doble. <br>
Pequeño fracmento de codigo para agregar los estudiantes al sistema de "EDD - GoDrive"<br>

<img align='center' src="https://user-images.githubusercontent.com/88565998/222050462-bec30caf-ebeb-43b8-955e-088c641cd7f4.png" width="500"><br>
Pequeño fracmento de codigo para rechazar los estudiantes del sistema definitivamente<br>
<img align='center' src="https://user-images.githubusercontent.com/88565998/222050723-d20d55e0-e14e-4ccc-a8c1-aa0c1a333360.png" width="350"><br>




De igual manera se podra observar por medio de reportes la cola de estudiantes que aun no se han aceptado. <br>
<img align='center' src="https://user-images.githubusercontent.com/88565998/222048047-273014bf-d57d-4bbe-ac74-f9bba0d2e752.png" width="500"><br>

De igual manera cuando se acepte o rechace algun estudiante aparecera en una pila la cual contendra la hora que se rechazo o acepto el estudiante.<br>



Tambien se podra observar las lista de los estudiantes que fueron asigandos al sistema.<br>

<img align='center' src="https://user-images.githubusercontent.com/88565998/222051698-602eb33c-646b-489d-93a5-eaea96faee1a.png" width="750"><br>

### Inision de sesion
En esta parte despues de ser aceptado el estudiante en el sistema podra iniciar sesion por medio de su carne y password. Se mostrala las validadciones que se hicieorn para verificar si las credenciales ingresadas son correctas o no.<br>
<img align='center' src="https://user-images.githubusercontent.com/88565998/222052856-bebe42b7-481f-4fe7-93dd-b319cdd1e4a4.png" width="500"><br>

### Ordenamiento
Tambien se tiene ordenan los estudiantes por el numero de carne que tienen.
<img align='center' src="https://user-images.githubusercontent.com/88565998/222053129-4e7ab847-67be-4b13-8f97-f7dd21e18837.png" width="500"><br>

### Graficar 
Tambien por medio de la herramienta de graphiz podemos visualizar los usuarios que se encuentran en la pila, cola y la lista doble.<br>
<img align='center' src="https://user-images.githubusercontent.com/88565998/222053550-dfd0d964-2ad5-4c6c-bcd9-40c41100d038.png" width="500"><br>

### JSON
Tambien se escribio todos los datos aceptados en un archivo json. <br>
<img align='center' src="https://user-images.githubusercontent.com/88565998/222053829-a20e58bd-a8cd-4484-8473-c756ab1ae814.png" width="350"><br>


