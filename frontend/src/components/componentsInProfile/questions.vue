<template>
  <div id="users-questions">
    <div class="collection">
      <router-link v-for="(question, id) in questions" :key="id" :to="'/question/' + question.id" class="collection-item">{{question.title}}</router-link>
    </div>
  </div>
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      questions: [],
    }
  },
  mounted() {
    // get user's questions
    axios.get(`/api/questions/${this.$route.params.id}`).then((res) => {
      for (let i = 0; i < res.data.length; i ++) {
        this.questions.push(res.data[i])
      }
    })
  },
}
</script>