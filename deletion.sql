-- --------------------
-- Clear all data
-- Order matters due to foreign keys

-- Option 1: Using DELETE in dependency order
DELETE FROM public.student_lessons;
DELETE FROM public.lesson_logs;
DELETE FROM public.lesson_schedule;
DELETE FROM public.teacher_assignments;
DELETE FROM public.students;
DELETE FROM public.classes;
DELETE FROM public.subjects;
DELETE FROM public.teacher;
DELETE FROM public.attendance_statuses;

-- Option 2: Using TRUNCATE with CASCADE (faster for large datasets)
TRUNCATE TABLE 
    public.student_lessons,
    public.lesson_logs,
    public.lesson_schedule,
    public.teacher_assignments,
    public.students,
    public.classes,
    public.subjects,
    public.teacher,
    public.attendance_statuses
    RESTART IDENTITY CASCADE;
