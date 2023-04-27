class nodoArbol {
    constructor(nombre, carne,password,carpetaRaiz){
        this.carne = carne;
        this.nombre = nombre;
        this.password = password;
        this.carpetaRaiz = carpetaRaiz;
    }
}

class Nodo{
    constructor(usuario){
        this.usuario = usuario;
        this.izquierdo = null;
        this.derecho = null;
        this.altura = 0;
    }
}

class Arbol {
    constructor(){
        this.raiz = null;
    }
    
    insertar(usuario){
        this.raiz = this.agregarUsuario(usuario, this.raiz);
    }

    agregarUsuario(usuario,raiz){
        if(raiz == null){
            return new Nodo(usuario);
        }
        if(usuario.carne < raiz.usuario.carne){
            raiz.izquierdo = this.agregarUsuario(usuario, raiz.izquierdo);
            if(this.altura(raiz.derecha)-this.altura(raiz.izquierda) == -2) {
                if(usuario.carnet < raiz.izquierda.usuario.carnet) {
                    raiz = this.reizq(raiz);
                } else {
                    raiz = this.rederizq(raiz);
                }
            }
        }else if(usuario.carne > raiz.usuario.carne){
            raiz.derecho = this.agregarUsuario(usuario, raiz.derecho);
            if(this.altura(raiz.derecha)-this.altura(raiz.izquierda) == 2) {
                if(usuario.carnet > raiz.derecha.usuario.carnet) {
                    raiz = this.reder(raiz);
                } else {
                    raiz = this.rederder(raiz);
                }
            }
        }else{
            raiz.usuario = usuario;
        }
        raiz.altura = this.maximo(this.altura(raiz.izquierda), this.altura(raiz.derecha))+1;
        return raiz;
    }   

    buscar(carne,password){
        var temp = this.raiz;
        while(temp != null){
            if(temp.usuario.carne == carne && temp.usuario.password == password){
                return true;
            }else if(carne < temp.usuario.carne){
                temp = temp.izquierdo;
            }else{
                temp = temp.derecho;
            }
        }
        return false;
    }

    maximo(a,b){
        if (a > b){
            return a;
        }
        return b;
    }

    altura(raiz){
        if(raiz == null){
            return -1;
        }
        return raiz.altura;
    }

    reizq(nodo) {
        var aux = nodo.izquierda;
        nodo.izquierda = aux.derecha;
        aux.izquierda = nodo;
        nodo.altura = this.maximo(this.altura(nodo.derecha), this.altura(nodo.izquierda))+1;
        aux.altura = this.maximo(this.altura(nodo.derecha), nodo.altura)+1;
        return aux;
    }

    reder(nodo) {
        var aux = nodo.derecha;
        nodo.derecha = aux.izquierda;
        aux.izquierda = nodo;
        nodo.altura = this.maximo(this.altura(nodo.derecha), this.altura(nodo.izquierda))+1;
        aux.altura = this.maximo(this.altura(nodo.derecha), nodo.altura)+1;
        return aux;
    }

    rederizq(nodo) {
        nodo.izquierda = this.reder(nodo);
        return this.reder(nodo);
    }

    rederder(nodo) {
        nodo.derecha = this.reizq(nodo);
        return this.reder(nodo);
    }
/* ------------ IN-ORDEN  ------------ */
    inOrden(nodo){
         if(nodo != null){
            this.inOrden(nodo.izquierdo);
            this.inOrden(nodo.derecho);
        }
    }

    5

/* ------------ PRE-ORDEN  ------------ */

}
