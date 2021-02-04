<template>
  <div>
    {{users}}
  </div>
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      users: [],
    }
  },
  mounted() {
    axios.get("/api/users").then((res) => {
      for (let i = 0; i < res.data.length; i ++) {
        console.log(res.data);
        this.users.push(res.data[i]);
      }
    })
  },
}
</script>