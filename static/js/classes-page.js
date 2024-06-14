document.addEventListener("DOMContentLoaded", () => {
    const urlParams = new URLSearchParams(window.location.search);
    const courseID = urlParams.get('courseID');

    const createClassLink = document.getElementById('create-class-link');
    createClassLink.href = `/class/create?courseID=${courseID}&name`;

    const statisticsLink = document.getElementById('statistics-link');
    statisticsLink.href = `/attendance/by_course?courseID=${courseID}`;

    fetch(`/api/web/courses`)
        .then(res => res.json())
        .then(data => {
            const courses = data.courses;

            courses.forEach(course => {
                if (course.id == courseID) {
                    console.log("Course: ",course.name)
                    createClassLink.href = `/class/create?courseID=${courseID}&name=${course.name}`
                    const mainHeaderElement = document.querySelector('.main-header');
                    const h1Element = mainHeaderElement.querySelector('h1');
                    h1Element.innerText = course.name;
                    window.courseName = course.name
                    mainHeaderElement.classList.remove('hidden');
                }
            })

            fetch(`/api/web/classes?courseID=${courseID}`)
                .then(response => response.json())
                .then(data => {
                    console.log(data)
                    const classesContainer = document.querySelector('.classes');
                    data.forEach(classInstance => {
                        const classContainer = document.createElement('div');
                        classContainer.classList.add('class-container');

                        const classElement = document.createElement('div');
                        classElement.classList.add('class');

                        const colorDiv = document.createElement('div');
                        colorDiv.classList.add('color');
                        const linkToClass = document.createElement('div');
                        linkToClass.classList.add('link-to-the-class-info');
                        const link = document.createElement('a');
                        link.setAttribute('href', `/attendance/by_class?classID=${classInstance.id}&courseID=${courseID}`);
                        link.textContent = `${courseName} - ${classInstance.id} - ${classInstance.room}`;
                        const startDate = document.createElement('div');
                        startDate.classList.add('date');
                        const startDateText = document.createElement('p');
                        startDateText.textContent = classInstance.starttime;

                        startDate.appendChild(startDateText)
                        linkToClass.appendChild(link)
                        classElement.appendChild(colorDiv)
                        classElement.appendChild(linkToClass)
                        classElement.appendChild(startDate)
                        classContainer.appendChild(classElement);
                        classesContainer.appendChild(classContainer);
                    })
                })
                .catch(error => console.error('Error fetching classes:', error));
        })
        .catch(error => console.error('Error fetching courses:', error));




});