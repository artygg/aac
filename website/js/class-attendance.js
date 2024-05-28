document.addEventListener('DOMContentLoaded', function() {
    const endCheckingButton = document.getElementById('end_checking_prematurely');

    // Add event listener to the "End checking prematurely" button
    endCheckingButton.addEventListener('click', function() {
        const table = document.getElementById('attendance-table');
        const rows = table.getElementsByTagName('tr');

        // Loop through each row in the table body
        for (let i = 1; i < rows.length; i++) {
            const presenceCell = rows[i].getElementsByTagName('td')[3];

            // Check if the cell contains "Mark excuse"
            if (presenceCell && presenceCell.innerText.trim() === 'Mark excuse') {
                // Change text to "Absent"
                presenceCell.innerText = 'Absent';

                // Change background color to red
                presenceCell.style.backgroundColor = '#FF5733'; // Red color
            }
        }

        // Change button text to "Save results"
        endCheckingButton.value = 'Save results';

        // Change button action to redirect to course_overview.html
        endCheckingButton.onclick = function() {
            window.location.href = '../course_overview.html';
        };
    });
});