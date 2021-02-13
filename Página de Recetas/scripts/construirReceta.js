const construirReceta = () => {
    // Obtenemos los parametros de la url
    const parametros = window.location.search

    //Creamos la instancia de URLSearchParams
    const urlParams = new URLSearchParams(parametros);

    const name = urlParams.get('name')
    const desc = urlParams.get('desc')
    const img = urlParams.get('img')
    const ingre = urlParams.get('ingre')
    const proc = urlParams.get('proc')

    // console.log(name,desc,img,ingre,pro)

    const mainReceta = document.getElementById('main-receta');
    if (mainReceta) {
        mainReceta.innerHTML = `
        <section class="about-us">
        <div class="content-wrapper receta">
            <article class="about-item">
                <div class="image-container">
                    <img src="${img}" alt="foto de receta">
                </div>
                <div>
                    <h2>${name}</h2>
                    <p>${desc}</p>
                </div>
            </article>
            <article class="about-item">
                <div class="image-container">
                    <img src="/images/undraw_Chef_cu0r.svg" alt="logito">
                </div>
                <div id="ingredientes" class="list-ingredientes">
                    <h2>Ingredientes</h2>
                </div>
            </article>
            <article>
                <h2>Procedimiento</h2>
                <p>${proc}</p>
            </article>
        </div>
        </section>
        `
    }

    crearIngredientes(ingre)
}

const crearIngredientes = (ingre) => {
    const ingredientes = ingre.split(',')
    const listIngredientes = document.createElement('ul');
    const listIngreFragment = document.createDocumentFragment()

    for (const ingrediente of ingredientes) {
        const li = document.createElement('li')
        listIngreFragment.append(li)
        li.textContent = ingrediente
    }

    const divIngredientes = document.getElementById('ingredientes');
    listIngredientes.appendChild(listIngreFragment)
    divIngredientes.appendChild(listIngredientes)
}

construirReceta()