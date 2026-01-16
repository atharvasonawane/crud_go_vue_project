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

import axios from "../axios"

export default {
  name: "AddStudent",

  data() {
    return {
      isEdit: false,
      student: {
        // id: null,
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

          if (this.isEdit) {
            await axios.put(
              "/students",
              this.student,
              { withCredentials: true }
            )
              ;
          }
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
            "/students",
            formData,
            {
              headers: { "Content-Type": "multipart/form-data" },
              withCredentials: true
            }
          );
          alert("Student added successfully")
        }

        // this.$router.push("/")
        this.$router.push("/students");
      } catch (error) {
        console.error("Save error:", error.response || error)
        alert("Error saving student")
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


  async mounted() {

    axios.get("http://localhost:5173/locations.json")
      .then(res => {
        this.locations = res.data;
      })
      .catch(err => console.error(err));

    if (this.$route.path === "/add-student") {
      this.isEdit = false;
      return;
    }


    try {
      const response = await axios.get("/student-detail", {
        withCredentials: true
      });
      console.log(response.data)

      this.student = response.data;
      if (this.student.dob) {
        this.student.dob = this.student.dob.split("T")[0];
      }

      const stateObj = this.locations.find(
        s => s.id == this.student.state
      );
      this.districts = stateObj ? stateObj.districts : [];

      const districtObj = this.districts.find(
        d => d.id == this.student.district
      );
      this.talukas = districtObj ? districtObj.talukas : [];
      this.isEdit = true;
    } catch (err) {

      this.isEdit = false;
    }
  }


}
</script>
