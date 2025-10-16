const rightNavLinksDiv = document.getElementsByClassName('right-nav-links')[0]
const burgerMenuDiv = document.getElementsByClassName('burger-menu')[0]
const closeNavButton = document.getElementsByClassName('close-nav-button')[0]
burgerMenuDiv.addEventListener('click', () => {
    rightNavLinksDiv.classList.add('open')
})
document.addEventListener('click', event => {
    if (event.target !== rightNavLinksDiv && !rightNavLinksDiv.contains(event.target) && 
        event.target !== burgerMenuDiv && !burgerMenuDiv.contains(event.target)) {
        rightNavLinksDiv.classList.remove('open')
    }
})
document.addEventListener('scroll', () => rightNavLinksDiv.classList.remove('open'))
closeNavButton.addEventListener('click', () => rightNavLinksDiv.classList.remove('open'))