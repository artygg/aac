document.addEventListener('DOMContentLoaded', () => {
    const urlParams = new URLSearchParams(window.location.search);
    const courseID = urlParams.get('courseID');
    const goBack = document.getElementById("go-back");
    goBack.href = `/classes?courseID=${courseID}`;

    fetch(`/api/web/courses`)
        .then(res => res.json())
        .then(data => {
            const courses = data.courses;
            courses.forEach(course => {
                if (course.id == courseID) {
                    console.log("Course: ", course.name);
                    const mainHeaderElement = document.querySelector('.main-header');
                    mainHeaderElement.classList.remove('hidden');
                    const h1Element = mainHeaderElement.querySelector('h1');
                    h1Element.innerText = course.name;
                }
            });
        })
        .catch(error => console.error('Error showing course name:', error));

    fetch(`/api/web/attendance/by_course?courseID=${courseID}`)
        .then(res => res.json())
        .then(data => {
            const attendanceData = data;
            const percentages = calculateStatusPercentage(attendanceData);
            const tbody = document.querySelector('tbody');

            Object.keys(percentages).forEach(studentID => {
                const studentData = attendanceData.find(entry => entry.student.id == studentID);
                const tr = document.createElement('tr');

                const studentIdCell = document.createElement('td');
                studentIdCell.textContent = studentData.student.id;
                tr.appendChild(studentIdCell);

                const firstNameCell = document.createElement('td');
                firstNameCell.textContent = studentData.student.first_name;
                tr.appendChild(firstNameCell);

                const lastNameCell = document.createElement('td');
                lastNameCell.textContent = studentData.student.last_name;
                tr.appendChild(lastNameCell);

                const presenceCell = document.createElement('td');
                presenceCell.textContent = percentages[studentID] + "%";
                tr.appendChild(presenceCell);

                tbody.appendChild(tr);
                if (percentages[studentID] >= 80) {
                    presenceCell.className = "green";
                } else {
                    presenceCell.className = "red";
                }
            });

            // Call the filter function every second
            setInterval(filterStudents, 1000);

            // Add input event listener to the search bar
            // Function to filter students based on the search term
            function filterStudents() {
                const searchTerm = document.getElementById('student-search').value.toLowerCase();
                const rows = tbody.getElementsByTagName('tr');
                for (let i = 0; i < rows.length; i++) {
                    const firstName = rows[i].getElementsByTagName('td')[1].textContent.toLowerCase();
                    const lastName = rows[i].getElementsByTagName('td')[2].textContent.toLowerCase();
                    console.log("Student: ", firstName, " ", lastName);
                    if (searchTerm === "" || firstName.includes(searchTerm) || lastName.includes(searchTerm)) {
                        rows[i].style.display = '';
                    } else {
                        rows[i].style.display = 'none';
                    }
                }
            }
        })
        .catch(error => console.error('Error fetching attendance by course:', error));

    function calculateStatusPercentage(items) {
        const counts = {};

        items.forEach(item => {
            const studentID = item.student.id;
            const status = parseInt(item.status);

            if (!counts[studentID]) {
                counts[studentID] = { total: 0, status1Count: 0 };
            }

            counts[studentID].total++;
            if (status === 1 || status === 3) {
                counts[studentID].status1Count++;
            }
        });

        const percentages = {};
        Object.keys(counts).forEach(studentID => {
            const count = counts[studentID];
            percentages[studentID] = Math.floor((count.status1Count / count.total) * 100);
        });

        return percentages;
    }
});