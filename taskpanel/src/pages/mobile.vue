<template>
  <v-container fluid id="taskpanel" class="fill-height d-flex flex-column">
    <v-dialog v-model="settings" :persistent="!anumber">
      <v-card class="ma-auto w-100" max-width="400" prepend-icon="mdi-settings">
        <v-card-title>Settings</v-card-title>
        <v-card-subtitle>Set your A# to track tasks</v-card-subtitle>
        <v-card-item>
          <v-text-field id="anum" ref="anum" :prepend-icon="result" v-model="anumber" @input="anumCheck" @keyup.enter="submitAnum" hint="Enter the A# from your student ID" label="A#"></v-text-field>
        </v-card-item>
        <v-card-actions>
          <v-btn class="ms-auto" text="Close" v-if="anumber" @click="settings = false"></v-btn>
          <v-spacer></v-spacer>
          <v-btn class="ms-auto" text="Save" @click="submitAnum"></v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
    <v-row id="filters" class="align-self-start d-flex w-100 flex-grow-0">
      <v-col cols="9" sm="4" md="7">
        <v-text-field id="search" v-model="search" clearable label="Search" hint="Search for tasks by name or description"></v-text-field>
      </v-col>
      <v-col cols="3" class="mt-2 d-flex d-sm-none">
        <v-btn @click="settings = true" variant="tonal">
          <v-icon>mdi-cog</v-icon>
        </v-btn>
      </v-col>
      <v-col cols="6" sm="3" md="2">
        <v-combobox clearable chips multiple label="Tags" v-model="selectedTags" :items="tasktags"></v-combobox>
      </v-col>
      <v-col cols="6" sm="3" md="2" class="d-flex align-self-start">
        <v-switch label="Show All" inset color="secondary" v-model="showall"></v-switch>
      </v-col>
      <v-col cols="1" class="mt-2 d-none d-sm-flex">
        <v-btn @click="settings = true" variant="tonal">
          <v-icon>mdi-cog</v-icon>
        </v-btn>
      </v-col>
    </v-row>
    <v-row align="center" justify="center" class="d-flex flex-row w-100">
      <v-col v-for="task in tasklist" class="d-flex flex-column" cols="12" sm="4" md="3" lg="2">
        <v-card class="task-card d-flex flex-column text-center" :class="{ 'selected' : selected == task.ID }" variant="tonal" @click="selectTask(task.ID)">
          <v-card-item>
            <v-card-title>{{ task.name }}</v-card-title>
            <v-card-subtitle v-if="workingdata[task.ID]">
              {{ workingdata[task.ID] }} {{ workingdata[task.ID] > 1 ? 'People' : 'Person' }} working
            </v-card-subtitle>
          </v-card-item>
          <v-card-text>
            {{ task.description }}
          </v-card-text>
        </v-card>
      </v-col>
      <v-col cols="12">
        <v-btn class="bigbutton" :class="{ 'selected' : selected == 0 }" variant="tonal" @click="selectTask(0)">
          Not Tracking Time
        </v-btn>
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts" setup>
  const focusFilter = [12, 13, 15, 19, 20, 23, 27, 28, 31, 33, 35, 39, 42, 43, 46, 48, 49, 57]
  const settings = ref(false)
  const loading = ref(false)
  const showall = ref(false)
  const selectedTags = ref([])
  const search = ref('')
  const anumber = ref('')
  const hash = ref('')
  const selected = ref(0)
  const result = ref('mdi-form-textbox')
  const taskdata = ref({})
  const workingdata = ref({})
  const anum = useTemplateRef('anum')
  const tasktags = computed(() => {
    const tags = new Set()
    for (const task of Array.from(taskdata.value)) {
      for (const tag of task.tags) {
        tags.add(tag.name)
      }
    }
    return Array.from(tags)
  })
  const tasklist = computed(() => {
    let tasks = Array.from(taskdata.value)
    if (!showall.value) {
      tasks = tasks.filter(task => focusFilter.includes(task.ID))
    }
    if (selectedTags.value.length > 0) {
      tasks = tasks.filter(task => task.tags.some(tag => selectedTags.value.includes(tag.name)))
    }
    if (search.value) {
      tasks = tasks.filter(task => (task.name + task.description).toUpperCase().includes(search.value.toUpperCase()))
    }
    return tasks
  })
  const selectTask = (taskID:number) => {
    if (selected.value == taskID) {
      return
    }
    selected.value = taskID
    if (selected.value > 0) {
      clockOn(selected.value)
    } else {
      clockOff()
    }
  }
  const getTasks = async () => {
    loading.value = true
    try {
      const response = await fetch('https://api.tesc.farm/tasks')
      if (!response.ok) {
        console.log(response.status)
      }
      taskdata.value = await response.json()
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
      const response = await fetch('https://api.tesc.farm/hours/working')
      if (!response.ok) {
        console.log(response.status)
      }
      taskdata.value.forEach(task => {
        workingdata.value[task.ID] = 0
      })
      const jsondata = await response.json()
      jsondata.forEach(punch => {
        workingdata.value[punch.task_id]++
        if (punch.worker.barcode == hash.value) {
          selected.value = punch.task_id
        }
      })
    } catch (e) {
      console.log(e)
    } finally {
      loading.value = false
    }
  }
  const submitAnum = () => {
    if (anumber.value == '') {
      return
    }
    localStorage.setItem('anumber', anumber.value)
    settings.value = false
    setHash()
  }
  const anumCheck = () => {
    if (anumber.value.length == 9 && anumber.value[0] == 'A' && !isNaN(Number(anumber.value.substring(1)))) {
      result.value = 'mdi-check-circle'
    } else {
      result.value = 'mdi-form-textbox'
    }
  }
  const clockOn = async (taskID:number) => {
    const data = {barcode: anumber.value, task: taskID}
    const response = await fetch('https://api.tesc.farm/hours/punch', {method: 'POST', body: JSON.stringify(data)})
    if (!response.ok) {
      console.log(response)
    }
    updateWorking()
  }
  const clockOff = async () => {
    const data = {barcode: anumber.value}
    const response = await fetch('https://api.tesc.farm/hours/punch', {method: 'POST', body: JSON.stringify(data)})
    if (!response.ok) {
      console.log(response)
    }
    updateWorking()
  }
  const setHash = async () => {
    const data = {barcode: anumber.value}
    const response = await fetch('https://api.tesc.farm/worker/lookup', {method: 'POST', body: JSON.stringify(data)})
    if (!response.ok) {
      console.log(response)
    }
    const jsondata = await response.json()
    hash.value = jsondata.barcode
  }
  let intervalID
  onMounted(() => {
    anumber.value = localStorage.getItem('anumber')
    if (!anumber.value) {
      settings.value = true
    } else {
      anumCheck()
      setHash()
    }
    getTasks()
    intervalID = setInterval(updateWorking, 60000)
  })
  onBeforeUnmount(() => {
    clearInterval(intervalID);
  });
</script>
