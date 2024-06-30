<script setup>
import ConnectXWelcome from "@/components/ConnectXWelcome.vue";
import LoginForm from "@/components/form/auth/LoginForm.vue";
</script>

<template>
  <div id="login-page-container">
    <ConnectXWelcome />
    <div id="login-container">
      <div id="form-container">
        <p id="logging-in-to-connectx">Logging in to ConnectX</p>
        <p id="enter-credential-text">Enter your credentials below</p>
        <br />
        <LoginForm />
        <p id="dont-have-account-text">
          Don't have an account?
          <RouterLink to="/auth/signup">Register Here</RouterLink>
        </p>
      </div>
    </div>
  </div>
</template>

<style scoped>
#login-page-container {
  display: flex;
  height: 100%;
  width: 100%;
  padding: 20px;
  box-sizing: border-box;
}

#login-container {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  flex-grow: 1;
}

#logging-in-to-connectx {
  font-size: 35px;
}

#enter-credential-text {
  font-size: 20px;
}

#dont-have-account-text {
  margin-top: 5px;
}

#form-container {
  width: 50%;
}

@media screen and (max-width: 1280px) {
  #form-container {
    width: 60%;
  }
}
</style>