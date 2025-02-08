<template>
  <v-data-table :item-value="ID" :items="items" :headers="headers" show-select return-object>
  </v-data-table>
</template>

<script lang="ts" setup>
import items from "@/assets/hours.json";
for (const hours of items) {
  hours.crop = "N/A";
  hours.worktype = "Other"
  if (hours.process) {
    hours.crop = hours.process.harvest.crop.name;
    hours.worktype = "Processing";
  } else if (hours.harvest) {
    hours.crop = hours.harvest.crop.name;
    hours.worktype = "Harvesting"
  }
}
const headers = [
  { title: 'Crop', key: 'crop'},
  { title: 'Work Type', key: 'worktype'},
  { title: 'Name', key: 'worker.name' },
  { title: 'Start', key: 'start' },
  { title: 'Hours', key: 'duration' },
];
</script>
