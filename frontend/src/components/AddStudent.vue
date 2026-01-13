<template>
  <div>
    <h2>{{ isEdit ? "Edit Student" : "Add Student" }}</h2>

    <form @submit.prevent="submitForm">
      <input v-model="student.studentName" placeholder="Student Name" required />

      <input v-model="student.address" placeholder="Address" />

      <select v-model="student.state" @change="onStateChange">
        <option disabled value="">Select State</option>
        <option v-for="s in locations" :key="s.id" :value="s.id">{{ s.name }}</option>
      </select>

      <select v-model="student.district" @change="onDistrictChange">
        <option disabled value="">Select District</option>
        <option v-for="d in districts" :key="d.id" :value="d.id">{{ d.name }}</option>
      </select>

      <select v-model="student.taluka">
        <option disabled value="">Select Taluka</option>
        <option v-for="t in talukas" :key="t.id" :value="t.id">{{ t.name }}</option>
      </select>

      <select v-model="student.gender">
        <option disabled value="">Select Gender</option>
        <option>Male</option>
        <option>Female</option>
      </select>

      <!-- DATE FIX -->
      <input type="date" v-model="student.dob" required />

      <input type="file" @change="handleFileUpload" />

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
        photoFile: null,
        handicapped: false,
        email: "",
        mobileNumber: "",
        bloodGroup: "",
      },
      locations: [],
      districts: [],
      talukas: []
    }
  },

  methods: {
    async submitForm() {
      try {

        if (!this.student.dob) {
          alert("Date of Birth is required")
          return
        }

        if (this.isEdit) {
          const studentToSend = {
            ...this.student,
            dob: this.student.dob
          }
          await axios.put(
            `http://localhost:8000/students/${this.student.id}`,
            studentToSend
          )
          alert("Student updated successfully")
        } else {
          const formData = new FormData()
          for (const key in this.student) {
            if (key === "photoFile" && this.student.photoFile) {
              formData.append("photo", this.student.photoFile)
            } else if (key !== "photoFile") {
              formData.append(key, this.student[key])
            }
          }

          await axios.post(
            "http://localhost:8000/students",
            formData,
            { headers: { "Content-Type": "multipart/form-data" } }
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
        const state = this.locations.find(s => s.id == this.student.state)
        this.districts = state ? state.districts : []

        const district = this.districts.find(d => d.id == this.student.district)
        this.talukas = district ? district.talukas : []
      } catch (error) {
        console.error(error)
        alert("Failed to load student data")
      }
    },
    handleFileUpload(event) {
      this.student.photoFile = event.target.files[0]; // store the selected file
    },

    onStateChange() {
      const state = this.locations.find(s => s.id == this.student.state)
      this.districts = state ? state.districts : []
      this.student.district = ""
      this.talukas = []
      this.student.taluka = ""
    },

    onDistrictChange() {
      const district = this.districts.find(d => d.id == this.student.district)
      this.talukas = district ? district.talukas : []
      this.student.taluka = ""
    }

  },


  mounted() {
    axios.get("/locations.json")
      .then(res => {
        this.locations = res.data

        const id = this.$route.params.id
        if (id) {
          this.isEdit = true
          this.fetchStudent(id)
        }
      })
      .catch(err => console.error(err))
  }

}
</script>
