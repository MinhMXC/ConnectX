<script setup>
import useBackendGet from "@/composables/useBackendGet.js";
import Image from "primevue/image";
import router from "@/router.js";
const { data: tuitionCenterData } = useBackendGet("/tuition_center");

</script>

<template>
  <p style="font-size: 30px; margin-top: 40px">Tutor Listing</p>
  <div class="white-elevation" :key="tuition_center.id" v-for="(tuition_center) in tuitionCenterData ?? []" style="padding: 20px; border-radius: 10px; margin-top: 30px">
    <div style="display: flex" @click="() => router.push(`/tuition_center/${tuition_center.id}`)">
      <Image :src="tuition_center.picture" width="400px"/>
      <div id="tutor-info">
        <p id="name">{{ tuition_center.name }}</p>
        <p class="other-info">Email: {{ tuition_center.email }}</p>
        <p class="other-info">Gender: {{ tuition_center.gender === null ? "Others" : tuition_center.gender ? "Female" : "Male" }}</p>
        <p class="other-info">Phone Number: {{ tuition_center.phone }}</p>
        <p class="other-info">Address: {{ tuition_center.address }}</p>
        <p class="other-info">Address Link: {{ tuition_center.address_link }}</p>
        <p class="other-info">Website: {{ tuition_center.website }}</p>
        <p class="other-info"><i>About us: </i><br /> {{ tuition_center.description }}</p>
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