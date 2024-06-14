document.addEventListener('DOMContentLoaded', function() {
    console.log('Personal area is loaded and ready.');

    document.getElementById('settings-form').addEventListener('submit', function(event) {
        event.preventDefault();
        alert('Settings saved successfully!');
    });
});