-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: May 27, 2024 at 02:15 PM
-- Server version: 10.4.32-MariaDB
-- PHP Version: 8.2.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `db`
--

-- --------------------------------------------------------

--
-- Table structure for table `attendances`
--

CREATE TABLE `attendances` (
  `ClassId` int(11) NOT NULL,
  `StudentId` int(11) NOT NULL,
  `Time` datetime NOT NULL,
  `Status` varchar(15) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `classes`
--

CREATE TABLE `classes` (
  `Id` int(11) NOT NULL,
  `CourseID` int(11) NOT NULL,
  `Room` varchar(10) NOT NULL,
  `StartTime` datetime NOT NULL,
  `EndTime` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `classes`
--

INSERT INTO `classes` (`Id`, `CourseID`, `Room`, `StartTime`, `EndTime`) VALUES
(1, 101, 'Room A', '2024-05-20 09:00:00', '2024-05-20 10:30:00'),
(2, 102, 'Room B', '2024-05-20 11:00:00', '2024-05-20 12:30:00');

-- --------------------------------------------------------

--
-- Table structure for table `classes-courses-bridge`
--

CREATE TABLE `classes-courses-bridge` (
  `ClassID` int(11) NOT NULL,
  `CourseID` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `courses`
--

CREATE TABLE `courses` (
  `id` int(11) NOT NULL,
  `Name` varchar(60) NOT NULL,
  `TeacherID` int(11) NOT NULL,
  `Year` int(11) NOT NULL,
  `StartDate` date NOT NULL,
  `EndDate` date NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `courses`
--

INSERT INTO `courses` (`id`, `Name`, `TeacherID`, `Year`, `StartDate`, `EndDate`) VALUES
(101, 'Mathematics', 1, 0, '0000-00-00', '0000-00-00'),
(102, 'Physics', 2, 0, '0000-00-00', '0000-00-00');

-- --------------------------------------------------------

--
-- Table structure for table `courses-groups-bridge`
--

CREATE TABLE `courses-groups-bridge` (
  `GroupID` varchar(20) NOT NULL,
  `Courseid` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `devices`
--

CREATE TABLE `devices` (
  `MAC` varchar(12) NOT NULL,
  `Key` int(60) NOT NULL,
  `Room` varchar(10) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `devices`
--

INSERT INTO `devices` (`MAC`, `Key`, `Room`) VALUES
('00:11:22:33', 123456, 'Room A'),
('66:77:88:99', 654321, 'Room B');

-- --------------------------------------------------------

--
-- Table structure for table `groups`
--

CREATE TABLE `groups` (
  `Id` varchar(20) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `groups`
--

INSERT INTO `groups` (`Id`) VALUES
('1'),
('2');

-- --------------------------------------------------------

--
-- Table structure for table `groups assign`
--

CREATE TABLE `groups assign` (
  `GroupID` varchar(20) DEFAULT NULL,
  `StudentID` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `groups assign`
--

INSERT INTO `groups assign` (`GroupID`, `StudentID`) VALUES
('1', 292573),
('1', 5139872),
('1', 5304377),
('2', 5320259),
('2', 5332370),
('2', 5357608);

-- --------------------------------------------------------

--
-- Table structure for table `students`
--

CREATE TABLE `students` (
  `Id` int(11) NOT NULL,
  `FirstName` varchar(60) NOT NULL,
  `LastName` varchar(60) NOT NULL,
  `Email` varchar(60) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `students`
--

INSERT INTO `students` (`Id`, `FirstName`, `LastName`, `Email`) VALUES
(292573, 'Artjoms', 'Grisajevs', 'artjoms.grisajevs@student.nhlstenden.com'),
(5139872, 'Sofronie', 'Albu', 'sofronie.albu@student.nhlstenden.com'),
(5304377, 'Nikita', 'Golovanov', 'nikita.golovanov@student.nhlstenden.com'),
(5320259, 'Yaroslav', 'Oleinychenko', 'yaroslav.oleinychenko@student.nhlstenden.com'),
(5332370, 'Ekaterina', 'Tarlykova', 'ekaterina.tarlykova@student.nhlstenden.com'),
(5357608, 'Fjodor', 'Smorodins', 'fjodor.smorodins@student.nhlstenden.com'),
(5415373, 'Sebastian', 'Serban', 'sebastian.serban@student.nhlstenden.com');

-- --------------------------------------------------------

--
-- Table structure for table `teachers`
--

CREATE TABLE `teachers` (
  `id` int(11) NOT NULL,
  `firstName` varchar(60) NOT NULL,
  `lastName` varchar(60) NOT NULL,
  `email` varchar(60) NOT NULL,
  `password` varchar(255) NOT NULL,
  `registrationDate` date NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `teachers`
--

INSERT INTO `teachers` (`id`, `firstName`, `lastName`, `email`, `password`, `registrationDate`) VALUES
(1, 'John', 'Doe', 'john.doe@nhlstenden.com', 'password123', '2024-01-15'),
(2, 'Jane', 'Smith', 'jane.smith@nhlstenden.com', 'smithjane2024', '2024-02-10'),
(3, 'Michael', 'Brown', 'michael.brown@nhlstenden.com', 'mikebrown2024', '2024-03-05');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `attendances`
--
ALTER TABLE `attendances`
  ADD PRIMARY KEY (`ClassId`,`StudentId`),
  ADD KEY `fk_student_id` (`StudentId`);

--
-- Indexes for table `classes`
--
ALTER TABLE `classes`
  ADD PRIMARY KEY (`Id`),
  ADD KEY `SubjectID` (`CourseID`),
  ADD KEY `Room` (`Room`);

--
-- Indexes for table `classes-courses-bridge`
--
ALTER TABLE `classes-courses-bridge`
  ADD PRIMARY KEY (`ClassID`,`CourseID`),
  ADD KEY `Lessonid` (`ClassID`),
  ADD KEY `Subjectid` (`CourseID`);

--
-- Indexes for table `courses`
--
ALTER TABLE `courses`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_TeacherID` (`TeacherID`);

--
-- Indexes for table `courses-groups-bridge`
--
ALTER TABLE `courses-groups-bridge`
  ADD PRIMARY KEY (`GroupID`,`Courseid`),
  ADD KEY `GroupID` (`GroupID`),
  ADD KEY `Subjectid` (`Courseid`);

--
-- Indexes for table `devices`
--
ALTER TABLE `devices`
  ADD PRIMARY KEY (`MAC`),
  ADD UNIQUE KEY `Room` (`Room`);

--
-- Indexes for table `groups`
--
ALTER TABLE `groups`
  ADD PRIMARY KEY (`Id`);

--
-- Indexes for table `groups assign`
--
ALTER TABLE `groups assign`
  ADD KEY `GroupID` (`GroupID`),
  ADD KEY `StudentID` (`StudentID`);

--
-- Indexes for table `students`
--
ALTER TABLE `students`
  ADD PRIMARY KEY (`Id`);

--
-- Indexes for table `teachers`
--
ALTER TABLE `teachers`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `courses`
--
ALTER TABLE `courses`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=103;

--
-- AUTO_INCREMENT for table `courses-groups-bridge`
--
ALTER TABLE `courses-groups-bridge`
  MODIFY `Courseid` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `students`
--
ALTER TABLE `students`
  MODIFY `Id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5415374;

--
-- AUTO_INCREMENT for table `teachers`
--
ALTER TABLE `teachers`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `attendances`
--
ALTER TABLE `attendances`
  ADD CONSTRAINT `attendances_ibfk_1` FOREIGN KEY (`ClassId`) REFERENCES `classes` (`Id`),
  ADD CONSTRAINT `attendances_ibfk_2` FOREIGN KEY (`StudentId`) REFERENCES `students` (`Id`),
  ADD CONSTRAINT `attendances_ibfk_3` FOREIGN KEY (`ClassId`) REFERENCES `classes` (`Id`),
  ADD CONSTRAINT `attendances_ibfk_4` FOREIGN KEY (`StudentId`) REFERENCES `students` (`Id`),
  ADD CONSTRAINT `attendances_ibfk_5` FOREIGN KEY (`ClassId`) REFERENCES `classes` (`Id`),
  ADD CONSTRAINT `attendances_ibfk_6` FOREIGN KEY (`StudentId`) REFERENCES `students` (`Id`),
  ADD CONSTRAINT `fk_lesson_id` FOREIGN KEY (`ClassId`) REFERENCES `classes` (`Id`),
  ADD CONSTRAINT `fk_student_id` FOREIGN KEY (`StudentId`) REFERENCES `students` (`Id`);

--
-- Constraints for table `classes`
--
ALTER TABLE `classes`
  ADD CONSTRAINT `classes_ibfk_1` FOREIGN KEY (`CourseID`) REFERENCES `courses` (`id`),
  ADD CONSTRAINT `classes_ibfk_2` FOREIGN KEY (`Room`) REFERENCES `devices` (`Room`);

--
-- Constraints for table `courses`
--
ALTER TABLE `courses`
  ADD CONSTRAINT `fk_TeacherID` FOREIGN KEY (`TeacherID`) REFERENCES `teachers` (`id`);

--
-- Constraints for table `courses-groups-bridge`
--
ALTER TABLE `courses-groups-bridge`
  ADD CONSTRAINT `courses-groups-bridge_ibfk_1` FOREIGN KEY (`GroupID`) REFERENCES `groups` (`Id`),
  ADD CONSTRAINT `courses-groups-bridge_ibfk_2` FOREIGN KEY (`Courseid`) REFERENCES `courses` (`id`);

--
-- Constraints for table `groups assign`
--
ALTER TABLE `groups assign`
  ADD CONSTRAINT `groups assign_ibfk_1` FOREIGN KEY (`GroupID`) REFERENCES `groups` (`Id`),
  ADD CONSTRAINT `groups assign_ibfk_2` FOREIGN KEY (`StudentID`) REFERENCES `students` (`Id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
