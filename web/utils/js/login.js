document.getElementById("login-form").addEventListener("submit", function(event) {
    event.preventDefault();
    var username = document.getElementById("username").value;
    var password = document.getElementById("password").value;
    
    var data = {
        username: username,
        password: password
    };

    // Realizar la solicitud POST al endpoint de inicio de sesión
    fetch('http://www.peluqueria.ps:8080/api/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    })
    .then(response => {
        if (response.ok) {
            // El inicio de sesión fue exitoso, puedes redirigir al usuario a otra página o realizar alguna acción adicional
            console.log("Inicio de sesión exitoso");
            // Redirigir a otra página
            window.location.href = "../index.html";
        } else {
            // El inicio de sesión falló, puedes mostrar un mensaje de error al usuario
            console.error("Inicio de sesión fallido");
            alert("Nombre de usuario o contraseña incorrectos");
        }
    })
    .catch(error => {
        // En caso de un error de red u otro tipo de error, puedes manejarlo aquí
        console.error("Error al intentar iniciar sesión:", error);
        alert("Se produjo un error al intentar iniciar sesión. Por favor, inténtalo de nuevo más tarde.");
    });
});

