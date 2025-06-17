const navBar = document.getElementById('top-navbar')

window.addEventListener('scroll', () => {
    const scrollTop = window.scrollY
    const windowHeight = window.innerHeight
    const documentHeight = document.documentElement.scrollHeight
    const triggerPoint = window.innerHeight * 0.85 // 60vh
    if (scrollTop + windowHeight >= documentHeight) {
        navBar.classList.remove('scrolled')
        navBar.classList.add('bottom')
    } else if (window.scrollY > triggerPoint) {
        navBar.classList.remove('bottom')
        navBar.classList.add('scrolled')
    } else {
        navBar.classList.remove('bottom')
        navBar.classList.remove('scrolled')
    }
})