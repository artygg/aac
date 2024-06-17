# Automated Attendance 

## File Organization
Group similar code into separate files.

## Package Naming
Use short, lower case, single-word names. No underscores or mixed caps.

## Imports
Group imports into standard library, third-party packages, and local packages. Use blank lines to separate groups.

```bash
import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"regexp"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"

)
```

## Variables and Constants
   - Use camelCase for variables and constants.
   - Use descriptive names, except for local variables with a very small scope.

```bash
var attendanceRecords []Attendance
```

## Structs

- Use camelCase for field names.
- Capitalize the first letter for exported fields or methods (to make them public).

```
type User struct {
    ID        int
    FirstName string
    LastName  string
}
```

## Functions
   - Use camelCase for function and method names.
   - Function and method names should be descriptive.

``` bash
func (teacher *Teacher) getCourses(db *sql.DB) error {
	qwery := fmt.Sprintf("SELECT `id`, `Name`, `TeacherID` FROM `courses` WHERE `TeacherID`= '%v'", teacher.Id)
	rows, err := db.Query(qwery)
	if err != nil {
		return err
	}
	return nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
```

## Error Handling
- Use error type for error handling.
- Handle errors immediately after calling a function that might return an error.

``` bash
tx, err := db.Begin()
if err != nil {
	log.Println("Error starting transaction:", err)
	return err
}
```

## Comments
   - Use comments to explain why something is done, not what is done.
   - Use complete sentences, starting with a capital letter.


## Indentation and Line Length
- Use tabs for indentation.
- There is no limit to the length of characters for a single coding line.

``` bash
query := fmt.Sprintf("SELECT a.ClassId, a.Status, a.Time, s.Id, s.FirstName, s.LastName, s.Email FROM attendances a INNER JOIN classes c ON a.ClassId = c.Id INNER JOIN students s ON a.StudentId = s.Id WHERE c.CourseID = %v ", courseID)
rows, err := db.Query(query)
if err != nil {
	log.Println("Error executing query:", err)
	return nil, fmt.Errorf("failed to execute query: %w", err)
}
defer func() {
	if err := rows.Close(); err != nil {
		log.Println("Error closing rows:", err)
	}
}()
```


