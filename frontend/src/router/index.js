
import { createRouter, createWebHistory } from 'vue-router';
import AddStudent from '../components/AddStudent.vue';
import StudentList from '../components/StudentList.vue';

const routes = [
  {
    path: "/",
    redirect: "/add-student",
  },
  {
    path: "/add-student",
    component: AddStudent,
  },
  {
    path: "/students",
    component: StudentList,
  },
    {
    path: "/edit-student/:id",
    component: AddStudent
  }

]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

export default router