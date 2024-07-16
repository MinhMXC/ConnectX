<script setup>

import Button from "primevue/button";
import router from "@/router.js";
import useBackendDelete from "@/composables/useBackendDelete.js";
import {computed, inject} from "vue";

const props = defineProps(['qualification', 'levelData']);
const qualification = computed(() => props.qualification);
const $cookies = inject("$cookies");
const { deleteFn } = useBackendDelete(`/qualification/${qualification.value.id}`);

</script>

<template>
  <div class="qualification-item-ctn white-elevation">
    <p class="qualification-details">{{qualification.name}}</p>
    <p class="qualification-details">{{qualification.description}}</p>
    <p :key="levelData" class="qualification-details">
      Level: {{levelData?.find(level => level.id === qualification.level_id)?.name}}
    </p>
    <p class="qualification-details">Attained: {{new Date(qualification.time).toDateString()}}</p>
    <div v-if="Number($cookies.get('user_id')) === qualification.tutor_id" style="margin-top: 10px">
      <Button @click="router.push(`/qualification/${qualification.id}/edit`)">
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
.qualification-details {
  font-size: 20px;
}

.qualification-item-ctn {
  margin-top: 20px;
  padding: 15px;
  border-radius: 10px;
}
</style>