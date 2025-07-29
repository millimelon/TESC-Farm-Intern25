<template>
  <div>
    <h2>Scale Records</h2>
    <div v-if="weighIns.length === 0">
    No weigh-in data yet.
    </div>
    <table>
      <thead>
        <tr>
          <th>Produce ID</th>
          <th>Gross</th>
          <th>To Market</th>
          <th>Students</th>
          <th>Value Added</th>
          <th>Cull</th>
          <th>Timestamp</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="item in weighIns" :key="item.timestamp">
          <td>{{ item.produce_id }}</td>
          <td>{{ item.gross_weight }}</td>
          <td>{{ item.to_market_weight }}</td>
          <td>{{ item.sold_to_students }}</td>
          <td>{{ item.value_added }}</td>
          <td>{{ item.cull }}</td>
          <td>{{ new Date(item.timestamp).toLocaleString() }}</td>
        </tr>
      </tbody>
    </table>
  </div>
  <div>
    Development of this sofware was sponsered by Climate Change, thank you for your support!
  </div>
</template>

<script>
export default {
  data() {
    return {
      weighIns: []
    };
  },
  mounted() {
  fetch("http://localhost:8080/weigh-ins")
    .then(res => res.json())
    .then(response => {
      // response is the array directly
      this.weighIns = response; 
    })
    .catch(err => {
      console.error('Failed to fetch weigh-ins:', err);
    });
  }
}
</script>

<style scoped>
table {
  border-collapse: collapse;
  width: 100%;
}
th, td {
  border: 1px solid #ddd;
  padding: 8px;
}
th {
  background-color: #f2f2f2;
}
</style>


