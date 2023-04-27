const arbol = localStorage.getItem("arbolEstudiante");
const avl = JSON.parse(arbol);
const tabla = localStorage.getItem("tabla");

function login() {
    const form = document.getElementById('login');
    form.addEventListener('submit', (e) => {
        e.preventDefault();
        let username = document.getElementById('user').value;
        let password = document.getElementById('password').value;
        
        if (username == 'admin' && password == 'admin'){
          alert('Bienvenido Administrador');
          window.location = 'admin.html';
        }else if (avl !== null){
            try{
              const estudiante = buscarEstudiante(avl.raiz, parseInt(username), password);
              if (estudiante != null){
                  alert('Bienvenido ' + estudiante.usuario.nombre);
                  window.location = 'estudiante.html';
              }else{
                  alert('El usuario o la contraseña son incorrectos');
              }
            }catch (e){
              alert('El usuario o la contraseña son incorrectos');
            } 
        }else{  
            alert('El usuario o la contraseña son incorrectos');
        }
    });
}

function buscarEstudiante(avl, carnet, password) {
    if (avl === null) {
      return null;
    } else if (carnet  < avl.usuario.carne) {
      return buscarEstudiante(avl.izquierda, carnet, password);
    } else if (carnet > avl.usuario.carne) {
      return buscarEstudiante(avl.derecho, carnet, password);
    } else if (password !== avl.usuario.password) {
      return null;
    } else {
      return avl;
    }
  }