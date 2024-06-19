-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Хост: db
-- Время создания: Июн 18 2024 г., 15:53
-- Версия сервера: 8.0.37
-- Версия PHP: 8.2.8

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- База данных: `aas`
--

-- --------------------------------------------------------

--
-- Структура таблицы `attendances`
--

CREATE TABLE `attendances` (
  `ClassId` int NOT NULL,
  `StudentId` int NOT NULL,
  `Time` datetime NOT NULL,
  `Status` varchar(15) COLLATE utf8mb4_general_ci NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Дамп данных таблицы `attendances`
--

INSERT INTO `attendances` (`ClassId`, `StudentId`, `Time`, `Status`) VALUES
(1, 5320259, '2024-06-18 10:21:02', '1');

-- --------------------------------------------------------

--
-- Структура таблицы `classes`
--

CREATE TABLE `classes` (
  `Id` int NOT NULL,
  `CourseID` int NOT NULL,
  `Room` varchar(10) COLLATE utf8mb4_general_ci NOT NULL,
  `StartTime` datetime NOT NULL,
  `EndTime` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Дамп данных таблицы `classes`
--

INSERT INTO `classes` (`Id`, `CourseID`, `Room`, `StartTime`, `EndTime`) VALUES
(1, 101, 'Room A', '2024-05-20 09:00:00', '2024-06-19 10:30:00'),
(2, 102, 'Room B', '2024-05-20 11:00:00', '2024-05-20 12:30:00');

-- --------------------------------------------------------

--
-- Структура таблицы `classes-groups-bridge`
--

CREATE TABLE `classes-groups-bridge` (
  `GroupID` varchar(20) COLLATE utf8mb4_general_ci NOT NULL,
  `Classid` int NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Структура таблицы `courses`
--

CREATE TABLE `courses` (
  `id` int NOT NULL,
  `Name` varchar(60) COLLATE utf8mb4_general_ci NOT NULL,
  `TeacherID` int NOT NULL,
  `Year` int NOT NULL,
  `StartDate` date NOT NULL,
  `EndDate` date NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Дамп данных таблицы `courses`
--

INSERT INTO `courses` (`id`, `Name`, `TeacherID`, `Year`, `StartDate`, `EndDate`) VALUES
(101, 'Project INNOVATE', 1, 0, '2024-06-17', '2024-06-21'),
(102, 'Physics', 2, 0, '0000-00-00', '0000-00-00'),
(103, 'PROFESSIONAL SKILLS', 1, 1, '2024-06-17', '2024-06-19'),
(104, 'Course1', 1, 1, '2024-06-17', '2024-06-18'),
(105, 'FJODORS', 1, 1, '2024-06-17', '2024-06-21'),
(106, 'Course 1', 4, 1, '2024-06-17', '2024-06-30');

-- --------------------------------------------------------

--
-- Структура таблицы `courses-groups-bridge`
--

CREATE TABLE `courses-groups-bridge` (
  `GroupID` varchar(20) COLLATE utf8mb4_general_ci NOT NULL,
  `Courseid` int NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Дамп данных таблицы `courses-groups-bridge`
--

INSERT INTO `courses-groups-bridge` (`GroupID`, `Courseid`) VALUES
('1', 103),
('2', 103),
('1', 104),
('2', 104),
('1', 105),
('2', 105),
('1', 106),
('2', 106);

-- --------------------------------------------------------

--
-- Структура таблицы `devices`
--

CREATE TABLE `devices` (
  `MAC` varchar(12) COLLATE utf8mb4_general_ci NOT NULL,
  `Key` int NOT NULL,
  `Room` varchar(10) COLLATE utf8mb4_general_ci NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Дамп данных таблицы `devices`
--

INSERT INTO `devices` (`MAC`, `Key`, `Room`) VALUES
('00:11:22:33', 123456, 'Room A'),
('66:77:88:99', 654321, 'Room B');

-- --------------------------------------------------------

--
-- Структура таблицы `groups`
--

CREATE TABLE `groups` (
  `Id` varchar(20) COLLATE utf8mb4_general_ci NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Дамп данных таблицы `groups`
--

INSERT INTO `groups` (`Id`) VALUES
('1'),
('2');

-- --------------------------------------------------------

--
-- Структура таблицы `groups assign`
--

CREATE TABLE `groups assign` (
  `GroupID` varchar(20) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `StudentID` int DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Дамп данных таблицы `groups assign`
--

INSERT INTO `groups assign` (`GroupID`, `StudentID`) VALUES
('1', 5139872),
('1', 5304377),
('2', 5320259),
('2', 5332370),
('2', 5357608);

-- --------------------------------------------------------

--
-- Структура таблицы `students`
--

CREATE TABLE `students` (
  `Id` int NOT NULL,
  `FirstName` varchar(60) COLLATE utf8mb4_general_ci NOT NULL,
  `LastName` varchar(60) COLLATE utf8mb4_general_ci NOT NULL,
  `Email` varchar(60) COLLATE utf8mb4_general_ci NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Дамп данных таблицы `students`
--

INSERT INTO `students` (`Id`, `FirstName`, `LastName`, `Email`) VALUES
(5139872, 'Sofronie', 'Albu', 'sofronie.albu@student.nhlstenden.com'),
(5292573, 'Artjoms', 'Grisajevs', 'artjoms.grisajevs@student.nhlstenden.com'),
(5304377, 'Nikita', 'Golovanov', 'nikita.golovanov@student.nhlstenden.com'),
(5320259, 'Yaroslav', 'Oleinychenko', 'yaroslav.oleinychenko@student.nhlstenden.com'),
(5332370, 'Ekaterina', 'Tarlykova', 'ekaterina.tarlykova@student.nhlstenden.com'),
(5357608, 'Fjodor', 'Smorodins', 'fjodor.smorodins@student.nhlstenden.com'),
(5415373, 'Sebastian', 'Serban', 'sebastian.serban@student.nhlstenden.com');

-- --------------------------------------------------------

--
-- Структура таблицы `teachers`
--

CREATE TABLE `teachers` (
  `id` int NOT NULL,
  `firstName` varchar(60) COLLATE utf8mb4_general_ci NOT NULL,
  `lastName` varchar(60) COLLATE utf8mb4_general_ci NOT NULL,
  `email` varchar(60) COLLATE utf8mb4_general_ci NOT NULL,
  `password` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  `registrationDate` date NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Дамп данных таблицы `teachers`
--

INSERT INTO `teachers` (`id`, `firstName`, `lastName`, `email`, `password`, `registrationDate`) VALUES
(1, 'John', 'Doe', 'john.doe@nhlstenden.com', '$2a$14$N7qhcsOTq1acHuIf4mnBQuyX/ivZxUFZHszWRJvAngESmYG00u/YS', '2024-01-15'),
(2, 'Jane', 'Smith', 'jane.smith@nhlstenden.com', '$2a$14$N7qhcsOTq1acHuIf4mnBQuyX/ivZxUFZHszWRJvAngESmYG00u/YS', '2024-02-10'),
(3, 'Michael', 'Brown', 'michael.brown@nhlstenden.com', '$2a$14$N7qhcsOTq1acHuIf4mnBQuyX/ivZxUFZHszWRJvAngESmYG00u/YS', '2024-03-05'),
(4, 'Artjoms', 'Grisajevs', 'agrisajevs@nhlstenden.com', '$2a$14$N7qhcsOTq1acHuIf4mnBQuyX/ivZxUFZHszWRJvAngESmYG00u/YS', '2024-06-18');

--
-- Индексы сохранённых таблиц
--

--
-- Индексы таблицы `attendances`
--
ALTER TABLE `attendances`
  ADD PRIMARY KEY (`ClassId`,`StudentId`),
  ADD KEY `fk_student_id` (`StudentId`);

--
-- Индексы таблицы `classes`
--
ALTER TABLE `classes`
  ADD PRIMARY KEY (`Id`),
  ADD KEY `SubjectID` (`CourseID`),
  ADD KEY `Room` (`Room`);

--
-- Индексы таблицы `classes-groups-bridge`
--
ALTER TABLE `classes-groups-bridge`
  ADD PRIMARY KEY (`GroupID`,`Classid`),
  ADD KEY `classes-groups-bridge-classesid` (`Classid`);

--
-- Индексы таблицы `courses`
--
ALTER TABLE `courses`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_TeacherID` (`TeacherID`);

--
-- Индексы таблицы `courses-groups-bridge`
--
ALTER TABLE `courses-groups-bridge`
  ADD PRIMARY KEY (`GroupID`,`Courseid`),
  ADD KEY `GroupID` (`GroupID`),
  ADD KEY `Subjectid` (`Courseid`);

--
-- Индексы таблицы `devices`
--
ALTER TABLE `devices`
  ADD PRIMARY KEY (`MAC`),
  ADD UNIQUE KEY `Room` (`Room`);

--
-- Индексы таблицы `groups`
--
ALTER TABLE `groups`
  ADD PRIMARY KEY (`Id`);

--
-- Индексы таблицы `groups assign`
--
ALTER TABLE `groups assign`
  ADD KEY `GroupID` (`GroupID`),
  ADD KEY `StudentID` (`StudentID`);

--
-- Индексы таблицы `students`
--
ALTER TABLE `students`
  ADD PRIMARY KEY (`Id`);

--
-- Индексы таблицы `teachers`
--
ALTER TABLE `teachers`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT для сохранённых таблиц
--

--
-- AUTO_INCREMENT для таблицы `classes`
--
ALTER TABLE `classes`
  MODIFY `Id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT для таблицы `classes-groups-bridge`
--
ALTER TABLE `classes-groups-bridge`
  MODIFY `Classid` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=116;

--
-- AUTO_INCREMENT для таблицы `courses`
--
ALTER TABLE `courses`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=107;

--
-- AUTO_INCREMENT для таблицы `courses-groups-bridge`
--
ALTER TABLE `courses-groups-bridge`
  MODIFY `Courseid` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=107;

--
-- AUTO_INCREMENT для таблицы `students`
--
ALTER TABLE `students`
  MODIFY `Id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5415374;

--
-- AUTO_INCREMENT для таблицы `teachers`
--
ALTER TABLE `teachers`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- Ограничения внешнего ключа сохраненных таблиц
--

--
-- Ограничения внешнего ключа таблицы `attendances`
--
ALTER TABLE `attendances`
  ADD CONSTRAINT `fk_class_id` FOREIGN KEY (`ClassId`) REFERENCES `classes` (`Id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  ADD CONSTRAINT `fk_student_id` FOREIGN KEY (`StudentId`) REFERENCES `students` (`Id`);

--
-- Ограничения внешнего ключа таблицы `classes`
--
ALTER TABLE `classes`
  ADD CONSTRAINT `classes_ibfk_1` FOREIGN KEY (`CourseID`) REFERENCES `courses` (`id`),
  ADD CONSTRAINT `classes_ibfk_2` FOREIGN KEY (`Room`) REFERENCES `devices` (`Room`);

--
-- Ограничения внешнего ключа таблицы `classes-groups-bridge`
--
ALTER TABLE `classes-groups-bridge`
  ADD CONSTRAINT `classes-groups-bridge-classesid` FOREIGN KEY (`Classid`) REFERENCES `classes` (`Id`),
  ADD CONSTRAINT `classes-groups-bridge-groupid` FOREIGN KEY (`GroupID`) REFERENCES `groups` (`Id`);

--
-- Ограничения внешнего ключа таблицы `courses`
--
ALTER TABLE `courses`
  ADD CONSTRAINT `fk_TeacherID` FOREIGN KEY (`TeacherID`) REFERENCES `teachers` (`id`);

--
-- Ограничения внешнего ключа таблицы `courses-groups-bridge`
--
ALTER TABLE `courses-groups-bridge`
  ADD CONSTRAINT `courses-groups-bridge_ibfk_1` FOREIGN KEY (`GroupID`) REFERENCES `groups` (`Id`),
  ADD CONSTRAINT `courses-groups-bridge_ibfk_2` FOREIGN KEY (`Courseid`) REFERENCES `courses` (`id`);

--
-- Ограничения внешнего ключа таблицы `groups assign`
--
ALTER TABLE `groups assign`
  ADD CONSTRAINT `groups assign_ibfk_1` FOREIGN KEY (`GroupID`) REFERENCES `groups` (`Id`),
  ADD CONSTRAINT `groups assign_ibfk_2` FOREIGN KEY (`StudentID`) REFERENCES `students` (`Id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
