<script setup>
import Button from "primevue/button";
import {inject, ref} from "vue";
import router from "@/router.js";
import Menu from "primevue/menu";
import useBackendGet from "@/composables/useBackendGet.js";
import useLogout from "@/composables/useLogout.js";

const $cookies = inject("$cookies");
const email = $cookies.get("email");
const user_type = $cookies.get("user_type");
const user_id = $cookies.get("user_id");
const { logout } = useLogout();

const user_type_string = user_type === "0" ? "User" : user_type === "1" ? "Tutor" : "Tuition Center";
let user_type_url = user_type === "0" ? "user" : user_type === "1" ? "tutor" : "tuition_center";

const menu = ref();
const items = ref([
  {
    label: 'Options',
    items: [
      {
        label: 'Profile',
        icon: 'pi pi-user',
        command: () => router.push(`/${user_type_url}/${user_id}`)
      },
      {
        label: 'Sign Out',
        icon: 'pi pi-sign-out',
        command: logout
      }
    ]
  }
]);

const toggle = (event) => {
  menu.value.toggle(event);
};
</script>

<template>
  <div id="app-bar" class="white-elevation">
    <p style="font-size: 35px; font-weight: 600" @click="router.push('/')">ConnectX</p>
    <div style="flex-grow: 1">&nbsp;</div>
    <p id="app-bar-user" class="white-elevation" v-if="email" @click="toggle">
      {{$cookies.get("email")}} | {{user_type_string}}
    </p>
    <Menu ref="menu" id="overlay_menu" :model="items" :popup="true" />
    <Button v-if="!email" @click="router.push('/auth/login')">Login</Button>
    <Button v-if="!email" @click="router.push('/auth/signup')">Sign Up</Button>
  </div>
</template>

<style scoped>
#app-bar {
  display: flex;
  align-items: center;
  box-sizing: border-box;
  padding-left: 30px;
  padding-right: 30px;
  gap: 10px;
  width: 100vw;
}

#app-bar-user {
  font-size: 20px;
  padding: 8px;
  margin: 0;
  border-radius: 10px;
}
</style>