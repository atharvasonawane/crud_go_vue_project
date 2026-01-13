<template>
    <div>
        <h2>Student List</h2>

        <table>
            <thead>
                <tr>
                    <th>ID</th>
                    <th>Name</th>
                    <th>Email</th>
                    <th>Mobile</th>
                    <th>Actions</th>
                </tr>
            </thead>

            <tbody>
                <tr v-for="student in students" :key="student.id">
                    <td>{{ student.id }}</td>
                    <td>{{ student.studentName }}</td>
                    <td>{{ student.email }}</td>
                    <td>{{ student.mobileNumber }}</td>
                    <td>

                        <button @click="editStudent(student.id)">Edit</button>

                        <button @click="deleteStudent(student.id)">Delete</button>
                    </td>
                </tr>
            </tbody>
        </table>
    </div>
</template>

<script>
import axios from "axios"

export default {
    name: "StudentList",
    data() {
        return {
            students: [],
        }
    },

    methods: {
        async fetchStudents() {
            try {
                const response = await axios.get("http://localhost:8000/students")
                this.students = response.data
            } catch (error) {
                console.error(error)
            }
        },

        async deleteStudent(id) {
            if (!confirm("Are you sure you want to delete this student?")) return

            try {
                await axios.delete(`http://localhost:8000/students/${id}`)
                this.fetchStudents()
            } catch (error) {
                console.error(error)
            }
        },
        editStudent(id) {
            this.$router.push(`/edit-student/${id}`)
        }
    },

    mounted() {
        this.fetchStudents()
    },
}
</script>
