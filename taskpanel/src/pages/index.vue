<template>
  <v-container>
    <h1>Tasks:</h1>
    <v-row align="stretch">
      <v-col v-for="task in taskdata" class="d-flex flex-column" cols="12" sm="6" md="4">
        <v-card class="task-card d-flex flex-column flex-grow-1" variant="tonal">
          <v-card-item>
            <v-card-title>{{ task.name }}</v-card-title>
            <v-card-subtitle v-if="workingdata[task.ID]">
              {{ workingdata[task.ID] }} People working
            </v-card-subtitle>
          </v-card-item>
          <v-card-text>
            {{ task.description }}
          </v-card-text>
          <v-spacer></v-spacer>
          <v-card-actions>
            <v-btn class="bigbutton" variant="tonal" @click="clockOn(task.ID)">
              Choose Task
            </v-btn>
          </v-card-actions>
        </v-card>
      </v-col>
    </v-row>
    <v-btn class="bigbutton" variant="tonal" @click="clockOff()">
      Clock Off
    </v-btn>
  </v-container>
</template>

<script lang="ts" setup>
  const loading = ref(false)
  const taskdata = ref({})
  const workingdata = ref({})
  const getTasks = async () => {
    loading.value = true
    try {
      const response = await fetch('https://jsondemo.tesc.farm/tasks');
      if (!response.ok) {
        console.log(response.status)
      }
      taskdata.value = await response.json();
      taskdata.value.forEach(task => {
        workingdata.value[task.ID] = 0
      })
    } catch (e) {
      console.log(e)
    } finally {
      updateWorking()
    }
  }
  const updateWorking = async () => {
    loading.value = true
    try {
      const response = await fetch('https://jsondemo.tesc.farm/hours/working');
      if (!response.ok) {
        console.log(response.status)
      }
      taskdata.value.forEach(task => {
        workingdata.value[task.ID] = 0
      })
      const jsondata = await response.json();
      jsondata.forEach(punch => {
        workingdata.value[punch.task_id]++
      })
    } catch (e) {
      console.log(e)
    } finally {
      loading.value = false;
    }
  }
  const clockOn = async (taskID) => {
    let anum = prompt("Enter your A#:")
    const data = {barcode: anum, task: taskID}
    const response = await fetch("https://jsondemo.tesc.farm/hours/punch", {method: "POST", body: JSON.stringify(data)})
    const jsondata = await response.json();
    updateWorking()
  }
  const clockOff = async () => {
    let anum = prompt("Enter your A#:")
    const data = {barcode: anum}
    const response = await fetch("https://jsondemo.tesc.farm/hours/punch", {method: "POST", body: JSON.stringify(data)})
    const jsondata = await response.json();
    updateWorking()
  }
  onMounted(getTasks);
</script>
