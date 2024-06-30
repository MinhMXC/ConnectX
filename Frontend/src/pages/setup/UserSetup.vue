<script setup>
import 'primeicons/primeicons.css'
import Back from "@/pages/Back.vue";
import UserSetupForm from "@/components/form/user/UserSetupForm.vue";
import router from "@/router.js";
import FullScreenCenteredContainer from "@/components/FullScreenCenteredContainer.vue";
</script>

<template>
  <FullScreenCenteredContainer>
    <div style="margin-bottom: 150px">
      <Back :navigate="() => router.push('/setup')" />
      <h1 id="user-setup-text">User Setup</h1>
      <p id="please-enter-text">Please enter these information so we could better improve your experience @ ConnectX!</p>

      <UserSetupForm />
    </div>
  </FullScreenCenteredContainer>
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