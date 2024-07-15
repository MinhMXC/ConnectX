<script setup>

import Image from 'primevue/image';
import useBackendGet from "@/composables/useBackendGet.js";
import {useRoute} from "vue-router";
import Button from "primevue/button";
import router from "@/router.js";
import useLogout from "@/composables/useLogout.js";
import {inject} from "vue";
import useBackendDelete from "@/composables/useBackendDelete.js";
import QualificationItem from "@/components/QualificationItem.vue";

const route = useRoute();
const { logout } = useLogout();
const { data, status } = useBackendGet(`/tutor/${route.params.id}`);
const { data: qualificationData } = useBackendGet(`/qualification/tutor/${route.params.id}`);
const { data: levelData } = useBackendGet(`/level`);
const $cookies = inject("$cookies");

console.log(qualificationData.value);

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

        <P class="other-info" style="margin-top: 30px">Qualifications: </P>
        <div id="qualification-ctn" v-if="qualificationData !== null">
          <QualificationItem
              v-for="(qualification, index) in qualificationData"
              :key="qualification.id"
              :qualification="qualification"
              :levelData="levelData"
          />
        </div>

        <div v-if="qualificationData === null">
          <p style="font-size: 20px">NONE</p>
        </div>

        <Button style="margin-top: 10px" @click="router.push('/qualification/create')">
          Add Qualification
        </Button>
      </div>
    </div>

    <div v-if="$cookies.get('email') === data.email" class="profile-button-ctn">
      <Button @click="editProfileOnClick">Edit Profile</Button>
      <Button @click="router.push('/auth/change_password')">Change Password</Button>
      <Button @click="logout">Logout</Button>
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