-- Populate Classes
INSERT INTO public.classes (grade, letter)
SELECT g, l
FROM generate_series(1, 11) AS g,
     unnest(ARRAY['A','B','C','D']) AS l;

-- Populate Teachers
INSERT INTO public.teacher (first_name, last_name, patronymic)
SELECT 'Teacher' || g || l, 'Lastname' || g || l, 'Patronymic' || g || l
FROM generate_series(1, 20) AS g,
     generate_series(1, 1) AS l;

-- Populate Subjects
INSERT INTO public.subjects (subject_name)
VALUES ('Math'), ('Physics'), ('Chemistry'), ('Biology'), ('History'), ('Literature');

-- Populate Students
INSERT INTO public.students (class_id, first_name, last_name, patronymic)
SELECT c.id, 'Student' || gs, 'Lastname' || gs, 'Patronymic' || gs
FROM generate_series(1, 500) AS gs
CROSS JOIN public.classes c
WHERE gs <= 20;  -- max 20 students per class

INSERT INTO public.student_lessons (student_id, lesson_id, grade, attendance_status)
SELECT s.id, l.id, trunc(random()*100)::int, 
       (ARRAY['P','A','L','E','S'])[floor(random()*5+1)::int]
FROM public.students s
CROSS JOIN public.lesson_logs l
WHERE l.date >= '2025-01-01' AND l.date <= '2025-06-30';