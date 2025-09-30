// 代码生成时间: 2025-09-30 23:01:24
package main

import (
    "buffalo"
    "github.com/gobuffalo/buffalo-pop/v2/pop"
    "github.com/unrolled/render"
)

// Student model represents a student entity
type Student struct {
    ID       uint   "db:id,pk"
    Name     string "db:name"
    Age      int    "db:age"
    Grade    string "db:grade"
    CreatedAt string "db:created_at"
    UpdatedAt string "db:updated_at"
}

// StudentService handles the business logic for student operations
type StudentService struct {
    DB *pop.Connection
    R  *render.Render
}

// NewStudentService creates a new instance of StudentService
func NewStudentService(db *pop.Connection, r *render.Render) *StudentService {
    return &StudentService{DB: db, R: r}
}

// CreateStudent adds a new student to the database
func (s *StudentService) CreateStudent(student *Student) error {
    // Validate student data (omitted for brevity)
    // ...
    
    // Use the DB connection to save the student
    err := s.DB.Create(student)
    if err != nil {
        return err
    }
    
    return nil
}

// GetStudent retrieves a student by ID
func (s *StudentService) GetStudent(id uint) (*Student, error) {
    var student Student
    err := s.DB.FindBy("id", id, &student)
    if err != nil {
        return nil, err
    }
    
    return &student, nil
}

// UpdateStudent updates an existing student in the database
func (s *StudentService) UpdateStudent(student *Student) error {
    // Validate student data (omitted for brevity)
    // ...
    
    // Use the DB connection to update the student
    err := s.DB.Update(student)
    if err != nil {
        return err
    }
    
    return nil
}

// DeleteStudent removes a student from the database
func (s *StudentService) DeleteStudent(id uint) error {
    var student Student
    err := s.DB.Destroy(student, id)
    if err != nil {
        return err
    }
    
    return nil
}

// StudentHandler handles HTTP requests related to students
type StudentHandler struct {
    Context buffalo.Context
    Service *StudentService
}

// ListStudents responds with a list of students
func (h *StudentHandler) ListStudents(c buffalo.Context) error {
    students := []Student{}
    err := h.Service.DB.All(&students)
    if err != nil {
        return c.Error(500, err)
    }
    
    return c.Render(200, r.Data(map[string]interface{}{"students": students}))
}

// ShowStudent responds with a single student
func (h *StudentHandler) ShowStudent(c buffalo.Context) error {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.Error(400, errors.New("Invalid student ID"))
    }
    
    student, err := h.Service.GetStudent(uint(id))
    if err != nil {
        if err == sql.ErrNoRows {
            return c.Error(404, errors.New("Student not found"))
        }
        return c.Error(500, err)
    }
    
    return c.Render(200, r.JSON(student))
}

// NewStudent responds with a blank form for a new student
func (h *StudentHandler) NewStudent(c buffalo.Context) error {
    return c.Render(200, r.Data(map[string]interface{}{"student": Student{}}))
}

// CreateStudent handles the POST request for creating a student
func (h *StudentHandler) CreateStudent(c buffalo.Context) error {
    var student Student
    if err := c.Bind(&student); err != nil {
        return c.Error(400, err)
    }
    if err := h.Service.CreateStudent(&student); err != nil {
        return c.Error(500, err)
    }
    
    return c.Render(201, r.JSON(student))
}

// EditStudent responds with a student for editing
func (h *StudentHandler) EditStudent(c buffalo.Context) error {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.Error(400, errors.New("Invalid student ID"))
    }
    
    student, err := h.Service.GetStudent(uint(id))
    if err != nil {
        if err == sql.ErrNoRows {
            return c.Error(404, errors.New("Student not found"))
        }
        return c.Error(500, err)
    }
    
    return c.Render(200, r.Data(map[string]interface{}{"student": student}))
}

// UpdateStudent handles the PUT request for updating a student
func (h *StudentHandler) UpdateStudent(c buffalo.Context) error {
    var student Student
    if err := c.Bind(&student); err != nil {
        return c.Error(400, err)
    }
    if err := h.Service.UpdateStudent(&student); err != nil {
        return c.Error(500, err)
    }
    
    return c.Render(200, r.JSON(student))
}

// DeleteStudent handles the DELETE request for deleting a student
func (h *StudentHandler) DeleteStudent(c buffalo.Context) error {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.Error(400, errors.New("Invalid student ID"))
    }
    
    if err := h.Service.DeleteStudent(uint(id)); err != nil {
        return c.Error(500, err)
    }
    
    return c.Render(200, r.Data(map[string]interface{}{"message": "Student deleted successfully"}))
}

func main() {
    app := buffalo.Automatic(buffalo.Options{
        Address: ":3000", // Serve at 0.0.0.0:3000
    })

    // Set up renderer to use HTML templates
    app.Renderer = render.New(render.Options{
        Directory: "templates",
        IsDevelopment: true,
    })

    // Set up database connection
    // Replace with your actual database credentials
    app.DB = pop.Connect("development")

    // Create a new instance of StudentService
    studentService := NewStudentService(app.DB, app.Renderer)

    // Create a new instance of StudentHandler
    studentHandler := StudentHandler{
        Service: studentService,
    }

    // Define routes
    app.GET("/students", studentHandler.ListStudents)
    app.GET("/students/{id}", studentHandler.ShowStudent)
    app.GET("/students/new", studentHandler.NewStudent)
    app.POST("/students", studentHandler.CreateStudent)
    app.GET("/students/{id}/edit", studentHandler.EditStudent)
    app.PUT("/students/{id}", studentHandler.UpdateStudent)
    app.DELETE("/students/{id}", studentHandler.DeleteStudent)

    // Start the application
    app.Serve()
}