// When the user scrolls the page, execute manageStickyClassForNav
window.onscroll = function() { manageStickyClassForNav() };
var navbar = document.getElementById("navbar");
var sticky = navbar.offsetTop;

// Add the sticky class to the navbar when you reach its scroll position.
// Remove "sticky" when you leave the scroll position
function manageStickyClassForNav() {
  if (window.pageYOffset >= sticky) {
    navbar.classList.add("sticky")
  } else {
    navbar.classList.remove("sticky");
  }
}
