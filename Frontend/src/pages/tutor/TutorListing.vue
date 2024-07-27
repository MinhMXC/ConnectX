<script setup>
import useBackendGet from "@/composables/useBackendGet.js";
import Image from "primevue/image";
import router from "@/router.js";

const { data: tutorData } = useBackendGet("/tutor");
</script>

<template>
  <p style="font-size: 30px; margin-top: 40px">Tutor Listing</p>
  <div class="white-elevation" :key="tutor.id" v-for="(tutor) in tutorData ?? []" style="padding: 20px; border-radius: 10px; margin-top: 30px">
    <div style="display: flex;" @click="() => router.push(`/tutor/${tutor.id}`)">
      <Image :src="tutor.picture" width="400px"/>
      <div id="tutor-info">
        <p id="name">{{ tutor.name }}</p>
        <p class="other-info">Email: {{ tutor.email }}</p>
        <p class="other-info">{{ tutor.age }} years old</p>
        <p class="other-info">Gender: {{ tutor.gender === null ? "Others" : tutor.gender ? "Female" : "Male" }}</p>
        <p class="other-info">Phone Number: {{ tutor.phone }}</p>
        <p class="other-info"><i>About me: </i><br /> {{ tutor.description }}</p>
      </div>
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