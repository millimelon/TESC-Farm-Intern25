<template>
  <div>
    <h2>Weigh-In Records</h2>
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
      this.weighIns = response.data;
    });
}
</script>
