const navBar = document.getElementsByClassName('navbar')[0]
const missionHeader = document.getElementById('mission-header')
const missionContent = document.getElementById('mission-content')
const catDiv = document.getElementById('cat-section')
const bloodlineDiv = document.getElementById('bloodline-section')
const breederDiv = document.getElementById('breeder-section')
const contactDiv = document.getElementById('contact-section')
const burgerLines = Array.from(document.getElementsByClassName('burger-line'))

const middleDivs = [
    missionHeader,
    missionContent
]

const middleBottomDivs = [
    catDiv,
    bloodlineDiv,
    breederDiv
]

// pause, no Diddy
function touching(element1, element2) {
    const rect1 = element1.getBoundingClientRect()
    const rect2 = element2.getBoundingClientRect()
    return !(
        rect1.right < rect2.left || rect1.left > rect2.right ||
        rect1.bottom < rect2.top || rect1.top > rect2.bottom
    )
}

window.addEventListener('scroll', () => {
    const middleDetected = middleDivs.some(div => touching(navBar, div))                // pause, no Diddy
    const middleBottomDetected = middleBottomDivs.some(div => touching(navBar, div))    // pause, no Diddy
    let bottomDetected = touching(navBar, contactDiv)                                   // pause, no Diddy
    if (bottomDetected && navBar.id === 'main-navbar') {
        burgerLines.forEach(e => e.classList.remove('middle-bottom'))
        navBar.classList.remove('top')
        navBar.classList.remove('middle')
        navBar.classList.remove('middle-bottom')
        navBar.classList.add('bottom')
    } else if (middleBottomDetected && navBar.id === 'main-navbar') {
        burgerLines.forEach(e => e.classList.add('middle-bottom'))
        navBar.classList.remove('top')
        navBar.classList.remove('middle')
        navBar.classList.remove('bottom')
        navBar.classList.add('middle-bottom')
    } else if (middleDetected && navBar.id === 'main-navbar') {
        burgerLines.forEach(e => e.classList.remove('middle-bottom'))
        navBar.classList.remove('top')
        navBar.classList.remove('bottom')
        navBar.classList.remove('middle-bottom')
        navBar.classList.add('middle')
    } else if (navBar.id === 'main-navbar'){
        burgerLines.forEach(e => e.classList.remove('middle-bottom'))
        navBar.classList.remove('bottom')
        navBar.classList.remove('middle')
        navBar.classList.remove('middle-bottom')
        navBar.classList.add('top')
    }
})
