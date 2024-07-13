<script setup>
import Button from 'primevue/button';
import {inject, ref} from "vue";
import useBackendPost from "@/composables/useBackendPost.js";
import FormStatus from "@/components/FormStatusText.vue";
import router from "@/router.js";
import CXInputText from "@/components/form_input/CXInputText.vue";
import CXPassword from "@/components/form_input/CXPassword.vue";

const email = ref("");
const password = ref("");
const $cookies = inject("$cookies");
const { post, data, status, loading } = useBackendPost("/user/login");

const loginOnClick = async () => {
  await post({ email: email.value, password: password.value });

  if (status.value === "Success") {
    $cookies.set("email", data.value.email);

    const user_type = data.value.user_type === 1
        ? "User"
        : data.value.user_type === 2
        ? "Tutor"
            : "Tuition Center";

    $cookies.set("user_type", user_type);
    if (data.value.user_type === -1) {
      await router.push("/setup");
    } else {
      await router.push("/");
    }
  }
};

</script>

<template>
  <div id="display-form-container">
    <CXInputText label="Email" v-model="email" />
    <CXPassword label="Password" v-model="password" :feedback="false" />
    <FormStatus :status="status" />
    <Button
        id="login-button"
        label="Login"
        :loading="loading"
        @click="loginOnClick"
    />
  </div>
</template>

<style scoped>
#display-form-container {
  display: flex;
  flex-direction: column;
  row-gap: 8px;
  width: 100%;
}

#login-button {
  margin-top: 10px;
}

.p-password {
  flex-direction: column;
}

@media screen and (max-width: 1280px) {
  #dis {
    display: none;
  }
}
</style>