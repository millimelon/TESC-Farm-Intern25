<template>
  <v-container fluid id="taskpanel" class="fill-height d-flex flex-column withfloat">
    <v-row id="filters" class="align-self-start d-flex w-100">
      <v-col cols="12" sm="4" md="8">
        <v-text-field id="search" v-model="search" clearable label="Search" hint="Search for tasks by name or description"></v-text-field>
      </v-col>
      <v-col cols="6" sm="4" md="2">
        <v-combobox clearable chips multiple label="Tags" v-model="selectedTags" :items="tasktags"></v-combobox>
      </v-col>
      <v-col cols="6" sm="4" md="2" class="d-flex align-self-start">
        <v-switch label="Show All" inset color="secondary" v-model="showall"></v-switch>
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
        <v-btn class="bigbutton" :class="{ 'selected' : selected == -1 }" variant="tonal" @click="selectTask(-1)">
          Stop Tracking Time
        </v-btn>
      </v-col>
    </v-row>
  </v-container>
  <div id="anumfloat" class="align-self-end" v-if="selected">
    <v-text-field id="anum" ref="anum" :prepend-icon="result" v-model="anumber" @input="anumCheck" @keyup.enter="submitAnum" @keydown.esc="selectTask(0)" hint="Enter the A# from your student ID" :label="selectedName" autocomplete="off"></v-text-field>
  </div>
</template>

<script lang="ts" setup>
  const focusFilter = [12, 13, 15, 19, 20, 23, 27, 28, 31, 33, 35, 39, 42, 43, 46, 48, 49, 57]
  const loading = ref(false)
  const showall = ref(false)
  const selectedTags = ref([])
  const search = ref('')
  const anumber = ref('')
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
    tasks.sort((a, b) => a.name.localeCompare(b.name))
    return tasks
  })
  const selectedName = computed(() => {
    if (selected.value == 0) {
      return "None"
    }
    if (selected.value < 0) {
      return "Stop Tracking Time"
    }
    return taskdata.value.find(task => task.ID === selected.value).name
  })
  const selectTask = (taskID:number) => {
    selected.value = taskID
    if (taskID == 0) {
      return
    }
    nextTick(() => {
      anum.value.focus()
    })
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
    if (selected.value == -1) {
      clockOff(anumber.value)
    } else {
      clockOn(anumber.value, selected.value)
    }
    anumber.value = ''
  }
  const anumCheck = (e:event) => {
    if (anumber.value.length > 8) {
      submitAnum()
      e.target.focus()
    }
  }
  const clockOn = async (anum:string, taskID:number) => {
    const data = {barcode: anum, task: taskID}
    const response = await fetch('https://api.tesc.farm/hours/punch', {method: 'POST', credentials: 'include', body: JSON.stringify(data)})
    if (response.ok) {
      result.value = 'mdi-check-circle'
    } else {
      result.value = 'mdi-alert-circle'
    }
    setTimeout(() => {
      result.value = 'mdi-form-textbox'
    }, 3000)
    updateWorking()
  }
  const clockOff = async (anum:string) => {
    const data = {barcode: anum}
    const response = await fetch('https://api.tesc.farm/hours/punch', {method: 'POST', credentials: 'include', body: JSON.stringify(data)})
    if (response.ok) {
      result.value = 'mdi-check-circle'
    } else {
      result.value = 'mdi-alert-circle'
    }
    setTimeout(() => {
      result.value = 'mdi-form-textbox'
    }, 3000)
    updateWorking()
  }
  let intervalID
  onMounted(() => {
    getTasks()
    intervalID = setInterval(updateWorking, 60000)
  })
  onBeforeUnmount(() => {
    clearInterval(intervalID);
  })
</script>
