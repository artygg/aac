document.getElementById('loginForm').addEventListener('submit', function(event) {
    event.preventDefault();

    var login = document.getElementById('email').value;
    var password = document.getElementById('password').value;

    var xhr = new XMLHttpRequest();
    xhr.open('POST', '/api/web/login', true);
    xhr.setRequestHeader('Content-Type', 'application/json;charset=UTF-8');

    xhr.onreadystatechange = function() {
        if (xhr.readyState === XMLHttpRequest.DONE) {
            var status = xhr.status;

            if (status === 0 || (status >= 200 && status < 400)) {
                window.location.href = '/courses';
            } else {
                console.log("Status is bad (", status, ")")
                document.getElementById('message').textContent = 'Invalid login or password';
            }
        }
    };

    xhr.send(JSON.stringify({
        email: login,
        password: password
    }));
});