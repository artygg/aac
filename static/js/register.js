document.getElementById('main-form').addEventListener('submit', function(event) {
    event.preventDefault();

    const firstName = document.getElementById('first_name').value;
    const lastName = document.getElementById('last_name').value;
    const email = document.getElementById('email').value;
    const password = document.getElementById('password').value;
    const validatePassword = document.getElementById('validate_password').value;

    const validationErrorMessage = document.getElementById('validation-error-message');

    validationErrorMessage.style.display = 'none';

    const emailDomain = "@nhlstenden.com";

    if (!firstName|| !lastName || !email || email === emailDomain) {
        validationErrorMessage.textContent = "Please fill out all fields";
        validationErrorMessage.classList.remove('hidden');
        console.log("Please fill out all fields")
        return;
    }

    if (!email.endsWith(emailDomain)) {
        validationErrorMessage.style.display = 'block';
        validationErrorMessage.innerText = 'Invalid email address format';
        return;
    }

    if (password !== validatePassword) {
        validationErrorMessage.style.display = 'block';
        validationErrorMessage.innerText = 'Passwords do not match';
        return;
    }

    const userData = {
        first_name: firstName,
        last_name: lastName,
        email: email,
        password: password
    };

    fetch('/api/web/teacher', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(userData)
    })
        .then(response => {
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            window.location.href = "/login";
        })
        .catch(error => {
            console.error('There was a problem with the registration request:', error);
        });
});