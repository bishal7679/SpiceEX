const wrapper = document.querySelector(".wrapper");
const loginlink = document.querySelector(".login-link");
const registerlink = document.querySelector(".register-link");
const btnPopup = document.querySelector(".btnlogin-popup");
const iconClose = document.querySelector(".icon-close");
const showcase = document.querySelector(".showcase");
const nxtbtn = document.querySelectorAll(".nextbutton");
const backbtn = document.querySelectorAll(".backbutton");
const nav_collpse = document.querySelector(".navbar-collapse");
const navbar_toggler = document.querySelector(".navbar-toggler");
const navbar = document.querySelector(".nav-link.prmry");
// const paymentbackbtn = document.querySelectorAll(".paymentbackbtn");

nxtbtn.forEach((button) => {
  button.addEventListener("click", () => {
    showcase.classList.add("active");
  });
});

backbtn.forEach((button) => {
  button.addEventListener("click", () => {
    showcase.classList.remove("active");
  });
});

registerlink.addEventListener("click", () => {
  wrapper.classList.add("active");
});

loginlink.addEventListener("click", () => {
  wrapper.classList.remove("active");
});

btnPopup.addEventListener("click", () => {
  wrapper.classList.add("active-popup");
  nav_collpse.classList.remove("show");
});

iconClose.addEventListener("click", () => {
  wrapper.classList.remove("active-popup");
});

navbar_toggler.addEventListener("click", () => {
  navbar.classList.remove("prmry");
  navbar.classList.add("tglle");
});

function Display() {
  Display.prototype.show = function (displayMessage) {
    let message = document.getElementById("message");
    message.innerHTML = `${displayMessage}`;
    setTimeout(function () {
      message.innerHTML = "";
    }, 5000);
  };
}

function fileValidation(e) {
  var fileInput = document.getElementById("file1");
  let display = new Display();

  var filePath = fileInput.value;

  // Allowing file type
  var allowedExtensions = /(\.jpg|\.jpeg|\.pdf)$/i;

  if (!allowedExtensions.exec(filePath)) {
    // alert("Invalid file type");
    display.show("Invalid file type! only *.pdf/*.jpg/*.jpeg supported");
    fileInput.value = "";
    return false;
  }
  e.preventDefault();
}

// -------------------------------------------------
