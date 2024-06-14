fetch('/api/web/groups')
    .then(response => response.json())
    .then(data => {
        const groupIds = data.map(group => group.id);
        printDropdownOptions(groupIds);
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

function printDropdownOptions(array) {
    const selectElement = document.getElementById("groupInput");
    array.forEach(option => {
        const optionElement = document.createElement("option");
        optionElement.textContent = option;
        optionElement.value = option;
        selectElement.appendChild(optionElement);
    });
}

function typeNameOfTheCourse() {
    let mylist = document.getElementById("courseNameInput");
    let field = document.getElementById("courseNameOutput");
    field.value = mylist.value;
}

function typeNameOfTheYear() {
    let mylist = document.getElementById("yearNameInput");
    let field = document.getElementById("yearNameOutput");
    field.value = mylist.value;
}

function dateFunction() {
    let startDate = document.getElementById("startDate");
    let endDate = document.getElementById("endDate");
    let text = document.getElementById("date");
    text.value = "From " + startDate.value.toString() + " To " + endDate.value.toString();
}

document.getElementById("form").addEventListener("submit", function(event) {
    event.preventDefault();

    const courseName = document.getElementById("courseNameInput").value;
    const year = parseInt(document.getElementById("yearNameInput").value, 10);
    const startDate = document.getElementById("startDate").value;
    const endDate = document.getElementById("endDate").value;
    const groups = document.getElementById("groupOutput").value.trim().split(" ");
    const validationErrorMessage = document.getElementById('validation-error-message');

    const courseData = {
        name: courseName,
        year: year,
        start_date: startDate,
        end_date: endDate,
        groups: groups
    };

    if (!courseName|| !year || !groups || !endDate || !startDate) {
        validationErrorMessage.textContent = "Please fill out all fields";
        validationErrorMessage.classList.remove('hidden');
        console.log("Please fill out all fields")
        return;
    }

    if (new Date(startDate) >= new Date(endDate)) {
        validationErrorMessage.textContent = "The start date should be before the end date";
        validationErrorMessage.classList.remove('hidden');
        console.log("The start date should be before the end date")
        return;
    }

    console.log('Submitting course data:', courseData);

    fetch('/api/web/course', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(courseData)
    })
        .then(response => {
            if (!response.ok) {
                return response.text().then(text => { throw new Error(text) });
            }
        })
        .then(data => {
            console.log('Course created:', data);
            window.location.href = "/courses";
        })
        .catch(error => console.error('Error creating course:', error));
});