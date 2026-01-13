<template>
  <div>
    <h2>{{ isEdit ? "Edit Student" : "Add Student" }}</h2>

    <form @submit.prevent="submitForm">
      <input v-model="student.studentName" placeholder="Student Name" required />

      <input v-model="student.address" placeholder="Address" />

      <input v-model="student.state" placeholder="State" />

      <input v-model="student.district" placeholder="District" />

      <input v-model="student.taluka" placeholder="Taluka" />

      <select v-model="student.gender">
        <option disabled value="">Select Gender</option>
        <option>Male</option>
        <option>Female</option>
      </select>

      <!-- DATE FIX -->
      <input type="date" v-model="student.dob" required />

      <input v-model="student.photo" placeholder="Photo filename" />

      <label>
        Handicapped
        <input type="checkbox" v-model="student.handicapped" />
      </label>

      <input v-model="student.email" placeholder="Email" />

      <input v-model="student.mobileNumber" placeholder="Mobile Number" />

      <input v-model="student.bloodGroup" placeholder="Blood Group" />

      <div class="link-container">
        <router-link to="/students" class="nav-link">Student List</router-link>
      </div>

      <br /><br />
      <button class="submit-button" type="submit">
        {{ isEdit ? "Update Student" : "Save Student" }}
      </button>
    </form>
  </div>
</template>

<script>
import axios from "axios"

export default {
  name: "AddStudent",

  data() {
    return {
      isEdit: false,
      student: {
        id: null,
        studentName: "",
        address: "",
        state: "",
        district: "",
        taluka: "",
        gender: "",
        dob: "",
        photo: "",
        handicapped: false,
        email: "",
        mobileNumber: "",
        bloodGroup: "",
      },
    }
  },

  methods: {
    async submitForm() {
      try {

        if (!this.student.dob) {
          alert("Date of Birth is required")
          return
        }

        // Ensure DOB is string (YYYY-MM-DD)
        const studentToSend = {
          ...this.student,
          dob: this.student.dob
        }

        if (this.isEdit) {
          await axios.put(
            `http://localhost:8000/students/${this.student.id}`,
            studentToSend
          )
          alert("Student updated successfully")
        } else {
          await axios.post(
            "http://localhost:8000/students",
            studentToSend
          )
          alert("Student added successfully")
        }

        this.$router.push("/")
      } catch (error) {
        console.error("Save error:", error.response || error)
        alert("Error saving student")
      }
    },

    async fetchStudent(id) {
      try {
        const response = await axios.get(
          `http://localhost:8000/students/${id}`
        )

        const data = response.data

        if (data.dob) {
          data.dob = data.dob.split("T")[0]
        }

        this.student = data
      } catch (error) {
        console.error(error)
        alert("Failed to load student data")
      }
    },
  },

  mounted() {
    const id = this.$route.params.id
    if (id) {
      this.isEdit = true
      this.fetchStudent(id)
    }
  },
}
</script>
