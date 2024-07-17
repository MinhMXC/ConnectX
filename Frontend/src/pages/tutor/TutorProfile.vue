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
import RateItem from "@/components/RateItem.vue";

const route = useRoute();
const { logout } = useLogout();
const { data, status } = useBackendGet(`/tutor/${route.params.id}`);
const { data: rateData } = useBackendGet(`/rate/tutor/${route.params.id}`);
const { data: qualificationData } = useBackendGet(`/qualification/tutor/${route.params.id}`);
const { data: levelData } = useBackendGet(`/level`);
const { data: subjectData } = useBackendGet("/subject");
const $cookies = inject("$cookies");

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

        <p class="other-info" style="margin-top: 30px">Qualifications: </p>
        <div id="qualification-ctn" v-if="qualificationData !== null">
          <QualificationItem
              v-for="(qualification) in qualificationData"
              :key="qualification.id"
              :qualification="qualification"
              :levelData="levelData"
          />
        </div>

        <div v-if="qualificationData === null">
          <p style="font-size: 20px">NONE</p>
        </div>

        <Button v-if="$cookies.get('email') === data.email" style="margin-top: 10px" @click="router.push('/qualification/create')">
          Add Qualification
        </Button>

        <p class="other-info" style="margin-top: 30px">Rate: </p>
        <div id="qualification-ctn" v-if="rateData !== null">
          <RateItem
              v-for="(rate) in rateData"
              :key="rate.id"
              :rate="rate"
              :levelData="levelData"
              :subjectData="subjectData"
          />
        </div>

        <div v-if="subjectData === null">
          <p style="font-size: 20px">NONE</p>
        </div>

        <Button v-if="$cookies.get('email') === data.email" style="margin-top: 10px" @click="router.push('/rate/create')">
          Add Rate
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