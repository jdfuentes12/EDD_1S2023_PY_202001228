class nodoArbol {
    constructor(nombre, carne,password,carpetaRaiz){
        this.izquierdo = null;
        this.derecho = null;
        this.carne = carne;
        this.nombre = nombre;
        this.password = password;
        this.carpetaRaiz = carpetaRaiz;
    }
}

class Arbol {
    constructor(){
        this.raiz = null;
    }
    
    insertarNodo(nodo, raiz){
        if(raiz === null){
            raiz = nodo;
        }else{
            if(raiz.carne === nodo.carne){
                raiz.carne = nodo.carne
            }else if(raiz.carne < nodo.carne){
                raiz.derecho = this.insertarNodo(nodo, raiz.derecho)
            }else{
                raiz.izquierdo = this.insertarNodo(nodo, raiz.izquierdo)
            }
        }
        return raiz;
    }

    insertar(objeto){
        console.log(objeto)
        this.raiz = this.insertarNodo(objeto,this.raiz)
    }

    recorridoPreorden(raiz){
        var cadena = ""
        if(raiz !== null){
            cadena = cadena + "\""
            cadena = cadena + raiz.carne
            cadena = cadena + "\""
            if(raiz.izquierdo !== null){
                cadena = cadena + " -> "
                cadena = cadena + this.recorridoPreorden(raiz.izquierdo)
            }
            if(raiz.derecho !== null){
                cadena = cadena + " -> "
                cadena = cadena + this.recorridoPreorden(raiz.derecho)
            }
        }
        return cadena
    }

    recorridoInorden(raiz){
        var cadena = ""
        if(raiz !== null){
            if(raiz.izquierdo !== null){
                cadena += this.recorridoInorden(raiz.izquierdo)
                cadena += " -> "
            }
            cadena += "\""
            cadena += raiz.carne
            cadena += "\""
            if(raiz.derecho !== null){
                cadena += " -> "
                cadena += this.recorridoInorden(raiz.derecho)
            }
        }
        return cadena
    }

    recorridoPostOrden(raiz){
        var cadena = ""
        if(raiz !== null){
            if(raiz.izquierdo !== null){
                cadena += this.recorridoPostOrden(raiz.izquierdo)
                cadena += " -> "
            }
            if(raiz.derecho !== null){
                cadena += this.recorridoPostOrden(raiz.derecho)
                cadena += " -> "
            }
            cadena += "\""
            cadena += raiz.carne
            cadena += "\""
        }
        return cadena
    }

    recorridoArbol(){
        console.log(this.recorridoInorden(this.raiz))
        console.log(this.recorridoPreOrden(this.raiz))
        console.log(this.recorridoPostOrden(this.raiz))
    }

    retornarcarneesArbol(raiz, id){
        var cadena = "";
        var numero = id + 1;
        if(raiz !== null){
            cadena += "\"";
            cadena += raiz.carne;
            cadena += "\" ;";
            if(!(raiz.izquierdo === null) && !(raiz.derecho === null)){
                cadena += " x" + numero + " [label=\"\",width=.1,style=invis];"
                cadena += "\"";
                cadena += raiz.carne;
                cadena += "\" -> ";
                cadena += this.retornarcarneesArbol(raiz.izquierdo, numero)
                cadena += "\"";
                cadena += raiz.carne;
                cadena += "\" -> ";
                cadena += this.retornarcarneesArbol(raiz.derecho, numero)
                cadena += "{rank=same" + "\"" + raiz.izquierdo.carne + "\"" + " -> " + "\"" + raiz.derecho.carne + "\""  + " [style=invis]}; "
            }else if(!(raiz.izquierdo === null) && (raiz.derecho === null)){
                cadena += " x" + numero + " [label=\"\",width=.1,style=invis];"
                cadena += "\"";
                cadena += raiz.carne;
                cadena += "\" -> ";
                cadena += this.retornarcarneesArbol(raiz.izquierdo, numero)
                cadena += "\"";
                cadena += raiz.carne;
                cadena += "\" -> ";
                cadena += "x" + numero + "[style=invis]";
                cadena += "{rank=same" + "\"" + raiz.izquierdo.carne + "\"" + " -> " + "x" + numero + " [style=invis]}; "
            }else if((raiz.izquierdo === null) && !(raiz.derecho === null)){
                cadena += " x" + numero + " [label=\"\",width=.1,style=invis];"
                cadena += "\"";
                cadena += raiz.carne;
                cadena += "\" -> ";
                cadena += "x" + numero + "[style=invis]";
                cadena += "; \"";
                cadena += raiz.carne;
                cadena += "\" -> ";
                cadena += this.retornarcarneesArbol(raiz.derecho, numero)
                cadena += "{rank=same" + " x" + numero + " -> \"" + raiz.derecho.carne + "\"" +  " [style=invis]}; "
            }
        }
        return cadena;
    }

    graficarArbol(){
        var cadena = ""
        if(this.raiz !== null){
            cadena += "digraph arbol {"
            cadena += this.retornarcarneesArbol(this.raiz, 0)
            cadena += "}"
        }
        return cadena
    }

    /** 
     * Contenido de graficar los diferentes recorridos del arbol
     */
    recorridosArbol(){
        console.log("Recorrido Pre-Orden")
        let url = 'https://quickchart.io/graphviz?graph=';
        let body = "digraph G { graph[label = \"Pre-Orden\" rankdir = LR labelloc = t]" + this.recorridoPreorden(this.raiz) + "}";
        $("#image1").attr("src", url + body);
        console.log("Recorrido In-Orden")
        body = "digraph G { graph[label = \"In-Orden\" rankdir = LR labelloc = t]" + this.recorridoInorden(this.raiz) + "}";
        $("#image2").attr("src", url + body);
        console.log("Recorrido Post-Orden")
        body = "digraph G { graph[label = \"Post-Orden\" rankdir = LR labelloc = t]" + this.recorridoPostOrden(this.raiz) + "}";
        $("#image3").attr("src", url + body);
    }

    eliminarTodo(){
        this.raiz = null;
    }

}

function agregarVarios(){
    let carne = document.getElementById("carne").value
    let carnees = carne.split(',')
    try{
        carnees.forEach(element => {
            arbolBinario.insertarcarne(element)
        });
    }catch(error){
        alert("Hubo un error")
    }
    refrescarArbol()
}

function agregarVariosNumeros(){
    let carne = document.getElementById("carne").value
    let carnees = carne.split(',')
    try{
        carnees.forEach(element => {
            arbolBinario.insertarcarne(parseInt(element))
        });
    }catch(error){
        alert("Hubo un error")
    }
    refrescarArbol();
}

function refrescarArbol(){
    let url = 'https://quickchart.io/graphviz?graph=';
    let body = arbolBinario.graficarArbol();
    $("#image").attr("src", url + body);
    document.getElementById("carne").value = "";
}

/**
 * Funcion para recorrer
 */
function recorrerArbol(){
    arbolBinario.recorridosArbol();
}

/**
 * Funcion para reiniciar el arbol
 */

function limpiar(){
    arbolBinario.eliminarTodo();
    let url = 'https://quickchart.io/graphviz?graph=digraph G { raiz }';
    $("#image").attr("src", url);
    document.getElementById("carne").value = "";
}