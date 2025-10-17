// Advanced School Management System Frontend
function advancedSchoolApp() {
    return {
        // API Configuration
        apiBaseUrl: 'http://localhost:8000/api/v1',
        
        // Application State
        activeTab: 'schedule',
        
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
        showScheduleForm: false,
        showAttendanceForm: false,
        selectedClass: '',
        selectedClassForAttendance: '',
        
        // Form Data
        newSchedule: {
            subject_id: '',
            weekday: '',
            number: '',
            class_id: '',
            teacher_id: ''
        },
        
        // Notification System
        notification: {
            show: false,
            type: 'success',
            message: ''
        },
        
        // Charts
        charts: {
            attendance: null,
            grade: null,
            class: null,
            teacher: null
        },
        
        // Initialize Application
        init() {
            this.loadAllData();
            this.$nextTick(() => {
                this.initializeCharts();
            });
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
        
        // Load All Data
        async loadAllData() {
            try {
                const [classesData, studentsData, teachersData, subjectsData, assignmentsData, schedulesData, lessonsData, studentLessonsData, attendanceData] = await Promise.all([
                    this.apiCall('/classes'),
                    this.apiCall('/students'),
                    this.apiCall('/teachers'),
                    this.apiCall('/subjects'),
                    this.apiCall('/teacher-assignments'),
                    this.apiCall('/lesson-schedules'),
                    this.apiCall('/lesson-logs'),
                    this.apiCall('/student-lessons'),
                    this.apiCall('/attendance-statuses')
                ]);
                
                this.classes = classesData.data || [];
                this.students = studentsData.data || [];
                this.teachers = teachersData.data || [];
                this.subjects = subjectsData.data || [];
                this.teacherAssignments = assignmentsData.data || [];
                this.lessonSchedules = schedulesData.data || [];
                this.lessonLogs = lessonsData.data || [];
                this.studentLessons = studentLessonsData.data || [];
                this.attendanceStatuses = attendanceData.data || [];
                
                this.showNotification('All data loaded successfully!', 'success');
            } catch (error) {
                console.error('Failed to load data:', error);
            }
        },
        
        // Schedule Management
        async loadSchedules() {
            try {
                const data = await this.apiCall('/lesson-schedules');
                this.lessonSchedules = data.data || [];
                this.showNotification('Schedules loaded successfully!', 'success');
            } catch (error) {
                console.error('Failed to load schedules:', error);
            }
        },
        
        async createSchedule() {
            try {
                await this.apiCall('/lesson-schedules', 'POST', this.newSchedule);
                this.showNotification('Schedule created successfully!', 'success');
                this.showScheduleForm = false;
                this.newSchedule = { subject_id: '', weekday: '', number: '', class_id: '', teacher_id: '' };
                this.loadSchedules();
            } catch (error) {
                console.error('Failed to create schedule:', error);
            }
        },
        
        getScheduleForSlot(lessonNumber, weekday) {
            return this.lessonSchedules.filter(schedule => 
                schedule.number === lessonNumber && schedule.weekday === weekday
            );
        },
        
        filterSchedulesByClass() {
            if (this.selectedClass) {
                return this.lessonSchedules.filter(schedule => schedule.class_id == this.selectedClass);
            }
            return this.lessonSchedules;
        },
        
        // Teacher Assignments
        async loadTeacherAssignments() {
            try {
                const data = await this.apiCall('/teacher-assignments');
                this.teacherAssignments = data.data || [];
                this.showNotification('Teacher assignments loaded!', 'success');
            } catch (error) {
                console.error('Failed to load teacher assignments:', error);
            }
        },
        
        async deleteAssignment(assignment) {
            if (confirm('Are you sure you want to delete this assignment?')) {
                try {
                    await this.apiCall(`/teacher-assignments/${assignment.id}`, 'DELETE');
                    this.showNotification('Assignment deleted successfully!', 'success');
                    this.loadTeacherAssignments();
                } catch (error) {
                    console.error('Failed to delete assignment:', error);
                }
            }
        },
        
        // Student Lessons & Attendance
        async loadStudentLessons() {
            try {
                const data = await this.apiCall('/student-lessons');
                this.studentLessons = data.data || [];
                this.showNotification('Student lessons loaded!', 'success');
            } catch (error) {
                console.error('Failed to load student lessons:', error);
            }
        },
        
        filterAttendanceByClass() {
            if (this.selectedClassForAttendance) {
                // Filter students by class, then filter lessons by those students
                const classStudents = this.students.filter(student => student.class_id == this.selectedClassForAttendance);
                const studentIds = classStudents.map(student => student.id);
                return this.studentLessons.filter(lesson => studentIds.includes(lesson.student_id));
            }
            return this.studentLessons;
        },
        
        async deleteStudentLesson(lesson) {
            if (confirm('Are you sure you want to delete this record?')) {
                try {
                    await this.apiCall(`/student-lessons/${lesson.id}`, 'DELETE');
                    this.showNotification('Record deleted successfully!', 'success');
                    this.loadStudentLessons();
                } catch (error) {
                    console.error('Failed to delete record:', error);
                }
            }
        },
        
        // Utility Functions
        getSubjectName(subjectId) {
            const subject = this.subjects.find(s => s.id === subjectId);
            return subject ? subject.subject_name : 'Unknown';
        },
        
        getClassName(classId) {
            const cls = this.classes.find(c => c.id === classId);
            return cls ? `Grade ${cls.grade}${cls.letter}` : 'Unknown';
        },
        
        getTeacherName(teacherId) {
            const teacher = this.teachers.find(t => t.id === teacherId);
            return teacher ? `${teacher.first_name} ${teacher.last_name}` : 'Unknown';
        },
        
        getStudentName(studentId) {
            const student = this.students.find(s => s.id === studentId);
            return student ? `${student.first_name} ${student.last_name}` : 'Unknown';
        },
        
        getAttendanceStatusName(code) {
            const status = this.attendanceStatuses.find(s => s.code === code);
            return status ? status.description : code;
        },
        
        getAttendanceColor(code) {
            const colors = {
                'P': 'bg-green-100 text-green-800',
                'A': 'bg-red-100 text-red-800',
                'L': 'bg-yellow-100 text-yellow-800',
                'E': 'bg-blue-100 text-blue-800',
                'S': 'bg-purple-100 text-purple-800'
            };
            return colors[code] || 'bg-gray-100 text-gray-800';
        },
        
        getClassesForTeacher(teacherId, subjectId) {
            return this.lessonSchedules.filter(schedule => 
                schedule.teacher_id === teacherId && schedule.subject_id === subjectId
            );
        },
        
        // Chart Initialization
        initializeCharts() {
            this.createAttendanceChart();
            this.createGradeChart();
            this.createClassChart();
            this.createTeacherChart();
        },
        
        createAttendanceChart() {
            const ctx = document.getElementById('attendanceChart').getContext('2d');
            const attendanceData = this.calculateAttendanceDistribution();
            
            this.charts.attendance = new Chart(ctx, {
                type: 'doughnut',
                data: {
                    labels: attendanceData.labels,
                    datasets: [{
                        data: attendanceData.data,
                        backgroundColor: [
                            '#10B981', // Green for Present
                            '#EF4444', // Red for Absent
                            '#F59E0B', // Yellow for Late
                            '#3B82F6', // Blue for Excused
                            '#8B5CF6'  // Purple for Sick
                        ]
                    }]
                },
                options: {
                    responsive: true,
                    maintainAspectRatio: false,
                    plugins: {
                        legend: {
                            position: 'bottom'
                        }
                    }
                }
            });
        },
        
        createGradeChart() {
            const ctx = document.getElementById('gradeChart').getContext('2d');
            const gradeData = this.calculateGradeDistribution();
            
            this.charts.grade = new Chart(ctx, {
                type: 'bar',
                data: {
                    labels: gradeData.labels,
                    datasets: [{
                        label: 'Number of Grades',
                        data: gradeData.data,
                        backgroundColor: '#3B82F6'
                    }]
                },
                options: {
                    responsive: true,
                    maintainAspectRatio: false,
                    scales: {
                        y: {
                            beginAtZero: true
                        }
                    }
                }
            });
        },
        
        createClassChart() {
            const ctx = document.getElementById('classChart').getContext('2d');
            const classData = this.calculateClassPerformance();
            
            this.charts.class = new Chart(ctx, {
                type: 'line',
                data: {
                    labels: classData.labels,
                    datasets: [{
                        label: 'Average Grade',
                        data: classData.data,
                        borderColor: '#10B981',
                        backgroundColor: 'rgba(16, 185, 129, 0.1)',
                        tension: 0.4
                    }]
                },
                options: {
                    responsive: true,
                    maintainAspectRatio: false,
                    scales: {
                        y: {
                            beginAtZero: true,
                            max: 5
                        }
                    }
                }
            });
        },
        
        createTeacherChart() {
            const ctx = document.getElementById('teacherChart').getContext('2d');
            const teacherData = this.calculateTeacherWorkload();
            
            this.charts.teacher = new Chart(ctx, {
                type: 'horizontalBar',
                data: {
                    labels: teacherData.labels,
                    datasets: [{
                        label: 'Number of Classes',
                        data: teacherData.data,
                        backgroundColor: '#8B5CF6'
                    }]
                },
                options: {
                    responsive: true,
                    maintainAspectRatio: false,
                    scales: {
                        x: {
                            beginAtZero: true
                        }
                    }
                }
            });
        },
        
        // Data Calculations
        calculateAttendanceDistribution() {
            const distribution = { P: 0, A: 0, L: 0, E: 0, S: 0 };
            
            this.studentLessons.forEach(lesson => {
                if (distribution.hasOwnProperty(lesson.attendance_status)) {
                    distribution[lesson.attendance_status]++;
                }
            });
            
            return {
                labels: ['Present', 'Absent', 'Late', 'Excused', 'Sick'],
                data: [distribution.P, distribution.A, distribution.L, distribution.E, distribution.S]
            };
        },
        
        calculateGradeDistribution() {
            const distribution = { 1: 0, 2: 0, 3: 0, 4: 0, 5: 0 };
            
            this.studentLessons.forEach(lesson => {
                if (lesson.grade && distribution.hasOwnProperty(lesson.grade)) {
                    distribution[lesson.grade]++;
                }
            });
            
            return {
                labels: ['1', '2', '3', '4', '5'],
                data: [distribution[1], distribution[2], distribution[3], distribution[4], distribution[5]]
            };
        },
        
        calculateClassPerformance() {
            const classAverages = {};
            
            this.classes.forEach(cls => {
                const classStudents = this.students.filter(student => student.class_id === cls.id);
                const studentIds = classStudents.map(student => student.id);
                const classLessons = this.studentLessons.filter(lesson => 
                    studentIds.includes(lesson.student_id) && lesson.grade
                );
                
                if (classLessons.length > 0) {
                    const totalGrade = classLessons.reduce((sum, lesson) => sum + lesson.grade, 0);
                    classAverages[`Grade ${cls.grade}${cls.letter}`] = (totalGrade / classLessons.length).toFixed(2);
                } else {
                    classAverages[`Grade ${cls.grade}${cls.letter}`] = 0;
                }
            });
            
            return {
                labels: Object.keys(classAverages),
                data: Object.values(classAverages).map(val => parseFloat(val))
            };
        },
        
        calculateTeacherWorkload() {
            const teacherWorkload = {};
            
            this.teachers.forEach(teacher => {
                const teacherSchedules = this.lessonSchedules.filter(schedule => schedule.teacher_id === teacher.id);
                teacherWorkload[`${teacher.first_name} ${teacher.last_name}`] = teacherSchedules.length;
            });
            
            return {
                labels: Object.keys(teacherWorkload),
                data: Object.values(teacherWorkload)
            };
        },
        
        // Event Handlers
        editAssignment(assignment) {
            this.showNotification('Edit functionality coming soon!', 'success');
        },
        
        editStudentLesson(lesson) {
            this.showNotification('Edit functionality coming soon!', 'success');
        },
        
        // Notification System
        showNotification(message, type = 'success') {
            this.notification = {
                show: true,
                type,
                message
            };
            
            setTimeout(() => {
                this.notification.show = false;
            }, 3000);
        }
    }
}
