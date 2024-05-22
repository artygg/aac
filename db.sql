

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `db`
--

-- --------------------------------------------------------

--
-- Table structure for table `attendance`
--

CREATE TABLE `attendance` (
  `ClassId` int(11) NOT NULL,
  `StudentId` int(11) NOT NULL,
  `Time` datetime NOT NULL,
  `Status` varchar(15) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `class`
--

CREATE TABLE `class` (
  `Id` int(11) NOT NULL,
  `CourseID` int(11) NOT NULL,
  `Room` varchar(10) NOT NULL,
  `StartTime` datetime NOT NULL,
  `EndTime` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `class`
--

INSERT INTO `class` (`Id`, `CourseID`, `Room`, `StartTime`, `EndTime`) VALUES
(1, 101, 'Room A', '2024-05-20 09:00:00', '2024-05-20 10:30:00'),
(2, 102, 'Room B', '2024-05-20 11:00:00', '2024-05-20 12:30:00');

-- --------------------------------------------------------

--
-- Table structure for table `course`
--

CREATE TABLE `course` (
  `id` int(11) NOT NULL,
  `Name` varchar(60) NOT NULL,
  `TeacherID` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `course`
--

INSERT INTO `course` (`id`, `Name`, `TeacherID`) VALUES
(101, 'Mathematics', 1),
(102, 'Physics', 2);

-- --------------------------------------------------------

--
-- Table structure for table `course-class-group-bridge`
--

CREATE TABLE `course-class-group-bridge` (
  `GroupId` varchar(20) NOT NULL,
  `Classid` int(11) NOT NULL,
  `Coursetid` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `course-group-bridge`
--

CREATE TABLE `course-group-bridge` (
  `GroupID` varchar(20) NOT NULL,
  `Courseid` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `device`
--

CREATE TABLE `device` (
  `MAC` varchar(12) NOT NULL,
  `Key` int(60) NOT NULL,
  `Room` varchar(10) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `device`
--

INSERT INTO `device` (`MAC`, `Key`, `Room`) VALUES
('00:11:22:33', 123456, 'Room A'),
('66:77:88:99', 654321, 'Room B');

-- --------------------------------------------------------

--
-- Table structure for table `group`
--

CREATE TABLE `group` (
  `Id` varchar(20) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `group`
--

INSERT INTO `group` (`Id`) VALUES
('1'),
('2');

-- --------------------------------------------------------

--
-- Table structure for table `group assign`
--

CREATE TABLE `group assign` (
  `GroupID` varchar(20) DEFAULT NULL,
  `StudentID` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `group assign`
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
-- Table structure for table `student`
--

CREATE TABLE `student` (
  `Id` int(11) NOT NULL,
  `FirstName` varchar(60) NOT NULL,
  `LastName` varchar(60) NOT NULL,
  `Email` varchar(60) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `student`
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
-- Table structure for table `teacher`
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
-- Dumping data for table `teacher`
--

INSERT INTO `teacher` (`id`, `firstName`, `lastName`, `email`, `password`, `registrationDate`) VALUES
(1, 'John', 'Doe', 'john.doe@nhlstenden.com', 'password123', '2024-01-15'),
(2, 'Jane', 'Smith', 'jane.smith@nhlstenden.com', 'smithjane2024', '2024-02-10'),
(3, 'Michael', 'Brown', 'michael.brown@nhlstenden.com', 'mikebrown2024', '2024-03-05');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `attendance`
--
ALTER TABLE `attendance`
  ADD PRIMARY KEY (`ClassId`,`StudentId`),
  ADD KEY `fk_student_id` (`StudentId`);

--
-- Indexes for table `class`
--
ALTER TABLE `class`
  ADD PRIMARY KEY (`Id`),
  ADD KEY `SubjectID` (`CourseID`),
  ADD KEY `Room` (`Room`);

--
-- Indexes for table `course`
--
ALTER TABLE `course`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_TeacherID` (`TeacherID`);

--
-- Indexes for table `course-class-group-bridge`
--
ALTER TABLE `course-class-group-bridge`
  ADD PRIMARY KEY (`GroupId`,`Classid`,`Coursetid`),
  ADD KEY `Lessonid` (`Classid`),
  ADD KEY `Subjectid` (`Coursetid`);

--
-- Indexes for table `course-group-bridge`
--
ALTER TABLE `course-group-bridge`
  ADD PRIMARY KEY (`GroupID`,`Courseid`),
  ADD KEY `GroupID` (`GroupID`),
  ADD KEY `Subjectid` (`Courseid`);

--
-- Indexes for table `device`
--
ALTER TABLE `device`
  ADD PRIMARY KEY (`MAC`),
  ADD UNIQUE KEY `Room` (`Room`);

--
-- Indexes for table `group`
--
ALTER TABLE `group`
  ADD PRIMARY KEY (`Id`);

--
-- Indexes for table `group assign`
--
ALTER TABLE `group assign`
  ADD KEY `GroupID` (`GroupID`),
  ADD KEY `StudentID` (`StudentID`);

--
-- Indexes for table `student`
--
ALTER TABLE `student`
  ADD PRIMARY KEY (`Id`);

--
-- Indexes for table `teacher`
--
ALTER TABLE `teacher`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `course`
--
ALTER TABLE `course`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=103;

--
-- AUTO_INCREMENT for table `course-group-bridge`
--
ALTER TABLE `course-group-bridge`
  MODIFY `Courseid` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `student`
--
ALTER TABLE `student`
  MODIFY `Id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5415374;

--
-- AUTO_INCREMENT for table `teacher`
--
ALTER TABLE `teacher`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `attendance`
--
ALTER TABLE `attendance`
  ADD CONSTRAINT `attendance_ibfk_1` FOREIGN KEY (`ClassId`) REFERENCES `class` (`Id`),
  ADD CONSTRAINT `attendance_ibfk_2` FOREIGN KEY (`StudentId`) REFERENCES `student` (`Id`),
  ADD CONSTRAINT `attendance_ibfk_3` FOREIGN KEY (`ClassId`) REFERENCES `class` (`Id`),
  ADD CONSTRAINT `attendance_ibfk_4` FOREIGN KEY (`StudentId`) REFERENCES `student` (`Id`),
  ADD CONSTRAINT `attendance_ibfk_5` FOREIGN KEY (`ClassId`) REFERENCES `class` (`Id`),
  ADD CONSTRAINT `attendance_ibfk_6` FOREIGN KEY (`StudentId`) REFERENCES `student` (`Id`),
  ADD CONSTRAINT `fk_lesson_id` FOREIGN KEY (`ClassId`) REFERENCES `class` (`Id`),
  ADD CONSTRAINT `fk_student_id` FOREIGN KEY (`StudentId`) REFERENCES `student` (`Id`);

--
-- Constraints for table `class`
--
ALTER TABLE `class`
  ADD CONSTRAINT `class_ibfk_1` FOREIGN KEY (`CourseID`) REFERENCES `course` (`id`),
  ADD CONSTRAINT `class_ibfk_2` FOREIGN KEY (`Room`) REFERENCES `device` (`Room`);

--
-- Constraints for table `course`
--
ALTER TABLE `course`
  ADD CONSTRAINT `fk_TeacherID` FOREIGN KEY (`TeacherID`) REFERENCES `teacher` (`id`);

--
-- Constraints for table `course-class-group-bridge`
--
ALTER TABLE `course-class-group-bridge`
  ADD CONSTRAINT `course-class-group-bridge_ibfk_1` FOREIGN KEY (`GroupId`) REFERENCES `group` (`Id`),
  ADD CONSTRAINT `course-class-group-bridge_ibfk_2` FOREIGN KEY (`Classid`) REFERENCES `class` (`Id`),
  ADD CONSTRAINT `course-class-group-bridge_ibfk_3` FOREIGN KEY (`Coursetid`) REFERENCES `course` (`id`);

--
-- Constraints for table `course-group-bridge`
--
ALTER TABLE `course-group-bridge`
  ADD CONSTRAINT `course-group-bridge_ibfk_1` FOREIGN KEY (`GroupID`) REFERENCES `group` (`Id`),
  ADD CONSTRAINT `course-group-bridge_ibfk_2` FOREIGN KEY (`Courseid`) REFERENCES `course` (`id`);

--
-- Constraints for table `group assign`
--
ALTER TABLE `group assign`
  ADD CONSTRAINT `group assign_ibfk_1` FOREIGN KEY (`GroupID`) REFERENCES `group` (`Id`),
  ADD CONSTRAINT `group assign_ibfk_2` FOREIGN KEY (`StudentID`) REFERENCES `student` (`Id`);

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
