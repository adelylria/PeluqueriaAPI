document.getElementById("register-form").addEventListener("submit", function(event) {
    event.preventDefault();
    var username = document.getElementById("username").value;
    var email = document.getElementById("email").value;
    var password = document.getElementById("password").value;
    // Aquí puedes hacer una solicitud AJAX para enviar el nombre de usuario, correo electrónico y contraseña al backend
    console.log("Username:", username);
    console.log("Email:", email);
    console.log("Password:", password);
});
