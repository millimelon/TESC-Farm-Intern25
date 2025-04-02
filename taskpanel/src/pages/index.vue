<template>
  <v-container id="taskpanel" fill-height>
    <v-row align="center" justify="center">
      <v-col v-for="task in taskdata" class="d-flex flex-column" cols="12" sm="6" md="4">
        <v-card class="task-card d-flex flex-column flex-grow-1" :class="{ 'selected' : selected == task.ID }" variant="tonal" @click="selectTask(task.ID)">
          <v-card-item>
            <v-card-title>{{ task.name }}</v-card-title>
            <v-card-subtitle v-if="workingdata[task.ID]">
              {{ workingdata[task.ID] }} People working
            </v-card-subtitle>
          </v-card-item>
          <v-card-text>
            {{ task.description }}
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
    <v-btn class="bigbutton" :class="{ 'selected' : selected == -1 }" variant="tonal" @click="selectTask(-1)">
      Clock Off
    </v-btn>
  </v-container>
  <div id="anumfloat" v-if="selected">
    <v-text-field id="anum" ref="anum" v-model="anumber" @input="anumCheck" hint="Enter the A# from your student ID" label="A#"></v-text-field>
  </div>
</template>

<script lang="ts" setup>
  const loading = ref(false)
  const selected = ref(0)
  const anumber = ref('')
  const taskdata = ref({})
  const workingdata = ref({})
  const anum = useTemplateRef('anum')
  const selectTask = (taskID:number) => {
    selected.value = taskID
    nextTick(() => {
      anum.value.focus()
    })
  }
  const getTasks = async () => {
    loading.value = true
    try {
      const response = await fetch('https://json.tesc.farm/tasks');
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
      const response = await fetch('https://json.tesc.farm/hours/working');
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
  const anumCheck = (e) => {
    if (anumber.value.length > 8) {
      if (selected.value == -1) {
        clockOff(anumber.value)
      } else {
        clockOn(anumber.value, selected.value)
      }
      anumber.value = ""
      e.target.focus()
    }
  }
  const clockOn = async (anum, taskID) => {
    const data = {barcode: anum, task: taskID}
    const response = await fetch("https://json.tesc.farm/hours/punch", {method: "POST", body: JSON.stringify(data)})
    const jsondata = await response.json();
    updateWorking()
  }
  const clockOff = async (anum) => {
    const data = {barcode: anum}
    const response = await fetch("https://json.tesc.farm/hours/punch", {method: "POST", body: JSON.stringify(data)})
    const jsondata = await response.json();
    updateWorking()
  }
  onMounted(getTasks);
</script>
