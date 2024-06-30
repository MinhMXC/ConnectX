<script setup>
import ConnectXWelcome from "@/components/ConnectXWelcome.vue";
import SignUpForm from "@/components/form/auth/SignUpForm.vue";
</script>

<template>
  <div id="signup-page-container">
    <ConnectXWelcome />
    <div id="signup-container">
      <div id="form-container">
        <p id="signup-on-connectx">Signing up on ConnectX</p>
        <p id="fill-in-credential">Fill in your account credentials below</p>
        <br />
        <SignUpForm />
      </div>
    </div>
  </div>
</template>

<style scoped>
#signup-page-container {
  display: flex;
  height: 100vh;
  width: 100vw;
  padding: 20px;
  box-sizing: border-box;
}

#signup-container {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  flex-grow: 1;
}

#signup-on-connectx {
  font-size: 35px;
}

#fill-in-credential {
  font-size: 20px;
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