<script setup>
import {ref} from "vue";
import CXInputText from "@/components/form_input/CXInputText.vue";
import CXInputNumber from "@/components/form_input/CXInputNumber.vue";
import CXSelect from "@/components/form_input/CXSelect.vue";
import CXTextarea from "@/components/form_input/CXTextarea.vue";
import Button from "primevue/button";

const name = ref("");
const age = ref();
const gender = ref("");
const handphone = ref();
const description = ref("");

const genderOptions = [ "Male", "Female", "Others" ]
</script>

<template>
  <div class="vertical-form-container">
    <CXInputText label="Name" v-model="name" />
    <CXInputNumber label="Age" v-model="age" />
    <CXSelect label="Gender" :options="genderOptions" v-model="gender" />
    <CXInputNumber label="Handphone" v-model="handphone" prefix="+65 " />
    <CXTextarea label="Description" v-model="description" :rows="10" placeholder="Tell users more about yourself" />

    <div style="display: flex; flex-direction: row-reverse">
      <Button id="submit-button" label="Submit" />
    </div>
  </div>
</template>

<style scoped>
.vertical-form-container {
  display: flex;
  flex-direction: column;
  gap: 10px;
}
</style>