<script setup>
import Image from 'primevue/image'
import { VueWriter } from "vue-writer";
const textArray = [' students', ' tutors', ' tuition centers', ' learning'];
</script>

<template>
  <div id="welcome-connectx-container">
    <h1 id="welcome-connectx">ConnextX</h1>
    <VueWriter :array="textArray" :typeSpeed="70" delay="1000">
      <p id="we-empower-text">We empower</p>
    </VueWriter>
  </div>
</template>

<style scoped>
#welcome-connectx-container {
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
  height: 100%;
  width: 40%;
  padding: 20px;
  border-radius: 5px;
  box-sizing: border-box;
  box-shadow: 0 5px 15px;
}

#welcome-connectx {
  font-size: 70px;
}

#we-empower-text {
  display: inline;
}

.is-typed {
  font-size: 30px;
}
</style>