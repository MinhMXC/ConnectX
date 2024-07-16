<script setup>

import Button from "primevue/button";
import router from "@/router.js";
import useBackendDelete from "@/composables/useBackendDelete.js";
import {computed, inject} from "vue";

const props = defineProps(['rate', 'levelData', 'subjectData']);
const rate = computed(() => props.rate);
const $cookies = inject("$cookies");
const { deleteFn } = useBackendDelete(`/rate/${rate.value.id}`);

const subject = computed(() => props.subjectData?.find((subject) => subject.id === props.rate.subject_id));

</script>

<template>
  <div class="rate-item-ctn white-elevation">
    <p class="rate-details">{{ `${subject?.name} (${levelData?.find((level) => level.id === subject?.level_id)?.name})` }}</p>
    <p class="rate-details">${{ rate.amount }} / hour</p>
    <p class="rate-details">Is Open: {{ rate.is_open }}</p>
    <div v-if="Number($cookies.get('user_id')) === rate.tutor_id || Number($cookies.get('user_id')) === rate.tuition_center_id"
         style="margin-top: 10px">
      <Button @click="router.push(`/rate/${rate.id}/edit`)">
        Edit
      </Button>
      <Button
          style="margin-left: 10px; background-color: rgb(200, 0, 0); border-color: rgb(200, 0, 0)"
          @click="async () => { await deleteFn(); router.go(0) }"
      >Delete</Button>
    </div>
  </div>
</template>

<style scoped>
.rate-details {
  font-size: 20px;
}

.rate-item-ctn {
  margin-top: 20px;
  padding: 15px;
  border-radius: 10px;
}
</style>