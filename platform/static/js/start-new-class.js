const urlParams = new URLSearchParams(window.location.search);
const urlCourseID = urlParams.get('courseID');
const urlCourseName = urlParams.get('name');
const goBack = document.getElementById("go-back")
goBack.href = `/classes?courseID=${urlCourseID}`

const subHeader = document.querySelector('.sub-header')
const header1 = subHeader.querySelector('h1')
header1.textContent = urlCourseName

fetch('/api/web/rooms')
    .then(response => response.json())
    .then(data2 => {
        const rooms = data2.map(spot => spot.room);
        console.log(rooms)
        printDropdownOptions(rooms, "roomSelector");
    })
    .catch(error => console.error('Error fetching rooms:', error));

fetch(`/api/web/groups/by_course?courseID=${urlCourseID}`)
    .then(response => response.json())
    .then(data => {
        const groupIds = data.map(group => group.id);
        printDropdownOptions(groupIds, "groupInput");
    })
    .catch(error => console.error('Error fetching groups:', error));


function chooseGroup() {
    let mylist = document.getElementById("groupInput");
    let field = document.getElementById("groupOutput");
    let option = mylist.options[mylist.selectedIndex].text;
    if (field.value.includes(option)) {
        field.value = field.value.replace(option + " ", "")
    } else {
        field.value += option + " ";
    }
    mylist.options[0].selected = 'selected';
}

function printDropdownOptions(array, elementId) {
    const selectElement = document.getElementById(elementId);
    array.forEach(option => {
        const optionElement = document.createElement("option");
        optionElement.textContent = option;
        optionElement.value = option;
        selectElement.appendChild(optionElement);
    });
}

function chooseRoom() {
    let mylist = document.getElementById("roomSelector");
    let field = document.getElementById("roomOutput");
    if (mylist.value != "Choose room") {
        field.value = mylist.value;
    }
}

function chooseEndTime() {
    let mylist = document.getElementById("endTime");
    let field = document.getElementById("endTimeOutput");
    field.value = mylist.value.replace("T", " ");
}

function chooseStartTime() {
    let mylist = document.getElementById("startTime");
    let field = document.getElementById("startTimeOutput");
    field.value = mylist.value.replace("T", " ");
}

document.getElementById("form").addEventListener("submit", function(event) {
    event.preventDefault();

    const room = document.getElementById("roomOutput").value;
    const courseID = parseInt(urlCourseID);
    const startTime = document.getElementById("startTimeOutput").value;
    const endTime = document.getElementById("endTimeOutput").value;
    const groups = document.getElementById("groupOutput").value.trim().split(" ");
    const validationErrorMessage = document.getElementById('validation-error-message');

    const classData = {
        course_id: courseID,
        start_time: startTime,
        end_time: endTime,
        room: room,
        groups: groups
    };

    if (!room || !groups || !endTime || !startTime) {
        validationErrorMessage.textContent = "Please fill out all fields";
        validationErrorMessage.classList.remove('hidden');
        console.log("Please fill out all fields")
        return;
    }

    if (new Date(startTime) >= new Date(endTime)) {
        validationErrorMessage.textContent = "The start time should be before the end time";
        validationErrorMessage.classList.remove('hidden');
        console.log("The start time should be before the end time")
        return;
    }

    console.log('Submitting course data:', classData);

    fetch('/api/web/class', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(classData)
    })
        .then(response => {
            if (!response.ok) {
                return response.text().then(text => { throw new Error(text) });
            }
        })
        .then(data => {
            console.log('Class created:', data);
            window.location.href = `/classes?courseID=${urlCourseID}`;
        })
        .catch(error => console.error('Error creating course:', error));
});