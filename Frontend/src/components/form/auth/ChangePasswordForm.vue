<script setup>
import CXPassword from "@/components/form_input/CXPassword.vue";
import {ref} from "vue";
import Button from "primevue/button";
import router from "@/router.js";
import FormStatusText from "@/components/FormStatusText.vue";
import useBackendPatch from "@/composables/useBackendPatch.js";

const password = ref("");
const confirmPassword = ref("");

const { patch, status, loading } = useBackendPatch("/user/change_password");

async function changePasswordOnClick() {
  await patch({
    password: password.value,
    confirm_password: confirmPassword.value
  });

  if (status.value === "Success") {
    router.back();
  }
}
</script>

<template>
  <div style="display: flex; flex-direction: column; min-width: 600px">
    <CXPassword label="New Password" v-model="password" />
    <CXPassword label="Confirm New Password" v-model="confirmPassword" />
    <FormStatusText id="form-status-text" :status="status" />
    <Button
        id="change-password-button"
        label="Change Password"
        :loading="loading"
        @click="changePasswordOnClick"
    />
  </div>
</template>

<style scoped>
#form-status-text {
  margin-top: 10px;
}

#change-password-button {
  margin-top: 10px;
}
</style>