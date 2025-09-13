-- --------------------
-- Classes
CREATE TABLE public.classes (
    id BIGSERIAL PRIMARY KEY,
    grade INT NOT NULL,
    letter CHAR(1) NOT NULL
);

-- --------------------
-- Students
CREATE TABLE public.students (
    id BIGSERIAL PRIMARY KEY,
    class_id BIGINT NOT NULL REFERENCES public.classes(id) ON DELETE CASCADE,
    first_name VARCHAR NOT NULL,
    last_name VARCHAR NOT NULL,
    patronymic VARCHAR
);

-- --------------------
-- Teacher
CREATE TABLE public.teacher (
    id BIGSERIAL PRIMARY KEY,
    first_name VARCHAR NOT NULL,
    last_name VARCHAR NOT NULL,
    patronymic VARCHAR
);

-- --------------------
-- Subjects
CREATE TABLE public.subjects (
    id BIGSERIAL PRIMARY KEY,
    subject_name VARCHAR NOT NULL
);

-- --------------------
-- TeacherAssignments
CREATE TABLE public.teacher_assignments (
    id BIGSERIAL PRIMARY KEY,
    teacher_id BIGINT NOT NULL REFERENCES public.teacher(id) ON DELETE CASCADE,
    subject_id BIGINT NOT NULL REFERENCES public.subjects(id) ON DELETE CASCADE
);

-- --------------------
-- LessonSchedule
CREATE TABLE public.lesson_schedule (
    id BIGSERIAL PRIMARY KEY,
    subject_id BIGINT NOT NULL REFERENCES public.subjects(id) ON DELETE CASCADE,
    class_id BIGINT NOT NULL REFERENCES public.classes(id) ON DELETE CASCADE,
    teacher_id BIGINT NOT NULL REFERENCES public.teacher(id) ON DELETE cascade,
    weekday INT NOT NULL, -- 1=Monday, 7=Sunday
    number INT NOT NULL -- lesson number in the day
    
);

-- --------------------
-- LessonOccurrences
CREATE TABLE public.lesson_logs (
    id BIGSERIAL PRIMARY KEY,
    subject_id BIGINT NOT NULL REFERENCES public.subjects(id) ON DELETE CASCADE,
    class_id BIGINT NOT NULL REFERENCES public.classes(id) ON DELETE CASCADE,
    teacher_id BIGINT NOT NULL REFERENCES public.teacher(id) ON DELETE cascade,
    date DATE NOT NULL,
    number INT NOT NULL
);

-- --------------------
-- AttendanceStatuses
CREATE TABLE public.attendance_statuses (
    code CHAR(1) PRIMARY KEY,
    description VARCHAR NOT NULL
);

-- Optional: insert common statuses
INSERT INTO public.attendance_statuses (code, description)
VALUES
('P', 'Present'),
('A', 'Absent'),
('L', 'Late'),
('E', 'Excused'),
('S', 'Sick');

-- --------------------
-- StudentLessons
CREATE TABLE public.student_lessons (
    id BIGSERIAL PRIMARY KEY,
    student_id BIGINT NOT NULL REFERENCES public.students(id) ON DELETE CASCADE,
    lesson_id BIGINT NOT NULL REFERENCES public.lesson_logs(id) ON DELETE CASCADE,
    attendance_status CHAR(1) NOT NULL REFERENCES public.attendance_statuses(code),
    grade INT
);
