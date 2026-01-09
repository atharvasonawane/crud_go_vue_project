
import { createRouter, createWebHistory } from 'vue-router';
import AddStudent from '../components/AddStudent.vue';
import StudentList from '../components/StudentList.vue';

const routes = [
    {
        path: '/',
        redirect: '/students'
    },
    {
        path: '/add-student',
        component: AddStudent
    },
    {
        path: '/students',
        component: StudentList
    },
]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

export default router