const navBar = document.getElementById('top-navbar')
const missionHeader = document.getElementById('mission-header')
const missionContent = document.getElementById('mission-content')
const catDiv = document.getElementById('cat-section')
const bloodlineDiv = document.getElementById('bloodline-section')
const breederDiv = document.getElementById('breeder-section')
const contactDiv = document.getElementById('contact-section')

const divsToCheck = [
    missionHeader,
    missionContent,
    catDiv,
    bloodlineDiv,
    breederDiv
]

function touching(element1, element2) {
    const rect1 = element1.getBoundingClientRect()
    const rect2 = element2.getBoundingClientRect()
    return !(
        rect1.right < rect2.left ||
        rect1.left > rect2.right ||
        rect1.bottom < rect2.top ||
        rect1.top > rect2.bottom
    )
}

window.addEventListener('scroll', () => {
    const middleDetected = divsToCheck.some(div => touching(navBar, div))
    let bottomDetected = touching(navBar, contactDiv)
    if (bottomDetected) {
        console.log('bottom')
        navBar.classList.remove('top')
        navBar.classList.remove('middle')
        navBar.classList.add('bottom')
    } else if (middleDetected) {
        console.log('middle')
        navBar.classList.remove('top')
        navBar.classList.remove('bottom')
        navBar.classList.add('middle')
    } else {
        console.log('top')
        navBar.classList.remove('bottom')
        navBar.classList.remove('middle')
        navBar.classList.add('top')
    }
})
