const viewReceta = () =>{
    const recetasRecientes = document.querySelectorAll(".reciente-item")
    const recetasDestacadas = document.querySelectorAll(".destacada-item")
    const recetasGeneral = document.querySelectorAll(".recetas-item")

    if (recetasRecientes) {
        recetasRecientes.forEach( e => {
            const nameReceta = e.querySelector('h3')
            const description = e.querySelector('p')
            const img = e.querySelector('img')
            const ingre = e.querySelector('.ingredientes')
            const proc = e.querySelector('.procedimiento')

            if (nameReceta && description && img && ingre && proc) {
                e.setAttribute('href', `pages/receta.html?name=${nameReceta.textContent}&desc=${description.textContent}&img=${img.src}&ingre=${ingre.textContent}&proc=${proc.textContent}`)
            }
        })
    }

    if (recetasDestacadas) {
        recetasDestacadas.forEach( e => {
            const nameReceta = e.querySelector('h3')
            const description = e.querySelector('p')
            const img = e.querySelector('img')
            const ingre = e.querySelector('.ingredientes')
            const proc = e.querySelector('.procedimiento')
            const link = e.querySelector('a')

            console.log(link)
            if (nameReceta && description && img && ingre && proc && link) {
                link.setAttribute('href', `pages/receta.html?name=${nameReceta.textContent}&desc=${description.textContent}&img=${img.src}&ingre=${ingre.textContent}&proc=${proc.textContent}`)
            }
        })
    }

    if (recetasGeneral) {
        recetasGeneral.forEach( e => {
            const nameReceta = e.querySelector('h3')
            const description = e.querySelector('p')
            const img = e.querySelector('img')
            const ingre = e.querySelector('.ingredientes')
            const proc = e.querySelector('.procedimiento')

            if (nameReceta && description && img && ingre && proc) {
                e.setAttribute('href', `receta.html?name=${nameReceta.textContent}&desc=${description.textContent}&img=${img.src}&ingre=${ingre.textContent}&proc=${proc.textContent}`)
            }
        })
    }
}

viewReceta()