//import {ArbolAVL} from "./Avl.js";
//import {Arbol} from "./Arbol.js";
let est_archivo;
let archivo_Estudiantes;


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
            
    };
    lector.readAsText(archivo);
}

let arbolEstudiante = new Arbol();

function carga_Masiva(e) {
    arbolEstudiante = new Arbol();
    //e.preventDefault();
    const formData = new FormData(e.target);
    const form = Object.fromEntries(formData);
    console.log(form);
    try{
        let fr = new FileReader();
        fr.readAsText(form.inputFile);
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
            }
        localStorage.setItem("arbolEstudiante", JSON.stringify(arbolEstudiante))
        }
    }catch(error){
        console.log(error);
        alert("Error al cargar el archivo...");
    }
}
