<script setup>
import Checkbox from "primevue/checkbox";
import {ref} from "vue";
import Button from "primevue/button";
import CXSelect from "@/components/form_input/CXSelect.vue";

const gender = ref("");
const isParent = ref(false);
const genderOptions = [ "Male", "Female", "Others" ]
</script>

<template>
  <CXSelect label="Gender" :options="genderOptions" v-model="gender" />

  <Checkbox id="is-parent-checkbox" v-model="isParent" binary />
  <label for="is-parent-checkbox">Are you a parent?</label>

  <br />

  <div style="display: flex; flex-direction: row-reverse">
    <Button id="submit-button" label="Submit" />
  </div>
</template>

<style scoped>
#is-parent-checkbox {
  margin-right: 10px;
  margin-top: 15px;
}

#submit-button {
  margin-top: 25px;
}
</style>