function tariff(id){
    var button = document.getElementsByClassName("button-tariff")[id];
    var block_text = document.getElementsByClassName("big-size-tariff")[id]
    var block_text_small = document.getElementsByClassName("small-size-tariff")[id]
    var block = document.getElementsByClassName("tariff-dynamic")[id]
    if(button.innerHTML == "Подробнее"){
        button.innerHTML ="Скрыть"
        block_text.style.display ="none"
        block_text_small.style.display ="block"
        block.style.height = "80px"
    }
    else{
        button.innerHTML = 'Подробнее';
        block_text_small.style.display ="none"
        block_text.style.display ="block"
        block.style.height ="450px"
        return
    }
}
function ScaleNewIn(s){
    s.style.marginLeft = "10%";
    s.style.width = "80%";
    s.style.height = "450px";
    s.style.scale = "1.1";
    body = document.getElementsByClassName("news-modal")[0];
    body.style.display = "block";
    middle = s.getElementsByTagName("div")[1];
    middle.style.display = "block";
    more = s.getElementsByClassName("more-info")[0];
    more.style.display = "block";
}
function ScaleNewOut(s){
    body = document.getElementsByClassName("news-modal")[0];
    body.style.display = "none";
    s.style.marginLeft = "10%";
    s.style.width = "80%";
    s.style.height = "300px";
    s.style.scale = "1.0";
    middle = s.getElementsByTagName("div")[1];
    middle.style.display = "none";
    more = s.getElementsByClassName("more-info")[0];
    more.style.display = "none";
}

function nav(){
     button = document.getElementsByClassName("ham")[0];
     mynav = document.getElementById("myNav");
     main = document.getElementsByClassName("main-body-container")[0];
    footer = document.getElementsByTagName("footer")[0];
    if(mynav.style.height == "100%"){
        mynav.style.height = "0%";
        main.style.display = "block";
        footer.style.display = "block";
        button.style.zIndex = "2";
    }else{
        mynav.style.height = "100%"
        button.style.zIndex = "1001";
        main.style.display = "none";
        footer.style.display = "none";
    }
}
// Get the modal
var modal = document.getElementById("myModal");

// Get the button that opens the modal
var btn = document.getElementById("myBtn");

// Get the <span> element that closes the modal
var span = document.getElementsByClassName("close")[0];

// When the user clicks on the button, open the modal
btn.onclick = function() {
    modal.style.display = "block";
}

// When the user clicks on <span> (x), close the modal
span.onclick = function() {
    modal.style.display = "none";
}

// When the user clicks anywhere outside of the modal, close it
window.onclick = function(event) {
    if (event.target == modal) {
        modal.style.display = "none";
    }
}

var slideIndex = 1;
showSlides(slideIndex);

// Next/previous controls
function plusSlides(n) {
    showSlides(slideIndex += n);
}

// Thumbnail image controls
function currentSlide(n) {
    showSlides(slideIndex = n);
}

function showSlides(n) {
    var i;
    var slides = document.getElementsByClassName("mySlides");
    var dots = document.getElementsByClassName("dot");
    if (n > slides.length) {slideIndex = 1}
    if (n < 1) {slideIndex = slides.length}
    for (i = 0; i < slides.length; i++) {
        slides[i].style.display = "none";
    }
    for (i = 0; i < dots.length; i++) {
        dots[i].className = dots[i].className.replace(" active", "");
    }
    slides[slideIndex-1].style.display = "block";
    dots[slideIndex-1].className += " active";
}
var slideIndex = 0;
showSlides();

function showSlides() {
    var i;
    var slides = document.getElementsByClassName("mySlides");
    for (i = 0; i < slides.length; i++) {
        slides[i].style.display = "none";
    }
    slideIndex++;
    if (slideIndex > slides.length) {slideIndex = 1}
    slides[slideIndex-1].style.display = "block";
    setTimeout(showSlides, 10000); // Change image every 2 seconds
}
