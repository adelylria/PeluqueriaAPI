document.getElementById("login-form").addEventListener("submit", function(event) {
    event.preventDefault();
    var username = document.getElementById("username").value;
    var password = document.getElementById("password").value;
    // Aquí puedes hacer una solicitud AJAX para enviar el nombre de usuario y la contraseña al backend
    console.log("Username:", username);
    console.log("Password:", password);
});
