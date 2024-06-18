# Automated Attendance 

# AAS Frontend Readme

## Overview
This repository contains the frontend for the AAC project, developed using JavaScript, HTML and CSS. The frontend 
creates the UI, and if the user is logged in as teacher, they have access to various the data and the information which 
they see, change and add.

## Prerequisites
- Back-end running
- A working internet connection

## Installation and usage
In order to have the access to an actual real data and information to which the teacher should have access to, the user need 
to copy all .go files from backend and paste to the frontend folder and do the same steps as installation and usage, with only difference 
that the command ` git checkout back-end ` needs to be changed to ` git checkout front-end `

## Main Structure
- `index.html`: Entry point to the website, the login page same as `signin.html`.
- `register.html`: The page where user can create a teacher account.
- Other pages are accessible only after login, where teachers can see, change and add information about courses, 
classes, groups, students and attendances.

## Detailed Structure
- `protected.html`: Protected page non-accessible to users who are not logged in.
- `courses-page.html`: Page that appears after user is logs in as a teacher, which contains list of courses. Teacher can 
filter between courses which have been in the past, which are in present and which will be in a future.
- `start-new-course`: In order to create a course, teacher should go to this page. The course should have name, year, 
start date, end date and groups for which this course will be available.
- `classes-page.html`: After clicking on one of the courses, the teacher would be transferred to the page with all the 
classes for this particular course.
- `start-new-class.html`: On this page the teacher is able to create a new class for the chosen course. The class should
have a room, start time, end time and groups which should be from those belonging to the course.
- `course-attendance-statistics.html`: This page gives an overview to the attendance of all students who have the chosen 
course. For each student it shows the percentage of the classes which the student has visited for the course. If the percentage
is 80 or above, the student attends enough.
- `class-attendance.html`: This page opens after teacher clicks on some class. There will be displayed a table where 
teachers can switch the attendance status for any student for whom this class is assigned to either "Present", "Absent", 
"Excused" or just leave as "Unmarked". Teacher can also click to "End the class prematurely", and by doing so end the class 
immediately amd mark "Unmarked" students as "Absent". Additionally, the timer will be ticking if the current time is within start time's and end time's boundaries of the class.