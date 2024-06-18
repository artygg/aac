document.addEventListener('DOMContentLoaded', function() {
    const urlParams = new URLSearchParams(window.location.search);
    const urlClassID = urlParams.get('classID');
    const urlCourseID = urlParams.get('courseID');
    const mainHeader = document.querySelector('.main-header')
    const courseName = mainHeader.querySelector('h1');
    const dateContainer = document.querySelector('.top-class-description')
    const date = dateContainer.querySelector('p');
    const bottomClassDescription = document.querySelector('.bottom-class-description');
    const paragraphs = bottomClassDescription.querySelectorAll('p');
    const classStatus = paragraphs[0]
    const timeElement = paragraphs[1]
    const goBack = document.getElementById("go-back")
    goBack.href = `/classes?courseID=${urlCourseID}`

    const endCheckingButton = document.getElementById('end_checking_prematurely');

    // Add event listener to the "End checking prematurely" button
    fetch(`/api/web/courses`)
        .then(res => res.json())
        .then(data => {
            const courses = data.courses;
            courses.forEach(course => {
                if (course.id === parseInt(urlCourseID)) {
                    courseName.textContent = course.name
                    courseName.classList.remove('invisible')
                }
            })
        })
        .catch(error => console.error('Error fetching courses:', error));

    fetch(`/api/web/classes?courseID=${urlCourseID}`)
        .then(res => res.json())
        .then(data => {
            data.forEach(entering => {
                if (entering.id === parseInt(urlClassID)) {
                    console.log(entering);
                    const classStartTime = new Date(entering.starttime);
                    const classEndTime = new Date(entering.endtime);
                    date.textContent = entering.starttime;
                    console.log(`${entering.starttime} and ${entering.endtime}`)
                    dateContainer.classList.remove('invisible');
                    setInterval(() => defineClassStatus(entering.starttime, entering.endtime), 200);

                    fetch(`/api/web/attendance/by_class?classID=${urlClassID}`)
                        .then(res => res.json())
                        .then(data => {
                            console.log(data)
                            const tbody = document.querySelector('tbody');
                            data.forEach(entry => {
                                const tr = document.createElement('tr');

                                const studentIdCell = document.createElement('td');
                                studentIdCell.textContent = entry.student.id;
                                tr.appendChild(studentIdCell);

                                const firstNameCell = document.createElement('td');
                                firstNameCell.textContent = entry.student.first_name;
                                tr.appendChild(firstNameCell);

                                const lastNameCell = document.createElement('td');
                                lastNameCell.textContent = entry.student.last_name;
                                tr.appendChild(lastNameCell);

                                generateAttendanceStatus(entry, tr, classStartTime, classEndTime);

                                tbody.appendChild(tr);
                            })
                            const rowsToRemove = tbody.children.length - data.length;
                            if (rowsToRemove > 0) {
                                for (let i = 0; i < rowsToRemove; i++) {
                                    tbody.removeChild(tbody.lastChild);
                                }
                            }
                        })
                        .catch(error => console.error('Error fetching courses:', error));

                    function markUnmarkedAsAbsent() {
                        const table = document.querySelector('table');
                        const rows = table.getElementsByTagName('tr');

                        for (let i = 1; i < rows.length; i++) {
                            const presenceCell = rows[i].getElementsByTagName('td')[3];
                            const select = presenceCell.querySelector('select');
                            let currentStatus;
                            if (select) {
                                currentStatus = select.value;
                            } else {
                                currentStatus = presenceCell.textContent;
                            }

                            if (currentStatus === "0" || currentStatus === "Unmarked") {
                                const studentIdCell = rows[i].getElementsByTagName('td')[0];
                                const studentId = studentIdCell.textContent;
                                const studentFirstName = rows[i].getElementsByTagName('td')[1].textContent;
                                const studentLastName = rows[i].getElementsByTagName('td')[2].textContent;

                                // Change text to "Absent" and background color to red

                                presenceCell.className = 'red';
                                presenceCell.textContent = 'Absent'



                                // Update status on server
                                updateStatus({
                                    id: parseInt(studentId),
                                    first_name: studentFirstName,
                                    last_name: studentLastName,
                                    email: ""
                                }, "2", urlClassID);
                            }

                        }

                        // Change button text to "Save results"
                        endCheckingButton.value = 'Save results';

                        endCheckingButton.onclick = function() {
                            window.location.href = '';
                        };
                    }

                    endCheckingButton.addEventListener('click', function() {
                        const input = {
                            class_id: parseInt(urlClassID)
                        };

                        fetch('/api/web/class/end', {
                            method: 'POST', // Specify the request method
                            headers: {
                                'Content-Type': 'application/json' // Indicate that the request body is in JSON format
                            },
                            body: JSON.stringify(input) // Convert the JavaScript object to a JSON string
                        })
                            .then(response => {
                                if (!response.ok) {
                                    // Handle any errors returned by the server
                                    throw new Error('Network response was not ok');
                                }
                            })
                            .catch(error => {
                                // Handle any errors that occurred during the fetch
                                console.error('There was a problem with the fetch operation:', error);
                            });
                        markUnmarkedAsAbsent();
                    });

                    let now = new Date()
                    const intervalID = setInterval(() => {
                        now = new Date()
                        if (now > classEndTime) {
                            console.log('Interval stopped!');
                            clearInterval(intervalID);
                            markUnmarkedAsAbsent();
                            console.log('Interval stopped');
                        }
                    }, 200);
                }
            })
        })
        .catch(error => console.error('Error fetching courses:', error));

    function updateStatus(student, newStatus, classId) {
        fetch(`/api/web/attendance`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                class_id: parseInt(classId),
                student: student,
                status: newStatus,
                time: "0000-00-00 00:00:00"
            })
        })
            .then(response => {
                if (!response.ok) {
                    throw new Error('Failed to update attendance status');
                } else {
                    console.log(`Status updated successfully to ${newStatus}`);
                }
            })
            .catch(error => {
                console.error('Error updating attendance status:', error);
            });
    }

    function setColors(cell, select, status) {
        if (status !== null) {
            switch (status) {
                case "1":
                    cell.className = "green"
                    if (select) select.className = "green"
                    break
                case "2":
                    cell.className = "red"
                    if (select) select.className = "red"
                    break
                case "3":
                    cell.className = "grey"
                    if (select) select.className = "grey"
                    break
                default:
                    cell.className = "white"
                    if (select) select.className = "white"
            }
        }
    }

    // Timer function to update the time every second
    function updateTimer(startingDateTime) {
        const startTime = new Date(startingDateTime);
        setInterval(() => {
            const now = new Date();
            const elapsed = new Date(now - startTime);
            const hours = String(elapsed.getUTCHours()).padStart(2, '0');
            const minutes = String(elapsed.getUTCMinutes()).padStart(2, '0');
            const seconds = String(elapsed.getUTCSeconds()).padStart(2, '0');
            timeElement.textContent = `Time of the class - ${hours}:${minutes}:${seconds}`;
        }, 200);
    }

    function generateAttendanceStatus(entry, tr, classStartTime, classEndTime) {
        const now = new Date();
        // console.log(`end time: ${classStartTime}`)
        if (now >= classStartTime && now <=classEndTime) {
            // console.log('You can change att');
            const presenceCell = document.createElement('td');
            presenceCell.className = 'white';

            const select = document.createElement('select');
            setColors(presenceCell, select, entry.status);
            select.innerHTML = `
                    <option value="0" ${entry.status === "0" ? "selected" : ""}>Unmarked</option>
                    <option value="1" ${entry.status === "1" ? "selected" : ""}>Present</option>
                    <option value="2" ${entry.status === "2" ? "selected" : ""}>Absent</option>
                    <option value="3" ${entry.status === "3" ? "selected" : ""}>Excused</option>
                `;

            // Append dropdown to cell
            presenceCell.appendChild(select);

            // Event listener for dropdown change
            select.addEventListener('change', function() {
                const newStatus = select.value;
                updateStatus(entry.student, newStatus, urlClassID);
                setColors(presenceCell, select, newStatus);
                console.log(select.value)
                setColors()
            });

            tr.appendChild(presenceCell);
        } else {
            // console.log('You cannot change att');
            // If class hasn't ended yet, add a regular cell with status and color
            const presenceCell = document.createElement('td');
            const statusActionMap = {
                "1": { textContent: "Present", className: "green" },
                "2": { textContent: "Absent", className: "red" },
                "3": { textContent: "Excused", className: "grey" }
            };

            const status = entry.status;
            const action = statusActionMap[status] || { textContent: "Unmarked", className: "white" };
            presenceCell.textContent = action.textContent;
            presenceCell.className = action.className;
            setColors(presenceCell, null, entry.status); // Pass null for select
            tr.appendChild(presenceCell);
        }
    }

    function defineClassStatus(start, end) {
        const now = new Date()
        const startTime = new Date(start)
        const endTime = new Date(end)
        if (now >= startTime && now <= endTime) {
            console.log("Class has started and is ongoing")
            window.onload = updateTimer(start);
            timeElement.classList.remove('invisible')
            classStatus.textContent = "Class has started and is ongoing"
            classStatus.classList.remove('invisible')
            return "ongoing"
        }
        else if (now > endTime) {
            console.log("Class has already passed")
            timeElement.classList.add('invisible');
            classStatus.classList.remove('invisible')
            classStatus.textContent = "Class has already passed"
            return "passed"
        } else {
            console.log("Class has not started yet")
            timeElement.classList.add('invisible');
            classStatus.classList.remove('invisible')
            classStatus.textContent = "Class has not started yet"
            return "upcoming";
        }
    }

});