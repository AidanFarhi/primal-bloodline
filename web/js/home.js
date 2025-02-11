// scroll effect to darken main image
const mainContainer = document.getElementById('main-container')
const section1Container = document.getElementById('section-1')
mainContainer.addEventListener('scroll', () => {
    const rect = section1Container.getBoundingClientRect()
    const viewportHeight = window.innerHeight
    const distanceFromTop = Math.max(0, rect.bottom)
    const darkness = 1 - Math.min(distanceFromTop / viewportHeight, 1)
    section1Container.style.backgroundColor = `rgba(0, 0, 0, ${darkness})`
})

// menu in-and-out effect
const menuBurger = document.getElementById('logout')
menuBurger.addEventListener('click', () => {
    const menu = document.getElementById('menu')
    menu.classList.toggle('open')
})
