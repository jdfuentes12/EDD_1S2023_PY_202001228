class Vertice {
    constructor(valor) {
        this.valor = valor;
        this.vecinos = [];
        this.visitado = false;
    }

    agregarVecino(vertice) {
        this.vecinos.push(vertice);
        vertice.vecinos.push(this);
    }
}

class Grafo {
    constructor() {
        this.vertices = [];
    }

    agregarVertice(valor) {
        const vertice = new Vertice(valor);
        this.vertices.push(vertice);
    }

    agregarArista(v1, v2) {
        const vertice1 = this.vertices.find((vertice) => vertice.valor === v1);
        const vertice2 = this.vertices.find((vertice) => vertice.valor === v2);
        vertice1.agregarVecino(vertice2);
    }

    _buscarProfundidad(vertice) {
        vertice.visitado = true;
        console.log(vertice.valor);

        for (const vecino of vertice.vecinos) {
            if (!vecino.visitado) {
                this._buscarProfundidad(vecino);
            }
        }
    }

    buscarProfundidad(valor) {
        const vertice = this.vertices.find((vertice) => vertice.valor === valor);
        this._buscarProfundidad(vertice);
    }
}

// Ejemplo de uso
const grafo = new Grafo();

grafo.agregarVertice(1);
grafo.agregarVertice(2);
grafo.agregarVertice(3);
grafo.agregarVertice(4);
grafo.agregarVertice(5);

grafo.agregarArista(1, 2);
grafo.agregarArista(1, 3);
grafo.agregarArista(2, 3);
grafo.agregarArista(2, 4);
grafo.agregarArista(3, 4);
grafo.agregarArista(4, 5);

grafo.buscarProfundidad(1);
