
function salir() {
    window.location.href = "index.html";
}

document.getElementById("admin").addEventListener('change', leerArchivo, false);

function leerArchivo(e) {
    var archivo = e.target.files[0];

    if (!archivo) {
        return;
    }
    var lector = new FileReader();
    lector.onload = function(e) {
        var contenido = e.target.result;
        console.log(contenido); 
        alert("Archivo cargado correctamente.");
            
    };
    lector.readAsText(archivo);
}

let arbolEstudiante = new Arbol();

function carga_Masiva(e) {
    e.preventDefault();
    arbolEstudiante = new Arbol();
    //e.preventDefault();
    const formData = new FormData(e.target);
    const form = Object.fromEntries(formData);
    console.log(form);
    try{
        let fr = new FileReader();
        fr.readAsText(form.inputFile);
        contador = 0;
        fr.onload = () => {
            var all_est = JSON.parse(fr.result).alumnos;
            var alumn = all_est;
            for(let i=0; i<alumn.length; i++) {
                let nombre = alumn[i].nombre;
                let carnet = alumn[i].carnet;
                let password = alumn[i].password;
                let Carpeta_Raiz = alumn[i].Carpeta_Raiz;
                let nuevo = new nodoArbol(nombre, carnet, password, Carpeta_Raiz);
                arbolEstudiante.insertar(nuevo);
                var fila 
                if(contador == 0){
                    fila = "<tr><td>"+carnet+"</td><td>"+nombre+"</td><td>"+password+"</td></tr>";
                    contador++;
                }else{
                    fila = fila + "<tr><td>"+carnet+"</td><td>"+nombre+"</td><td>"+password+"</td></tr>";
                }
                localStorage.setItem("tabla",fila)
                document.getElementById("tablaEstudiantes").innerHTML = fila;
            }
            
            alert("Carga masiva exitosa."); 
            localStorage.setItem("arbolEstudiante", JSON.stringify(arbolEstudiante))
        }
        
    }catch(error){
        console.log(error);
        alert("Error al cargar el archivo...");
    }
}

function acualizar(){
    var tabla = localStorage.getItem("tabla");
    document.getElementById("tablaEstudiantes").innerHTML = tabla;
}

function impresion(){
    arbolEstudiante.impresion();
}

function buscar(usuario, password){
    var validar = false;
    validar = arbolEstudiante.buscar(usuario, password);
    if (validar){
        return true;
    }
    return false;

}


