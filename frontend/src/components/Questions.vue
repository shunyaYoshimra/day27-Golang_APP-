<template>
  <div id="questions">
    <select v-model="about" @change="SearchByAbout()">
      <option value disabled selected>トピック検索(Search by topics)</option>
      <option value="classes">授業(classes)</option>
      <option value="seminer">ゼミ(seminer)</option>
      <option value="club activities">部活/サークル(club activities)</option>
      <option value="study abroad">留学(study abroad)</option>
      <option value="look for a job">就活(look for a job)</option>
      <option value="off capmus activities">校外活動(off capmus activities)</option>
      <option value="part time job">バイト(part time job)</option>
      <option value="other">その他(other)</option>
    </select>
    <label>Materialize Select</label>
    <div class="collection">
      <router-link
        v-for="(question, id) in questions"
        :key="id"
        :to="{ name: 'question', params: {id: question.id}}"
        class="collection-item"
      >
        {{question.title}}
        <span v-if="question.status">解決済み(solved)</span>
      </router-link>
    </div>
  </div>
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      questions: [],
      about: ""
    };
  },
  mounted() {
    // jquery for select
    $(document).ready(function() {
      $("select").formSelect();
    });
    // get questions
    axios.get("/api/questions").then(res => {
      for (let i = 0; i < res.data.length; i++) {
        this.questions.push(res.data[i]);
      }
    });
  },
  methods: {
    SearchByAbout() {
      this.questions = [];
      axios
        .get(`/api/searched_questions/${this.about}`)
        .then(res => {
          for (let i = 0; i < res.data.length; i++) {
            this.questions.push(res.data[i]);
          }
        })
        .catch(err => {
          console.log(err.response);
        });
    }
  }
};
</script>

<style lang="scss" scoped>
#questions {
  width: 600px;
  margin: 0 auto;

  .collection {
    position: relative;
    word-break: break-all;
    span {
      position: absolute;
      right: 20px;
      color: #333;
    }
  }
}
@media (max-width: 480px) {
  #questions {
    width: 98%;
    margin: auto;
  }
}
</style>