<template>
  <v-container fluid id="taskpanel" class="fill-height d-flex flex-column">
    <v-row id="filters" class="align-self-start d-flex flex-column w-100">
      <v-col cols="12">
        <v-text-field id="search" v-model="search" label="Search" hint="Search for tasks by name, tag, or description"></v-text-field>
        <v-btn v-for="tag in tags">{{ tag.name }}</v-btn>
        <v-btn v-if="hidden" @click="showall">Show All</v-btn>
        <v-btn v-if="filters" @click="resetfilters">Reset</v-btn>
      </v-col>
    </v-row>
    <v-row align="center" justify="center" class="d-flex flex-row w-100">
      <v-col v-for="task in taskdata" class="d-flex flex-column" cols="12" sm="4" md="3" lg="2">
        <v-card class="task-card d-flex flex-column text-center" :class="{ 'selected' : selected == task.ID }" variant="tonal" @click="selectTask(task.ID)">
          <template v-slot:prepend>
            <v-img src="@/assets/placeholder.png" class="taskicon"></v-img>
          </template>
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
      <v-col cols="12">
        <v-btn class="bigbutton" :class="{ 'selected' : selected == -1 }" variant="tonal" @click="selectTask(-1)">
          Clock Off
        </v-btn>
      </v-col>
    </v-row>
  </v-container>
  <div id="anumfloat" class="align-self-end" v-if="selected">
    <v-text-field id="anum" ref="anum" v-model="anumber" @input="anumCheck" @keyup.enter="submitAnum" @keydown.esc="selectTask(0)" hint="Enter the A# from your student ID" :label="selectedName"></v-text-field>
  </div>
</template>

<script lang="ts" setup>
  const loading = ref(false)
  const selected = ref(0)
  const anumber = ref('')
  const search = ref('')
  const response = ref('')
  const taskdata = ref({})
  const workingdata = ref({})
  const anum = useTemplateRef('anum')
  const selectedName = computed(() => {
    if (selected.value == 0) {
      return "None"
    }
    if (selected.value < 0) {
      return "Clock Off"
    }
    return taskdata.value.find(task => task.ID === selected.value).name
  })
  const selectTask = (taskID:number) => {
    selected.value = taskID
    if (taskID == 0) {
      console.log('DESELECT')
      return
    }
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
    const response = await fetch('https://json.tesc.farm/hours/punch', {method: 'POST', body: JSON.stringify(data)})
    const jsondata = await response.json();
    updateWorking()
  }
  const clockOff = async (anum:string) => {
    const data = {barcode: anum}
    const response = await fetch('https://json.tesc.farm/hours/punch', {method: 'POST', body: JSON.stringify(data)})
    const jsondata = await response.json();
    updateWorking()
  }
  onMounted(getTasks);
</script>
