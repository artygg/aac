

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `aas_database`
--

-- --------------------------------------------------------

--
-- Table structure for table `attendance`
--

CREATE TABLE `attendance` (
  `LessonId` int(11) NOT NULL,
  `StudentId` int(11) NOT NULL,
  `Time` datetime NOT NULL
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

-- --------------------------------------------------------

--
-- Table structure for table `group`
--

CREATE TABLE `group` (
  `Id` varchar(20) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `group assign`
--

CREATE TABLE `group assign` (
  `GroupID` varchar(20) DEFAULT NULL,
  `StudentID` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `lesson`
--

CREATE TABLE `lesson` (
  `Id` int(11) NOT NULL,
  `SubjectID` int(11) NOT NULL,
  `Room` varchar(10) NOT NULL,
  `StartTime` datetime NOT NULL,
  `EndTime` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

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

-- --------------------------------------------------------

--
-- Table structure for table `subject`
--

CREATE TABLE `subject` (
  `id` int(11) NOT NULL,
  `Name` varchar(60) NOT NULL,
  `TeacherID` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `subject-group-bridge`
--

CREATE TABLE `subject-group-bridge` (
  `GroupID` varchar(20) NOT NULL,
  `Subjectid` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `subject-lesson-group-bridge`
--

CREATE TABLE `subject-lesson-group-bridge` (
  `GroupId` varchar(20) NOT NULL,
  `Lessonid` int(11) NOT NULL,
  `Subjectid` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

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
-- Indexes for dumped tables
--

--
-- Indexes for table `attendance`
--
ALTER TABLE `attendance`
  ADD PRIMARY KEY (`LessonId`,`StudentId`),
  ADD KEY `fk_student_id` (`StudentId`);

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
-- Indexes for table `lesson`
--
ALTER TABLE `lesson`
  ADD PRIMARY KEY (`Id`),
  ADD KEY `SubjectID` (`SubjectID`),
  ADD KEY `Room` (`Room`);

--
-- Indexes for table `student`
--
ALTER TABLE `student`
  ADD PRIMARY KEY (`Id`);

--
-- Indexes for table `subject`
--
ALTER TABLE `subject`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_TeacherID` (`TeacherID`);

--
-- Indexes for table `subject-group-bridge`
--
ALTER TABLE `subject-group-bridge`
  ADD PRIMARY KEY (`GroupID`,`Subjectid`),
  ADD KEY `GroupID` (`GroupID`),
  ADD KEY `Subjectid` (`Subjectid`);

--
-- Indexes for table `subject-lesson-group-bridge`
--
ALTER TABLE `subject-lesson-group-bridge`
  ADD PRIMARY KEY (`GroupId`,`Lessonid`,`Subjectid`),
  ADD KEY `Lessonid` (`Lessonid`),
  ADD KEY `Subjectid` (`Subjectid`);

--
-- Indexes for table `teacher`
--
ALTER TABLE `teacher`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `student`
--
ALTER TABLE `student`
  MODIFY `Id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `subject`
--
ALTER TABLE `subject`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `subject-group-bridge`
--
ALTER TABLE `subject-group-bridge`
  MODIFY `Subjectid` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `teacher`
--
ALTER TABLE `teacher`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `attendance`
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
-- Constraints for table `group assign`
--
ALTER TABLE `group assign`
  ADD CONSTRAINT `group assign_ibfk_1` FOREIGN KEY (`GroupID`) REFERENCES `group` (`Id`),
  ADD CONSTRAINT `group assign_ibfk_2` FOREIGN KEY (`StudentID`) REFERENCES `student` (`Id`);

--
-- Constraints for table `lesson`
--
ALTER TABLE `lesson`
  ADD CONSTRAINT `lesson_ibfk_1` FOREIGN KEY (`SubjectID`) REFERENCES `subject` (`id`),
  ADD CONSTRAINT `lesson_ibfk_2` FOREIGN KEY (`Room`) REFERENCES `device` (`Room`);

--
-- Constraints for table `subject`
--
ALTER TABLE `subject`
  ADD CONSTRAINT `fk_TeacherID` FOREIGN KEY (`TeacherID`) REFERENCES `teacher` (`id`);

--
-- Constraints for table `subject-group-bridge`
--
ALTER TABLE `subject-group-bridge`
  ADD CONSTRAINT `subject-group-bridge_ibfk_1` FOREIGN KEY (`GroupID`) REFERENCES `group` (`Id`),
  ADD CONSTRAINT `subject-group-bridge_ibfk_2` FOREIGN KEY (`Subjectid`) REFERENCES `subject` (`id`);

--
-- Constraints for table `subject-lesson-group-bridge`
--
ALTER TABLE `subject-lesson-group-bridge`
  ADD CONSTRAINT `subject-lesson-group-bridge_ibfk_1` FOREIGN KEY (`GroupId`) REFERENCES `group` (`Id`),
  ADD CONSTRAINT `subject-lesson-group-bridge_ibfk_2` FOREIGN KEY (`Lessonid`) REFERENCES `lesson` (`Id`),
  ADD CONSTRAINT `subject-lesson-group-bridge_ibfk_3` FOREIGN KEY (`Subjectid`) REFERENCES `subject` (`id`);

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
