<script setup>
import {ref} from "vue";
import CXInputText from "@/components/form_input/CXInputText.vue";
import CXInputNumber from "@/components/form_input/CXInputNumber.vue";
import CXTextarea from "@/components/form_input/CXTextarea.vue";
import Button from "primevue/button";

const name = ref("");
const handphone = ref();
const address = ref("");
const addressLink = ref("");
const description = ref("");
const website = ref("");
</script>

<template>
  <div class="vertical-form-container">
    <CXInputText label="Name" v-model="name" />
    <CXInputNumber label="Handphone" v-model="handphone" prefix="+65 " />
    <CXInputText label="Address" v-model="address" />
    <CXInputText label="Address Link" v-model="addressLink" />
    <CXTextarea label="Website" v-model="website" :rows="1" placeholder="So users can find out more about you" />
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