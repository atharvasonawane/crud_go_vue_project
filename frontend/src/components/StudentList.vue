<template>
    <div>
        <h2>Student List</h2>

        <table>
            <thead>
                <tr>
                    <th>ID</th>
                    <th>Photo</th>
                    <th>Name</th>
                    <th>Email</th>
                    <th>Mobile</th>
                    <th>Actions</th>

                </tr>
            </thead>

            <tbody>
                <tr v-for="student in students" :key="student.id">
                    <td>{{ student.id }}</td>
                    <td>
                        <img v-if="student.photo" :src="`http://localhost:8000/uploads/${student.photo}`"
                            alt="Student Photo" width="50">
                    </td>
                    <td>{{ student.studentName }}</td>
                    <td>{{ student.email }}</td>
                    <td>{{ student.mobileNumber }}</td>
                    <td class="action-buttons">
                        <button @click="editStudent(student)">Edit</button>
                        <button @click="deleteStudent(student.id)">Delete</button>
                    </td>
                </tr>
            </tbody>
        </table>
        <br />
        <div class="link-container">
            <router-link to="/add-student" class="nav-link">Add Student</router-link>
        </div>

        <br />
        <button @click="downloadPDF">
            Download PDF
        </button>

    </div>
</template>

<script>
// import axios from "axios"
import axios from "../axios"

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
                const response = await axios.get("/students")
                this.students = response.data
            } catch (error) {
                console.error(error)
            }
        },

        async deleteStudent(id) {
            if (!confirm("Are you sure you want to delete this student?")) return

            try {

                await axios.post(
                    "/select-student",
                    { student_id: id },
                    { withCredentials: true }
                )

                await axios.delete("/students", {
                    withCredentials: true
                })

                this.fetchStudents()
            } catch (error) {
                console.error(error)
            }
        },
        // editStudent(id) {
        //     this.$router.push(`/edit-student/${id}`)
        // }
        async editStudent(student) {
            await axios.post(
                "/select-student",
                { student_id: student.id },
                { withCredentials: true }
            );

            this.$router.push("/edit-student");
            console.log(student);
        },

        downloadPDF(){
            fetch("http://localhost:8000/students/pdf")
            .then(response =>{
                if(!response.ok){
                    throw new Error("Failed to download pdf")
                }
                return response.blob()
            })
            .then(blob =>{
                const url = window.URL.createObjectURL(blob)
                const a = document.createElement("a")
                a.href = url
                a.download = "students.pdf"
                document.body.appendChild(a)
                a.click()
                a.remove()
                window.URL.revokeObjectURL(url)
            })
            .catch(error => {
                console.log(error)
                alert("PDF download failed")
            })
        }
    },

    mounted() {
        this.fetchStudents()
    },
}
</script>
