<template>
  <div>
    <h2>Add Student</h2>

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

      <input type="date" v-model="student.dob" />

      <input v-model="student.photo" placeholder="Photo filename" />

      <label>
        Handicapped
        <input type="checkbox" v-model="student.handicapped" />
      </label>

      <input v-model="student.email" placeholder="Email" />

      <input v-model="student.mobileNumber" placeholder="Mobile Number" />

      <input v-model="student.bloodGroup" placeholder="Blood Group" />

      <br /><br />
      <button type="submit">Save Student</button>
    </form>
  </div>
</template>

<script>
import axios from "axios"

export default {
  name: "AddStudent",
  data() {
    return {
      student: {
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
        await axios.post("http://localhost:8000/students", this.student)
        alert("Student added successfully")

        // clear form
        this.student = {
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
        }
      } catch (error) {
        alert("Error adding student")
        console.error(error)
      }
    },
  },
}
</script>
