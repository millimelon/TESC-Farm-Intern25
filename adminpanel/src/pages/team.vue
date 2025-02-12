<template>
  <h1 class="v-text-h1" style="text-align: center; padding: 20px">
    Student Hours
  </h1>
  <v-data-table
    style="max-width: 90%; margin: auto; margin-bottom: 33px"
    :item-value="ID"
    :items="items"
    :headers="headers"
    show-select
    return-object
  />
</template>

<script setup lang="ts">
import items from "@/assets/hours.json";

definePage({
  meta: {
    title: "Student Hours",
  },
});

const headers = [
  { title: "Name", key: "worker.name" },
  { title: "Hours", key: "duration" },
];
</script>
