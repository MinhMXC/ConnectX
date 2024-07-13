<script setup>
import Button from "primevue/button";
import {inject} from "vue";
import router from "@/router.js";

const $cookies = inject("$cookies");
const email = $cookies.get("email");
</script>

<template>
  <div id="app-bar" class="white-elevation">
    <p style="font-size: 35px; font-weight: 600">ConnectX</p>
    <div style="flex-grow: 1">&nbsp;</div>
    <p style="font-size: 20px" v-if="email">
      {{$cookies.get("email")}} | {{$cookies.get("user_type")}}
    </p>
    <Button v-if="!email" @click="router.push('/auth/login')">Login</Button>
    <Button v-if="!email" @click="router.push('/auth/signup')">Sign Up</Button>
  </div>
</template>

<style scoped>
#app-bar {
  display: flex;
  align-items: center;
  box-sizing: border-box;
  padding-left: 20px;
  padding-right: 20px;
  gap: 10px;
  width: 100vw;
}
</style>