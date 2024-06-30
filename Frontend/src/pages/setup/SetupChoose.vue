<script setup>
import Button from 'primevue/button'
import studentImg from '../../assets/student.png';
import teacherImg from '../../assets/teacher.png';
import tuitionCenterImg from '../../assets/tuition_center.png'
import router from "@/router.js";
import FullScreenCenteredContainer from "@/components/FullScreenCenteredContainer.vue";
</script>

<template>
  <FullScreenCenteredContainer>
    <div id="setup-container">
      <p id="welcome-to-connectx-text">Welcome to ConnectX!</p>
      <p id="please-tell-us-text">Please tell us what type of user you are!</p>
      <div id="user-types-container">
        <div class="user-type-item">
          <img class="user-type-image" :src="studentImg" alt="student" />
          <Button class="user-type-button" label="Student / Parent" @click="router.push('/setup/user')" />
        </div>
        <div class="user-type-item">
          <img class="user-type-image" :src="teacherImg" alt="student" />
          <Button class="user-type-button" label="Tutor" @click="router.push('/setup/tutor')" />
        </div>
        <div class="user-type-item">
          <img class="user-type-image" :src="tuitionCenterImg" alt="student" />
          <Button class="user-type-button" label="Tuition Center" @click="router.push('/setup/tuition_center')" />
        </div>
      </div>
    </div>
  </FullScreenCenteredContainer>
</template>

<style scoped>

#setup-container {
  width: 90%;
  margin-bottom: 100px;
}

#welcome-to-connectx-text {
  font-size: 70px;
  font-weight: 500;
}

#please-tell-us-text {
  font-size: 30px;
}

#user-types-container {
  display: flex;
  margin-top: 8vh;
  gap: 30px;
}

.user-type-item {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  width: 33.33%;
}

.user-type-image {
  width: 15vw;
}

.user-type-button {
  margin-top: 10px;
  width: 100%;
}
</style>