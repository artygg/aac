-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Хост: mysql
-- Время создания: Май 20 2024 г., 18:50
-- Версия сервера: 11.1.2-MariaDB-1:11.1.2+maria~ubu2204
-- Версия PHP: 8.2.10

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- База данных: `db`
--

-- --------------------------------------------------------

--
-- Структура таблицы `attendance`
--

CREATE TABLE `attendance` (
  `LessonId` int(11) NOT NULL,
  `StudentId` int(11) NOT NULL,
  `Time` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Структура таблицы `device`
--

CREATE TABLE `device` (
  `MAC` varchar(12) NOT NULL,
  `Key` int(60) NOT NULL,
  `Room` varchar(10) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Дамп данных таблицы `device`
--

INSERT INTO `device` (`MAC`, `Key`, `Room`) VALUES
('00:11:22:33', 123456, 'Room A'),
('66:77:88:99', 654321, 'Room B');

-- --------------------------------------------------------

--
-- Структура таблицы `group`
--

CREATE TABLE `group` (
  `Id` varchar(20) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Дамп данных таблицы `group`
--

INSERT INTO `group` (`Id`) VALUES
('1'),
('2');

-- --------------------------------------------------------

--
-- Структура таблицы `group assign`
--

CREATE TABLE `group assign` (
  `GroupID` varchar(20) DEFAULT NULL,
  `StudentID` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Дамп данных таблицы `group assign`
--

INSERT INTO `group assign` (`GroupID`, `StudentID`) VALUES
('1', 292573),
('1', 5139872),
('1', 5304377),
('2', 5320259),
('2', 5332370),
('2', 5357608);

-- --------------------------------------------------------

--
-- Структура таблицы `lesson`
--

CREATE TABLE `lesson` (
  `Id` int(11) NOT NULL,
  `SubjectID` int(11) NOT NULL,
  `Room` varchar(10) NOT NULL,
  `StartTime` datetime NOT NULL,
  `EndTime` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Дамп данных таблицы `lesson`
--

INSERT INTO `lesson` (`Id`, `SubjectID`, `Room`, `StartTime`, `EndTime`) VALUES
(1, 101, 'Room A', '2024-05-20 09:00:00', '2024-05-20 10:30:00'),
(2, 102, 'Room B', '2024-05-20 11:00:00', '2024-05-20 12:30:00');

-- --------------------------------------------------------

--
-- Структура таблицы `student`
--

CREATE TABLE `student` (
  `Id` int(11) NOT NULL,
  `FirstName` varchar(60) NOT NULL,
  `LastName` varchar(60) NOT NULL,
  `Email` varchar(60) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Дамп данных таблицы `student`
--

INSERT INTO `student` (`Id`, `FirstName`, `LastName`, `Email`) VALUES
(292573, 'Artjoms', 'Grisajevs', 'artjoms.grisajevs@student.nhlstenden.com'),
(5139872, 'Sofronie', 'Albu', 'sofronie.albu@student.nhlstenden.com'),
(5304377, 'Nikita', 'Golovanov', 'nikita.golovanov@student.nhlstenden.com'),
(5320259, 'Yaroslav', 'Oleinychenko', 'yaroslav.oleinychenko@student.nhlstenden.com'),
(5332370, 'Ekaterina', 'Tarlykova', 'ekaterina.tarlykova@student.nhlstenden.com'),
(5357608, 'Fjodor', 'Smorodins', 'fjodor.smorodins@student.nhlstenden.com'),
(5415373, 'Sebastian', 'Serban', 'sebastian.serban@student.nhlstenden.com');

-- --------------------------------------------------------

--
-- Структура таблицы `subject`
--

CREATE TABLE `subject` (
  `id` int(11) NOT NULL,
  `Name` varchar(60) NOT NULL,
  `TeacherID` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Дамп данных таблицы `subject`
--

INSERT INTO `subject` (`id`, `Name`, `TeacherID`) VALUES
(101, 'Mathematics', 1),
(102, 'Physics', 2);

-- --------------------------------------------------------

--
-- Структура таблицы `subject-group-bridge`
--

CREATE TABLE `subject-group-bridge` (
  `GroupID` varchar(20) NOT NULL,
  `Subjectid` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Структура таблицы `subject-lesson-group-bridge`
--

CREATE TABLE `subject-lesson-group-bridge` (
  `GroupId` varchar(20) NOT NULL,
  `Lessonid` int(11) NOT NULL,
  `Subjectid` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Структура таблицы `teacher`
--

CREATE TABLE `teacher` (
  `id` int(11) NOT NULL,
  `firstName` varchar(60) NOT NULL,
  `lastName` varchar(60) NOT NULL,
  `email` varchar(60) NOT NULL,
  `password` varchar(255) NOT NULL,
  `registrationDate` date NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Дамп данных таблицы `teacher`
--

INSERT INTO `teacher` (`id`, `firstName`, `lastName`, `email`, `password`, `registrationDate`) VALUES
(1, 'John', 'Doe', 'john.doe@nhlstenden.com', 'password123', '2024-01-15'),
(2, 'Jane', 'Smith', 'jane.smith@nhlstenden.com', 'smithjane2024', '2024-02-10'),
(3, 'Michael', 'Brown', 'michael.brown@nhlstenden.com', 'mikebrown2024', '2024-03-05');

--
-- Индексы сохранённых таблиц
--

--
-- Индексы таблицы `attendance`
--
ALTER TABLE `attendance`
  ADD PRIMARY KEY (`LessonId`,`StudentId`),
  ADD KEY `fk_student_id` (`StudentId`);

--
-- Индексы таблицы `device`
--
ALTER TABLE `device`
  ADD PRIMARY KEY (`MAC`),
  ADD UNIQUE KEY `Room` (`Room`);

--
-- Индексы таблицы `group`
--
ALTER TABLE `group`
  ADD PRIMARY KEY (`Id`);

--
-- Индексы таблицы `group assign`
--
ALTER TABLE `group assign`
  ADD KEY `GroupID` (`GroupID`),
  ADD KEY `StudentID` (`StudentID`);

--
-- Индексы таблицы `lesson`
--
ALTER TABLE `lesson`
  ADD PRIMARY KEY (`Id`),
  ADD KEY `SubjectID` (`SubjectID`),
  ADD KEY `Room` (`Room`);

--
-- Индексы таблицы `student`
--
ALTER TABLE `student`
  ADD PRIMARY KEY (`Id`);

--
-- Индексы таблицы `subject`
--
ALTER TABLE `subject`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_TeacherID` (`TeacherID`);

--
-- Индексы таблицы `subject-group-bridge`
--
ALTER TABLE `subject-group-bridge`
  ADD PRIMARY KEY (`GroupID`,`Subjectid`),
  ADD KEY `GroupID` (`GroupID`),
  ADD KEY `Subjectid` (`Subjectid`);

--
-- Индексы таблицы `subject-lesson-group-bridge`
--
ALTER TABLE `subject-lesson-group-bridge`
  ADD PRIMARY KEY (`GroupId`,`Lessonid`,`Subjectid`),
  ADD KEY `Lessonid` (`Lessonid`),
  ADD KEY `Subjectid` (`Subjectid`);

--
-- Индексы таблицы `teacher`
--
ALTER TABLE `teacher`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT для сохранённых таблиц
--

--
-- AUTO_INCREMENT для таблицы `student`
--
ALTER TABLE `student`
  MODIFY `Id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5415374;

--
-- AUTO_INCREMENT для таблицы `subject`
--
ALTER TABLE `subject`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=103;

--
-- AUTO_INCREMENT для таблицы `subject-group-bridge`
--
ALTER TABLE `subject-group-bridge`
  MODIFY `Subjectid` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT для таблицы `teacher`
--
ALTER TABLE `teacher`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- Ограничения внешнего ключа сохраненных таблиц
--

--
-- Ограничения внешнего ключа таблицы `attendance`
--
ALTER TABLE `attendance`
  ADD CONSTRAINT `attendance_ibfk_1` FOREIGN KEY (`LessonId`) REFERENCES `lesson` (`Id`),
  ADD CONSTRAINT `attendance_ibfk_2` FOREIGN KEY (`StudentId`) REFERENCES `student` (`Id`),
  ADD CONSTRAINT `attendance_ibfk_3` FOREIGN KEY (`LessonId`) REFERENCES `lesson` (`Id`),
  ADD CONSTRAINT `attendance_ibfk_4` FOREIGN KEY (`StudentId`) REFERENCES `student` (`Id`),
  ADD CONSTRAINT `attendance_ibfk_5` FOREIGN KEY (`LessonId`) REFERENCES `lesson` (`Id`),
  ADD CONSTRAINT `attendance_ibfk_6` FOREIGN KEY (`StudentId`) REFERENCES `student` (`Id`),
  ADD CONSTRAINT `fk_lesson_id` FOREIGN KEY (`LessonId`) REFERENCES `lesson` (`Id`),
  ADD CONSTRAINT `fk_student_id` FOREIGN KEY (`StudentId`) REFERENCES `student` (`Id`);

--
-- Ограничения внешнего ключа таблицы `group assign`
--
ALTER TABLE `group assign`
  ADD CONSTRAINT `group assign_ibfk_1` FOREIGN KEY (`GroupID`) REFERENCES `group` (`Id`),
  ADD CONSTRAINT `group assign_ibfk_2` FOREIGN KEY (`StudentID`) REFERENCES `student` (`Id`);

--
-- Ограничения внешнего ключа таблицы `lesson`
--
ALTER TABLE `lesson`
  ADD CONSTRAINT `lesson_ibfk_1` FOREIGN KEY (`SubjectID`) REFERENCES `subject` (`id`),
  ADD CONSTRAINT `lesson_ibfk_2` FOREIGN KEY (`Room`) REFERENCES `device` (`Room`);

--
-- Ограничения внешнего ключа таблицы `subject`
--
ALTER TABLE `subject`
  ADD CONSTRAINT `fk_TeacherID` FOREIGN KEY (`TeacherID`) REFERENCES `teacher` (`id`);

--
-- Ограничения внешнего ключа таблицы `subject-group-bridge`
--
ALTER TABLE `subject-group-bridge`
  ADD CONSTRAINT `subject-group-bridge_ibfk_1` FOREIGN KEY (`GroupID`) REFERENCES `group` (`Id`),
  ADD CONSTRAINT `subject-group-bridge_ibfk_2` FOREIGN KEY (`Subjectid`) REFERENCES `subject` (`id`);

--
-- Ограничения внешнего ключа таблицы `subject-lesson-group-bridge`
--
ALTER TABLE `subject-lesson-group-bridge`
  ADD CONSTRAINT `subject-lesson-group-bridge_ibfk_1` FOREIGN KEY (`GroupId`) REFERENCES `group` (`Id`),
  ADD CONSTRAINT `subject-lesson-group-bridge_ibfk_2` FOREIGN KEY (`Lessonid`) REFERENCES `lesson` (`Id`),
  ADD CONSTRAINT `subject-lesson-group-bridge_ibfk_3` FOREIGN KEY (`Subjectid`) REFERENCES `subject` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
