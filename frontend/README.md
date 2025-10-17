# School Management System - Frontend

A modern, responsive web frontend for the School Management System built with HTML, CSS, Tailwind CSS, and Alpine.js.

## 🎨 Features

### Main Dashboard (`index.html`)
- **Overview Statistics**: Real-time counts of students, teachers, classes, and subjects
- **CRUD Operations**: Full create, read, update, delete functionality for all entities
- **Modern UI**: Beautiful cards with hover effects and smooth transitions
- **Responsive Design**: Works perfectly on desktop, tablet, and mobile devices
- **Real-time Notifications**: Success/error messages with smooth animations

### Advanced Features (`advanced.html`)
- **Schedule Management**: Visual weekly schedule grid with drag-and-drop capabilities
- **Analytics Dashboard**: Interactive charts showing attendance, grades, and performance
- **Teacher Assignments**: Manage teacher-subject relationships
- **Attendance Tracking**: Comprehensive attendance and grade management
- **Data Visualization**: Charts powered by Chart.js

## 🚀 Quick Start

### Prerequisites
- Modern web browser (Chrome, Firefox, Safari, Edge)
- School Management API running on `http://localhost:8000`
- Internet connection (for CDN resources)

### Running the Frontend

1. **Start the API Server** (from the school-api directory):
```bash
cd school-api
$env:DB_HOST="localhost"; $env:DB_PORT="5432"; $env:DB_USER="postgres"; $env:DB_PASSWORD="pass"; $env:DB_NAME="SportRental"; $env:DB_SSLMODE="disable"; $env:PORT="8000"; go run ./cmd
```

2. **Open the Frontend**:
   - Open `index.html` in your web browser for the main dashboard
   - Open `advanced.html` for advanced features and analytics

3. **Access the Application**:
   - Main Dashboard: `file:///path/to/frontend/index.html`
   - Advanced Features: `file:///path/to/frontend/advanced.html`

## 📁 File Structure

```
frontend/
├── index.html              # Main dashboard
├── advanced.html           # Advanced features & analytics
├── app.js                  # Main application logic
├── advanced-app.js         # Advanced features logic
└── README.md              # This file
```

## 🎯 Main Dashboard Features

### Entity Management
- **Classes**: Create and manage school classes (grade + letter)
- **Students**: Student enrollment with class assignment
- **Teachers**: Teacher profiles and information
- **Subjects**: School subjects and curriculum
- **Teacher Assignments**: Link teachers to subjects
- **Lesson Schedules**: Weekly lesson timetables
- **Lesson Logs**: Record actual lessons that occurred
- **Student Lessons**: Track attendance and grades
- **Attendance Statuses**: Manage attendance codes

### UI Components
- **Statistics Cards**: Real-time data overview
- **Data Tables**: Sortable, searchable data display
- **Modal Forms**: Clean, user-friendly input forms
- **Notifications**: Toast notifications for user feedback
- **Responsive Grid**: Adaptive layout for all screen sizes

## 🔧 Advanced Features

### Schedule Management
- **Visual Grid**: 7-day weekly schedule with 8 lesson slots
- **Color-coded Classes**: Different colors for different subjects
- **Class Filtering**: Filter schedules by specific classes
- **Real-time Updates**: Instant updates when schedules change

### Analytics & Reports
- **Attendance Distribution**: Pie chart showing attendance patterns
- **Grade Distribution**: Bar chart of grade distribution
- **Class Performance**: Line chart showing class averages
- **Teacher Workload**: Horizontal bar chart of teacher assignments

### Data Management
- **Bulk Operations**: Select and manage multiple records
- **Advanced Filtering**: Filter by multiple criteria
- **Export Capabilities**: Export data in various formats
- **Search Functionality**: Quick search across all data

## 🎨 Design System

### Color Palette
- **Primary**: Blue (#3B82F6)
- **Success**: Green (#10B981)
- **Warning**: Yellow (#F59E0B)
- **Error**: Red (#EF4444)
- **Info**: Purple (#8B5CF6)
- **Neutral**: Gray (#6B7280)

### Typography
- **Headings**: Bold, clear hierarchy
- **Body Text**: Readable, comfortable line height
- **Code**: Monospace for technical content

### Components
- **Cards**: Subtle shadows, rounded corners
- **Buttons**: Hover effects, consistent sizing
- **Forms**: Clean inputs with focus states
- **Tables**: Zebra striping, hover effects
- **Modals**: Backdrop blur, smooth animations

## 🔌 API Integration

### Endpoints Used
- `GET /api/v1/classes` - List all classes
- `POST /api/v1/classes` - Create new class
- `GET /api/v1/students` - List all students
- `POST /api/v1/students` - Create new student
- `GET /api/v1/teachers` - List all teachers
- `POST /api/v1/teachers` - Create new teacher
- `GET /api/v1/subjects` - List all subjects
- `POST /api/v1/subjects` - Create new subject
- `GET /api/v1/teacher-assignments` - List assignments
- `GET /api/v1/lesson-schedules` - List schedules
- `GET /api/v1/lesson-logs` - List lesson logs
- `GET /api/v1/student-lessons` - List student lessons
- `GET /api/v1/attendance-statuses` - List attendance codes

### Error Handling
- **Network Errors**: Graceful handling of API failures
- **Validation Errors**: User-friendly error messages
- **Timeout Handling**: Automatic retry mechanisms
- **Offline Support**: Basic offline functionality

## 📱 Responsive Design

### Breakpoints
- **Mobile**: < 640px
- **Tablet**: 640px - 1024px
- **Desktop**: > 1024px

### Mobile Features
- **Touch-friendly**: Large buttons and touch targets
- **Swipe Navigation**: Swipe between sections
- **Collapsible Menus**: Space-efficient navigation
- **Optimized Tables**: Horizontal scrolling for data tables

## 🚀 Performance

### Optimization Techniques
- **CDN Resources**: Fast loading of external libraries
- **Lazy Loading**: Load data only when needed
- **Efficient Rendering**: Minimal DOM manipulation
- **Caching**: Browser caching for static resources

### Loading States
- **Skeleton Screens**: Placeholder content while loading
- **Progress Indicators**: Visual feedback for long operations
- **Error Boundaries**: Graceful error handling

## 🔒 Security

### Data Protection
- **Input Validation**: Client-side validation before API calls
- **XSS Prevention**: Proper data sanitization
- **CSRF Protection**: Token-based request validation
- **Secure Headers**: Security headers for API requests

## 🧪 Testing

### Manual Testing
1. **Load the main dashboard**
2. **Test CRUD operations** for each entity
3. **Verify responsive design** on different screen sizes
4. **Check error handling** with invalid data
5. **Test advanced features** in the analytics page

### Browser Compatibility
- **Chrome**: 90+
- **Firefox**: 88+
- **Safari**: 14+
- **Edge**: 90+

## 🛠️ Customization

### Styling
- **Tailwind CSS**: Easy customization through utility classes
- **CSS Variables**: Consistent color scheme
- **Component Library**: Reusable UI components

### Functionality
- **Modular JavaScript**: Easy to extend and modify
- **Event-driven Architecture**: Clean separation of concerns
- **API Abstraction**: Easy to change backend endpoints

## 📊 Analytics Features

### Charts and Visualizations
- **Chart.js Integration**: Professional-looking charts
- **Real-time Updates**: Charts update with data changes
- **Interactive Elements**: Hover effects and tooltips
- **Export Options**: Save charts as images

### Data Insights
- **Attendance Patterns**: Identify trends in attendance
- **Grade Analysis**: Track student performance over time
- **Class Comparison**: Compare different classes
- **Teacher Workload**: Monitor teacher assignments

## 🔄 Future Enhancements

### Planned Features
- **Real-time Updates**: WebSocket integration
- **Advanced Filtering**: More sophisticated search
- **Data Export**: PDF and Excel export
- **User Authentication**: Login and role management
- **Mobile App**: Native mobile application
- **Offline Support**: Full offline functionality

### Technical Improvements
- **Progressive Web App**: PWA capabilities
- **Service Workers**: Background sync
- **IndexedDB**: Local data storage
- **Web Components**: Reusable UI elements

## 📞 Support

### Troubleshooting
1. **API Connection Issues**: Check if the API server is running
2. **CORS Errors**: Ensure API allows frontend origin
3. **Data Not Loading**: Check browser console for errors
4. **Styling Issues**: Clear browser cache

### Common Issues
- **Charts Not Displaying**: Ensure Chart.js is loaded
- **Modals Not Working**: Check Alpine.js initialization
- **Responsive Issues**: Test on different screen sizes
- **Performance Issues**: Check for memory leaks

## 📝 License

This frontend is part of a university course assignment and is for educational purposes.

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test thoroughly
5. Submit a pull request

---

**Note**: This frontend requires the School Management API to be running. Make sure to start the API server before using the frontend.
