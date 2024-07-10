<script setup>

import Image from 'primevue/image';
import useBackendGet from "@/composables/useBackendGet.js";
import {useRoute} from "vue-router";
import Button from "primevue/button";
import router from "@/router.js";

const route = useRoute();
const { data, status } = useBackendGet(`/tutor/${route.params.id}`);

const editProfileOnClick = () => {
  router.push(`/tutor/${route.params.id}/edit`);
};
</script>

<template>
  <p class="not-found-text" v-if="status !== 'Success'">{{ status }}</p>
  <div class="white-elevation rounded profile-ctn" v-if="status === 'Success'">
    <div style="display: flex">
      <Image :src="data.picture"/>
      <div id="tutor-info">
        <p id="name">{{ data.name }}</p>
        <p class="other-info">Email: {{ data.email }}</p>
        <p class="other-info">{{ data.age }} years old</p>
        <p class="other-info">Gender: {{ data.gender === null ? "Others" : data.gender ? "Female" : "Male" }}</p>
        <p class="other-info">Phone Number: {{ data.phone }}</p>

        <p class="other-info"><i>About me: </i><br /> {{ data.description }}</p>
      </div>
    </div>

    <div style="margin-top: 20px; display: flex; flex-direction: column">
      <Button @click="editProfileOnClick">Edit Profile</Button>
    </div>
  </div>
</template>

<style scoped>
#name {
  font-size: 40px;
  font-weight: 600;
}

.other-info {
  font-size: 25px;
}

#tutor-info {
  margin-left: 20px;
}
</style>