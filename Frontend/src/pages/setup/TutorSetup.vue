<script setup>
import 'primeicons/primeicons.css'
import Back from "@/pages/Back.vue";
import router from "@/router.js";
import TutorSetupForm from "@/components/form/tutor/TutorSetupForm.vue";
</script>

<template>
  <div style="margin-bottom: 150px">
    <Back :navigate="() => router.push('/setup')" />
    <h1 id="user-setup-text">Tutor Setup</h1>
    <p id="please-enter-text">Please enter these information so we could better improve your experience @ ConnectX!</p>

    <TutorSetupForm />
  </div>
</template>

<style scoped>
#user-setup-text {
  margin-top: 40px;
}
#please-enter-text {
  font-size: 18px;
  margin-bottom: 40px;
}
</style>