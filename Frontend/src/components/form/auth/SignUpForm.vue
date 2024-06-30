<script setup>
import Button from 'primevue/button'
import {ref} from "vue";
import FormStatusText from "@/components/FormStatusText.vue";
import useBackendPost from "@/composables/useBackendPost.js";
import CXPassword from "@/components/form_input/CXPassword.vue";
import CXInputText from "@/components/form_input/CXInputText.vue";

const email = ref("");
const password = ref("");
const confirmPassword = ref("");

const { post, data, status, loading } = useBackendPost("/user")

const signupOnClick = async () => {
  await post({ email: email.value, password: password.value, confirm_password: confirmPassword.value });
  if (status.value === "Success") {
    email.value = "";
    password.value = "";
    confirmPassword.value = "";
  }
}
</script>

<template>
  <div id="display-form-container">
    <CXInputText label="Email" v-model="email" />
    <CXPassword label="Password" v-model="password" :feedback="true" />
    <CXPassword label="Confirm Password" v-model="confirmPassword" :feedback="true" />
    <FormStatusText :status="status" />
    <p v-if="status === 'Success'">
      Thank you for registering an account at Connect X. <br />
      Please
      <RouterLink to="/auth/login">click here</RouterLink>
      to go back to the login page
      and login to your new account
    </p>
    <Button
        id="signup-button"
        label="Sign Up"
        :loading="loading"
        @click="signupOnClick()"
    />
  </div>
</template>

<style scoped>
#display-form-container {
  display: flex;
  flex-direction: column;
  row-gap: 8px;
}

#signup-button {
  margin-top: 10px;
}
</style>