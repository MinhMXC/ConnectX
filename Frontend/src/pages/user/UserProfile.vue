<script setup>

import Image from 'primevue/image';
import useBackendGet from "@/composables/useBackendGet.js";
import {useRoute} from "vue-router";

const route = useRoute();
const { data, status } = useBackendGet(`/user/${route.params.id}`)

</script>

<template>
  <p class="not-found-text" v-if="status !== 'Success'">{{ status }}</p>
  <div class="white-elevation rounded profile-ctn" style="display: flex" v-if="status === 'Success'">
    <Image :src="data.picture"/>
    <div id="user-info">
      <p id="username">{{ data.username }}</p>
      <p class="other-info">Email: {{ data.email }}</p>
      <p class="other-info">Gender: {{ data.gender === null ? "Others" : data.gender ? "Female" : "Male" }}</p>
      <p class="other-info">Role: {{ data.is_parent ? "Parent" : "Student" }}</p>
    </div>
  </div>
</template>

<style scoped>
#username {
  font-size: 40px;
  font-weight: 600;
}

.other-info {
  font-size: 25px;
}

#user-info {
  margin-left: 20px;
}
</style>