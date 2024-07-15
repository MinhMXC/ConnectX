<script setup>
import AppBar from "@/components/AppBar.vue";
</script>

<template>
  <AppBar :key="$route.fullPath" />
  <RouterView :key="$route.fullPath" />
</template>

<style>
#app {
  display: flex;
  align-items: center;
  flex-direction: column;
  min-width: 100vw;
  min-height: 100vh;
  box-sizing: border-box;
}

p {
  margin: 0;
  padding: 0;
}

.white-elevation {
  box-sizing: border-box;
  box-shadow: 0 5px 15px;
}

.rounded {
  border-radius: 10px;
}

.ctn-ctr-row {
  display: flex;
  justify-content: center;
  align-items: center;
}

.ctn-ctr-col {
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
}

.ctn-ctr-col-no-align {
  display: flex;
  justify-content: center;
  flex-direction: column;
}

.not-found-text {
  margin-top: 50px;
  font-size: 40px;
}

.profile-ctn {
  padding: 20px;
  width: 80%;
  min-width: 600px;
  margin-top: 50px;
}

.profile-button-ctn {
  margin-top: 20px;
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.vertical-form-container {
  display: flex;
  flex-direction: column;
  gap: 10px;
}
</style>
