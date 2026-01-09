<template>
  <div>
    <nav>
      <router-link to="/add-student">Add Student</router-link>
      <br/>
      <router-link to="/students">Student List</router-link>
    </nav>
    
    <router-view/>
  </div>
</template>

<script>
export default {
  name: "App",
}
</script>
