<script setup>
import Button from 'primevue/button';
import Image from 'primevue/image';
import useBackendGet from "@/composables/useBackendGet.js";
import {useRoute, useRouter} from "vue-router";

const route = useRoute();
const router = useRouter();
const { data, status } = useBackendGet(`/user/${route.params.id}`);

const editProfileOnClick = () => {
  router.push(`/user/${route.params.id}/edit`);
};
</script>

<template>
  <p class="not-found-text" v-if="status !== 'Success'">{{ status }}</p>
  <div class="white-elevation rounded profile-ctn" style="display: flex; flex-direction: column" v-if="status === 'Success'">
    <div style="display: flex">
      <Image :src="data.picture"/>
      <div id="user-info">
        <p id="username">{{ data.username }}</p>
        <p class="other-info">Email: {{ data.email }}</p>
        <p class="other-info">Gender: {{ data.gender === null ? "Others" : data.gender ? "Female" : "Male" }}</p>
        <p class="other-info">Role: {{ data.is_parent ? "Parent" : "Student" }}</p>
      </div>
    </div>

    <div style="margin-top: 20px; display: flex; flex-direction: column">
      <Button @click="editProfileOnClick">Edit Profile</Button>
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