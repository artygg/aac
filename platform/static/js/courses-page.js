function previousCourses() {
    fetch('/api/web/courses')
        .then(response => response.json())
        .then(data => {
            console.log(data)
            const courses = data.courses;
            const coursesContainer = document.querySelector('.courses');

            // Clear existing courses (if any)
            coursesContainer.innerHTML = '';

            courses.forEach(course => {
                console.log('Course: ', course.id)
                const x = Date.now();
                const y = course.end_date;
                if (x > new Date(y)) {
                    const courseContainer = document.createElement('div');
                    courseContainer.classList.add('course-container');

                    const courseElement = document.createElement('div');
                    courseElement.classList.add('course');

                    const colorDiv = document.createElement('div');
                    colorDiv.classList.add('color');
                    const linkToCourse = document.createElement('div');
                    linkToCourse.classList.add('link-to-the-course');
                    const link = document.createElement('a');
                    link.setAttribute('href', `/classes?courseID=${course.id}`);
                    link.textContent = course.name;
                    const dropdown = document.createElement('div');
                    dropdown.classList.add('dropdown');
                    const dropdownText = document.createElement('p');
                    dropdownText.classList.add('dropdownText');
                    dropdownText.textContent = "Teacher";
                    linkToCourse.appendChild(link);
                    dropdown.appendChild(dropdownText);

                    courseElement.appendChild(colorDiv);
                    courseElement.appendChild(linkToCourse);
                    courseElement.appendChild(dropdown);
                    courseContainer.appendChild(courseElement);
                    coursesContainer.appendChild(courseContainer);
                }
            });
        })
        .catch(error => console.error('Error fetching courses:', error));
}

function futureCourses() {
    fetch('/api/web/courses')
        .then(response => response.json())
        .then(data => {
            console.log(data)
            const courses = data.courses;
            const coursesContainer = document.querySelector('.courses');

            // Clear existing courses (if any)
            coursesContainer.innerHTML = '';

            courses.forEach(course => {
                console.log('Course: ', course.id)
                const x = Date.now();
                const y = course.start_date;
                if (x < new Date(y)) {
                    const courseContainer = document.createElement('div');
                    courseContainer.classList.add('course-container');

                    const courseElement = document.createElement('div');
                    courseElement.classList.add('course');

                    const colorDiv = document.createElement('div');
                    colorDiv.classList.add('color');
                    const linkToCourse = document.createElement('div');
                    linkToCourse.classList.add('link-to-the-course');
                    const link = document.createElement('a');
                    link.setAttribute('href', `/classes?courseID=${course.id}`);
                    link.textContent = course.name;
                    const dropdown = document.createElement('div');
                    dropdown.classList.add('dropdown');
                    const dropdownText = document.createElement('p');
                    dropdownText.classList.add('dropdownText');
                    dropdownText.textContent = "Teacher";
                    linkToCourse.appendChild(link);
                    dropdown.appendChild(dropdownText);

                    courseElement.appendChild(colorDiv);
                    courseElement.appendChild(linkToCourse);
                    courseElement.appendChild(dropdown);
                    courseContainer.appendChild(courseElement);
                    coursesContainer.appendChild(courseContainer);
                }
            });
        })
        .catch(error => console.error('Error fetching courses:', error));
}

function currentCourses() {
    fetch('/api/web/courses')
        .then(response => response.json())
        .then(data => {
            console.log(data)
            const courses = data.courses;
            const coursesContainer = document.querySelector('.courses');

            // Clear existing courses (if any)
            coursesContainer.innerHTML = '';

            courses.forEach(course => {
                console.log('Course: ', course.id)
                const x = Date.now();
                const y = course.start_date;
                const z = course.end_date
                if (x >= new Date(y) && x <= new Date(z)) {
                    const courseContainer = document.createElement('div');
                    courseContainer.classList.add('course-container');

                    const courseElement = document.createElement('div');
                    courseElement.classList.add('course');

                    const colorDiv = document.createElement('div');
                    colorDiv.classList.add('color');
                    const linkToCourse = document.createElement('div');
                    linkToCourse.classList.add('link-to-the-course');
                    const link = document.createElement('a');
                    link.setAttribute('href', `/classes?courseID=${course.id}`);
                    link.textContent = course.name;
                    const dropdown = document.createElement('div');
                    dropdown.classList.add('dropdown');
                    const dropdownText = document.createElement('p');
                    dropdownText.classList.add('dropdownText');
                    dropdownText.textContent = "Teacher";
                    linkToCourse.appendChild(link);
                    dropdown.appendChild(dropdownText);

                    courseElement.appendChild(colorDiv);
                    courseElement.appendChild(linkToCourse);
                    courseElement.appendChild(dropdown);
                    courseContainer.appendChild(courseElement);
                    coursesContainer.appendChild(courseContainer);
                }
            });
        })
        .catch(error => console.error('Error fetching courses:', error));
    var elems = document.querySelectorAll(".selected");

    [].forEach.call(elems, function(el) {
        el.classList.remove("selected");
    });

    var element = document.getElementById("current");
    element.classList.add("selected");
}