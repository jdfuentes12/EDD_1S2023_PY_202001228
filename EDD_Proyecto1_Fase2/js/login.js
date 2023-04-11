const form = document.getElementById('login');

form.addEventListener('submit', (e) => {
    e.preventDefault();
    let username = document.getElementById('user').value;
    let password = document.getElementById('password').value;
    
    if (username === 'admin' && password === 'admin') {
        alert('Bienvenido Administrador' );
        window.location = 'admin.html';
    }else {
        alert('El usuario o la contraseña son incorrectos');
        
    }
});