let avlTree = new treeAVL();

function CargaMasiva(e) {
    avlTree = new treeAVL();
    e.preventDefault();
    const formData = new FormData(e.target);
    const form = Object.fromEntries(formData);

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
                let nuevo = new Persona(nombre, carnet, password, Carpeta_Raiz);
                avlTree.insert(nuevo);
            }

            $('#studentsTable tbody').html(
                all_est.map((item, index) => {
                    return(`
                        <tr>
                            <th>${item.carnet}</th>
                            <td>${item.nombre}</td>
                            <td>${item.password}</td>
                        </tr>
                    `);
                }).join('')
            )
            
            localStorage.setItem("avlTree", JSON.stringify(avlTree))
            alert('Archivo cargado correctamente...');
            //console.log(avlTree)
        }
    }catch(error){
        console.log(error);
        alert("Error al cargar el archivo...");
    }
}