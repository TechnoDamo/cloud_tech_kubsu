// School Management System Frontend
function schoolApp() {
    return {
        // API Configuration
        apiBaseUrl: 'https://d5ds9jru50sk4al4aiaa.svoluuab.apigw.yandexcloud.net/api/v1',
        
        // Application State
        stats: {
            students: 0,
            teachers: 0,
            classes: 0,
            subjects: 0
        },
        
        // Data Storage
        classes: [],
        students: [],
        teachers: [],
        subjects: [],
        teacherAssignments: [],
        lessonSchedules: [],
        lessonLogs: [],
        studentLessons: [],
        attendanceStatuses: [],
        
        // UI State
        showClassForm: false,
        showStudentForm: false,
        showTeacherForm: false,
        showSubjectForm: false,
        showScheduleForm: false,
        showAttendanceForm: false,
        showDataModal: false,
        currentDataTitle: '',
        currentDataHeaders: [],
        currentData: [],
        
        // Form Data
        newClass: {
            grade: '',
            letter: ''
        },
        newStudent: {
            class_id: '',
            first_name: '',
            last_name: '',
            patronymic: ''
        },
        newTeacher: {
            first_name: '',
            last_name: '',
            patronymic: ''
        },
        newSubject: {
            subject_name: ''
        },
        newSchedule: {
            subject_id: '',
            weekday: '',
            number: '',
            class_id: '',
            teacher_id: ''
        },
        newAttendance: {
            student_id: '',
            lesson_id: '',
            grade: '',
            attendance_status: ''
        },
        
        // Notification System
        notification: {
            show: false,
            type: 'success',
            message: ''
        },
        
        // Initialize Application
        init() {
            this.loadStats();
            this.loadClasses();
            this.loadAttendanceStatuses();
        },
        
        // API Helper Functions
        async apiCall(endpoint, method = 'GET', data = null) {
            try {
                const options = {
                    method,
                    headers: {
                        'Content-Type': 'application/json',
                    }
                };
                
                if (data) {
                    options.body = JSON.stringify(data);
                }
                
                const response = await fetch(`${this.apiBaseUrl}${endpoint}`, options);
                
                if (!response.ok) {
                    throw new Error(`HTTP error! status: ${response.status}`);
                }
                
                return await response.json();
            } catch (error) {
                console.error('API call failed:', error);
                this.showNotification(`API Error: ${error.message}`, 'error');
                throw error;
            }
        },
        
        // Statistics Loading
        async loadStats() {
            try {
                const [classesData, studentsData, teachersData, subjectsData] = await Promise.all([
                    this.apiCall('/classes'),
                    this.apiCall('/students'),
                    this.apiCall('/teachers'),
                    this.apiCall('/subjects')
                ]);
                
                this.stats.classes = classesData.data?.length || 0;
                this.stats.students = studentsData.data?.length || 0;
                this.stats.teachers = teachersData.data?.length || 0;
                this.stats.subjects = subjectsData.data?.length || 0;
            } catch (error) {
                console.error('Failed to load stats:', error);
            }
        },
        
        // Classes Management
        async loadClasses() {
            try {
                const data = await this.apiCall('/classes');
                this.classes = data.data || [];
                this.showDataModal = true;
                this.currentDataTitle = 'Classes';
                this.currentDataHeaders = ['ID', 'Grade', 'Letter'];
                this.currentData = this.classes.map(cls => ({
                    id: cls.id,
                    Grade: cls.grade,
                    Letter: cls.letter
                }));
            } catch (error) {
                console.error('Failed to load classes:', error);
            }
        },
        
        async createClass() {
            try {
                await this.apiCall('/classes', 'POST', this.newClass);
                this.showNotification('Class created successfully!', 'success');
                this.showClassForm = false;
                this.newClass = { grade: '', letter: '' };
                this.loadStats();
                this.loadClasses();
            } catch (error) {
                console.error('Failed to create class:', error);
            }
        },
        
        // Students Management
        async loadStudents() {
            try {
                const data = await this.apiCall('/students');
                this.students = data.data || [];
                this.showDataModal = true;
                this.currentDataTitle = 'Students';
                this.currentDataHeaders = ['ID', 'Class ID', 'First Name', 'Last Name', 'Patronymic'];
                this.currentData = this.students.map(student => ({
                    id: student.id,
                    'Class ID': student.class_id,
                    'First Name': student.first_name,
                    'Last Name': student.last_name,
                    'Patronymic': student.patronymic || '-'
                }));
            } catch (error) {
                console.error('Failed to load students:', error);
            }
        },
        
        async createStudent() {
            try {
                await this.apiCall('/students', 'POST', this.newStudent);
                this.showNotification('Student created successfully!', 'success');
                this.showStudentForm = false;
                this.newStudent = { class_id: '', first_name: '', last_name: '', patronymic: '' };
                this.loadStats();
                this.loadStudents();
            } catch (error) {
                console.error('Failed to create student:', error);
            }
        },
        
        // Teachers Management
        async loadTeachers() {
            try {
                const data = await this.apiCall('/teachers');
                this.teachers = data.data || [];
                this.showDataModal = true;
                this.currentDataTitle = 'Teachers';
                this.currentDataHeaders = ['ID', 'First Name', 'Last Name', 'Patronymic'];
                this.currentData = this.teachers.map(teacher => ({
                    id: teacher.id,
                    'First Name': teacher.first_name,
                    'Last Name': teacher.last_name,
                    'Patronymic': teacher.patronymic || '-'
                }));
            } catch (error) {
                console.error('Failed to load teachers:', error);
            }
        },
        
        async createTeacher() {
            try {
                await this.apiCall('/teachers', 'POST', this.newTeacher);
                this.showNotification('Teacher created successfully!', 'success');
                this.showTeacherForm = false;
                this.newTeacher = { first_name: '', last_name: '', patronymic: '' };
                this.loadStats();
                this.loadTeachers();
            } catch (error) {
                console.error('Failed to create teacher:', error);
            }
        },
        
        // Subjects Management
        async loadSubjects() {
            try {
                const data = await this.apiCall('/subjects');
                this.subjects = data.data || [];
                this.showDataModal = true;
                this.currentDataTitle = 'Subjects';
                this.currentDataHeaders = ['ID', 'Subject Name'];
                this.currentData = this.subjects.map(subject => ({
                    id: subject.id,
                    'Subject Name': subject.subject_name
                }));
            } catch (error) {
                console.error('Failed to load subjects:', error);
            }
        },
        
        async createSubject() {
            try {
                await this.apiCall('/subjects', 'POST', this.newSubject);
                this.showNotification('Subject created successfully!', 'success');
                this.showSubjectForm = false;
                this.newSubject = { subject_name: '' };
                this.loadStats();
                this.loadSubjects();
            } catch (error) {
                console.error('Failed to create subject:', error);
            }
        },
        
        // Teacher Assignments Management
        async loadTeacherAssignments() {
            try {
                const data = await this.apiCall('/teacher-assignments');
                this.teacherAssignments = data.data || [];
                this.showDataModal = true;
                this.currentDataTitle = 'Teacher Assignments';
                this.currentDataHeaders = ['ID', 'Teacher ID', 'Subject ID'];
                this.currentData = this.teacherAssignments.map(assignment => ({
                    id: assignment.id,
                    'Teacher ID': assignment.teacher_id,
                    'Subject ID': assignment.subject_id
                }));
            } catch (error) {
                console.error('Failed to load teacher assignments:', error);
            }
        },
        
        // Lesson Schedules Management
        async loadLessonSchedules() {
            try {
                const data = await this.apiCall('/lesson-schedules');
                this.lessonSchedules = data.data || [];
                this.showDataModal = true;
                this.currentDataTitle = 'Lesson Schedules';
                this.currentDataHeaders = ['ID', 'Subject ID', 'Weekday', 'Number', 'Class ID', 'Teacher ID'];
                this.currentData = this.lessonSchedules.map(schedule => ({
                    id: schedule.id,
                    'Subject ID': schedule.subject_id,
                    'Weekday': this.getWeekdayName(schedule.weekday),
                    'Number': schedule.number,
                    'Class ID': schedule.class_id,
                    'Teacher ID': schedule.teacher_id
                }));
            } catch (error) {
                console.error('Failed to load lesson schedules:', error);
            }
        },
        
        // Lesson Logs Management
        async loadLessonLogs() {
            try {
                const data = await this.apiCall('/lesson-logs');
                this.lessonLogs = data.data || [];
                this.showDataModal = true;
                this.currentDataTitle = 'Lesson Logs';
                this.currentDataHeaders = ['ID', 'Subject ID', 'Date', 'Number', 'Class ID', 'Teacher ID'];
                this.currentData = this.lessonLogs.map(log => ({
                    id: log.id,
                    'Subject ID': log.subject_id,
                    'Date': log.date,
                    'Number': log.number,
                    'Class ID': log.class_id,
                    'Teacher ID': log.teacher_id
                }));
            } catch (error) {
                console.error('Failed to load lesson logs:', error);
            }
        },
        
        // Student Lessons Management
        async loadStudentLessons() {
            try {
                const data = await this.apiCall('/student-lessons');
                this.studentLessons = data.data || [];
                this.showDataModal = true;
                this.currentDataTitle = 'Student Lessons (Attendance & Grades)';
                this.currentDataHeaders = ['ID', 'Student ID', 'Lesson ID', 'Grade', 'Attendance Status'];
                this.currentData = this.studentLessons.map(lesson => ({
                    id: lesson.id,
                    'Student ID': lesson.student_id,
                    'Lesson ID': lesson.lesson_id,
                    'Grade': lesson.grade || '-',
                    'Attendance Status': this.getAttendanceStatusName(lesson.attendance_status)
                }));
            } catch (error) {
                console.error('Failed to load student lessons:', error);
            }
        },
        
        // Attendance Statuses Management
        async loadAttendanceStatuses() {
            try {
                const data = await this.apiCall('/attendance-statuses');
                this.attendanceStatuses = data.data || [];
            } catch (error) {
                console.error('Failed to load attendance statuses:', error);
            }
        },
        
        // Utility Functions
        getWeekdayName(weekday) {
            const weekdays = ['Sunday', 'Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday'];
            return weekdays[weekday] || 'Unknown';
        },
        
        getAttendanceStatusName(code) {
            const status = this.attendanceStatuses.find(s => s.code === code);
            return status ? status.description : code;
        },
        
        // CRUD Operations
        async editItem(item) {
            // Implementation for editing items
            this.showNotification('Edit functionality coming soon!', 'success');
        },
        
        async deleteItem(item) {
            if (confirm('Are you sure you want to delete this item?')) {
                try {
                    // Determine the endpoint based on the data type
                    let endpoint = '';
                    if (this.currentDataTitle === 'Classes') endpoint = '/classes';
                    else if (this.currentDataTitle === 'Students') endpoint = '/students';
                    else if (this.currentDataTitle === 'Teachers') endpoint = '/teachers';
                    else if (this.currentDataTitle === 'Subjects') endpoint = '/subjects';
                    else if (this.currentDataTitle === 'Teacher Assignments') endpoint = '/teacher-assignments';
                    else if (this.currentDataTitle === 'Lesson Schedules') endpoint = '/lesson-schedules';
                    else if (this.currentDataTitle === 'Lesson Logs') endpoint = '/lesson-logs';
                    else if (this.currentDataTitle === 'Student Lessons (Attendance & Grades)') endpoint = '/student-lessons';
                    
                    await this.apiCall(`${endpoint}/${item.id}`, 'DELETE');
                    this.showNotification('Item deleted successfully!', 'success');
                    
                    // Reload the current data
                    if (this.currentDataTitle === 'Classes') this.loadClasses();
                    else if (this.currentDataTitle === 'Students') this.loadStudents();
                    else if (this.currentDataTitle === 'Teachers') this.loadTeachers();
                    else if (this.currentDataTitle === 'Subjects') this.loadSubjects();
                    else if (this.currentDataTitle === 'Teacher Assignments') this.loadTeacherAssignments();
                    else if (this.currentDataTitle === 'Lesson Schedules') this.loadLessonSchedules();
                    else if (this.currentDataTitle === 'Lesson Logs') this.loadLessonLogs();
                    else if (this.currentDataTitle === 'Student Lessons (Attendance & Grades)') this.loadStudentLessons();
                    
                } catch (error) {
                    console.error('Failed to delete item:', error);
                }
            }
        },
        
        // Notification System
        showNotification(message, type = 'success') {
            this.notification = {
                show: true,
                type,
                message
            };
            
            // Auto-hide after 3 seconds
            setTimeout(() => {
                this.notification.show = false;
            }, 3000);
        }
    }
}
