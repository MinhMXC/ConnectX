<script setup>
import 'primeicons/primeicons.css';
import Back from "@/pages/Back.vue";
import router from "@/router.js";
import FullScreenCenteredContainer from "@/components/FullScreenCenteredContainer.vue";
import ChangePasswordForm from "@/components/form/auth/ChangePasswordForm.vue";
</script>

<template>
  <FullScreenCenteredContainer>
    <div style="margin-bottom: 150px">
      <Back :navigate="() => router.back()" />
      <h1 id="changing-password-text">Changing Password</h1>

      <ChangePasswordForm />
    </div>
  </FullScreenCenteredContainer>
</template>

<style scoped>
#changing-password-text {
  margin-top: 20px;
  margin-bottom: 40px;
}
</style>