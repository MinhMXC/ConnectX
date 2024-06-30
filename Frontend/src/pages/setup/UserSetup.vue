<script setup>
import 'primeicons/primeicons.css'
import Back from "@/pages/Back.vue";
import UserSetupForm from "@/components/form/user/UserSetupForm.vue";
import router from "@/router.js";
</script>

<template>
  <div style="display: flex; align-items: center; justify-content: center; width: 100%">
    <div style="margin-bottom: 150px">
      <Back :navigate="() => router.push('/setup')" />
      <h1 id="user-setup-text">User Setup</h1>
      <p id="please-enter-text">Please enter these information so we could better improve your experience @ ConnectX!</p>

      <UserSetupForm />
    </div>
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